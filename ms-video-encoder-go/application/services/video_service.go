package services

import (
	"context"
	"encoder/application/repositories"
	"encoder/domain"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"cloud.google.com/go/storage"
)


type VideoService struct {
	Video	*domain.Video
	VideoRepository repositories.VideoRepository	
}

func NewVideoService() VideoService {
	return VideoService{}
}

func (v *VideoService) Download(bucketName string) error {
	
	ctx := context.Background()	
	client, err := storage.NewClient(ctx)

	if err!= nil {
    return err
  }

	bucket := client.Bucket(bucketName)
	obj := bucket.Object(v.Video.FilePath)
	reader, err := obj.NewReader(ctx)

	if err!= nil {
    return err
  }
	defer reader.Close()
	body, err := ioutil.ReadAll(reader)
	if err!= nil {
    return err
  }

	file, err := os.Create(os.Getenv("LOCAL_STORAGE_PATH") + string(os.PathSeparator) + v.Video.ID + ".mp4")
	if err!= nil {
    return err
  }

	_, err = file.Write(body)
	if err!= nil {
    return err
  }
	defer file.Close()

	log.Printf("video %v has been stored", v.Video.ID)
	return nil
}

func (v *VideoService) Fragment() error {

	path := os.Getenv("LOCAL_STORAGE_PATH")
	separator := string(os.PathSeparator)
	videoID := v.Video.ID
	
	err := os.MkdirAll(path + separator + videoID, os.ModePerm)
  if err!= nil {
    return err
  }
	
	source := path + separator + videoID + ".mp4"
	target := path + separator + videoID + ".frag"

	cmd := exec.Command("mp4fragment", source, target)
	output, err := cmd.CombinedOutput()

  if err!= nil {
		return err
	}
	
	printOutput(output)
	return nil
}

func printOutput(out []byte) {
	if len(out) > 0 {
		log.Printf("========> Output: %s\n", string(out))
	}
}
