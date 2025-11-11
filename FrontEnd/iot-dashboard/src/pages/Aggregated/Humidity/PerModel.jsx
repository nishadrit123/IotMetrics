import React from "react";
import DataTable from "../../../components/DataTable";

function Humidity() {
  const humidityColumns = [
    "model",
    "uniqTrend",
    "countManufacturer",
  ];

  return (
    <div className="container mt-4">
      <DataTable
        apiBaseUrl="http://localhost:8080/v1/humidity/aggregation/model"
        columns={humidityColumns}
      />
    </div>
  );
}

export default Humidity;
