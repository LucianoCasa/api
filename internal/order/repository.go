package order

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository interface {
	List() ([]Order, error)
	Create(order *Order) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) List() ([]Order, error) {
	var orders []Order
	err := r.db.Find(&orders).Error
	return orders, err
}

func (r *repository) Create(order *Order) error {
	return r.db.Create(order).Error
}

func InitMySQL(dsn string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
