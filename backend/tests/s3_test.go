package tests

import (
	"backend/internal/s3service"
	"context"
	"os"
	"testing"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockS3Client is a mock of the S3ClientAPI
type MockS3Client struct {
	mock.Mock
}

func (m *MockS3Client) ListObjectsV2(ctx context.Context, input *s3.ListObjectsV2Input, opts ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
	args := m.Called(ctx, input, opts)
	return args.Get(0).(*s3.ListObjectsV2Output), args.Error(1)
}

type MockPresigner struct {
	mock.Mock
}

func (m *MockPresigner) GetObject(bucketName string, objectKey string, lifetimeSecs int64) (*v4.PresignedHTTPRequest, error) {
	args := m.Called(bucketName, objectKey, lifetimeSecs)
	return args.Get(0).(*v4.PresignedHTTPRequest), args.Error(1)
}

func (m *MockPresigner) PutObject(bucketName string, objectKey string, lifetimeSecs int64) (*v4.PresignedHTTPRequest, error) {
	args := m.Called(bucketName, objectKey, lifetimeSecs)
	return args.Get(0).(*v4.PresignedHTTPRequest), args.Error(1)
}

func (m *MockPresigner) DeleteObject(bucketName string, objectKey string) (*v4.PresignedHTTPRequest, error) {
	args := m.Called(bucketName, objectKey)
	return args.Get(0).(*v4.PresignedHTTPRequest), args.Error(1)
}

func TestGetPhotoshootYears(t *testing.T) {
	mockS3Client := new(MockS3Client)
	mockPresigner := new(MockPresigner)

	s3Service := s3service.NewMockService(mockS3Client, mockPresigner)

	expectedYears := []string{"2020-2021", "2021-2022", "2022-2023", "2023-2024"}
	mockS3Client.On("ListObjectsV2", mock.Anything, mock.AnythingOfType("*s3.ListObjectsV2Input"), mock.Anything).Return(&s3.ListObjectsV2Output{
		CommonPrefixes: []types.CommonPrefix{
			{Prefix: aws.String("photoshoots/2020-2021/")},
			{Prefix: aws.String("photoshoots/2021-2022/")},
			{Prefix: aws.String("photoshoots/2022-2023/")},
			{Prefix: aws.String("photoshoots/2023-2024/")},
		},
	}, nil)

	years, err := s3Service.GetPhotoshootYears(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedYears, years)

	mockS3Client.AssertExpectations(t)
}

// integration test
func TestGetPhotoshootYearsIntegration(t *testing.T) {
	if os.Getenv("RUN_INTEGRATION_TESTS") != "true" {
		t.Skip("Skipping integration test")
	}

	s3Service := s3service.NewService()

	years, err := s3Service.GetPhotoshootYears(context.Background())
	assert.NoError(t, err)

	// checking if years is not empty
	assert.NotEmpty(t, years)
	// checking values
	assert.Contains(t, years, "2020-2021")
}

func TestGetPhotoshootEvents(t *testing.T) {
	mockS3Client := new(MockS3Client)
	mockPresigner := new(MockPresigner)

	s3Service := s3service.NewMockService(mockS3Client, mockPresigner)

	// define mock response
	mockS3Client.On("ListObjectsV2", mock.Anything, mock.AnythingOfType("*s3.ListObjectsV2Input"), mock.Anything).Return(&s3.ListObjectsV2Output{
		CommonPrefixes: []types.CommonPrefix{
			{Prefix: aws.String("photoshoots/2020-2021/Event1")},
			{Prefix: aws.String("photoshoots/2020-2021/Event2")},
		},
	}, nil)

	// call function to test
	events, err := s3Service.GetPhotoshootEvents(context.Background(), "2020-2021")
	assert.NoError(t, err)
	assert.Equal(t, []string{"Event1", "Event2"}, events)

	mockS3Client.AssertExpectations(t)
}

func TestGetPhotoshootEventsIntegration(t *testing.T) {
	if os.Getenv("RUN_INTEGRATION_TESTS") != "true" {
		t.Skip("Skipping integration test")
	}

	s3Service := s3service.NewService()

	expectedYear := "2020-2021"
	expectedEvents := []string{"AASA 2020", "CNY 2021"}

	events, err := s3Service.GetPhotoshootEvents(context.Background(), expectedYear)
	assert.NoError(t, err)

	// checking if events is not empty
	assert.NotEmpty(t, events)
	// checking values
	assert.Equal(t, expectedEvents, events)
}

func TestListPhotoshootPhotos(t *testing.T) {
	mockS3Client := new(MockS3Client)
	mockPresigner := new(MockPresigner)

	s3Service := s3service.NewMockService(mockS3Client, mockPresigner)

	mockS3Client.On("ListObjectsV2", mock.Anything, mock.AnythingOfType("*s3.ListObjectsV2Input"), mock.Anything).Return(&s3.ListObjectsV2Output{
		Contents: []types.Object{
			{Key: aws.String("photoshoots/2020-2021/Event1/photo1.jpg")},
			{Key: aws.String("photoshoots/2020-2021/Event1/photo2.jpg")},
		},
	}, nil)

	photos, err := s3Service.ListPhotoshootPhotos(context.Background(), "2020-2021", "Event1")
	assert.NoError(t, err)
	assert.Equal(t, []string{"photo1.jpg", "photo2.jpg"}, photos)

	mockS3Client.AssertExpectations(t)
}

func TestListPhotoshootPhotosIntegration(t *testing.T) {
	if os.Getenv("RUN_INTEGRATION_TESTS") != "true" {
		t.Skip("Skipping integration test")
	}

	s3Service := s3service.NewService()

	expectedYear := "2020-2021"
	expectedEvent := "AASA 2020"
	expectedPhotos := "DSC01330.jpg"

	photos, err := s3Service.ListPhotoshootPhotos(context.Background(), expectedYear, expectedEvent)
	assert.NoError(t, err)

	// checking if photos is not empty
	assert.NotEmpty(t, photos)
	// checking values
	assert.Contains(t, photos, expectedPhotos)
}

func TestGetPhotoshootPhotos(t *testing.T) {
	mockS3Client := new(MockS3Client)
	mockPresigner := new(MockPresigner)

	s3Service := s3service.NewMockService(mockS3Client, mockPresigner)

	year := "2020-2021"
	event := "AASA 2020"
	bucket := os.Getenv("S3_BUCKET_NAME")

	// mock s3 response
	mockS3Client.On("ListObjectsV2", mock.Anything, mock.AnythingOfType("*s3.ListObjectsV2Input"), mock.Anything).Return(&s3.ListObjectsV2Output{
		Contents: []types.Object{
			{Key: aws.String("photoshoots/2020-2021/AASA 2020/DSC01330.jpg")},
			{Key: aws.String("photoshoots/2020-2021/AASA 2020/DSC01331.jpg")},
		},
	}, nil)

	// mock presigner response
	mockPresigner.On("GetObject", bucket, "photoshoots/2020-2021/AASA 2020/DSC01330.jpg", int64(900)).Return(&v4.PresignedHTTPRequest{URL: "https://presigned.url/DSC01330.jpg"}, nil)
	mockPresigner.On("GetObject", bucket, "photoshoots/2020-2021/AASA 2020/DSC01331.jpg", int64(900)).Return(&v4.PresignedHTTPRequest{URL: "https://presigned.url/DSC01331.jpg"}, nil)

	// call function to test
	urls, err := s3Service.GetPhotoshootPhotos(context.Background(), year, event)
	assert.NoError(t, err)
	assert.Equal(t, []string{"https://presigned.url/DSC01330.jpg", "https://presigned.url/DSC01331.jpg"}, urls)

	mockS3Client.AssertExpectations(t)
	mockPresigner.AssertExpectations(t)
}

func TestGetPhotoshootPhotosIntegration(t *testing.T) {
	if os.Getenv("RUN_INTEGRATION_TESTS") != "true" {
		t.Skip("Skipping integration test")
	}

	s3Service := s3service.NewService()

	expectedYear := "2020-2021"
	expectedEvent := "AASA 2020"

	photos, err := s3Service.GetPhotoshootPhotos(context.Background(), expectedYear, expectedEvent)
	assert.NoError(t, err)

	// checking if photos is not empty
	assert.NotEmpty(t, photos)
	// checking values
	for _, photo := range photos {
		assert.Contains(t, photo, "https://")
	}
}

func TestGenerateEventImageUploadURL(t *testing.T) {
	mockS3Client := new(MockS3Client)
	mockPresigner := new(MockPresigner)

	s3Service := s3service.NewMockService(mockS3Client, mockPresigner)

	eventID := "123"
	filename := "photo.jpg"
	bucket := os.Getenv("S3_BUCKET_NAME")
	objectKey := "events/123/photo.jpg"
	lifetimeSecs := int64(900)

	expectedURL := "https://presigned.url/photo.jpg"

	// mock presigner response
	mockPresigner.On("PutObject", bucket, objectKey, lifetimeSecs).Return(&v4.PresignedHTTPRequest{URL: expectedURL}, nil)

	// call function to test
	url, err := s3Service.GenerateEventImageUploadURL(eventID, filename, lifetimeSecs)
	assert.NoError(t, err)
	assert.Equal(t, expectedURL, url)

	mockPresigner.AssertExpectations(t)
}

func TestGenerateEventImageUploadURLIntegration(t *testing.T) {
	if os.Getenv("RUN_INTEGRATION_TESTS") != "true" {
		t.Skip("Skipping integration test")
	}

	s3Service := s3service.NewService()

	expectedEventID := "123"
	expectedFilename := "photo.jpg"
	expectedLifetimeSecs := int64(900)

	url, err := s3Service.GenerateEventImageUploadURL(expectedEventID, expectedFilename, expectedLifetimeSecs)
	assert.NoError(t, err)

	// checking if url is not empty
	assert.NotEmpty(t, url)
	// checking values
	assert.Contains(t, url, "https://")
}
