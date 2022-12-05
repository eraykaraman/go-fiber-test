package user

type Service interface {
	Get(id uint) (*Model, error)
	Create(model Model) (uint, error)
}

type service struct {
	respository Respository
}

var _ Service = service{}

func NewService(respository Respository) Service {
	return service{respository: respository}
}

func (s service) Get(id uint) (*Model, error) {
	return s.respository.Get(id)
}

func (s service) Create(model Model) (uint, error) {
	return s.respository.Create(model)
}
