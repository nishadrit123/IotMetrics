Identity / Metadata (static fields)
These describe the simulated device itself — they never change once initialized

| **Name**       | **Type**    | **Explanation**                                               |
| -------------- | ----------- | ------------------------------------------------------------- |
| `Id`           | `string`    | Unique reading ID (UUID) generated per data point             |
| `DeviceId`     | `string`    | Unique identifier for the GPS device                          |
| `DeviceName`   | `string`    | Human-readable name of the device                             |
| `Location`     | `string`    | Physical installation location (optional — e.g., “Vehicle A”) |
| `Model`        | `string`    | GPS module model (e.g., “GNSS-XYZ”)                           |
| `Manufacturer` | `string`    | GPS module manufacturer                                       |
| `InstallDate`  | `time.Time` | When the GPS device was deployed                              |


Configuration / Behavior Parameters (static per run)
Defines how it behaves when generating data

| **Name**          | **Type**        | **Explanation**                                                  |
| ----------------- | --------------- | ---------------------------------------------------------------- |
| `UpdateInterval`  | `time.Duration` | Frequency of GPS readings                                        |
| `SpeedNoiseLevel` | `float64`       | Random variation added to simulated speed                        |
| `CoordNoiseLevel` | `float64`       | Small random deviation for latitude/longitude per tick           |
| `MaxSpeed`        | `float64`       | Maximum possible speed for the device (for realistic simulation) |
| `DriftRate`       | `float64`       | Slow drift in coordinates to simulate movement or GPS inaccuracy |


Dynamic State (changes every run)
Tracks the current live condition of the simulated device

| **Name**    | **Type**    | **Explanation**                               |
| ----------- | ----------- | --------------------------------------------- |
| `Latitude`  | `float64`   | Current latitude of the device                |
| `Longitude` | `float64`   | Current longitude of the device               |
| `Altitude`  | `float64`   | Current altitude (optional)                   |
| `Speed`     | `float64`   | Current speed (m/s or km/h)                   |
| `Heading`   | `float64`   | Direction of movement (degrees)               |
| `NextRead`  | `time.Time` | When the next GPS reading should be generated |
| `IsMoving`  | `bool`      | Whether the device is currently moving        |
