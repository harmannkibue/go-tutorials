package UberDi

type PersonService struct {
	config     *Config
	repository *PersonRepository
}

func (service *PersonService) FindAll() []*Person {
	if service.config.Enabled {
		return service.repository.FindAll()
	}

	return []*Person{}
}

func NewPersonService(config *Config, repository *PersonRepository) *PersonService {
	return &PersonService{config: config, repository: repository}
}
