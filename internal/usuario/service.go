package usuario

import (
	"context"
	"log"
)

type (
	Filters struct {
		FirstName string
		LastName  string
	}

	Service interface {
		LoginAsync(ctx context.Context, correo string) (interface{}, error)
		//Create(ctx context.Context, firstName, lastName, email, phone string) (*domain.USUARIO, error)
		/*Get(ctx context.Context, id int64) (*domain.USUARIO, error)
		GetAll(ctx context.Context, filters Filters, offset, limit int) ([]domain.USUARIO, error)
		Delete(ctx context.Context, id int64) error
		Update(ctx context.Context, id int64, firstName *string) error
		Count(ctx context.Context, filters Filters) (int, error)*/
	}
	service struct {
		log  *log.Logger
		repo Repository
	}
)

func NewService(log *log.Logger, repo Repository) Service {
	return &service{
		log:  log,
		repo: repo,
	}
}

func (s service) LoginAsync(ctx context.Context, correo string) (interface{}, error) {
	if err := s.repo.LoginAsync(ctx, correo); err != nil {
		return nil, err
	}

	return nil, nil
}

/*
func (s service) Create(ctx context.Context, firstName, lastName, email, phone string) (*domain.USUARIO, error) {
	user := domain.USUARIO{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}

	if err := s.repo.Create(ctx, &user); err != nil {
		return nil, err
	}

	return &user, nil
}*/
/*
func (s service) GetAll(ctx context.Context, filters Filters, offset, limit int) ([]domain.USUARIO, error) {

	users, err := s.repo.GetAll(ctx, filters, offset, limit)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s service) Get(ctx context.Context, id int64) (*domain.USUARIO, error) {
	user, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s service) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s service) Update(ctx context.Context, id int64, firstName *string) error {
	return s.repo.Update(ctx, id, firstName)
}

func (s service) Count(ctx context.Context, filters Filters) (int, error) {
	return s.repo.Count(ctx, filters)
}
*/
