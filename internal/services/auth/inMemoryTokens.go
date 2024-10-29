package auth

import (
	"errors"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func (s *AuthTravelerService) UpdateTokens(id int, refreshToken string) (string, string, error) {
	traveler, err := s.GetTravelerById(id)
	if err != nil {
		return "", "", err
	}

	if _, exists := s.tokens.refreshTokens[traveler.Email]; !exists {
		return "", "", errors.New("tokens not found")
	}

	accessToken, err := s.GenerateToken(id)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = s.GenerateRefreshToken(id)
	if err != nil {
		return "", "", err
	}

	s.tokens.accessTokens[traveler.Email] = accessToken
	s.tokens.refreshTokens[traveler.Email] = refreshToken
	return s.tokens.accessTokens[traveler.Email], s.tokens.refreshTokens[traveler.Email], nil
}

func (s *AuthTravelerService) StoreTokens(email, accessToken, refreshToken string) {
	s.tokens.accessTokens[email] = accessToken
	s.tokens.refreshTokens[email] = refreshToken
}

// Получение токенов из хранилища
func (s *AuthTravelerService) GetTokens(email string) (string, string, error) {
	accessToken, accessOk := s.tokens.accessTokens[email]
	refreshToken, refreshOk := s.tokens.refreshTokens[email]

	if !accessOk || !refreshOk {
		return "", "", errors.New("tokens not found")
	}
	return accessToken, refreshToken, nil
}

func (s *AuthTravelerService) GenerateToken(id int) (string, error) {
	traveler, err := s.GetTravelerById(id)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		Subject:   strconv.Itoa(traveler.ID),
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthTravelerService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims["sub"].(string), nil
}

func (t *AuthTravelerService) GenerateRefreshToken(id int) (string, error) {
	traveler, err := t.GetTravelerById(id)
	if err != nil {
		return "", err
	}

	// Создаем refresh токен с длительным сроком действия (например, 7 дней)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(refreshTokenTTL).Unix(),
		Subject:   strconv.Itoa(traveler.ID),
	})
	return refreshToken.SignedString([]byte(signingKey))
}
