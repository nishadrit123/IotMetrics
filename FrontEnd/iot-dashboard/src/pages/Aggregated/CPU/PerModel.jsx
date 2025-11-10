import React from "react";
import DataTable from "../../../components/DataTable";

function CPU() {
  const cpuColumns = [
    "model",
    "uniqFrequency",
    "countNoiseLevel",
  ];

  return (
    <div className="container mt-4">
      <DataTable
        apiBaseUrl="http://localhost:8080/v1/cpu/aggregation/model"
        columns={cpuColumns}
      />
    </div>
  );
}

export default CPU;
