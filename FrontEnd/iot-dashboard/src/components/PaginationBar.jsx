import React from "react";
import { Pagination } from "react-bootstrap";

const PaginationBar = ({ totalPages = 10, currentPage = 1 }) => {
  const lastPage = totalPages;
  const pageItems = [];

  // show first three pages
  for (let i = 1; i <= 3 && i <= lastPage; i++) {
    pageItems.push(
      <Pagination.Item key={i} active={i === currentPage}>
        {i}
      </Pagination.Item>
    );
  }

  // show ellipsis if more pages
  if (lastPage > 4) {
    pageItems.push(<Pagination.Ellipsis key="ellipsis" disabled />);
    pageItems.push(<Pagination.Item key={lastPage}>{lastPage}</Pagination.Item>);
  }

  return (
    <div className="d-flex justify-content-center mt-4">
      <Pagination>
        <Pagination.Prev disabled={currentPage === 1} />
        {pageItems}
        <Pagination.Next disabled={currentPage === lastPage} />
      </Pagination>
    </div>
  );
};

export default PaginationBar;
