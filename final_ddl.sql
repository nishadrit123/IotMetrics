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

INSERT INTO cpu_metadata (device_id, hostname, loc, model, core_count, frequency) VALUES
('CPU-9f3a1c72', 'alpha-core-01', 'Mumbai', 'Intel Xeon E5-2690 v4', 4, 2.4),
('CPU-47b8d2ef', 'alpha-core-02', 'Pune', 'AMD Ryzen Threadripper PRO 5995WX', 6, 3.1),
('CPU-a13f5e90', 'beta-node-01', 'Bangalore', 'Intel Xeon E5-2690 v4', 8, 2.9),
('CPU-6b2c4d11', 'beta-node-02', 'Delhi', 'Intel Core i9-12900K', 12, 3.5),
('CPU-f09a7b34', 'gamma-cpu-01', 'Bangalore', 'AMD Ryzen 9 7950X', 16, 3.8),
('CPU-3e7c9a52', 'gamma-cpu-02', 'Pune', 'Intel Xeon Platinum 8280', 10, 2.6),
('CPU-b4f1d689', 'delta-engine-01', 'Mumbai', 'AMD Ryzen 9 7950X', 8, 2.8),
('CPU-28d6a9c3', 'delta-engine-02', 'Bangalore', 'Intel Xeon Platinum 8280', 6, 3.3),
('CPU-c8e2f934', 'epsilon-core-01', 'Pune', 'Intel Core i9-12900K', 4, 3.0),
('CPU-5a7b3d80', 'zeta-node-01', 'Bangalore', 'AMD Ryzen 7 5800X3D', 12, 2.5),
('CPU-d49f8a15', 'eta-cpu-01', 'Mumbai', 'Intel Xeon E5-2690 v4', 8, 3.7),
('CPU-7c2d1e63', 'theta-core-01', 'Delhi', 'Intel Xeon Platinum 8280', 14, 4.0),
('CPU-fb8a6d27', 'iota-engine-01', 'Mumbai', 'AMD Ryzen 7 5800X3D', 6, 2.2),
('CPU-9a3e2b74', 'kappa-node-01', 'Pune', 'Intel Xeon E5-2690 v4', 8, 3.2),
('CPU-1f6d4e58', 'lambda-cpu-01', 'Mumbai', 'AMD Ryzen Threadripper PRO 5995WX', 16, 3.9);

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
) ENGINE = MergeTree()
PARTITION BY updated_at
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
layout(HASHED()); 

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

CREATE TABLE cpu_daily_summary (
    loc String,
    day String,
    avgCurrentUsage AggregateFunction(avg, Float64),
    maxSpikeMagnitude AggregateFunction(max, Float64),
    avgCPUTemperature AggregateFunction(avg, Float64),
    countRecords AggregateFunction(count, UInt64)
)
ENGINE = AggregatingMergeTree()
PARTITION BY day
ORDER BY (loc, day);

CREATE MATERIALIZED VIEW cpu_daily_refresh_mv
REFRESH EVERY 24 HOUR 
TO cpu_daily_summary AS
SELECT
    dictGetString('cpu_metadatadict', 'loc', device_id) AS loc,
    formatDateTime(updated_at, '%Y%m%d') AS day,
    avgState(current_usage) AS avgCurrentUsage,
    maxState(spike_magnitude) AS maxSpikeMagnitude,
    avgState(cpu_temperature) AS avgCPUTemperature,
    countState() AS countRecords
FROM cpu
GROUP BY loc, day;

--------------------------------------- Refresh MV 1 ---------------------------------------

-- CREATE DATABASE metrics;
-- USE metrics;
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

INSERT INTO gps_metadata (device_id, loc, model, manufacturer, install_date) VALUES
('GPS-1001A', 'Vehicle A', 'ProGPS-8', 'NavTech', '2021-01-10'),
('GPS-1002B', 'Vehicle B', 'SmartGPS-900', 'GeoDynamics', '2020-05-22'),
('GPS-1003C', 'Truck 1', 'MicroGPS-20', 'AeroInstruments', '2022-02-05'),
('GPS-1004D', 'Truck 2', 'SmartGPS-900', 'SatelliteCorp', '2023-03-18'),
('GPS-1005E', 'Drone 1', 'MicroGPS-20', 'PrecisionGPS', '2021-08-07'),
('GPS-1006F', 'Drone 2', 'AeroNav-1', 'AeroInstruments', '2022-06-15'),
('GPS-1007G', 'Boat 1', 'AeroNav-1', 'GeoDynamics', '2020-09-28'),
('GPS-1008H', 'Boat 2', 'MicroGPS-20', 'SatelliteCorp', '2021-12-01'),
('GPS-1009I', 'Lab Test Unit 1', 'SmartGPS-900', 'GeoDynamics', '2023-01-12'),
('GPS-1010J', 'Lab Test Unit 2', 'MicroGPS-20', 'PrecisionGPS', '2024-04-09'),
('GPS-1011K', 'Warehouse Vehicle 1', 'MicroGPS-20', 'NavTech', '2022-07-22'),
('GPS-1012L', 'Warehouse Vehicle 2', 'SmartGPS-900', 'GeoDynamics', '2020-10-16'),
('GPS-1013M', 'Field Unit 1', 'ProGPS-8', 'AeroInstruments', '2021-05-02'),
('GPS-1014N', 'Field Unit 2', 'MicroGPS-20', 'GeoDynamics', '2023-05-27'),
('GPS-1015O', 'Delivery Van 1', 'ProGPS-8', 'NavTech', '2024-01-05');

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


