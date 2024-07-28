package service

import (
	"regexp"
)

type ServiceHealth struct {
	Name    string `json:"name"`
	Health  string `json:"health"`
	Message string `json:"message,omitempty"`
	Status  string `json:"status"`
}

func (s *ServiceHealth) ProvideSystemctlStatus() error {
	status, err := GetServiceStatus(s.Name)
	if err != nil {
		return err
	}
	strToReplace := s.Name + "\\["
	m := regexp.MustCompile(strToReplace)
	replace := "\n\n" + s.Name + "["
	status = m.ReplaceAllString(status, replace)
	// loc := m.FindStringIndex(status)
	// fmt.Println("loc", loc)
	// if loc != nil {
	// 	before := status[:loc[0]]
	// 	after := status[loc[1]:]
	// 	replaced := strings.Replace(status[loc[0]:loc[1]], strToReplace, "<pre>"+strToReplace, 1)
	// 	status = before + replaced + after + "</pre>"
	// }
	s.Status = status
	return nil
}
