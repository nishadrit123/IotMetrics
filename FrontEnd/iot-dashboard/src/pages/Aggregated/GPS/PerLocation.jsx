import React from "react";
import DataTable from "../../../components/DataTable";

function GPS() {
  const gpsColumns = [
    "loc",
    "maxLongitude",
    "avgLatitude",
    "minDriftRate",
  ];

  return (
    <div className="container mt-4">
      <DataTable
        apiBaseUrl="http://localhost:8080/v1/gps/aggregation/location"
        columns={gpsColumns}
      />
    </div>
  );
}

export default GPS;
