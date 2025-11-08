import { NavLink, useLocation } from "react-router-dom";

function SubTabs() {
  const location = useLocation();

  const getSubTabs = () => {
    if (location.pathname.startsWith("/static")) {
      return [
        { path: "/static/cpu", label: "CPU" },
        { path: "/static/temperature", label: "Temperature" },
        { path: "/static/pressure", label: "Pressure" },
        { path: "/static/humidity", label: "Humidity" },
        { path: "/static/gps", label: "GPS" },
      ];
    } else if (location.pathname.startsWith("/aggregated")) {
      return [
        { path: "/aggregated/cpu", label: "CPU" },
        { path: "/aggregated/temperature", label: "Temperature" },
        { path: "/aggregated/pressure", label: "Pressure" },
        { path: "/aggregated/humidity", label: "Humidity" },
        { path: "/aggregated/gps", label: "GPS" },
      ];
    } else if (location.pathname.startsWith("/correlated")) {
      return [
        { path: "/correlated/correlation", label: "Correlation" },
        { path: "/correlated/rolling-gps", label: "Rolling GPS" },
      ];
    }
    return [];
  };

  const subTabs = getSubTabs();

  if (subTabs.length === 0) return null;

  return (
    <ul className="nav nav-tabs bg-white px-3">
      {subTabs.map((tab) => (
        <li className="nav-item" key={tab.path}>
          <NavLink
            to={tab.path}
            className={({ isActive }) =>
              `nav-link ${isActive ? "active fw-semibold" : ""}`
            }
          >
            {tab.label}
          </NavLink>
        </li>
      ))}
    </ul>
  );
}

export default SubTabs;
