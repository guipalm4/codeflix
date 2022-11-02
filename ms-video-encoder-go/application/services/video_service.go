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

var StoragePath = os.Getenv("LOCAL_STORAGE_PATH")
var Separator = string(os.PathSeparator)

type VideoService struct {
	Video           *domain.Video
	VideoRepository repositories.VideoRepository
}

func NewVideoService() VideoService {
	return VideoService{}
}

func (v *VideoService) Download(bucketName string) error {

	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	if err != nil {
		return err
	}

	bucket := client.Bucket(bucketName)
	obj := bucket.Object(v.Video.FilePath)
	reader, err := obj.NewReader(ctx)

	if err != nil {
		return err
	}
	defer reader.Close()
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	file, err := os.Create(getVideoPath(v) + ".mp4")
	if err != nil {
		return err
	}

	_, err = file.Write(body)
	if err != nil {
		return err
	}
	defer file.Close()

	log.Printf("video %v has been stored", v.Video.ID)
	return nil
}

func (v *VideoService) Fragment() error {

	err := os.MkdirAll(getVideoPath(v), os.ModePerm)
	if err != nil {
		return err
	}

	source := getVideoPath(v) + ".mp4"
	target := getVideoPath(v) + ".frag"

	cmd := exec.Command("mp4fragment", source, target)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return err
	}

	printOutput(output)
	return nil
}

func (v *VideoService) Encode() error {

	cmdArgs := []string{}
	cmdArgs = append(cmdArgs, getVideoPath(v)+".frag")
	cmdArgs = append(cmdArgs, "--use-segment-timeline")
	cmdArgs = append(cmdArgs, "-o")
	cmdArgs = append(cmdArgs, getVideoPath(v))
	cmdArgs = append(cmdArgs, "-f")
	cmdArgs = append(cmdArgs, "--exec-dir")
	cmdArgs = append(cmdArgs, "/opt/bento4/bin/")
	cmd := exec.Command("mp4dash", cmdArgs...)

	output, err := cmd.CombinedOutput()

	if err != nil {
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

func getVideoPath(v *VideoService) string {
	return StoragePath + Separator + v.Video.ID
}
