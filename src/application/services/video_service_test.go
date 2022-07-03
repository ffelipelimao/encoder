package services_test

import (
	"log"
	"testing"
	"time"

	"github.com/ffelipelimao/encoder/src/application/repositories"
	"github.com/ffelipelimao/encoder/src/application/services"
	"github.com/ffelipelimao/encoder/src/domain"
	"github.com/ffelipelimao/encoder/src/framework/database"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func init() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatalf("error loading .env file")
	}
}

func prepare() (*domain.Video, repositories.VideoRepositoryDB) {
	db := database.NewDbTest()

	/* TODO: Close connection
	defer db.Close()*/

	video := domain.NewVideo()
	video.ID = uuid.NewString()
	video.FilePath = "test.mp4"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDB{DB: db}
	return video, repo
}

func TestVideoServiceDownload(t *testing.T) {
	video, repo := prepare()

	videoService := services.NewVideoService()
	videoService.Video = video
	videoService.VideoRepository = repo

	err := videoService.Download("bucketname")
	require.Nil(t, err)

	err = videoService.Fragment()
	require.Nil(t, err)

	err = videoService.Encode()
	require.Nil(t, err)

	err = videoService.Finish()
	require.Nil(t, err)

}
