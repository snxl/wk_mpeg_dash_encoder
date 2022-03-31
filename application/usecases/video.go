package usecases

import (
	"github.com/snxl/wk_mpeg_dash_encoder/application/repositories"
	"github.com/snxl/wk_mpeg_dash_encoder/domain"
	"cloud.google.com/go/storage"
)

type Video struct {
	Video *domain.Video
	VideoRepository repositories.VideoRepository
}

func NewVideo () Video {
	return Video{}
}

func (v *Video) Download(bucketName string) error {
 return nil
}