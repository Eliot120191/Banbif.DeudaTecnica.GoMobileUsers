package usuario

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Eliot120191/Banbif.DeudaTecnica.GoDomain/domain"
	"gorm.io/gorm"
)

type Repository interface {
	LoginAsync(ctx context.Context, correo string) error
	/*
		Create(ctx context.Context, user *domain.USUARIO) error
		GetAll(ctx context.Context, filters Filters, offset, limit int) ([]domain.USUARIO, error)
		Get(ctx context.Context, id int64) (*domain.USUARIO, error)
		Delete(ctx context.Context, id int64) error
		Update(ctx context.Context, id int64, correo *string) error
		Count(ctx context.Context, filters Filters) (int, error)*/
}

type repo struct {
	log *log.Logger
	db  *gorm.DB
}

func NewRepo(log *log.Logger, db *gorm.DB) Repository {
	return &repo{
		log: log,
		db:  db,
	}
}

func (repo *repo) LoginAsync(ctx context.Context, correo string) error {
	var user = domain.USUARIO{Correo: correo}
	if err := repo.db.WithContext(ctx).First(&user).Error; err != nil {
		repo.log.Println(err)
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound{"error"}
		}
		return nil
	}

	/*
		if err := repo.db.WithContext(ctx).Create(user).Error; err != nil {
			repo.log.Println(err)
			return err
		}
		repo.log.Println("user created with id: ", user.BASE.ID)*/
	return nil
}

/*
func (repo *repo) Create(ctx context.Context, user *domain.USUARIO) error {
	if err := repo.db.WithContext(ctx).Create(user).Error; err != nil {
		repo.log.Println(err)
		return err
	}
	repo.log.Println("user created with id: ", user.BASE.ID)
	return nil
}

func (repo *repo) GetAll(ctx context.Context, filters Filters, offset, limit int) ([]domain.USUARIO, error) {
	var u []domain.USUARIO

	tx := repo.db.WithContext(ctx).Model(&u)
	tx = applyFilters(tx, filters)
	tx = tx.Limit(limit).Offset(offset)
	result := tx.Order("created_at desc").Find(&u)
	if result.Error != nil {
		repo.log.Println(result.Error)
		return nil, result.Error
	}
	return u, nil

}

func (repo *repo) Get(ctx context.Context, id int64) (*domain.USUARIO, error) {
	user := domain.USUARIO{BASE: domain.BASE{ID: id}}

	if err := repo.db.WithContext(ctx).First(&user).Error; err != nil {
		repo.log.Println(err)
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound{"error"}
		}
		return nil, err
	}
	return &user, nil
}

func (repo *repo) Delete(ctx context.Context, id int64) error {
	user := domain.USUARIO{BASE: domain.BASE{ID: id}}

	result := repo.db.WithContext(ctx).Delete(&user)

	if result.Error != nil {
		repo.log.Println(result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		repo.log.Printf("user %s doesn't exists", id)
		return ErrNotFound{"error"}
	}
	return nil
}

func (repo *repo) Update(ctx context.Context, id int64, correo *string) error {

	values := make(map[string]interface{})

	if correo != nil {
		values["correo"] = *correo
	}

	result := repo.db.WithContext(ctx).Model(&domain.USUARIO{}).Where("id = ?", id).Updates(values)

	if result.Error != nil {
		repo.log.Println(result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		repo.log.Printf("user %s doesn't exists", id)
		return ErrNotFound{"error"}
	}

	return nil
}*/

func (repo *repo) Count(ctx context.Context, filters Filters) (int, error) {
	var count int64
	tx := repo.db.WithContext(ctx).Model(domain.USUARIO{})
	tx = applyFilters(tx, filters)
	if err := tx.Count(&count).Error; err != nil {
		repo.log.Println(err)
		return 0, err
	}

	return int(count), nil
}

func applyFilters(tx *gorm.DB, filters Filters) *gorm.DB {

	if filters.FirstName != "" {
		filters.FirstName = fmt.Sprintf("%%%s%%", strings.ToLower(filters.FirstName))
		tx = tx.Where("lower(first_name) like ?", filters.FirstName)
	}
	if filters.LastName != "" {
		filters.LastName = fmt.Sprintf("%%%s%%", strings.ToLower(filters.LastName))
		tx = tx.Where("lower(last_name) like ?", filters.LastName)
	}

	return tx
}
