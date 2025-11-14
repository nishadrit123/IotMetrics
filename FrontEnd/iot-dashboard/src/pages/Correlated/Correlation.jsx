import React, { useState } from "react";
import Heatmap from "../../components/HeatMap";

export default function Correlation() {
    const [locInput, setLocInput] = useState("");

    const handleKeyDown = (e) => {
        if (e.key === "Enter") {
            const newLocs = locInput.trim() === ""
                ? ["all"]
                : locInput.split(",").map((x) => x.trim());

            window.dispatchEvent(
                new CustomEvent("heatmap-update", { detail: newLocs })
            );
        }
    };

    return (
        <div className="container mt-4">
            <h3>Correlation Heatmap</h3>

            {/* City Input */}
            <div className="mb-3" style={{ maxWidth: "300px" }}>
                <label className="form-label">Locations (comma separated)</label>
                <input
                    type="text"
                    className="form-control"
                    placeholder='Example: Pune, Mumbai'
                    value={locInput}
                    onChange={(e) => setLocInput(e.target.value)}
                    onKeyDown={handleKeyDown}
                />
            </div>

            <Heatmap />
        </div>
    );
}
