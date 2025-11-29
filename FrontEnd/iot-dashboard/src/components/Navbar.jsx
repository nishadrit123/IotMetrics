import { NavLink, useLocation } from "react-router-dom";
import { FiLogOut } from "react-icons/fi"; // Using react-icons for a logout icon

function Navbar() {
  const location = useLocation();

  const isActiveGroup = (group) => location.pathname.startsWith(`/${group}`);

  const handleLogout = async () => {
    try {
      await fetch("http://localhost:8080/v1/logout", {
        method: "POST",
        credentials: "include",
      });
      // Redirect to login
      window.location.href = "/authentication/login";
    } catch (err) {
      console.error("Logout failed:", err);
      window.location.href = "/authentication/login";
    }
  };

  return (
    <nav className="navbar navbar-expand-lg navbar-dark bg-dark px-4 py-2 shadow-sm">
      <a className="navbar-brand fw-bold text-uppercase" href="#">
        IoT Metrics Dashboard
      </a>

      <button
        className="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarNav"
        aria-controls="navbarNav"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span className="navbar-toggler-icon"></span>
      </button>

      <div className="collapse navbar-collapse" id="navbarNav">
        <ul className="navbar-nav ms-auto">
          <li className="nav-item me-5">
            <NavLink
              to="/static/cpu"
              className={`nav-link ${
                isActiveGroup("static") ? "active fw-bold" : ""
              }`}
            >
              Static Data
            </NavLink>
          </li>

          <li className="nav-item me-5">
            <NavLink
              to="/aggregated/cpu"
              className={`nav-link ${
                isActiveGroup("aggregated") ? "active fw-bold" : ""
              }`}
            >
              Aggregated Data
            </NavLink>
          </li>

          <li className="nav-item me-3">
            <NavLink
              to="/correlated/correlation"
              className={`nav-link ${
                isActiveGroup("correlated") ? "active fw-bold" : ""
              }`}
            >
              Correlated Data
            </NavLink>
          </li>

          {/* Logout icon */}
          <li className="nav-item">
            <button
              className="btn btn-link nav-link p-0"
              onClick={handleLogout}
              title="Logout"
            >
              <FiLogOut size={20} />
            </button>
          </li>
        </ul>
      </div>
    </nav>
  );
}

export default Navbar;
