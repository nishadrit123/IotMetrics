import React from "react";
import DataTable from "../../../components/DataTable";

function Humidity() {
  const humidityColumns = [
    "loc",
    "day",
    "avgCurrentHumidity",
    "maxSpikeMagnitude",
    "sumBaselineHumidity",
    "countRecords",
  ];

  return (
    <div className="container mt-4">
      <DataTable
        apiBaseUrl="http://localhost:8080/v1/humidity/dailyaggregation/location"
        columns={humidityColumns}
      />
    </div>
  );
}

export default Humidity;
