package services

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"localx/internal/repository"
	"mime/multipart"
	"net/http"
)

var ErrAvatarNotFound = errors.New("avatar not found")

type AvatarService struct {
	userRepo   *repository.TravelerRepository
	storageURL string
	bucket     string
	authToken  string
}

// Изменение конструктора AvatarService, чтобы принимать userRepo
func NewAvatarService(userRepo *repository.TravelerRepository, storageURL, bucket, authToken string) *AvatarService {
	return &AvatarService{
		userRepo:   userRepo,
		storageURL: storageURL,
		bucket:     bucket,
		authToken:  authToken,
	}
}

func (s *AvatarService) UploadAvatar(userID int, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	// Чтение файла
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	fileKey := fmt.Sprintf("avatars/%d/avatar.png", userID)
	uploadURL := fmt.Sprintf("%s/object/%s/%s", s.storageURL, s.bucket, fileKey)

	// Создание HTTP-запроса на загрузку
	req, err := http.NewRequest("POST", uploadURL, bytes.NewReader(fileData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", fileHeader.Header.Get("Content-Type"))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.authToken))

	// Отправка запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to upload avatar: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("upload failed with status: %s, response: %s", resp.Status, body)
	}

	avatarURL := fmt.Sprintf("%s/storage/v1/object/public/%s/%s", s.storageURL, s.bucket, fileKey)

	err = s.userRepo.SetProfilePicture(context.Background(), int64(userID), avatarURL)
	if err != nil {
		return "", fmt.Errorf("failed to update profile picture in database: %w", err)
	}

	// Возвращаем URL аватара
	return avatarURL, nil
}

func (s *AvatarService) DeleteAvatar(userID int) error {
	// Generate the file key and URL for the avatar in storage
	fileKey := fmt.Sprintf("avatars/%d/avatar.png", userID)
	url := fmt.Sprintf("%s/object/%s/%s", s.storageURL, s.bucket, fileKey)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create delete request: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.authToken))

	// Send the DELETE request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to delete avatar: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body for additional error info
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("delete failed with status: %s, response: %s", resp.Status, body)
	}

	// Clear the profile_pic field in the database
	if err := s.userRepo.ClearProfilePicture(context.Background(), int64(userID)); err != nil {
		return fmt.Errorf("failed to clear profile picture in database: %w", err)
	}

	return nil
}

func (s *AvatarService) GetAvatarURL(userID int) (string, error) {
	// Retrieve traveler data from the database using the user ID
	traveler, err := s.userRepo.GetTravelerByID(context.Background(), int64(userID))
	if err != nil {
		return "", fmt.Errorf("failed to retrieve avatar from database: %w", err)
	}

	// Check if the avatar URL is present in the profile_pic field
	if traveler == nil || traveler.ProfilePictureURL == "" {
		return "", ErrAvatarNotFound
	}

	// Return the avatar URL from the database
	return traveler.ProfilePictureURL, nil
}
