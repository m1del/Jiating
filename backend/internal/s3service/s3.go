package s3service

import (
	"backend/loggers"
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Service interface
type Service interface {
	GetYears() ([]string, error)
}

// S3ClientAPI defines the methods used from the S3 client.
type S3ClientAPI interface {
	ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
}

// service struct
type service struct {
	s3Client S3ClientAPI
}

// NewService creates a new instance of the service with the provided S3 client.
func NewService(s3Client S3ClientAPI) Service {
	return &service{
		s3Client: s3Client,
	}
}

// newAWSConfig creates and returns a new AWS configuration.
func NewAWSConfig() (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(os.Getenv("AWS_REGION")),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")),
	)
	if err != nil {
		loggers.Error.Printf("failed to load AWS config: %v", err)
		return aws.Config{}, err
	}
	return cfg, nil
}
