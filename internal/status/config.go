package status

type Config struct {
	Name           string `json:"name"`
	HealthcheckUrl string `json:"healthcheck_url,omitempty"`
	Daemon         bool   `json:"daemon"`
}
