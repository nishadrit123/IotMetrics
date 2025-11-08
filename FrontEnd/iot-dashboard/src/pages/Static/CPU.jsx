import { useEffect, useState } from "react";
import axios from "axios";
import PaginationBar from "../../components/PaginationBar";

function CPU() {
  const [cpuData, setCpuData] = useState([]);
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  useEffect(() => {
    const fetchData = async () => {
      try {
        setLoading(true);
        setError("");
        const res = await axios.get("http://localhost:8080/v1/cpu/statistics", {
          params: { page: currentPage },
        });

        const tableData = res.data.data.data;
        setCpuData(tableData);
        setTotalPages(res.data.data.total_pages);
      } catch (err) {
        console.log(err) 
        setError("Failed to fetch CPU statistics.");
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, [currentPage]);

  const handlePageChange = (page) => {
    if (page >= 1 && page <= totalPages && page !== currentPage) {
      setCurrentPage(page);
    }
  };

  return (
    <div className="container mt-4">
      {loading && <div className="text-center text-muted">Loading data...</div>}
      {error && <div className="alert alert-danger">{error}</div>}

      {!loading && !error && cpuData.length > 0 && (
        <div className="table-responsive">
          <table className="table table-striped table-bordered align-middle w-100">
            <thead className="table-dark">
              <tr>
                <th>Device ID</th>
                <th>Location</th>
                <th>Model</th>
                <th>Host Name</th>
                <th>Core Count</th>
                <th>Frequency (GHz)</th>
                <th>Baseline Usage</th>
                <th>Spike Probability</th>
                <th>Spike Magnitude</th>
                <th>Noise Level</th>
                <th>Current Usage</th>
                <th>Temperature (°C)</th>
                <th>Is Spiking</th>
                <th>Last Spike Time</th>
                <th>Updated At</th>
              </tr>
            </thead>
            <tbody>
              {cpuData.map((item) => (
                <tr key={item.id}>
                  <td>{item.device_id}</td>
                  <td>{item.loc}</td>
                  <td>{item.model}</td>
                  <td>{item.host_name}</td>
                  <td>{item.core_count}</td>
                  <td>{item.frequency_ghz}</td>
                  <td>{item.baseline_usage}</td>
                  <td>{item.spike_probability}</td>
                  <td>{item.spike_magnitude}</td>
                  <td>{item.noise_level}</td>
                  <td>{item.current_usage}</td>
                  <td>{item.temperature_celsius}</td>
                  <td>{item.is_spiking ? "Yes" : "No"}</td>
                  <td>{new Date(item.last_spike_time).toLocaleString()}</td>
                  <td>{new Date(item.updated_at).toLocaleString()}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}

      <PaginationBar totalPages={totalPages} currentPage={1} />
    </div>
  );
}

export default CPU;
