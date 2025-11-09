import React from "react";
import DataTable from "../../components/DataTable";

function CPU() {
  // Define which columns you want to display for CPU
  const cpuColumns = [
    "device_id",
    "loc",
    "model",
    "host_name",
    "core_count",
    "frequency_ghz",
    "baseline_usage",
    "spike_probability",
    "spike_magnitude",
    "noise_level",
    "current_usage",
    "temperature_celsius",
    "is_spiking",
    "last_spike_time",
    "updated_at",
  ];

  return (
    <div className="container mt-4">
      <DataTable
        apiBaseUrl="http://localhost:8080/v1/cpu/statistics"
        columns={cpuColumns}
      />
    </div>
  );
}

export default CPU;
