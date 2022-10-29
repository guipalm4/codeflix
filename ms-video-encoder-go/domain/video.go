package domain

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type Video struct {
	ID         string    `valid:"uuid" json:"encoded_video_folder" gorm:"type:uuid;primary_key"`
	ResourceID string    `valid:"notnull" json:"resource_id" gorm:"type:varchar(255)"`
	FilePath   string    `valid:"notnull" json:"file_path" gorm:"type:varchar(255)"`
	CreatedAt  time.Time `valid:"-" json:"-"`
	Jobs       []*Job    `valid:"-" json:"-" gorm:"ForeignKey:VideoID"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewVideo() *Video {
	return &Video{}
}

func (video *Video) Validate() error {
	_, err := govalidator.ValidateStruct(video)

	if err != nil {
		return err
	}
	return nil
}
