package user

import "gorm.io/gorm"

type Respository interface {
	Get(id uint) (*Model, error)
	Create(model Model) (uint, error)
	Migration() error
}

type respository struct {
	db *gorm.DB
}

var _ Respository = respository{}

//check receiver == source
func NewRepository(db *gorm.DB) Respository {
	return respository{db: db}
}

func (repo respository) Get(id uint) (*Model, error) {
	model := &Model{ID: id}
	err := repo.db.First(model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (repo respository) Create(model Model) (uint, error) {
	err := repo.db.Create(&model).Error
	if err != nil {
		return 0, err
	}
	return model.ID, nil
}

func (repo respository) Migration() error {
	return repo.db.AutoMigrate(&Model{})
}
