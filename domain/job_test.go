package domain_test

import (
	uuid "github.com/satori/go.uuid"
	"github.com/snxl/wk_mpeg_dash_encoder/domain"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewJob(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "Path"
	video.CreatedAt = time.Now()

	job, err := domain.NewJob("path", "Converted", video)

	require.NotNil(t, job)
	require.Nil(t, err)
}
