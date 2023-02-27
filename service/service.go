package service

type Service struct {
	Name      string
	IPAddress string
	Port      int
}

type ServiceCommand struct {
}

func NewService(name string) *Service {
	return &Service{
		Name: name,
	}
}
