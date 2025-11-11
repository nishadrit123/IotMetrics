import React from "react";
import DataTable from "../../../components/DataTable";

function Temperature() {
  const temperatureColumns = [
    "loc",
    "day",
    "avgCurrentTemperature",
    "maxSpikeMagnitude",
    "sumBaselineTemperature",
    "countRecords",
  ];

  return (
    <div className="container mt-4">
      <DataTable
        apiBaseUrl="http://localhost:8080/v1/temperature/dailyaggregation/location"
        columns={temperatureColumns}
      />
    </div>
  );
}

export default Temperature;
