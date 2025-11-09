import React from "react";
import DataTable from "../../components/DataTable";

function Temperature() {
  const temperatureColumns = [
    "device_id",
    "loc",
    "model",
    "manufacturer",
    "install_date",
    "baseline_temperature",
    "spike_probability",
    "spike_magnitude",
    "noise_level",
    "current_temperature",
    'drift_rate',
    "is_spiking",
    'trend',
    "last_spike_time",
    "updated_at",
  ];

  return (
    <div className="container mt-4">
      <DataTable
        apiBaseUrl="http://localhost:8080/v1/temperature/statistics"
        columns={temperatureColumns}
      />
    </div>
  );
}

export default Temperature;
