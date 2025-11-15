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

// Since CH's window function (over) is being used for rolling GPS 
// which does not support direct use of 'where' or 'having' clauses 
// in same query and hence for simplicity, advanced search is not
// implemented for 'rolling_avg' and 'delta' columns here. 
// However, this still can be achieved by using subquery like: 

// SELECT * FROM (
  // SELECT model, day, avgMerge(avgSpeed), avg(avgMerge(avgSpeed)) OVER (
  // 	PARTITION BY model  
  // 	ORDER BY day  
  //     ROWS BETWEEN CURRENT ROW AND 2 FOLLOWING 
  // ) AS temp, (avgMerge(avgSpeed) - temp) AS roll FROM gps_daily_summary GROUP BY (model, day)
// ) WHERE temp > 9;

// Backend implementation for handling this is not done here yet.