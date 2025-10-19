CREATE DATABASE metrics;
USE metrics;
-- DROP DATABASE metrics
-- SHOW TABLES FROM metrics; 

CREATE TABLE temperature_metadata (
    device_id String,
    loc String,
    model String,
    manufacturer String,
    install_date Date,
) ENGINE = MergeTree()
ORDER BY (loc, model);

CREATE TABLE temperature (
    id String,
    device_name String,
    device_id String,
    baseline_temperature Float64,
    spike_probability Float64,
    spike_magnitude Float64,
    noise_level Float64,
    updated_interval DateTime,
    drift_rate Float64,
    current_temperature Float64,
    is_spiking Bool, 
    last_spike_time DateTime,
    next_read_time DateTime,
    trend String, 
    updated_at DateTime DEFAULT now(),
) ENGINE = MergeTree()
PARTITION BY updated_at
ORDER BY (device_id, updated_at);

CREATE DICTIONARY temperature_metadatadict (
    device_id String,
    loc String,
    model String,
    manufacturer String,
    install_date Date,
)
PRIMARY KEY device_id
source (CLICKHOUSE(table 'temperature_metadata'))
lifetime(0) -- no updates required as metadata will always be static
layout(HASHED()); 

--------------------------------------- Incremental MV 1 ---------------------------------------

CREATE TABLE TEMPERATURE_PER_LOCATION (
    loc String, 
    maxSpikeMagnitude AggregateFunction(max, Float64),
    avgCurrentTemperature AggregateFunction(avg, Float64),
    minDriftRate AggregateFunction(min, Float64)
) ENGINE = AggregatingMergeTree()
ORDER BY loc; 

CREATE MATERIALIZED VIEW temperature_per_location_mv 
TO TEMPERATURE_PER_LOCATION AS
SELECT 
    dictGetString('temperature_metadatadict', 'loc', device_id) AS loc,
    maxState(spike_magnitude) AS maxSpikeMagnitude,
    avgState(current_temperature) AS avgCurrentTemperature,
    minState(drift_rate) AS minDriftRate
FROM temperature
GROUP BY loc;

--------------------------------------- Incremental MV 1 ---------------------------------------



--------------------------------------- Incremental MV 2 ---------------------------------------

CREATE TABLE TEMPERATURE_PER_MODEL (
    model String, 
    uniqTrend AggregateFunction(uniq, String),
    countManufacturer AggregateFunction(count, Float64)
) ENGINE = AggregatingMergeTree()
ORDER BY model;

CREATE MATERIALIZED VIEW temperature_per_model_mv
TO TEMPERATURE_PER_MODEL AS
SELECT 
    dictGetString('temperature_metadatadict', 'model', device_id) AS model,
    uniqState(trend) AS uniqTrend,
    countState(dictGetString('temperature_metadatadict', 'manufacturer', device_id)) AS countManufacturer
FROM temperature
GROUP BY model;

--------------------------------------- Incremental MV 2 ---------------------------------------


--------------------------------------- Refresh MV 1 ---------------------------------------

CREATE TABLE temperature_daily_summary (
    loc String,
    day String,
    avgCurrentTemperature AggregateFunction(avg, Float64),
    maxSpikeMagnitude AggregateFunction(max, Float64),
    sumBaselineTemperature AggregateFunction(sum, Float64),
    countRecords AggregateFunction(count, UInt64)
)
ENGINE = AggregatingMergeTree()
PARTITION BY day
ORDER BY (loc, day);

CREATE MATERIALIZED VIEW temperature_daily_refresh_mv
REFRESH EVERY 24 HOUR 
TO temperature_daily_summary AS
SELECT
    dictGetString('temperature_metadatadict', 'loc', device_id) AS loc,
    formatDateTime(updated_at, '%Y%m%d') AS day,
    avgState(current_temperature) AS avgCurrentTemperature,
    maxState(spike_magnitude) AS maxSpikeMagnitude,
    sumState(baseline_temperature) AS sumBaselineTemperature,
    countState() AS countRecords
FROM temperature
GROUP BY loc, day;

--------------------------------------- Refresh MV 1 ---------------------------------------
