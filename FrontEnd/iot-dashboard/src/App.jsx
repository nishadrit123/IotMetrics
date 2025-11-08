import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import Navbar from "./components/Navbar";
import SubTabs from "./components/SubTabs";
import PaginationBar from "./components/PaginationBar";

// Static Data Pages
import CPU from "./pages/Static/CPU";
import Temperature from "./pages/Static/Temperature";
import Pressure from "./pages/Static/Pressure";
import Humidity from "./pages/Static/Humidity";
import GPS from "./pages/Static/GPS";

// Aggregated Data Pages
import CPUAgg from "./pages/Aggregated/CPU";
import TempAgg from "./pages/Aggregated/Temperature";
import PressureAgg from "./pages/Aggregated/Pressure";
import HumidityAgg from "./pages/Aggregated/Humidity";
import GPSAgg from "./pages/Aggregated/GPS";

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

        {/* Page Content */}
        <div className="container py-4">
          <Routes>
            {/* Default redirect */}
            <Route path="/" element={<Navigate to="/static/cpu" replace />} />

            {/* Static Data */}
            <Route path="/static/cpu" element={<CPU />} />
            <Route path="/static/temperature" element={<Temperature />} />
            <Route path="/static/pressure" element={<Pressure />} />
            <Route path="/static/humidity" element={<Humidity />} />
            <Route path="/static/gps" element={<GPS />} />

            {/* Aggregated Data */}
            <Route path="/aggregated/cpu" element={<CPUAgg />} />
            <Route path="/aggregated/temperature" element={<TempAgg />} />
            <Route path="/aggregated/pressure" element={<PressureAgg />} />
            <Route path="/aggregated/humidity" element={<HumidityAgg />} />
            <Route path="/aggregated/gps" element={<GPSAgg />} />

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
