import React from "react";
import DataTable from "../../../components/DataTable";

function Humidity() {
  const humidityColumns = [
    "loc",
    "maxSpikeMagnitude",
    "avgCurrentHumidity",
    "minDriftRate",
  ];

  return (
    <div className="container mt-4">
      <DataTable
        apiBaseUrl="http://localhost:8080/v1/humidity/aggregation/location"
        columns={humidityColumns}
      />
    </div>
  );
}

export default Humidity;
