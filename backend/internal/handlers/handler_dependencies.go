package handlers

import "backend/internal/s3service"

type HandlerDependencies struct {
	S3Service s3service.Service
}
