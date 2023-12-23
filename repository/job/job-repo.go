package jobrepo

import "gorm.io/gorm"

type JobRepo struct {
	DB *gorm.DB
}

func NewJobRepo(db *gorm.DB) *JobRepo {
	return &JobRepo{DB : db}
}

func (r *JobRepo) FindAll() {}

func (r *JobRepo) FindByPk() {}

func (r *JobRepo) Save() {}

func (r *JobRepo) Update() {}

func (r *JobRepo) Delete() {}
