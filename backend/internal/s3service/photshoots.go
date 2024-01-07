package s3service

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

func (s *service) GetYears() ([]string, error) {
	bucket := os.Getenv("S3_BUCKET_NAME")
	prefix := "photoshoots/"

	// call to list
	output, err := s.s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket:    aws.String(bucket),
		Prefix:    aws.String(prefix),
		Delimiter: aws.String("/"),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list objects from s3: %v", err)
	}

	var years []string
	for _, content := range output.CommonPrefixes {
		year := strings.TrimPrefix(*content.Prefix, prefix)
		year = strings.TrimSuffix(year, "/")
		years = append(years, year)
	}

	return years, nil
}

func (s *service) GetEvents(year string) ([]string, error) {
	bucket := os.Getenv("S3_BUCKET_NAME")
	prefix := fmt.Sprintf("photoshoots/%s/", year)

	output, err := s.s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket:    aws.String(bucket),
		Prefix:    aws.String(prefix),
		Delimiter: aws.String("/"),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list objects from s3: %v", err)
	}

	var events []string
	for _, content := range output.CommonPrefixes {
		event := strings.TrimPrefix(*content.Prefix, prefix)
		event = strings.TrimSuffix(event, "/")
		events = append(events, event)
	}

	return events, nil
}
