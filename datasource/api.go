package datasource

type WaterTemperature struct {
	Type           string
	Temp_water     float32
	Formatted_Time string
	Alias          string
	Ts             int
	Latitude       float32
	Longitude      float32
	Gmap           string
}

type TemperatureLocation struct {
	Temperature float32 `json:"Temperature"`
	Location    string  `json:"Location"`
}
