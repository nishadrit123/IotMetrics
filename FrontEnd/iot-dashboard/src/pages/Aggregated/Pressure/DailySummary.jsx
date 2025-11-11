import React from "react";
import DataTable from "../../../components/DataTable";

function Pressure() {
  const pressureColumns = [
    "loc",
    "day",
    "avgCurrentPressure",
    "maxSpikeMagnitude",
    "sumBaselinePressure",
    "countRecords",
  ];

  return (
    <div className="container mt-4">
      <DataTable
        apiBaseUrl="http://localhost:8080/v1/pressure/dailyaggregation/location"
        columns={pressureColumns}
      />
    </div>
  );
}

export default Pressure;
