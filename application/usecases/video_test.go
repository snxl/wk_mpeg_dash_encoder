package usecases_test

import (
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"github.com/snxl/wk_mpeg_dash_encoder/application/repositories"
	"github.com/snxl/wk_mpeg_dash_encoder/application/usecases"
	"github.com/snxl/wk_mpeg_dash_encoder/domain"
	"github.com/snxl/wk_mpeg_dash_encoder/framework/database"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
	"time"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func prepare() (*domain.Video, repositories.VideoRepositoryDb) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = os.Getenv("BUCKET_FILE_TEST")
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}

	return video, repo
}

func TestDownloadAndFragment(t *testing.T) {

	videoClass, videoRepo := prepare()

	video := usecases.NewVideo()

	video.Video = videoClass
	video.VideoRepository = videoRepo

	err := video.Download(os.Getenv("BUCKET_TEST"))
	require.Nil(t, err)

	err = video.Fragment()
	require.Nil(t, err)

	err = video.Encode()
	require.Nil(t, err)

	err = video.Finish()
	require.Nil(t, err)
}
