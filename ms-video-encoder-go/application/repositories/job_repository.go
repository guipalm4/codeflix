package repositories

import (
	"encoder/domain"
	"fmt"

	"github.com/jinzhu/gorm"
)

type JobRepository interface {
	Insert(job *domain.Job) (*domain.Job, error)
	Update(job *domain.Job) (*domain.Job, error)
	Find(id string) (*domain.Job, error)
}

type JobrepositoryDb struct {
	Db *gorm.DB
}

func NewJobRepositoryDb(db *gorm.DB) *JobrepositoryDb {
	return &JobrepositoryDb{Db: db}
}

func (repo JobrepositoryDb) Insert(job *domain.Job) (*domain.Job, error) {

	err := repo.Db.Create(job).Error

	if err != nil {
		return nil, err
	}

	return job, nil
}

func (repo JobrepositoryDb) Find(id string) (*domain.Job, error) {
	var job domain.Job
  repo.Db.Preload("Video").First(&job, "id =?", id)

	if job.ID == "" {
		return nil, fmt.Errorf("job does not exist")
	}

	return &job, nil
}

func (repo JobrepositoryDb) Update(job *domain.Job) (*domain.Job, error) {
	err := repo.Db.Save(&job).Error

	if err!= nil {
    return nil, err
  }
	return job, nil
}