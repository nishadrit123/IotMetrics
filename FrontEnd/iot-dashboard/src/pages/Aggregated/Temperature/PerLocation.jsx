import React from "react";
import DataTable from "../../../components/DataTable";

function Temperature() {
  const temperatureColumns = [
    "loc",
    "maxSpikeMagnitude",
    "avgCurrentTemperature",
    "minDriftRate",
  ];

  return (
    <div className="container mt-4">
      <DataTable
        apiBaseUrl="http://localhost:8080/v1/temperature/aggregation/location"
        columns={temperatureColumns}
      />
    </div>
  );
}

export default Temperature;
