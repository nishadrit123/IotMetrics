CREATE DATABASE metrics;
USE metrics;
-- DROP DATABASE metrics
-- SHOW TABLES FROM metrics; 

CREATE TABLE gps_metadata (
    device_id String,
    loc String,
    model String,
    manufacturer String,
    install_date Date,
) ENGINE = MergeTree()
ORDER BY (loc, model);

CREATE TABLE gps (
    id String,
    device_name String,
    device_id String,
    updated_interval DateTime,
    drift_rate Float64,
    latitude Float64,
    longitude Float64,
    altitude Float64,
    speed Float64,
    heading Float64,
    is_moving Bool, 
    next_read_time DateTime,
    updated_at DateTime DEFAULT now(),
) ENGINE = MergeTree()
PARTITION BY updated_at
ORDER BY (device_id, updated_at);

CREATE DICTIONARY gps_metadatadict (
    device_id String,
    loc String,
    model String,
    manufacturer String,
    install_date Date,
)
PRIMARY KEY device_id
source (CLICKHOUSE(table 'gps_metadata'))
lifetime(0) -- no updates required as metadata will always be static
layout(HASHED()); 

--------------------------------------- Incremental MV 1 ---------------------------------------

CREATE TABLE GPS_PER_LOCATION (
    loc String, 
    maxLongitude AggregateFunction(max, Float64),
    avgLatitude AggregateFunction(avg, Float64),
    minDriftRate AggregateFunction(min, Float64)
) ENGINE = AggregatingMergeTree()
ORDER BY loc; 

CREATE MATERIALIZED VIEW gps_per_location_mv 
TO GPS_PER_LOCATION AS
SELECT 
    dictGetString('gps_metadatadict', 'loc', device_id) AS loc,
    maxState(longitude) AS maxLongitude,
    avgState(latitude) AS avgLatitude,
    minState(drift_rate) AS minDriftRate
FROM gps
GROUP BY loc;

--------------------------------------- Incremental MV 1 ---------------------------------------



--------------------------------------- Incremental MV 2 ---------------------------------------

CREATE TABLE GPS_PER_MODEL (
    model String, 
    maxHeading AggregateFunction(max, Float64),
    countManufacturer AggregateFunction(count, Float64)
) ENGINE = AggregatingMergeTree()
ORDER BY model;

CREATE MATERIALIZED VIEW gps_per_model_mv
TO GPS_PER_MODEL AS
SELECT 
    dictGetString('gps_metadatadict', 'model', device_id) AS model,
    maxState(heading) AS maxHeading,
    countState(dictGetString('gps_metadatadict', 'manufacturer', device_id)) AS countManufacturer
FROM gps
GROUP BY model;

--------------------------------------- Incremental MV 2 ---------------------------------------


--------------------------------------- Refresh MV 1 ---------------------------------------

CREATE TABLE gps_daily_summary (
    model String,
    day String,
    avgSpeed AggregateFunction(avg, Float64),
    maxAltitude AggregateFunction(max, Float64),
    sumDriftRate AggregateFunction(sum, Float64),
    countRecords AggregateFunction(count, UInt64)
)
ENGINE = AggregatingMergeTree()
PARTITION BY day
ORDER BY (model, day);

CREATE MATERIALIZED VIEW gps_daily_refresh_mv
REFRESH EVERY 24 HOUR 
TO gps_daily_summary AS
SELECT
    dictGetString('gps_metadatadict', 'model', device_id) AS model,
    formatDateTime(updated_at, '%Y%m%d') AS day,
    avgState(speed) AS avgSpeed,
    maxState(altitude) AS maxAltitude,
    sumState(drift_rate) AS sumDriftRate,
    countState() AS countRecords
FROM gps
GROUP BY model, day;

--------------------------------------- Refresh MV 1 ---------------------------------------
