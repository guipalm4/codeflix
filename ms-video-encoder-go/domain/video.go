package domain

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type Video struct {
	ID         string    `valid:"uuid" json:"id"`
	ResourceID string    `valid:"notnull" json:"resource_id"`
	FilePath   string    `valid:"notnull" json:"file_path"`
	CreatedAt  time.Time `valid:"-" json:"created_at"`
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
