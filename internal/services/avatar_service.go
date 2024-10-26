package services

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type AvatarService struct {
	storageURL string
	bucket     string
	authToken  string
}

func NewAvatarService(storageURL, bucket, authToken string) *AvatarService {
	return &AvatarService{
		storageURL: storageURL,
		bucket:     bucket,
		authToken:  authToken,
	}
}

func (s *AvatarService) UploadAvatar(userID int, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	// Новый URL для Supabase загрузки
	fileKey := fmt.Sprintf("avatars/%d/%s", userID, fileHeader.Filename)
	url := fmt.Sprintf("%s/object/%s/%s", s.storageURL, s.bucket, fileKey)

	req, err := http.NewRequest("POST", url, bytes.NewReader(fileData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Устанавливаем необходимые заголовки, включая токен
	req.Header.Set("Content-Type", fileHeader.Header.Get("Content-Type"))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.authToken))

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
	return avatarURL, nil
}
