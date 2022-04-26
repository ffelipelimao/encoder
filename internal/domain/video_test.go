package domain_test

import (
	"testing"
	"time"

	"github.com/ffelipelimao/encoder/internal/domain"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()

	err := video.Validate()
	require.Error(t, err)
}

func TestVideoIDIsNotUUID(t *testing.T) {
	video := domain.NewVideo()
	video.ID = "fail"
	video.ResourceID = "fail"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewString()
	video.ResourceID = uuid.NewString()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Nil(t, err)
}
