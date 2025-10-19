CREATE DATABASE metrics;
USE metrics;
-- DROP DATABASE metrics
-- SHOW TABLES FROM metrics; 

CREATE TABLE humidity_metadata (
    device_id String,
    loc String,
    model String,
    manufacturer String,
    install_date Date,
) ENGINE = MergeTree()
ORDER BY (loc, model);

CREATE TABLE humidity (
    id String,
    device_name String,
    device_id String,
    baseline_humidity Float64,
    spike_probability Float64,
    spike_magnitude Float64,
    noise_level Float64,
    updated_interval DateTime,
    drift_rate Float64,
    current_humidity Float64,
    is_spiking Bool, 
    last_spike_time DateTime,
    next_read_time DateTime,
    trend String, 
    updated_at DateTime DEFAULT now(),
) ENGINE = MergeTree()
PARTITION BY updated_at
ORDER BY (device_id, updated_at);

CREATE DICTIONARY humidity_metadatadict (
    device_id String,
    loc String,
    model String,
    manufacturer String,
    install_date Date,
)
PRIMARY KEY device_id
source (CLICKHOUSE(table 'humidity_metadata'))
lifetime(0) -- no updates required as metadata will always be static
layout(HASHED()); 

--------------------------------------- Incremental MV 1 ---------------------------------------

CREATE TABLE HUMIDITY_PER_LOCATION (
    loc String, 
    maxSpikeMagnitude AggregateFunction(max, Float64),
    avgCurrentHumidity AggregateFunction(avg, Float64),
    minDriftRate AggregateFunction(min, Float64)
) ENGINE = AggregatingMergeTree()
ORDER BY loc; 

CREATE MATERIALIZED VIEW humidity_per_location_mv 
TO HUMIDITY_PER_LOCATION AS
SELECT 
    dictGetString('humidity_metadatadict', 'loc', device_id) AS loc,
    maxState(spike_magnitude) AS maxSpikeMagnitude,
    avgState(current_humidity) AS avgCurrentHumidity,
    minState(drift_rate) AS minDriftRate
FROM humidity
GROUP BY loc;

--------------------------------------- Incremental MV 1 ---------------------------------------



--------------------------------------- Incremental MV 2 ---------------------------------------

CREATE TABLE HUMIDITY_PER_MODEL (
    model String, 
    uniqTrend AggregateFunction(uniq, Float64),
    countManufacturer AggregateFunction(count, Float64)
) ENGINE = AggregatingMergeTree()
ORDER BY model;

CREATE MATERIALIZED VIEW humidity_per_model_mv
TO HUMIDITY_PER_MODEL AS
SELECT 
    dictGetString('humidity_metadatadict', 'model', device_id) AS model,
    uniqState(trend) AS uniqTrend,
    countState(dictGetString('humidity_metadatadict', 'manufacturer', device_id)) AS countManufacturer
FROM humidity
GROUP BY model;

--------------------------------------- Incremental MV 2 ---------------------------------------


--------------------------------------- Refresh MV 1 ---------------------------------------

CREATE TABLE humidity_daily_summary (
    loc String,
    day String,
    avgCurrentHumidity AggregateFunction(avg, Float64),
    maxSpikeMagnitude AggregateFunction(max, Float64),
    sumBaselineHumidity AggregateFunction(sum, Float64),
    countRecords AggregateFunction(count, UInt64)
)
ENGINE = AggregatingMergeTree()
PARTITION BY day
ORDER BY (loc, day);

CREATE MATERIALIZED VIEW humidity_daily_refresh_mv
REFRESH EVERY 24 HOUR 
TO humidity_daily_summary AS
SELECT
    dictGetString('humidity_metadatadict', 'loc', device_id) AS loc,
    formatDateTime(updated_at, '%Y%m%d') AS day,
    avgState(current_humidity) AS avgCurrentHumidity,
    maxState(spike_magnitude) AS maxSpikeMagnitude,
    sumState(baseline_humidity) AS sumBaselineHumidity,
    countState() AS countRecords
FROM humidity
GROUP BY loc, day;

--------------------------------------- Refresh MV 1 ---------------------------------------