-- CREATE DATABASE metrics;
-- USE metrics;
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

INSERT INTO humidity_metadata (device_id, loc, model, manufacturer, install_date) VALUES
('HUM-9001A', 'Mumbai', 'HX-200', 'EnviroTech', '2021-03-10'),
('HUM-8423B', 'Pune', 'HX-210', 'AeroDynamics', '2020-08-22'),
('HUM-7310C', 'Bangalore', 'HumidPro-5', 'ThermoLogic', '2022-02-05'),
('HUM-6598D', 'Delhi', 'DrySense-X1', 'NanoInstruments', '2023-04-19'),
('HUM-5246E', 'Bangalore', 'MoistTrack-300', 'ClimaEdge', '2021-11-07'),
('HUM-4875F', 'Delhi', 'HX-210', 'NanoInstruments', '2022-06-15'),
('HUM-3762G', 'Pune', 'DrySense-X1', 'AeroDynamics', '2020-09-28'),
('HUM-2691H', 'Bangalore', 'MoistTrack-300', 'ThermoLogic', '2021-12-01'),
('HUM-1539I', 'Delhi', 'DrySense-X1', 'AeroDynamics', '2023-01-18'),
('HUM-0452J', 'Pune', 'HX-210', 'NanoInstruments', '2024-03-09'),
('HUM-9238K', 'Pune', 'HX-210', 'EnviroTech', '2022-07-22'),
('HUM-8725L', 'Mumbai', 'HX-200', 'AeroDynamics', '2020-10-16'),
('HUM-7603M', 'Mumbai', 'HX-210', 'EnviroTech', '2021-05-02'),
('HUM-6421N', 'Pune', 'DrySense-X1', 'ClimaEdge', '2023-05-27'),
('HUM-5087O', 'Delhi', 'HumidPro-5', 'AeroDynamics', '2024-01-05');

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
    uniqTrend AggregateFunction(uniq, String),
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


-- CREATE DATABASE metrics;
-- USE metrics;
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

INSERT INTO pressure_metadata (device_id, loc, model, manufacturer, install_date) VALUES
('PRS-1001A', 'Mumbai', 'PressureTrack-8', 'EnviroTech', '2021-01-10'),
('PRS-1002B', 'Delhi', 'LabPress-33', 'AeroDynamics', '2020-05-22'),
('PRS-1003C', 'Bangalore', 'LabPress-33', 'ThermoLogic', '2022-02-05'),
('PRS-1004D', 'Pune', 'PressureMax-200', 'NanoInstruments', '2023-03-18'),
('PRS-1005E', 'Bangalore', 'EnviroPress-50', 'AeroDynamics', '2021-08-07'),
('PRS-1006F', 'Pune', 'EnviroPress-50', 'AeroDynamics', '2022-06-15'),
('PRS-1007G', 'Bangalore', 'PressureMax-200', 'HydroLabs', '2020-09-28'),
('PRS-1008H', 'Delhi', 'AeroPressure-1', 'EnviroTech', '2021-12-01'),
('PRS-1009I', 'Mumbai', 'LabPress-33', 'ThermoLogic', '2023-01-12'),
('PRS-1010J', 'Pune', 'EnviroPress-50', 'NanoInstruments', '2024-04-09'),
('PRS-1011K', 'Delhi', 'PressureTrack-8', 'AeroDynamics', '2022-07-22'),
('PRS-1012L', 'Mumbai', 'EnviroPress-50', 'ThermoLogic', '2020-10-16'),
('PRS-1013M', 'Bangalore', 'PressureMax-200', 'EnviroTech', '2021-05-02'),
('PRS-1014N', 'Bangalore', 'LabPress-33', 'HydroLabs', '2023-05-27'),
('PRS-1015O', 'Mumbai', 'PressureTrack-8', 'EnviroTech', '2024-01-05');

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
    uniqTrend AggregateFunction(uniq, String),
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


