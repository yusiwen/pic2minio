package main

import (
	"context"
	"crypto/sha1"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// IsValidUrl tests a string to determine if it is a well-structured url or not.
func IsValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

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
	var output []string
	for _, file := range argsWithoutProg {
		if IsValidUrl(file) {
			h := sha1.New()
			h.Write([]byte(file))
			base := fmt.Sprintf("%x", h.Sum(nil)) + ".png"

			func() {
				//Get the response bytes from the url
				response, err := http.Get(file)
				if err != nil {
					log.Fatalln(err)
				}
				defer response.Body.Close()

				if _, err := minioClient.PutObject(ctx, bucket, baseDir+"/"+base,
					response.Body, response.ContentLength, minio.PutObjectOptions{}); err != nil {
					log.Fatalln(err)
				}
				output = append(output, fmt.Sprintf("https://%s/%s/%s/%s", endpoint, bucket, baseDir, base))
			}()

		} else {
			base := filepath.Base(file)
			if _, err :=
				minioClient.FPutObject(ctx, bucket, baseDir+"/"+base,
					file, minio.PutObjectOptions{}); err != nil {
				log.Fatalln(err)
			}
			output = append(output, fmt.Sprintf("https://%s/%s/%s/%s", endpoint, bucket, baseDir, base))
		}
	}

	for _, s := range output {
		fmt.Println(s)
	}
}
