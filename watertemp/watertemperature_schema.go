package watertemp

type WaterTemperatureSchema struct {
	Type          string
	Temp_water    float32
	FormattedTime string
	Alias         string
	Ts            int
	Latitude      float32
	Longitude     float32
	Gmap          string
}
