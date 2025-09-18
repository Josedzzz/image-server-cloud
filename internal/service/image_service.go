package service

import (
	"encoding/base64"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"image-server/internal/model"
)

type ImageService interface {
	LoadRandomImages(dir string, count int) ([]model.ImageData, error)
}

type imageService struct{}

func NewImageService() ImageService {
	// Initialize random seed properly
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return &imageService{}
}

func (s *imageService) LoadRandomImages(dir string, count int) ([]model.ImageData, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var imageFiles []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		ext := filepath.Ext(entry.Name())
		if ext == ".jpg" || ext == ".jpeg" || ext == ".png" {
			imageFiles = append(imageFiles, entry.Name())
		}
	}

	if len(imageFiles) == 0 {
		return []model.ImageData{}, nil
	}

	if count > len(imageFiles) {
		count = len(imageFiles)
	}

	selected := make([]model.ImageData, 0, count)
	selectedIndices := make(map[int]bool)

	for len(selected) < count && len(selected) < len(imageFiles) {
		idx := rand.Intn(len(imageFiles))
		if !selectedIndices[idx] { // ← FIXED: Changed to !selectedIndices[idx]
			selectedIndices[idx] = true

			imgPath := filepath.Join(dir, imageFiles[idx])
			imgData, err := os.ReadFile(imgPath)
			if err != nil {
				continue
			}

			mimeType := http.DetectContentType(imgData)

			base64Data := base64.StdEncoding.EncodeToString(imgData)
			dataURI := "data:" + mimeType + ";base64," + base64Data

			selected = append(selected, model.ImageData{
				Name: imageFiles[idx],
				Data: dataURI,
			})
		}
	}

	return selected, nil
}

