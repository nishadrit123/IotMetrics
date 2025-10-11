Identity / Metadata (static fields)
These describe the simulated device itself — they never change once initialized

| **Name**       | **Type**    | **Explanation**                                     |
| -------------- | ----------- | --------------------------------------------------- |
| `Id`           | `string`    | Unique reading ID (UUID) generated per data point   |
| `DeviceId`     | `string`    | Unique identifier for the humidity sensor device    |
| `Location`     | `string`    | Physical location (e.g., “Warehouse-Zone-A”)        |
| `Model`        | `string`    | Sensor model (e.g., “HX-200”)                       |
| `Manufacturer` | `string`    | Device manufacturer name                            |
| `InstallDate`  | `time.Time` | Date when the sensor was installed                  |
| `Unit`         | `string`    | Measurement unit — always “%RH” (Relative Humidity) |


Configuration / Behavior Parameters (static per run)
Defines how it behaves when generating data

| **Name**           | **Type**        | **Explanation**                                                            |
| ------------------ | --------------- | -------------------------------------------------------------------------- |
| `BaselineHumidity` | `float64`       | Normal operating humidity level                                            |
| `SpikeProbability` | `float64`       | Probability (0–1) of a sudden humidity jump/drop                           |
| `SpikeMagnitude`   | `float64`       | Maximum deviation (%) during a spike                                       |
| `NoiseLevel`       | `float64`       | Small random variation added to simulate sensor noise                      |
| `UpdateInterval`   | `time.Duration` | Time between consecutive readings                                          |
| `DriftRate`        | `float64`       | Gradual long-term change in humidity due to seasonal/environmental effects |


Dynamic State (changes every run)
Tracks the current live condition of the simulated device

| **Name**          | **Type**    | **Explanation**                                                   |
| ----------------- | ----------- | ----------------------------------------------------------------- |
| `CurrentHumidity` | `float64`   | Latest measured humidity value                                    |
| `IsSpiking`       | `bool`      | Whether the sensor is currently in a spike state                  |
| `LastSpikeTime`   | `time.Time` | Timestamp of the last spike event                                 |
| `NextRead`        | `time.Time` | When the next reading will be generated                           |
| `Trend`           | `string`    | Indicates short-term direction — “rising”, “falling”, or “stable” |
