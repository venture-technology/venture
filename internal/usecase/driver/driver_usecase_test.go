package driver

import (
	"context"
	"fmt"
	"testing"

	"github.com/venture-technology/venture/mocks"
)

func TestDriverUseCase_GetGallery(t *testing.T) {
	mock := mocks.NewIAwsRepository(t)
	cnh := "55276739722"
	mock.On("ListImagesAtS3", context.Background(), fmt.Sprintf("%s/gallery", cnh)).Return([]string{"image1", "image2"}, nil)
	service := NewDriverUseCase(nil, mock, nil)

	links, err := service.GetGallery(context.Background(), &cnh)

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if len(links) != 2 {
		t.Errorf("Expected 2 links, got %d", len(links))
	}
}
