CREATE DATABASE metrics;
USE metrics;
-- DROP DATABASE metrics
-- SHOW TABLES FROM metrics; 

CREATE TABLE pressure_metadata (
    device_id String,
    loc String,
    model String,
    manufacturer String,
    install_date Date,
) ENGINE = MergeTree()
ORDER BY (loc, model);

CREATE TABLE pressure (
    id String,
    device_name String,
    device_id String,
    baseline_pressure Float64,
    spike_probability Float64,
    spike_magnitude Float64,
    noise_level Float64,
    updated_interval DateTime,
    drift_rate Float64,
    current_pressure Float64,
    is_spiking Bool, 
    last_spike_time DateTime,
    next_read_time DateTime,
    trend String, 
    updated_at DateTime DEFAULT now(),
) ENGINE = MergeTree()
PARTITION BY updated_at
ORDER BY (device_id, updated_at);

CREATE DICTIONARY pressure_metadatadict (
    device_id String,
    loc String,
    model String,
    manufacturer String,
    install_date Date,
)
PRIMARY KEY device_id
source (CLICKHOUSE(table 'pressure_metadata'))
lifetime(0) -- no updates required as metadata will always be static
layout(HASHED()); 

--------------------------------------- Incremental MV 1 ---------------------------------------

CREATE TABLE PRESSURE_PER_LOCATION (
    loc String, 
    maxSpikeMagnitude AggregateFunction(max, Float64),
    avgCurrentPressure AggregateFunction(avg, Float64),
    minDriftRate AggregateFunction(min, Float64)
) ENGINE = AggregatingMergeTree()
ORDER BY loc; 

CREATE MATERIALIZED VIEW pressure_per_location_mv 
TO PRESSURE_PER_LOCATION AS
SELECT 
    dictGetString('pressure_metadatadict', 'loc', device_id) AS loc,
    maxState(spike_magnitude) AS maxSpikeMagnitude,
    avgState(current_pressure) AS avgCurrentPressure,
    minState(drift_rate) AS minDriftRate
FROM pressure
GROUP BY loc;

--------------------------------------- Incremental MV 1 ---------------------------------------



--------------------------------------- Incremental MV 2 ---------------------------------------

CREATE TABLE PRESSURE_PER_MODEL (
    model String, 
    uniqTrend AggregateFunction(uniq, Float64),
    countManufacturer AggregateFunction(count, Float64)
) ENGINE = AggregatingMergeTree()
ORDER BY model;

CREATE MATERIALIZED VIEW pressure_per_model_mv
TO PRESSURE_PER_MODEL AS
SELECT 
    dictGetString('pressure_metadatadict', 'model', device_id) AS model,
    uniqState(trend) AS uniqTrend,
    countState(dictGetString('pressure_metadatadict', 'manufacturer', device_id)) AS countManufacturer
FROM pressure
GROUP BY model;

--------------------------------------- Incremental MV 2 ---------------------------------------


--------------------------------------- Refresh MV 1 ---------------------------------------

CREATE TABLE pressure_daily_summary (
    loc String,
    day String,
    avgCurrentPressure AggregateFunction(avg, Float64),
    maxSpikeMagnitude AggregateFunction(max, Float64),
    sumBaselinePressure AggregateFunction(sum, Float64),
    countRecords AggregateFunction(count, UInt64)
)
ENGINE = AggregatingMergeTree()
PARTITION BY day
ORDER BY (loc, day);

CREATE MATERIALIZED VIEW pressure_daily_refresh_mv
REFRESH EVERY 24 HOUR 
TO pressure_daily_summary AS
SELECT
    dictGetString('pressure_metadatadict', 'loc', device_id) AS loc,
    formatDateTime(updated_at, '%Y%m%d') AS day,
    avgState(current_pressure) AS avgCurrentPressure,
    maxState(spike_magnitude) AS maxSpikeMagnitude,
    sumState(baseline_pressure) AS sumBaselinePressure,
    countState() AS countRecords
FROM pressure
GROUP BY loc, day;

--------------------------------------- Refresh MV 1 ---------------------------------------
