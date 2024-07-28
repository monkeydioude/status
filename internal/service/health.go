package service

type ServiceHealth struct {
	Name    string `json:"name"`
	Health  string `json:"health"`
	Message string `json:"message,omitempty"`
	Status  string `json:"status"`
}

func (s *ServiceHealth) ProvideSystemctlStatus() error {
	status, err := GetServiceStatus(s.Name)
	s.Status = status
	return err
}
