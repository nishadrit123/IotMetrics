import React from "react";
import DataTable from "../../../components/DataTable";

function CPU() {
  const cpuColumns = [
    "loc",
    "maxSpikeMagnitude",
    "avgCurrentUsage",
    "totalCPUTemperature",
  ];

  return (
    <div className="container mt-4">
      <DataTable
        apiBaseUrl="http://localhost:8080/v1/cpu/aggregation/location"
        columns={cpuColumns}
      />
    </div>
  );
}

export default CPU;
