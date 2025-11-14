import React from "react";
import DataTable from "../../components/DataTable";

function GPS() {
  const gpsColumns = [
    "model",
    "day",
    "avgSpeed",
    "rolling_avg",
    "delta",
  ];

  return (
    <div className="container mt-4">
      <DataTable
        apiBaseUrl="http://localhost:8080/v1/gps/delta"
        columns={gpsColumns}
      />
    </div>
  );
}

export default GPS;
