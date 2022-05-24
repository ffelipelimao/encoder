package domain_test

import (
	"testing"
	"time"

	"github.com/ffelipelimao/encoder/src/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewJob(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewString()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	job, err := domain.NewJob("path", "converted", video)
	require.NotNil(t, job)
	require.Nil(t, err)
}
