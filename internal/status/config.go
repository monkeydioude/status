package status

type ServiceHealth struct {
	Name    string `json:"name"`
	Health  string `json:"health"`
	Message string `json:"message,omitempty"`
}

type Config struct {
	Name           string `json:"name"`
	HealthcheckUrl string `json:"healthcheck_url,omitempty"`
}
