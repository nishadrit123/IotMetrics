import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import Navbar from "./components/Navbar";
import SubTabs from "./components/SubTabs";
import SubTabsLevel2 from "./components/SubTabsLevel2";
import PaginationBar from "./components/PaginationBar";

// Login Page
import Login from "./pages/Login";

// Static Data Pages
import CPU from "./pages/Static/CPU";
import Temperature from "./pages/Static/Temperature";
import Pressure from "./pages/Static/Pressure";
import Humidity from "./pages/Static/Humidity";
import GPS from "./pages/Static/GPS";

// Aggregated Data Pages
import CPUAggLoc from "./pages/Aggregated/CPU/PerLocation";
import CPUAggModel from "./pages/Aggregated/CPU/PerModel";
import CPUAggDaily from "./pages/Aggregated/CPU/DailySummary";

import TempAggLoc from "./pages/Aggregated/Temperature/PerLocation";
import TempAggModel from "./pages/Aggregated/Temperature/PerModel";
import TempAggDaily from "./pages/Aggregated/Temperature/DailySummary";

import PressureAggLoc from "./pages/Aggregated/Pressure/PerLocation";
import PressureAggModel from "./pages/Aggregated/Pressure/PerModel";
import PressureAggDaily from "./pages/Aggregated/Pressure/DailySummary";

import HumidityAggLoc from "./pages/Aggregated/Humidity/PerLocation";
import HumidityAggModel from "./pages/Aggregated/Humidity/PerModel";
import HumidityAggDaily from "./pages/Aggregated/Humidity/DailySummary";

import GPSAggLoc from "./pages/Aggregated/GPS/PerLocation";
import GPSAggModel from "./pages/Aggregated/GPS/PerModel";
import GPSAggDaily from "./pages/Aggregated/GPS/DailySummary";

// Correlated Data Pages
import Correlation from "./pages/Correlated/Correlation";
import RollingGPS from "./pages/Correlated/RollingGPS";

function App() {
  return (
    <Router>
      <div className="bg-light min-vh-100">
        {/* Top Navbar (Main Tabs) */}
        <Navbar />

        {/* Sub Tabs (Dynamic per main tab) */}
        <SubTabs />
        <SubTabsLevel2 />

        {/* Page Content */}
        <div className="container py-4">
          <Routes>
            {/* Default redirect */}
            <Route path="/" element={<Navigate to="/authentication/login" replace />} />

            {/* Login path */}
            <Route path="/authentication/login" element={<Login />} />

            {/* Static Data */}
            <Route path="/static/cpu" element={<CPU />} />
            <Route path="/static/temperature" element={<Temperature />} />
            <Route path="/static/pressure" element={<Pressure />} />
            <Route path="/static/humidity" element={<Humidity />} />
            <Route path="/static/gps" element={<GPS />} />

            {/* Aggregated Data */}

            <Route path="/aggregated/cpu" element={<Navigate to="/aggregated/cpu/location" replace />} />
            <Route path="/aggregated/cpu/location" element={<CPUAggLoc />} />
            <Route path="/aggregated/cpu/model" element={<CPUAggModel />} />
            <Route path="/aggregated/cpu/daily" element={<CPUAggDaily />} />

            <Route path="/aggregated/temperature" element={<Navigate to="/aggregated/temperature/location" replace />} />
            <Route path="/aggregated/temperature/location" element={<TempAggLoc />} />
            <Route path="/aggregated/temperature/model" element={<TempAggModel />} />
            <Route path="/aggregated/temperature/daily" element={<TempAggDaily />} />

            <Route path="/aggregated/pressure" element={<Navigate to="/aggregated/pressure/location" replace />} />
            <Route path="/aggregated/pressure/location" element={<PressureAggLoc />} />
            <Route path="/aggregated/pressure/model" element={<PressureAggModel />} />
            <Route path="/aggregated/pressure/daily" element={<PressureAggDaily />} />

            <Route path="/aggregated/humidity" element={<Navigate to="/aggregated/humidity/location" replace />} />
            <Route path="/aggregated/humidity/location" element={<HumidityAggLoc />} />
            <Route path="/aggregated/humidity/model" element={<HumidityAggModel />} />
            <Route path="/aggregated/humidity/daily" element={<HumidityAggDaily />} />

            <Route path="/aggregated/gps" element={<Navigate to="/aggregated/gps/location" replace />} />
            <Route path="/aggregated/gps/location" element={<GPSAggLoc />} />
            <Route path="/aggregated/gps/model" element={<GPSAggModel />} />
            <Route path="/aggregated/gps/daily" element={<GPSAggDaily />} />
            
            {/* Correlated Data */}
            <Route path="/correlated/correlation" element={<Correlation />} />
            <Route path="/correlated/rolling-gps" element={<RollingGPS />} />
          </Routes>
        </div>
      </div>
    </Router>
  );
}

export default App;
