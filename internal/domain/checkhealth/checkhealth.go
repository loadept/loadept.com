package checkhealth

type CheckHealth struct {
	Status         string          `json:"status"`
	Timestamp      string          `json:"timestamp"`
	Timezone       string          `json:"timezone"`
	UptimeS        string          `json:"uptime_seconds"`
	ResponseTimeMs string          `json:"response_time_ms"`
	Services       ServicesChecked `json:"services"`
}

type ServicesChecked struct {
	Redis    string `json:"redis"`
	Database string `json:"database"`
}
