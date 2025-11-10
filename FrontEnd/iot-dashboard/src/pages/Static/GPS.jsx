import React from "react";
import DataTable from "../../components/DataTable";

function GPS() {
  const gpsColumns = [
    "device_id",
    "loc",
    "model",
    "manufacturer",
    "install_date",
    "latitude",
    "longitude",
    "altitude",
    "speed",
    "heading",
    "is_moving",
    "updated_at",
  ];

  return (
    <div className="container mt-4">
      <DataTable
        apiBaseUrl="http://localhost:8080/v1/gps/statistics"
        columns={gpsColumns}
      />
    </div>
  );
}

export default GPS;
