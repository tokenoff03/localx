package auth

import "errors"

func (storage *AuthTravelerService) UpdateTokens(email, accessToken, refreshToken string) error {
	if _, exists := storage.tokens.accessTokens[email]; !exists {
		return errors.New("tokens not found")
	}
	storage.tokens.accessTokens[email] = accessToken
	storage.tokens.refreshTokens[email] = refreshToken
	return nil
}

func (storage *AuthTravelerService) StoreTokens(email, accessToken, refreshToken string) {
	storage.tokens.accessTokens[email] = accessToken
	storage.tokens.refreshTokens[email] = refreshToken
}

// Получение токенов из хранилища
func (storage *AuthTravelerService) GetTokens(email string) (string, string, error) {
	accessToken, accessOk := storage.tokens.accessTokens[email]
	refreshToken, refreshOk := storage.tokens.refreshTokens[email]

	if !accessOk || !refreshOk {
		return "", "", errors.New("tokens not found")
	}
	return accessToken, refreshToken, nil
}
