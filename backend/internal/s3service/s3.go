package s3service

import (
	"backend/loggers"
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Service interface
type Service interface {
	GetYears(ctx context.Context) ([]string, error)
	GetEvents(ctx context.Context, year string) ([]string, error)
	ListPhotos(ctx context.Context, year, event string) ([]string, error)
	GetPhotos(ctx context.Context, year, event string) ([]string, error)
	GetPresignedURL(bucket, key string, lifetimeSecs int64) (string, error)
}

// S3ClientAPI defines the methods used from the S3 client.
type S3ClientAPI interface {
	ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
}

// PresignerAPI defines the methods used from the presigner.
type PresignerAPI interface {
	GetObject(bucketName string, objectKey string, lifetimeSecs int64) (*v4.PresignedHTTPRequest, error)
}

// service struct
type service struct {
	s3Client  S3ClientAPI
	presigner PresignerAPI
}

// NewMockService creates a new instance of the service with provided mock clients
func NewMockService(s3Client S3ClientAPI, presigner PresignerAPI) Service {
	return &service{
		s3Client:  s3Client,
		presigner: presigner,
	}
}

// NewService creates an instance intended for production
func NewService() Service {
	cfg, err := NewAWSConfig()
	if err != nil {
		loggers.Error.Printf("failed to load AWS config: %v", err)
		return nil
	}
	s3Client := s3.NewFromConfig(cfg)
	presigner := s3.NewPresignClient(s3Client)
	return &service{
		s3Client:  s3Client,
		presigner: &Presigner{PresignClient: presigner},
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

func (s *service) GetPresignedURL(bucket, key string, lifetimeSecs int64) (string, error) {
	request, err := s.presigner.GetObject(bucket, key, lifetimeSecs)
	if err != nil {
		return "", err
	}
	return request.URL, nil
}
