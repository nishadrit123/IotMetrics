import React, { useState, useEffect } from "react";
import { Table, Spinner } from "react-bootstrap";
import PaginationBar from "./PaginationBar";

const DataTable = ({ apiBaseUrl, columns }) => {
  const [data, setData] = useState([]);
  const [totalPages, setTotalPages] = useState(1);
  const [currentPage, setCurrentPage] = useState(1);
  const [loading, setLoading] = useState(false);
  const [orderBy, setOrderBy] = useState("");
  const [sort, setSort] = useState("asc");

  // 🧠 Format column headers (CPU_Model → CPU Model)
  const formatColumnName = (name) => {
    return name
      .replace(/_/g, " ")
      .replace(/\b\w/g, (char) => char.toUpperCase());
  };

  // 🧩 Fetch data with pagination + sorting
  const fetchData = async (page = 1, order = orderBy, sortDir = sort) => {
    try {
      setLoading(true);
      let url = `${apiBaseUrl}?page=${page}`;
      if (order) {
        url += `&order=${order}`;
        if (sortDir === "desc") url += `&sort=desc`;
      }

      const res = await fetch(url);
      const result = await res.json();
      console.log("API Response:", result);

      // ✅ Handle deeply nested data safely
      const tableData =
        result?.data?.data?.data || result?.data?.data || result?.data || [];
      const total =
        result?.data?.data?.total_pages ||
        result?.data?.total_pages ||
        result?.total_pages ||
        1;

      setData(tableData);
      setTotalPages(total);
      setCurrentPage(page);
    } catch (err) {
      console.error("Fetch error:", err);
    } finally {
      setLoading(false);
    }
  };

  // 🔁 Refetch when page/sort changes
  useEffect(() => {
    fetchData(currentPage, orderBy, sort);
  }, [currentPage, orderBy, sort]);

  // ⬆️⬇️ Handle sort logic
  const handleSort = (column) => {
    if (orderBy === column) {
      setSort(sort === "asc" ? "desc" : "asc");
    } else {
      setOrderBy(column);
      setSort("asc");
    }
  };

  const renderSortArrows = (column) => {
    if (orderBy !== column) return "↕️";
    return sort === "asc" ? "↑" : "↓";
  };

  return (
    <div className="px-3" style={{ width: "100%", overflowX: "auto" }}>
      {loading ? (
        <div className="text-center my-5">
          <Spinner animation="border" />
        </div>
      ) : (
        <>
          <div style={{ overflowX: "auto" }}>
            <Table
              bordered
              hover
              responsive
              className="align-middle text-center table-sm custom-table"
              style={{
                tableLayout: "fixed",
                width: "100%",
              }}
            >
              <thead className="table-dark">
                <tr>
                  {columns.map((col) => (
                    <th
                      key={col}
                      onClick={() => handleSort(col)}
                      style={{
                        cursor: "pointer",
                        whiteSpace: "normal", // allows wrapping
                        wordBreak: "break-word",
                        verticalAlign: "middle",
                        padding: "8px 4px",
                        fontSize: "13px",
                        lineHeight: "1.2",
                      }}
                    >
                      <div
                        style={{
                          display: "flex",
                          flexDirection: "column",
                          alignItems: "center",
                        }}
                      >
                        <span>{formatColumnName(col)}</span>
                        <span style={{ fontSize: "10px", opacity: 0.7 }}>
                          {renderSortArrows(col)}
                        </span>
                      </div>
                    </th>
                  ))}
                </tr>
              </thead>
              <tbody>
                {data.length > 0 ? (
                  data.map((row, idx) => (
                    <tr key={idx}>
                      {columns.map((key) => (
                        <td
                          key={key}
                          style={{
                            whiteSpace: "nowrap",
                            overflow: "hidden",
                            textOverflow: "ellipsis",
                            maxWidth: "160px",
                            fontSize: "13px",
                            padding: "6px 8px",
                          }}
                          title={row[key]}
                        >
                          {row[key] ?? "-"}
                        </td>
                      ))}
                    </tr>
                  ))
                ) : (
                  <tr>
                    <td colSpan={columns.length} className="text-center">
                      No data available
                    </td>
                  </tr>
                )}
              </tbody>
            </Table>
          </div>

          <PaginationBar
            currentPage={currentPage}
            totalPages={totalPages}
            onPageChange={(page) => fetchData(page, orderBy, sort)}
          />
        </>
      )}
    </div>
  );
};

export default DataTable;
