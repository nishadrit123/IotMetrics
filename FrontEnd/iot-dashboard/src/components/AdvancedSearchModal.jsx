import React, { useState, useEffect } from "react";
import { Modal, Button, Form, Row, Col } from "react-bootstrap";

const AdvancedSearchModal = ({
  show,
  handleClose,
  columns,
  onApply,
  defaultFilters = [],
}) => {
  const [filters, setFilters] = useState([]);

  // 🧠 Initialize filters when modal opens or defaults change
  useEffect(() => {
    if (show) {
      if (defaultFilters.length > 0) {
        // clone array to avoid mutation issues
        setFilters(JSON.parse(JSON.stringify(defaultFilters)));
      } else {
        setFilters([{ field: "", operator: "=", value: "" }]);
      }
    }
  }, [show, defaultFilters]);

  const addFilter = () => {
    setFilters([...filters, { field: "", operator: "=", value: "" }]);
  };

  const removeFilter = (index) => {
    const updated = filters.filter((_, i) => i !== index);
    setFilters(updated.length ? updated : [{ field: "", operator: "=", value: "" }]);
  };

  const updateFilter = (index, key, value) => {
    const updated = [...filters];
    updated[index][key] = value;
    setFilters(updated);
  };

  const applyFilters = () => {
    const validFilters = filters.filter((f) => f.field && f.value);
    const filterString = validFilters
      .map((f) => `${f.field}${f.operator}${f.value}`)
      .join(":");
    onApply(filterString);
    handleClose();
  };

  return (
    <Modal show={show} onHide={handleClose} size="lg" centered>
      <Modal.Header closeButton>
        <Modal.Title>Advanced Search</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        {filters.map((filter, index) => (
          <Row key={index} className="mb-2 align-items-center">
            <Col xs={4}>
              <Form.Select
                value={filter.field}
                onChange={(e) => updateFilter(index, "field", e.target.value)}
              >
                <option value="">Select Column</option>
                {columns.map((col) => (
                  <option key={col} value={col}>
                    {col.replace(/_/g, " ").replace(/\b\w/g, (char) => char.toUpperCase())}
                  </option>
                ))}
              </Form.Select>
            </Col>

            <Col xs={3}>
              <Form.Select
                value={filter.operator}
                onChange={(e) => updateFilter(index, "operator", e.target.value)}
              >
                <option value="=">=</option>
                <option value="!=">!=</option>
                <option value=">">&gt;</option>
                <option value="<">&lt;</option>
                <option value=">=">&gt;=</option>
                <option value="<=">&lt;=</option>
                <option value="like">like</option>
              </Form.Select>
            </Col>

            <Col xs={4}>
              <Form.Control
                type="text"
                placeholder="Value"
                value={filter.value}
                onChange={(e) => updateFilter(index, "value", e.target.value)}
              />
            </Col>

            <Col xs={1} className="text-center">
              <Button
                variant="outline-danger"
                size="sm"
                onClick={() => removeFilter(index)}
              >
                ✖
              </Button>
            </Col>
          </Row>
        ))}
        <div className="d-flex justify-content-start mt-3">
          <Button variant="outline-secondary" size="sm" onClick={addFilter}>
            ➕ Add Filter
          </Button>
        </div>
      </Modal.Body>
      <Modal.Footer>
        <Button variant="secondary" onClick={handleClose}>
          Cancel
        </Button>
        <Button variant="primary" onClick={applyFilters}>
          Apply
        </Button>
      </Modal.Footer>
    </Modal>
  );
};

export default AdvancedSearchModal;
