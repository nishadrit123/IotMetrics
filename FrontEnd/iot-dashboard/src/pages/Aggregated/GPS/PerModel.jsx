import React from "react";
import DataTable from "../../../components/DataTable";

function GPS() {
  const gpsColumns = [
    "model",
    "maxHeading",
    "countManufacturer",
  ];

  return (
    <div className="container mt-4">
      <DataTable
        apiBaseUrl="http://localhost:8080/v1/gps/aggregation/model"
        columns={gpsColumns}
      />
    </div>
  );
}

export default GPS;
