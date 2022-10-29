package domain

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
	
)

type Job struct {
	ID               string    `valid:"uuid" json:"job_id" gorm:"type:uuid;primary_key"`
	OutputBucketPath string    `valid:"notnull" json:"output_bucket_path"`
	Status           string    `valid:"notnull" json:"status"`
	Video            *Video    `valid:"-" json:"video"` 
	VideoID 				 string 	 `valid:"-" json:"video_id" gorm:"column:video_id;type:uuid;notnull"`
	Error            string    `valid:"-" json:"error"`
	CreatedAt        time.Time `valid:"-" json:"created_at"`
	UpdatedAt        time.Time `valid:"-" json:"updated_at"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewJob(output string, status string, video *Video) (*Job, error) {
	job := Job{
		OutputBucketPath: output,
		Status: status,
		Video: video,
	}
	job.prepare()
	err := job.Validate()

	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (job *Job) prepare() {
	job.ID = uuid.NewV4().String()
	job.CreatedAt = time.Now()
	job.UpdatedAt = time.Now()
}


func (job *Job) Validate() error {
	_, err := govalidator.ValidateStruct(job)

	if err != nil {
		return err
	}
	return nil
}
