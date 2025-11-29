package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"iot/data_simulator/common"
	"iot/internal/env"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gopkg.in/square/go-jose.v2"
)

const (
	NumDevices = 5

	CPUChanSize  = 500
	CPUBatchSize = 100

	TemperatureChanSize  = 500
	TemperatureBatchSize = 100

	HumidityChanSize  = 500
	HumidityBatchSize = 100

	PressureChanSize  = 500
	PressureBatchSize = 100

	GPSChanSize  = 500
	GPSBatchSize = 100

	DefaultURL = "/static/cpu"
)

type MetricsType interface {
	GenerateData(metricsChan chan common.Metrics)
}

func VerifyIDToken(idToken string) (map[string]interface{}, error) {
	// --- Step 1: Decode JWT header without validating ---
	parts := strings.Split(idToken, ".")
	if len(parts) < 2 {
		return nil, errors.New("invalid JWT format")
	}

	headerJSON, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, fmt.Errorf("cannot decode header: %v", err)
	}

	var header struct {
		Kid string `json:"kid"`
		Alg string `json:"alg"`
	}
	if err := json.Unmarshal(headerJSON, &header); err != nil {
		return nil, fmt.Errorf("cannot parse header: %v", err)
	}

	kid := header.Kid
	if kid == "" {
		return nil, errors.New("kid missing in token header")
	}

	// --- Step 2: Fetch JWKS from Okta ---
	jwksURL := env.GetString("OKTA_ISSUER", "") + "/v1/keys"
	resp, err := http.Get(jwksURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var jwks struct {
		Keys []json.RawMessage `json:"keys"`
	}
	json.Unmarshal(body, &jwks)

	// --- Step 3: Find matching key ---
	var foundKey jose.JSONWebKey
	for _, k := range jwks.Keys {
		var key jose.JSONWebKey
		json.Unmarshal(k, &key)
		if key.KeyID == kid {
			foundKey = key
			break
		}
	}

	if foundKey.Key == nil {
		return nil, errors.New("kid not found in JWKS")
	}

	// --- Step 4: Verify token signature with keyfunc ---
	parsed, err := jwt.Parse(idToken, func(token *jwt.Token) (interface{}, error) {
		return foundKey.Key, nil
	})
	if err != nil {
		return nil, err
	}

	if !parsed.Valid {
		return nil, errors.New("id token is invalid")
	}

	// --- Step 5: Return claims ---
	return parsed.Claims.(jwt.MapClaims), nil
}

func CreateInternalJWT(email, sub string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"sub":   sub,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(env.GetString("JWT_SECRET", "")))
}
