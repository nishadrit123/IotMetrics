import React, { useState } from "react";
import { Modal, Button, Form, Row, Col } from "react-bootstrap";

const AdvancedSearchModal = ({ show, handleClose, columns, onApply }) => {
  const [filters, setFilters] = useState([
    { column: "", operator: "=", value: "" },
  ]);

  const handleAddFilter = () => {
    setFilters([...filters, { column: "", operator: "=", value: "" }]);
  };

  const handleRemoveFilter = (index) => {
    const updated = filters.filter((_, i) => i !== index);
    setFilters(updated);
  };

  const handleChange = (index, key, value) => {
    const updated = [...filters];
    updated[index][key] = value;
    setFilters(updated);
  };

  const handleApply = () => {
    const filterString = filters
      .filter((f) => f.column && f.operator && f.value)
      .map((f) => `${f.column}${f.operator}${f.value}`)
      .join(":");
    onApply(filterString);
    handleClose();
  };

  const formatColumnName = (name) => {
    return name
      .replace(/_/g, " ")
      .replace(/\b\w/g, (char) => char.toUpperCase());
  };

  return (
    <Modal show={show} onHide={handleClose} centered size="lg">
      <Modal.Header closeButton>
        <Modal.Title>Advanced Search</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        {filters.map((filter, index) => (
          <Row className="mb-3" key={index}>
            <Col md={4}>
              <Form.Select
                value={filter.column}
                onChange={(e) => handleChange(index, "column", e.target.value)}
              >
                <option value="">Select Column</option>
                {columns.map((col) => (
                  <option key={col} value={col}>
                    {formatColumnName(col)}
                  </option>
                ))}
              </Form.Select>
            </Col>
            <Col md={2}>
              <Form.Select
                value={filter.operator}
                onChange={(e) =>
                  handleChange(index, "operator", e.target.value)
                }
              >
                <option value="=">=</option>
                <option value="~">~</option>
                <option value=">">&gt;</option>
                <option value="<">&lt;</option>
              </Form.Select>
            </Col>
            <Col md={4}>
              <Form.Control
                type="text"
                placeholder="Value"
                value={filter.value}
                onChange={(e) => handleChange(index, "value", e.target.value)}
              />
            </Col>
            <Col md={2}>
              <Button
                variant="outline-danger"
                onClick={() => handleRemoveFilter(index)}
                disabled={filters.length === 1}
              >
                ❌
              </Button>
            </Col>
          </Row>
        ))}

        <Button variant="outline-primary" onClick={handleAddFilter}>
          ➕ Add Filter
        </Button>
      </Modal.Body>
      <Modal.Footer>
        <Button variant="secondary" onClick={handleClose}>
          Cancel
        </Button>
        <Button variant="primary" onClick={handleApply}>
          Apply
        </Button>
      </Modal.Footer>
    </Modal>
  );
};

export default AdvancedSearchModal;
