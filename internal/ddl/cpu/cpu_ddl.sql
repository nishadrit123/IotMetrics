CREATE DATABASE metrics;
USE metrics;
-- DROP DATABASE metrics
-- SHOW TABLES FROM metrics; 

CREATE TABLE cpu_metadata (
    device_id String,
    hostname String,
    loc String,
    model String,
    core_count Int64,
    frequency Float64,
) ENGINE = MergeTree()
ORDER BY (loc, model);

CREATE TABLE cpu (
    id String,
    device_name String,
    device_id String,
    baseline_usage Float64,
    spike_probability Float64,
    spike_magnitude Float64,
    noise_level Float64,
    updated_interval DateTime,
    current_usage Float64,
    cpu_temperature Float64,
    is_spiking Bool, 
    last_spike_time DateTime,
    next_read_time DateTime,
    updated_at DateTime DEFAULT now(),
) ENGINE = MergeTree();
PARTITION BY toYYYYMM(updated_at)
ORDER BY (device_id, updated_at);

CREATE DICTIONARY cpu_metadatadict (
    device_id String,
    hostname String,
    loc String,
    model String,
    core_count Int64,
    frequency Float64
)
PRIMARY KEY device_id
source (CLICKHOUSE(table 'cpu_metadata'))
lifetime(0) -- no updates required as metadata will always be static
layout(FLAT()); 

--------------------------------------- Incremental MV 1 ---------------------------------------

CREATE TABLE CPU_PER_LOCATION (
    loc String, 
    maxSpikeMagnitude AggregateFunction(max, Float64),
    avgCurrentUsage AggregateFunction(avg, Float64),
    totalCPUTemperature AggregateFunction(sum, Float64)
) ENGINE = AggregatingMergeTree()
ORDER BY loc; 

CREATE MATERIALIZED VIEW cpu_per_location_mv 
TO CPU_PER_LOCATION AS
SELECT 
    dictGetString('cpu_metadatadict', 'loc', device_id) AS loc,
    maxState(spike_magnitude) AS maxSpikeMagnitude,
    avgState(current_usage) AS avgCurrentUsage,
    sumState(cpu_temperature) AS totalCPUTemperature
FROM cpu
GROUP BY loc;

--------------------------------------- Incremental MV 1 ---------------------------------------



--------------------------------------- Incremental MV 2 ---------------------------------------

CREATE TABLE CPU_PER_MODEL (
    model String, 
    uniqFrequency AggregateFunction(uniq, Float64),
    countNoiseLevel AggregateFunction(count, Float64)
) ENGINE = AggregatingMergeTree()
ORDER BY model;

CREATE MATERIALIZED VIEW cpu_per_model_mv
TO CPU_PER_MODEL AS
SELECT 
    dictGetString('cpu_metadatadict', 'model', device_id) AS model,
    uniqState(dictGetFloat64('cpu_metadatadict', 'frequency', device_id)) AS uniqFrequency,
    countState(noise_level) AS countNoiseLevel
FROM cpu
GROUP BY model;

--------------------------------------- Incremental MV 2 ---------------------------------------



--------------------------------------- Refresh MV 1 ---------------------------------------

CREATE TABLE cpu_hourly_summary (
    loc String,
    hour DateTime,
    avgCurrentUsage AggregateFunction(avg, Float64),
    maxSpikeMagnitude AggregateFunction(max, Float64),
    avgCPUTemperature AggregateFunction(avg, Float64),
    countRecords AggregateFunction(count, UInt64)
)
ENGINE = AggregatingMergeTree()
PARTITION BY toYYYYMMDD(hour)
ORDER BY (loc, hour);

CREATE MATERIALIZED VIEW cpu_hourly_refresh_mv
REFRESH EVERY 1 HOUR
TO cpu_hourly_summary AS
SELECT
    dictGetString('cpu_metadatadict', 'loc', device_id) AS loc,
    toStartOfHour(updated_at) AS hour,
    avgState(current_usage) AS avgCurrentUsage,
    maxState(spike_magnitude) AS maxSpikeMagnitude,
    avgState(cpu_temperature) AS avgCPUTemperature,
    countState() AS countRecords
FROM cpu
GROUP BY loc, hour;

--------------------------------------- Refresh MV 1 ---------------------------------------



--------------------------------------- Refresh MV 2 ---------------------------------------

CREATE TABLE cpu_daily_summary (
    loc String,
    day Date,
    avgCurrentUsage AggregateFunction(avg, Float64),
    maxSpikeMagnitude AggregateFunction(max, Float64),
    avgCPUTemperature AggregateFunction(avg, Float64),
    countRecords AggregateFunction(count, UInt64)
)
ENGINE = AggregatingMergeTree()
PARTITION BY toYYYYMM(day)
ORDER BY (loc, day);

CREATE MATERIALIZED VIEW cpu_daily_refresh_mv
REFRESH EVERY 24 HOUR 
TO cpu_daily_summary AS
SELECT
    dictGetString('cpu_metadatadict', 'loc', device_id) AS loc,
    toDate(updated_at) AS day,
    avgState(current_usage) AS avgCurrentUsage,
    maxState(spike_magnitude) AS maxSpikeMagnitude,
    avgState(cpu_temperature) AS avgCPUTemperature,
    countState() AS countRecords
FROM cpu
GROUP BY loc, day;

--------------------------------------- Refresh MV 2 ---------------------------------------
