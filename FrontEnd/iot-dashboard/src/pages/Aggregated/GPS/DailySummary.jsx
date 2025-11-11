import React from "react";
import DataTable from "../../../components/DataTable";

function GPS() {
  const gpsColumns = [
    "model",
    "day",
    "avgSpeed",
    "maxAltitude",
    "sumDriftRate",
    "countRecords",
  ];

  return (
    <div className="container mt-4">
      <DataTable
        apiBaseUrl="http://localhost:8080/v1/gps/dailyaggregation/model"
        columns={gpsColumns}
      />
    </div>
  );
}

export default GPS;
