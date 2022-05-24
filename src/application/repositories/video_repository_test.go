package repositories_test

import (
	"testing"
	"time"

	"github.com/ffelipelimao/encoder/src/application/repositories"
	"github.com/ffelipelimao/encoder/src/domain"
	"github.com/ffelipelimao/encoder/src/framework/database"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestVideoRepository(t *testing.T) {
	db := database.NewDbTest()

	/* TODO: Close connection
	defer db.Close()*/

	video := domain.NewVideo()
	video.ID = uuid.NewString()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDB{DB: db}
	repo.Insert(video)

	v, err := repo.Find(video.ID)

	require.NotEmpty(t, v.ID)
	require.Nil(t, err)
	require.Equal(t, v.ID, video.ID)
}
