package services

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"cloud.google.com/go/storage"
	"github.com/ffelipelimao/encoder/src/application/repositories"
	"github.com/ffelipelimao/encoder/src/domain"
)

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

	bkt := client.Bucket(bucketName)
	obj := bkt.Object(v.Video.FilePath)

	r, err := obj.NewReader(ctx)
	if err != nil {
		return err
	}

	defer r.Close()

	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	f, err := os.Create(os.Getenv("localStoragePath") + "/" + v.Video.ID + ".mp4")
	if err != nil {
		return err
	}

	_, err = f.Write(body)
	if err != nil {
		return err
	}

	defer f.Close()

	log.Printf("video %s has been stored", v.Video.ID)

	return nil
}

func (v *VideoService) Fragment() error {
	localStoragePath := v.getLocalStoragePath()

	err := os.Mkdir(localStoragePath, os.ModePerm)
	if err != nil {
		return err
	}

	source := localStoragePath + ".mp4"
	target := localStoragePath + ".mp4"

	cmd := exec.Command("mp4fragment", source, target)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	printOutput(output)

	return nil

}

func (v *VideoService) Encode() error {

	localStoragePath := v.getLocalStoragePath()
	cmdArgs := []string{}
	cmdArgs = append(cmdArgs, localStoragePath+".frag")
	cmdArgs = append(cmdArgs, "--use-segment-timeline")
	cmdArgs = append(cmdArgs, "-o")
	cmdArgs = append(cmdArgs, localStoragePath)
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

func (v *VideoService) Finish() error {
	localStoragePath := v.getLocalStoragePath()

	err := os.Remove(localStoragePath + ".mp4")
	if err != nil {
		log.Println("error removing mp4", v.Video.ID)
		return err
	}

	err = os.Remove(localStoragePath + ".frag")
	if err != nil {
		log.Println("error removing frag", v.Video.ID)
		return err
	}

	err = os.RemoveAll(localStoragePath)
	if err != nil {
		log.Println("error removing removing dir", v.Video.ID)
		return err
	}

	log.Println("file has been removed", v.Video.ID)

	return nil

}

func printOutput(out []byte) {
	if len(out) > 0 {
		log.Printf("=====> Output: %s\n", string(out))
	}
}

func (v *VideoService) getLocalStoragePath() string {
	return os.Getenv("localStoragePath") + "/" + v.Video.ID
}
