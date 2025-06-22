package order

type Service interface {
	ListOrders() ([]Order, error)
	CreateOrder(o *Order) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) ListOrders() ([]Order, error) {
	return s.repo.List()
}

func (s *service) CreateOrder(o *Order) error {
	return s.repo.Create(o)
}
