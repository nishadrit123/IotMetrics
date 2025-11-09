import { useState, useEffect } from "react";
import axios from "axios";

export default function usePaginatedData(baseUrl) {
  const [data, setData] = useState([]);
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    const fetchPage = async () => {
      try {
        setLoading(true);
        const res = await axios.get(`${baseUrl}?page=${currentPage}`);
        setData(res.data.data.data || []); // 👈 correct nested extraction
        setTotalPages(res.data.data.total_pages || 1);
      } catch (err) {
        console.error("Error fetching data:", err);
      } finally {
        setLoading(false);
      }
    };

    fetchPage();
  }, [baseUrl, currentPage]);

  const handlePageChange = (page) => {
    if (page >= 1 && page <= totalPages) {
      setCurrentPage(page);
    }
  };

  return { data, currentPage, totalPages, handlePageChange, loading };
}
