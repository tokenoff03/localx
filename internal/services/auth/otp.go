package auth

import (
	"fmt"
	"math/rand"
	"time"
)

func (s *AuthTravelerService) GenerateAndSaveOTP(email string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Генерация случайного кода
	code := fmt.Sprintf("%06d", rand.Intn(1000000))

	// Сохранение кода с указанием времени истечения (например, 5 минут)
	s.otpStore[email] = otpData{
		code:      code,
		expiresAt: time.Now().Add(5 * time.Minute),
	}
	fmt.Println(s.otpStore)
	return code, nil
}

func (s *AuthTravelerService) ValidateOTP(email, inputCode string) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Получаем данные для указанного почты
	data, exists := s.otpStore[email]
	if !exists {
		return false, fmt.Errorf("OTP not found")
	}

	// Проверяем, не истек ли код
	if time.Now().After(data.expiresAt) {
		delete(s.otpStore, email) // Удаляем истекший код
		return false, fmt.Errorf("OTP expired")
	}

	// Сравниваем введенный код с сохраненным
	if data.code != inputCode {
		return false, fmt.Errorf("Invalid OTP")
	}

	// Удаляем код после успешной проверки
	delete(s.otpStore, email)

	return true, nil
}

// Очистка истекших кодов (опционально)
func (s *AuthTravelerService) CleanExpiredOTPs() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for email, data := range s.otpStore {
		if time.Now().After(data.expiresAt) {
			delete(s.otpStore, email)
		}
	}
}

//Можно запустить отдельную гоу рутину которая будет очишать мапу
// go func() {
// 	for {
// 		time.Sleep(1 * time.Minute)
// 		authService.CleanExpiredOTPs()
// 	}
// }()
