package store

const PageSize = 10

var Tables = []string{
	"cpu_daily_summary", "temperature_daily_summary",
	"pressure_daily_summary", "humidity_daily_summary",
}

var TableMap map[string]string
var DictMap map[string]string

func init() {
	TableMap = make(map[string]string)
	TableMap["cpu_daily_summary"] = "avgCurrentUsage"
	TableMap["temperature_daily_summary"] = "avgCurrentTemperature"
	TableMap["pressure_daily_summary"] = "avgCurrentPressure"
	TableMap["humidity_daily_summary"] = "avgCurrentHumidity"

	DictMap = make(map[string]string)
	DictMap["hostname"] = "String"
	DictMap["loc"] = "String"
	DictMap["model"] = "String"
	DictMap["manufacturer"] = "String"
	DictMap["core_count"] = "Int64"
	DictMap["frequency"] = "Float64"
	DictMap["install_date"] = "Date"
}
