package user

type Service interface {
	CreateUser(req CreateUserRequest) (*User, error)
	GetUserByID(id uint) (*User, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) CreateUser(req CreateUserRequest) (*User, error) {
	user := &User{
		Name:  req.Name,
		Email: req.Email,
	}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) GetUserByID(id uint) (*User, error) {
	return s.repo.GetByID(id)
}
