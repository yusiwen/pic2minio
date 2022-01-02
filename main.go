package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	ctx := context.Background()
	config, err := InitConfig()
	if err != nil {
		log.Fatalln(err)
	}
	endpoint := config.EndPoint
	accessKeyID := config.AccessKey
	secretAccessKey := config.SecretKey
	bucket := config.Bucket
	baseDir := config.BaseDir

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: true,
	})
	if err != nil {
		log.Fatalln(err)
	}

	argsWithoutProg := os.Args[1:]
	output := []string{}
	for _, file := range argsWithoutProg {
		base := filepath.Base(file)
		if _, err :=
			minioClient.FPutObject(ctx, bucket, baseDir+base,
				file, minio.PutObjectOptions{}); err != nil {
			log.Fatalln(err)
		}
		output = append(output, fmt.Sprintf("https://%s/%s/%s/%s", endpoint, bucket, baseDir, base))
	}

	for _, s := range output {
		fmt.Println(s)
	}
}
