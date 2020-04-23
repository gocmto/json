package mypkg

// https://yandex.ru/dev/weather/doc/dg/concepts/forecast-info-docpage/

type StructWeather struct {
	Now      int
	NowDt    string
	Info     StructInfo
	Fact     StructFact
	Forecast StructForecast
}

/**
  "lat": 43.5575,
  "lon": 43.85083,
  "url": "https://yandex.ru/pogoda/?lat=43.5575&lon=43.85083"
*/
type StructInfo struct {
	Lat float32
	Lon float32
	Url string
}

/**
  "temp": 7,
  "feels_like": 5,
  "icon": "ovc",
  "condition": "overcast",
  "wind_speed": 2,
  "wind_gust": 3.6,
  "wind_dir": "n",
  "pressure_mm": 736,
  "pressure_pa": 982,
  "humidity": 76,
  "daytime": "d",
  "polar": false,
  "season": "spring",
  "obs_time": 1587548516
*/
type StructFact struct {
	Temp       uint8
	FeelsLike  uint8
	Icon       string
	Condition  string
	WindSpeed  float32
	WindGust   float32
	WindDir    string
	PressureMm int8
	PressurePa int8
	Humidity   uint8
	Daytime    string
	Polar      bool
	Season     string
	ObsTime    int
}

/**
  "date": "2020-04-22",
  "date_ts": 1587502800,
  "week": 17,
  "sunrise": "05:10",
  "sunset": "18:56",
  "moon_code": 7,
  "moon_text": "decreasing-moon",
  "parts": [{
*/
type StructForecast struct {
	Date     string
	DateTs   int
	Week     uint8
	sunrise  string
	sunset   string
	MoonCode uint8
	MoonText string
	Parts    []StructPart `json:"parts"`
}

/**
  "part_name": "evening",
  "temp_min": 6,
  "temp_max": 7,
  "temp_avg": 7,
  "feels_like": 4,
  "icon": "ovc_-ra",
  "condition": "overcast-and-light-rain",
  "daytime": "n",
  "polar": false,
  "wind_speed": 2.5,
  "wind_gust": 8.6,
  "wind_dir": "sw",
  "pressure_mm": 737,
  "pressure_pa": 983,
  "humidity": 79,
  "prec_mm": 0.6,
  "prec_period": 360,
  "prec_prob": 80
*/
type StructPart struct {
	PartName   string
	TempMin    uint8
	TempMax    uint8
	TempAvg    uint8
	FeelsLike  uint8
	Icon       string
	Condition  string
	Daytime    string
	Polar      bool
	WindSpeed  float32
	WindGust   float32
	WindDir    string
	PressureMm int8
	PressurePa int8
	Humidity   uint8
	PrecMm     float32
	PrecPeriod float32
	PrecProb   float32
}
