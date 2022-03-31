

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

func TestJobRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)

	require.Nil(t, err)

	repoJob := repositories.JobRepositoryDb{Db: db}
	repoJob.Insert(job)

	jobRes, err := repoJob.Find(job.ID)
	require.NotEmpty(t, jobRes.ID)
	require.Nil(t, err)
	require.Equal(t, jobRes.ID, job.ID)
	require.Equal(t, jobRes.Video.ID, video.ID)
}

func TestJobRepositoryDbUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)

	require.Nil(t, err)

	repoJob := repositories.JobRepositoryDb{Db: db}
	repoJob.Insert(job)

	job.Status = "Complete"

	repoJob.Update(job)

	jobRes, err := repoJob.Find(job.ID)
	require.NotEmpty(t, jobRes.ID)
	require.Nil(t, err)
	require.Equal(t, jobRes.Status, job.Status)
}