package tests

import (
	"backend/internal/s3service"
	"backend/internal/server"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler(t *testing.T) {
	s := &server.Server{}
	server := httptest.NewServer(http.HandlerFunc(s.HelloWorldHandler))
	defer server.Close()
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp.Body.Close()
	// Assertions
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}
	expected := "{\"message\":\"Hello World\"}"
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body. Err: %v", err)
	}
	if expected != string(body) {
		t.Errorf("expected response body to be %v; got %v", expected, string(body))
	}
}

// MockS3Client is a mock of the S3ClientAPI
type MockS3Client struct {
	mock.Mock
}

func (m *MockS3Client) ListObjectsV2(ctx context.Context, input *s3.ListObjectsV2Input, opts ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
	args := m.Called(ctx, input, opts)
	return args.Get(0).(*s3.ListObjectsV2Output), args.Error(1)
}

func TestGetYears(t *testing.T) {
	mockS3Client := new(MockS3Client)
	s3Service := s3service.NewService(mockS3Client)

	expectedYears := []string{"2020-2021", "2021-2022", "2022-2023", "2023-2024"}
	mockS3Client.On("ListObjectsV2", mock.Anything, mock.AnythingOfType("*s3.ListObjectsV2Input"), mock.Anything).Return(&s3.ListObjectsV2Output{
		CommonPrefixes: []types.CommonPrefix{
			{Prefix: aws.String("photoshoots/2020-2021/")},
			{Prefix: aws.String("photoshoots/2021-2022/")},
			{Prefix: aws.String("photoshoots/2022-2023/")},
			{Prefix: aws.String("photoshoots/2023-2024/")},
		},
	}, nil)

	years, err := s3Service.GetYears()
	assert.NoError(t, err)
	assert.Equal(t, expectedYears, years)

	mockS3Client.AssertExpectations(t)
}

// integration test
func TestGetYearsIntegration(t *testing.T) {
	if os.Getenv("RUN_INTEGRATION_TESTS") != "true" {
		t.Skip("Skipping integration test")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(os.Getenv("AWS_REGION")),
	)
	assert.NoError(t, err)

	s3Client := s3.NewFromConfig(cfg)
	s3Service := s3service.NewService(s3Client)

	years, err := s3Service.GetYears()
	assert.NoError(t, err)

	// checking if years is not empty
	assert.NotEmpty(t, years)
	// checking values
	assert.Contains(t, years, "2020-2021")
}
