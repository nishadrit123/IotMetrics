import React, { useEffect, useState } from "react";

export default function Heatmap() {
    const [matrix, setMatrix] = useState([]);
    const [labels, setLabels] = useState(["CPU", "Temperature", "Pressure", "Humidity"]);

    const fetchHeatmap = async (locs = ["all"]) => {
        try {
            const res = await fetch("http://localhost:8080/v1/heatmap", {
                credentials: "include",
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({ locs })
            });

            const json = await res.json();
            setMatrix(json.data.data);
        } catch (err) {
            console.error("Heatmap fetch error:", err);
        }
    };

    // Default load
    useEffect(() => {
        fetchHeatmap();
    }, []);

    // Listen for loc updates
    useEffect(() => {
        const handler = (e) => fetchHeatmap(e.detail);
        window.addEventListener("heatmap-update", handler);
        return () => window.removeEventListener("heatmap-update", handler);
    }, []);

    const getColor = (value) => {
        const green = Math.floor(255 * value);
        const red = 255 - green;
        return `rgb(${red},${green},80)`;
    };

    return (
        <div className="mt-4">
            {matrix.length === 0 ? (
                <p>No heatmap data</p>
            ) : (
                <table className="table table-bordered text-center">
                    <thead>
                        <tr>
                            <th></th>
                            {labels.map((label) => (
                                <th key={label}>{label}</th>
                            ))}
                        </tr>
                    </thead>
                    <tbody>
                        {matrix.map((row, rowIndex) => (
                            <tr key={rowIndex}>
                                <th>{labels[rowIndex]}</th>
                                {row.map((value, colIndex) => (
                                    <td
                                        key={colIndex}
                                        style={{
                                            backgroundColor: getColor(value),
                                            color: "black",
                                            fontWeight: "bold"
                                        }}
                                    >
                                        {value.toFixed(2)}
                                    </td>
                                ))}
                            </tr>
                        ))}
                    </tbody>
                </table>
            )}
        </div>
    );
}
