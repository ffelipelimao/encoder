package repositories

import (
	"fmt"

	"github.com/ffelipelimao/encoder/src/domain"
	"gorm.io/gorm"
)

type JobRepository interface {
	Insert(video *domain.Job) (*domain.Job, error)
	Find(id string) (*domain.Job, error)
	Update(video *domain.Job) (*domain.Job, error)
}

type JobRepositoryDB struct {
	DB *gorm.DB
}

func (repo JobRepositoryDB) Insert(job *domain.Job) (*domain.Job, error) {

	err := repo.DB.Create(job).Error
	if err != nil {
		return nil, err
	}

	return job, nil
}

func (repo JobRepositoryDB) Find(id string) (*domain.Job, error) {
	var job domain.Job
	repo.DB.Preload("Video").First(&job, "id=?", id)

	if job.ID == "" {
		return nil, fmt.Errorf("video does not exist")
	}

	return &job, nil
}

func (repo JobRepositoryDB) Update(job *domain.Job) (*domain.Job, error) {

	err := repo.DB.Save(&job).Error
	if err != nil {
		return nil, err
	}

	return job, nil
}
