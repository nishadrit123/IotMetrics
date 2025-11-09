import usePaginatedData from "../../hooks/usePaginatedData";
import PaginationBar from "../../components/PaginationBar";

function CPU() {
  const {
    data,
    currentPage,
    totalPages,
    handlePageChange,
    loading,
  } = usePaginatedData("http://localhost:8080/v1/cpu/statistics");

  return (
    <div className="container mt-4">
      {loading && <div className="text-center text-muted">Loading data...</div>}

      {!loading && data.length > 0 && (
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
              {data.map((item) => (
                <tr key={item.id}>
                  <td>{item.device_id}</td>
                  <td>{item.loc}</td>
                  <td title={item.model}>{item.model.slice(0, 25)}...</td>
                  <td title={item.host_name}>{item.host_name.slice(0, 20)}...</td>
                  <td>{item.core_count}</td>
                  <td>{item.frequency_ghz}</td>
                  <td>{item.baseline_usage}</td>
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

      <PaginationBar
        totalPages={totalPages}
        currentPage={currentPage}
        onPageChange={handlePageChange}
      />
    </div>
  );
}

export default CPU;
