import React, { useState, useEffect } from "react";
import { Table, Spinner, Button } from "react-bootstrap";
import { useSearchParams } from "react-router-dom";
import PaginationBar from "./PaginationBar";
import AdvancedSearchModal from "./AdvancedSearchModal";

const DataTable = ({ apiBaseUrl, columns }) => {
  const [searchParams, setSearchParams] = useSearchParams();

  const [data, setData] = useState([]);
  const [totalPages, setTotalPages] = useState(1);
  const [loading, setLoading] = useState(false);

  // URL Params
  const [currentPage, setCurrentPage] = useState(Number(searchParams.get("page")) || 1);
  const [orderBy, setOrderBy] = useState(searchParams.get("order") || "");
  const [sort, setSort] = useState(searchParams.get("sort") || "asc");
  const [filter, setFilter] = useState(searchParams.get("filter") || "");
  const [showModal, setShowModal] = useState(false);

  // Rolling GPS extra params
  const [preceding, setPreceding] = useState(1);
  const [following, setFollowing] = useState(0);

  const isRolling = apiBaseUrl.includes("/gps/delta");

  // Convert filter string into objects
  const parseFilterString = (filterString) => {
    if (!filterString) return [];
    const parts = filterString.split(":");
    return parts.map((p) => {
      const match = p.match(/^([^<>=~]+)([<>=~]{1})(.+)$/);
      if (!match) return { field: "", operator: "=", value: "" };
      return { field: match[1], operator: match[2], value: match[3] };
    });
  };

  const defaultFilters = parseFilterString(filter);

  const formatColumnName = (name) =>
    name.replace(/_/g, " ").replace(/\b\w/g, (char) => char.toUpperCase());

  // Fetch data
  const fetchData = async (page = 1, order = orderBy, sortDir = sort, filterVal = filter) => {
    try {
      setLoading(true);

      let url = `${apiBaseUrl}?page=${page}`;
      if (order) {
        url += `&order=${order}`;
        if (sortDir === "desc") url += `&sort=desc`;
      }
      if (filterVal) {
        url += `&filter=${encodeURIComponent(filterVal)}`;
      }

      const res = await fetch(url, {
        method: isRolling ? "POST" : "GET",
        headers: {
          "Content-Type": "application/json",
        },
        body: isRolling
          ? JSON.stringify({
              preceding,
              following,
            })
          : null,
      });

      const result = await res.json();
      console.log("API Response:", result);

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

  // Update URL params
  useEffect(() => {
    const params = {};
    if (orderBy) params.order = orderBy;
    if (sort) params.sort = sort;
    if (filter) params.filter = filter;
    if (currentPage) params.page = currentPage;
    setSearchParams(params);
  }, [orderBy, sort, filter, currentPage, setSearchParams]);

  // Refetch on state change
  useEffect(() => {
    fetchData(currentPage, orderBy, sort, filter);
  }, [currentPage, orderBy, sort, filter]);

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
      <div className="d-flex justify-content-end mb-2">
        <Button variant="outline-info" onClick={() => setShowModal(true)}>
          🔍 Advanced Search
        </Button>
      </div>

      {/* --- Rolling GPS inputs --- */}
      {apiBaseUrl.includes("/gps/delta") && (
        <div className="d-flex gap-3 mb-3">
          <div>
            <label className="me-2">Preceding</label>
            <input
              type="number"
              value={preceding}
              onChange={(e) => setPreceding(Number(e.target.value))}
              onKeyDown={(e) => e.key === "Enter" && fetchData(1)}
              className="form-control d-inline-block"
              style={{ width: "90px" }}
            />
          </div>

          <div>
            <label className="me-2">Following</label>
            <input
              type="number"
              value={following}
              onChange={(e) => setFollowing(Number(e.target.value))}
              onKeyDown={(e) => e.key === "Enter" && fetchData(1)}
              className="form-control d-inline-block"
              style={{ width: "90px" }}
            />
          </div>
        </div>
      )}

      {/* Rolling GPS will not support advanced search for last 2 columns */}
      <AdvancedSearchModal
        show={showModal}
        handleClose={() => setShowModal(false)}
        columns={isRolling ? columns.slice(0, -2) : columns}
        onApply={(filterString) => setFilter(filterString)}
        defaultFilters={defaultFilters}
      />

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
              style={{ tableLayout: "fixed", width: "100%" }}
            >
              <thead className="table-dark">
                <tr>
                  {columns.map((col) => (
                    <th
                      key={col}
                      onClick={() => handleSort(col)}
                      style={{
                        cursor: "pointer",
                        whiteSpace: "normal",
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
                          title={String(row[key])}
                        >
                          {String(row[key]) === "undefined" ? "false" : String(row[key])}
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
            onPageChange={(page) => setCurrentPage(page)}
          />
        </>
      )}
    </div>
  );
};

export default DataTable;
