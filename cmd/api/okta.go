package main

import (
	"encoding/json"
	"fmt"
	"io"
	"iot/internal/env"
	"net/http"
	"net/url"
	"strings"
)

type CredsPayload struct {
	Email string `json:"email"`
}

type OktaResponse struct {
	RedirectURL string `json:"redirect_url"`
}

type OktaTokenResponse struct {
	AccessToken string `json:"access_token"`
	IDToken     string `json:"id_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func (app *application) OktaLogin(w http.ResponseWriter, r *http.Request) {
	var payload CredsPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	oktaDomain := env.GetString("OKTA_DOMAIN", "")
	clientID := env.GetString("OKTA_CLIENT_ID", "")
	redirectURI := env.GetString("REDIRECT_URI", "")

	// Build authorize URL for Authorization Code Flow
	q := url.Values{}
	q.Set("client_id", clientID)
	q.Set("redirect_uri", redirectURI)
	q.Set("response_type", "code")  // IMPORTANT
	q.Set("response_mode", "query") // IMPORTANT
	q.Set("scope", "openid profile email")
	q.Set("state", "xyz123")           // should be randomized in real apps
	q.Set("nonce", "abc123")           // should be randomized in real apps
	q.Set("login_hint", payload.Email) // pre-fill email input

	authURL := fmt.Sprintf("https://%s/oauth2/default/v1/authorize?%s", oktaDomain, q.Encode())

	resp := OktaResponse{RedirectURL: authURL}
	if err := app.jsonResponse(w, http.StatusOK, resp); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) OktaCallBack(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	if code == "" {
		app.badRequestResponse(w, r, fmt.Errorf("authorization code missing"))
		return
	}

	// Validate state
	if state != "xyz123" {
		app.badRequestResponse(w, r, fmt.Errorf("invalid state"))
		return
	}

	// Exchange authorization code for tokens
	tokenURL := env.GetString("OKTA_ISSUER", "") + "/v1/token"

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)
	data.Set("redirect_uri", env.GetString("REDIRECT_URI", ""))
	data.Set("client_id", env.GetString("OKTA_CLIENT_ID", ""))
	data.Set("client_secret", env.GetString("OKTA_CLIENT_SECRET", ""))

	req, _ := http.NewRequestWithContext(ctx, "POST", tokenURL, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		app.internalServerError(w, r, fmt.Errorf("token exchange failed: %s", string(body)))
		return
	}

	var tok OktaTokenResponse
	json.Unmarshal(body, &tok)

	// Validate ID token (signature, exp, nonce, issuer)
	claims, err := VerifyIDToken(tok.IDToken)
	if err != nil {
		app.internalServerError(w, r, fmt.Errorf("invalid ID token: %v", err))
		return
	}

	email, _ := claims["email"].(string)
	sub, _ := claims["sub"].(string)

	// Create internal JWT (your own token)
	myjwt, err := CreateInternalJWT(email, sub)
	if err != nil {
		app.internalServerError(w, r, fmt.Errorf("failed to create internal JWT: %v", err))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    myjwt,
		HttpOnly: true,
		Secure:   false, // change to true in production (https)
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		MaxAge:   3600,
	})

	// Redirect FE with internal JWT
	redirectURL := fmt.Sprintf(
		"http://localhost:5173/static/cpu?",
	)

	http.Redirect(w, r, redirectURL, http.StatusFound)
}

func (app *application) OktaLogout(w http.ResponseWriter, r *http.Request) {
	// Clear the cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1, // delete
	})

	// Redirect to Okta logout endpoint
	oktaDomain := env.GetString("OKTA_DOMAIN", "")
	clientId := env.GetString("OKTA_CLIENT_ID", "")

	logoutURL := oktaDomain +
		"/oauth2/default/v1/logout?" +
		"id_token_hint=&" +
		"client_id=" + clientId +
		"&post_logout_redirect_uri=" +
		url.QueryEscape(env.GetString("FRONTEND_URL", "http://localhost:5173")+"/v1/authentication/login")

	http.Redirect(w, r, logoutURL, http.StatusTemporaryRedirect)
}
