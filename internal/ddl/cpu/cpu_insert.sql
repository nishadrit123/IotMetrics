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


select * from cpu;  
select spike_magnitude, updated_at from cpu where dictGetString('cpu_metadatadict', 'loc', device_id) = 'Mumbai' order by updated_at;
  
select loc, maxMerge(maxSpikeMagnitude), avgMerge(avgCurrentUsage), sumMerge(totalCPUTemperature)
from CPU_PER_LOCATION group by loc; 

select model, uniqMerge(uniqFrequency), countMerge(countNoiseLevel) from CPU_PER_MODEL group by model; 

select loc, day, maxMerge(maxSpikeMagnitude), avgMerge(avgCurrentUsage), avgMerge(avgCPUTemperature), countMerge(countRecords)
from cpu_daily_summary group by (loc, day); 

select loc, minute, maxMerge(maxSpikeMagnitude), avgMerge(avgCurrentUsage), avgMerge(avgCPUTemperature), countMerge(countRecords)
from cpu_minute_summary group by (loc, minute);  

select * from system.view_refreshes where database = 'metrics'; 