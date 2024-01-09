package s3service

import (
	"fmt"
	"log"
	"os"
	"time"
)

func (s *service) GeneratePresignedUploadURL(event, filename string, lifetimeSecs int64) (string, error) {
	startTime := time.Now()
	bucket := os.Getenv("S3_BUCKET_NAME")
	prefix := fmt.Sprintf("events/%s/%s", event, filename)

	req, err := s.presigner.PutObject(bucket, prefix, lifetimeSecs)
	if err != nil {
		return "", fmt.Errorf("failed to get presigned url: %v", err)
	}

	elapsedTime := time.Since(startTime)
	log.Printf("GenerateUploaderURL took %s", elapsedTime)

	return req.URL, nil
}

func (s *service) DevGeneratePresignedUploadURL(event, filename string, lifetimeSecs int64) (string, error) {
	startTime := time.Now()
	bucket := os.Getenv("S3_BUCKET_NAME")
	prefix := fmt.Sprintf("testing/%s/%s", event, filename)

	req, err := s.presigner.PutObject(bucket, prefix, lifetimeSecs)
	if err != nil {
		return "", fmt.Errorf("failed to get presigned url: %v", err)
	}

	elapsedTime := time.Since(startTime)
	log.Printf("GenerateUploaderURL took %s", elapsedTime)

	return req.URL, nil
}
