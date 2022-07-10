package usecases_test

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/snxl/wk_mpeg_dash_encoder/application/usecases"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func TestVideoUpload(t *testing.T) {

	video, repo := prepare()

	videoService := usecases.NewVideo()
	videoService.Video = video
	videoService.VideoRepository = repo
	fmt.Print(videoService)
	err := videoService.Download(os.Getenv("BUCKET_TEST"))
	require.Nil(t, err)

	err = videoService.Fragment()
	require.Nil(t, err)

	err = videoService.Encode()
	require.Nil(t, err)

	videoUpload := usecases.NewVideoUpload()
	videoUpload.OutputBucket = os.Getenv("BUCKET_TEST")
	videoUpload.VideoPath = os.Getenv("LOCAL_STORAGE_PATH") + "/" + video.ID

	doneUpload := make(chan string)
	go videoUpload.ProcessUpload(50, doneUpload)

	result := <-doneUpload
	require.Equal(t, result, "uploaded completed")

	err = videoService.Finish()
	require.Nil(t, err)

}
