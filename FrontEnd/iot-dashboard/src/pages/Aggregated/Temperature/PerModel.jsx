import React from "react";
import DataTable from "../../../components/DataTable";

function Temperature() {
  const temperatureColumns = [
    "model",
    "uniqTrend",
    "countManufacturer",
  ];

  return (
    <div className="container mt-4">
      <DataTable
        apiBaseUrl="http://localhost:8080/v1/temperature/aggregation/model"
        columns={temperatureColumns}
      />
    </div>
  );
}

export default Temperature;
