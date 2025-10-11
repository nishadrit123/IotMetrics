Identity / Metadata (static fields)
These describe the simulated device itself — they never change once initialized

| **Name**       | **Type**  | **Purpose**                                                             |
| -------------- | --------- | ----------------------------------------------------------------------- |
| `Id`           | string    | Unique ID for CH                                                        |
| `DeviceId`     | string    | Unique ID for the CPU device (e.g. `"cpu-001"`)                         |
| `Hostname`     | string    | Logical name of the device or host machine                              |
| `Location`     | string    | Where the device is (optional — e.g. "Bangalore-DC1")                   |
| `Model`        | string    | CPU model or family (e.g. "Intel i7-8700")                              |
| `CoreCount`    | int       | Number of cores for simulation realism                                  |
| `Frequency`    | float64   | Base frequency (e.g. 3.2 GHz)                                           |


Configuration / Behavior Parameters (static per run)
Defines how it behaves when generating data

| **Name**           | **Type**      | **Purpose**                                       |
| ------------------ | ------------- | ------------------------------------------------- |
| `BaselineUsage`    | float64       | The base CPU utilization (%) — e.g. 15.0          |
| `SpikeProbability` | float64       | Chance (0–1) of spike per tick                    |
| `SpikeMagnitude`   | float64       | Max % spike (e.g. 40 means spike up to +40%)      |
| `NoiseLevel`       | float64       | Random noise amplitude around baseline (e.g. ±5%) |
| `UpdateInterval`   | time.Duration | Sleep before new metric generation (how often)    |


Dynamic State (changes every run)
Tracks the current live condition of the simulated device

| **Name**        | **Type**  | **Purpose**                                   |
| --------------- | --------- | --------------------------------------------- |
| `CurrentUsage`  | float64   | Current CPU usage (%) — updated every tick    |
| `CPUTemperature`| float64   | CPU temperature correlated to load            |
| `IsSpiking`     | bool      | Whether currently in a high-load period       |
| `LastSpikeTime` | time.Time | When the last spike started                   |
| `NextRead`      | time.Time | When the next reading should be produced      |
