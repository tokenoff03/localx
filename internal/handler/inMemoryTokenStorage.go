package handler

import "errors"

type InMemoryTokenStorage struct {
	accessTokens  map[string]string // access token по phoneNumber
	refreshTokens map[string]string // refresh token по phoneNumber
}

func NewInMemoryTokenStorage() *InMemoryTokenStorage {
	return &InMemoryTokenStorage{
		accessTokens:  make(map[string]string),
		refreshTokens: make(map[string]string),
	}
}

func (storage *InMemoryTokenStorage) UpdateTokens(phoneNumber, accessToken, refreshToken string) error {
	if _, exists := storage.accessTokens[phoneNumber]; !exists {
		return errors.New("tokens not found")
	}
	storage.accessTokens[phoneNumber] = accessToken
	storage.refreshTokens[phoneNumber] = refreshToken
	return nil
}

func (storage *InMemoryTokenStorage) StoreTokens(phoneNumber, accessToken, refreshToken string) {
	storage.accessTokens[phoneNumber] = accessToken
	storage.refreshTokens[phoneNumber] = refreshToken
}

// Получение токенов из хранилища
func (storage *InMemoryTokenStorage) GetTokens(phoneNumber string) (string, string, error) {
	accessToken, accessOk := storage.accessTokens[phoneNumber]
	refreshToken, refreshOk := storage.refreshTokens[phoneNumber]

	if !accessOk || !refreshOk {
		return "", "", errors.New("tokens not found")
	}
	return accessToken, refreshToken, nil
}
