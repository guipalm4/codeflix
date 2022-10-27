package domain

import (
	"time"
)

type Job struct {
	ID               string `json:"id"`
	OutputBucketPath string `json:"output_bucket_path"`
	Status           string `json:"status"`
	Video            *Video
	Error            string    `json:"error"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
