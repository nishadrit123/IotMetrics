import React from "react";
import DataTable from "../../components/DataTable";

function Humidity() {
  const humidityColumns = [
    "device_id",
    "loc",
    "model",
    "manufacturer",
    "install_date",
    "baseline_humidity",
    "spike_probability",
    "spike_magnitude",
    "noise_level",
    "current_humidity",
    'drift_rate',
    "is_spiking",
    'trend',
    "last_spike_time",
    "updated_at",
  ];

  return (
    <div className="container mt-4">
      <DataTable
        apiBaseUrl="http://localhost:8080/v1/humidity/statistics"
        columns={humidityColumns}
      />
    </div>
  );
}

export default Humidity;
