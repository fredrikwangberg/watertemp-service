package datasource

type WaterTemperatureSchema struct {
	Type           string
	Temp_water     float32
	Formatted_Time string
	Alias          string
	Ts             int
	Latitude       float32
	Longitude      float32
	Gmap           string
}
