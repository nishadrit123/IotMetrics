import React from "react";
import DataTable from "../../components/DataTable";

function Pressure() {
  const pressureColumns = [
    "device_id",
    "loc",
    "model",
    "manufacturer",
    "install_date",
    "baseline_pressure",
    "spike_probability",
    "spike_magnitude",
    "noise_level",
    "current_pressure",
    'drift_rate',
    "is_spiking",
    'trend',
    "last_spike_time",
    "updated_at",
  ];

  return (
    <div className="container mt-4">
      <DataTable
        apiBaseUrl="http://localhost:8080/v1/pressure/statistics"
        columns={pressureColumns}
      />
    </div>
  );
}

export default Pressure;
