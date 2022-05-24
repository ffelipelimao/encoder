package repositories

import (
	"fmt"

	"github.com/ffelipelimao/encoder/src/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VideoRepository interface {
	Insert(video *domain.Video) (*domain.Video, error)
	Find(id string) (*domain.Video, error)
}

func NewVideoRepository(db *gorm.DB) *VideoRepositoryDB {
	return &VideoRepositoryDB{
		DB: db,
	}
}

type VideoRepositoryDB struct {
	DB *gorm.DB
}

func (repo VideoRepositoryDB) Insert(video *domain.Video) (*domain.Video, error) {
	if video.ID == "" {
		video.ID = uuid.NewString()
	}

	err := repo.DB.Create(video).Error
	if err != nil {
		return nil, err
	}

	return video, nil
}

func (repo VideoRepositoryDB) Find(id string) (*domain.Video, error) {
	var video domain.Video
	repo.DB.Preload("Jobs").First(&video, "id=?", id)

	if video.ID == "" {
		return nil, fmt.Errorf("video does not exist")
	}

	return &video, nil
}
