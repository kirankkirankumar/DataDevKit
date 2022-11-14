package repository

/*import "github.com/kirankkirankumar/gqlgen-ddk/graph/model"

type VideoRepository interface {
	CreateVideo(video *model.Video) error
	GetVideos() ([]*model.Video, error)
	MigrateVideos() error
}

func (r *repository) CreateVideo(video *model.Video) error {

	err := r.db.Create(video).Error

	return err
}

func (r *repository) GetVideos() ([]*model.Video, error) {

	videos := make([]*model.Video, 0)
	err := r.db.Find(videos).Error

	return videos, err
}

func (r *repository) MigrateVideos() error {

	err := r.db.AutoMigrate(&model.Video{})

	return err
}
*/