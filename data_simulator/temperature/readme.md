Identity / Metadata (static fields)
These describe the simulated device itself — they never change once initialized

| **Name**       | **Type**    | **Explanation**                                     |
| -------------- | ----------- | --------------------------------------------------- |
| `Id`           | `string`    | Unique reading ID (UUID) generated per data point   |
| `DeviceId`     | `string`    | Unique identifier for the temperature sensor device |
| `Location`     | `string`    | Physical location (e.g., “Mumbai/Lab-1”)            |
| `Model`        | `string`    | Sensor model (e.g., “TX-90”)                        |
| `Manufacturer` | `string`    | Device manufacturer name                            |
| `InstallDate`  | `time.Time` | When the sensor was installed                       |
| `Unit`         | `string`    | Temperature unit — “°C” or “°F”                     |


Configuration / Behavior Parameters (static per run)
Defines how it behaves when generating data

| **Name**           | **Type**        | **Explanation**                                              |
| ------------------ | --------------- | ------------------------------------------------------------ |
| `BaselineTemp`     | `float64`       | Average stable temperature baseline                          |
| `SpikeProbability` | `float64`       | Probability (0–1) of a sudden temperature anomaly per tick   |
| `SpikeMagnitude`   | `float64`       | Max deviation (°C) during a spike                            |
| `NoiseLevel`       | `float64`       | Small random noise added to each reading                     |
| `UpdateInterval`   | `time.Duration` | Time between consecutive readings                            |
| `DriftRate`        | `float64`       | Gradual drift in temperature over time                       |
|                    |                 | (e.g., sensor aging or environmental change)                 |


Dynamic State (changes every run)
Tracks the current live condition of the simulated device

| **Name**        | **Type**    | **Explanation**                                                   |
| --------------- | ----------- | ----------------------------------------------------------------- |
| `CurrentTemp`   | `float64`   | Latest measured temperature                                       |
| `IsSpiking`     | `bool`      | Whether the sensor is currently in a spike state                  |
| `LastSpikeTime` | `time.Time` | Timestamp of last spike event                                     |
| `NextRead`      | `time.Time` | When the next reading should be generated                         |
| `Trend`         | `string`    | Indicates short-term direction — “rising”, “falling”, or “stable” |
