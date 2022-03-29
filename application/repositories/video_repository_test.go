package repositories_test

import (
	uuid "github.com/satori/go.uuid"
	"github.com/snxl/wk_mpeg_dash_encoder/application/repositories"
	"github.com/snxl/wk_mpeg_dash_encoder/domain"
	"github.com/snxl/wk_mpeg_dash_encoder/framework/database"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewVideoRepository(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	v, err := repo.Find(video.ID)

	require.NotEmpty(t, v.ID)
	require.Nil(t, err)
	require.Equal(t, v.ID, video.ID)
}
