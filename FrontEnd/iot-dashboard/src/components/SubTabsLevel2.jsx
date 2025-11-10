import { NavLink, useLocation } from "react-router-dom";

function SubTabsLevel2() {
  const location = useLocation();

  // Only render for /aggregated/*
  if (!location.pathname.startsWith("/aggregated")) return null;

  // Extract which metric (cpu, temperature, etc.) you're on
  const parts = location.pathname.split("/");
  const metric = parts[2]; // e.g., "cpu", "temperature", etc.

  const subTabs = [
    { path: `/aggregated/${metric}/location`, label: "Per Location" },
    { path: `/aggregated/${metric}/model`, label: "Per Model" },
    { path: `/aggregated/${metric}/daily`, label: "Daily Summary" },
  ];

  return (
    <ul className="nav nav-pills bg-light px-3 border-bottom justify-content-center">
      {subTabs.map((tab) => (
        <li className="nav-item" key={tab.path}>
          <NavLink
            to={tab.path}
            className={({ isActive }) =>
              `nav-link py-2 px-3 ${isActive ? "active fw-semibold" : ""}`
            }
          >
            {tab.label}
          </NavLink>
        </li>
      ))}
    </ul>
  );
}

export default SubTabsLevel2;