-- CREATE DATABASE metrics;
-- USE metrics;
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

INSERT INTO temperature_metadata (device_id, loc, model, manufacturer, install_date) VALUES
('TEMP-9012A', 'Mumbai', 'TX-90', 'ThermoTech', '2021-06-10'),
('TEMP-8347B', 'Delhi', 'TX-85', 'EnviroSense', '2020-11-25'),
('TEMP-7231C', 'Bangalore', 'ThermIQ-10', 'HeatMaster', '2022-01-08'),
('TEMP-6458D', 'Pune', 'NanoTemp-X1', 'AeroDynamics', '2023-03-15'),
('TEMP-5193E', 'Mumbai', 'TX-85', 'HeatMaster', '2021-09-04'),
('TEMP-4872F', 'Pune', 'NanoTemp-X1', 'HeatMaster', '2022-07-19'),
('TEMP-3920G', 'Bangalore', 'ThermIQ-10', 'EnviroSense', '2020-05-30'),
('TEMP-2756H', 'Mumbai', 'TX-90', 'HeatMaster', '2021-12-12'),
('TEMP-1683I', 'Bangalore', 'NanoTemp-X1', 'EnviroSense', '2023-02-20'),
('TEMP-0549J', 'Delhi', 'TX-85', 'AeroDynamics', '2024-04-02'),
('TEMP-9234K', 'Delhi', 'ThermIQ-10', 'ThermoTech', '2022-08-17'),
('TEMP-8721L', 'Mumbai', 'TX-90', 'EnviroSense', '2020-10-06'),
('TEMP-7563M', 'Pune', 'TX-85', 'HeatMaster', '2021-02-14'),
('TEMP-6312N', 'Delhi', 'ThermIQ-10', 'AeroDynamics', '2023-06-21'),
('TEMP-5024O', 'Mumbai', 'TX-90', 'ThermoTech', '2024-01-11');

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


select loc, maxMerge(maxSpikeMagnitude), avgMerge(avgCurrentUsage), sumMerge(totalCPUTemperature)
from CPU_PER_LOCATION group by loc; 

select model, maxMerge(maxHeading), countMerge(countManufacturer)
from GPS_PER_MODEL group by model; 

select loc, maxMerge(maxSpikeMagnitude), avgMerge(avgCurrentHumidity), minMerge(minDriftRate)
from HUMIDITY_PER_LOCATION group by loc; 

select loc, maxMerge(maxSpikeMagnitude), avgMerge(avgCurrentPressure), minMerge(minDriftRate)
from PRESSURE_PER_LOCATION group by loc; 

select loc, maxMerge(maxSpikeMagnitude), avgMerge(avgCurrentTemperature), minMerge(minDriftRate)
from TEMPERATURE_PER_LOCATION group by loc;


select loc, day, maxMerge(maxSpikeMagnitude), avgMerge(avgCurrentUsage), avgMerge(avgCPUTemperature), countMerge(countRecords)
from cpu_daily_summary group by (loc, day); 

select model, day, avgMerge(avgSpeed), maxMerge(maxAltitude), sumMerge(sumDriftRate), countMerge(countRecords)
from gps_daily_summary group by (model, day); 

select loc, day, maxMerge(maxSpikeMagnitude), avgMerge(avgCurrentHumidity), sumMerge(sumBaselineHumidity), countMerge(countRecords)
from humidity_daily_summary group by (loc, day); 

select loc, day, maxMerge(maxSpikeMagnitude), avgMerge(avgCurrentPressure), sumMerge(sumBaselinePressure), countMerge(countRecords)
from pressure_daily_summary group by (loc, day); 

select loc, day, maxMerge(maxSpikeMagnitude), avgMerge(avgCurrentTemperature), sumMerge(sumBaselineTemperature), countMerge(countRecords)
from temperature_daily_summary group by (loc, day); 


select * from system.view_refreshes where database = 'metrics' 

select corr(acu, act) from ( select loc, day, avgMerge(avgCurrentUsage) as acu, avgMerge(avgCurrentTemperature) as act from
cpu_daily_summary JOIN temperature_daily_summary using (loc, day) where loc in ('Pune', 'Bangalore') group by (loc, day) );

SELECT model, day, avgMerge(avgSpeed), avg(avgMerge(avgSpeed)) OVER (
	PARTITION BY model  
	ORDER BY day  
    ROWS BETWEEN CURRENT ROW AND 2 FOLLOWING 
) as temp, avgMerge(avgSpeed) - temp from gps_daily_summary group by (model, day);  