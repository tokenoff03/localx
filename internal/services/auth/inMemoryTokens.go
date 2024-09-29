package auth

import "errors"

func (storage *AuthTravelerService) UpdateTokens(phoneNumber, accessToken, refreshToken string) error {
	if _, exists := storage.tokens.accessTokens[phoneNumber]; !exists {
		return errors.New("tokens not found")
	}
	storage.tokens.accessTokens[phoneNumber] = accessToken
	storage.tokens.refreshTokens[phoneNumber] = refreshToken
	return nil
}

func (storage *AuthTravelerService) StoreTokens(phoneNumber, accessToken, refreshToken string) {
	storage.tokens.accessTokens[phoneNumber] = accessToken
	storage.tokens.refreshTokens[phoneNumber] = refreshToken
}

// Получение токенов из хранилища
func (storage *AuthTravelerService) GetTokens(phoneNumber string) (string, string, error) {
	accessToken, accessOk := storage.tokens.accessTokens[phoneNumber]
	refreshToken, refreshOk := storage.tokens.refreshTokens[phoneNumber]

	if !accessOk || !refreshOk {
		return "", "", errors.New("tokens not found")
	}
	return accessToken, refreshToken, nil
}
