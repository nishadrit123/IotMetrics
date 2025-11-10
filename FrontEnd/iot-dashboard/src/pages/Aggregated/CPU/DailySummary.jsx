import React from "react";
import DataTable from "../../../components/DataTable";

function CPU() {
  const cpuColumns = [
    "loc",
    "day",
    "avgCurrentUsage",
    "maxSpikeMagnitude",
    "avgCPUTemperature",
    "countRecords",
  ];

  return (
    <div className="container mt-4">
      <DataTable
        apiBaseUrl="http://localhost:8080/v1/cpu/dailyaggregation/location"
        columns={cpuColumns}
      />
    </div>
  );
}

export default CPU;
