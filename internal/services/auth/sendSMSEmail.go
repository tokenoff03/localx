package auth

import (
	"github.com/MakMoinee/go-mith/pkg/email"
)

func (t *AuthTravelerService) SendEmail(to string, subject string, body string) error {
	emailService := email.NewEmailService(587, "smtp.gmail.com", "prince.123317@gmail.com", "ytpv btno vnxt tbqs")

	_, err := emailService.SendEmail(to, subject, body)
	if err != nil {
		return err
	}

	return nil
}

// const (
// 	emailFrom       = "prince.123317@gmail.com" // Укажите адрес вашей электронной почты
// 	credentialsFile = "../client_gmail.json"    // Путь к вашему JSON-файлу с учетными данными
// )

// var token *oauth2.Token

// func (t *AuthTravelerService) SendEmail(to string, subject string, body string) error {
// 	ctx := context.Background()

// 	b, err := os.ReadFile(credentialsFile)
// 	if err != nil {
// 		return fmt.Errorf("unable to read client secret file: %v", err)
// 	}

// 	config, err := google.ConfigFromJSON(b, gmail.GmailSendScope)
// 	if err != nil {
// 		return fmt.Errorf("unable to parse client secret file to config: %v", err)
// 	}

// 	client := getClient(ctx, config)
// 	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
// 	if err != nil {
// 		return fmt.Errorf("unable to create Gmail service: %v", err)
// 	}

// 	message := &gmail.Message{
// 		Raw: encodeWeb64String([]byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", emailFrom, to, subject, body))),
// 	}

// 	_, err = srv.Users.Messages.Send("me", message).Do()
// 	if err != nil {
// 		return fmt.Errorf("unable to send message: %v", err)
// 	}
// 	return nil
// }

// // encodeWeb64String кодирует строку в формат base64 для Gmail API
// func encodeWeb64String(b []byte) string {
// 	s := base64.StdEncoding.EncodeToString(b)
// 	s = strings.ReplaceAll(s, "+", "-")
// 	s = strings.ReplaceAll(s, "/", "_")
// 	return s
// }

// // getClient получает токен OAuth 2.0
// func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
// 	if token == nil {
// 		token = getNewToken(config) // Получите новый токен
// 	}
// 	return config.Client(ctx, token)
// }

// func getNewToken(config *oauth2.Config) *oauth2.Token {
// 	// Генерируем URL для авторизации
// 	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
// 	fmt.Printf("Перейдите по следующему URL для авторизации:\n%s\n", authURL)

// 	// Запрашиваем код авторизации
// 	var code string
// 	fmt.Print("Введите код: ")
// 	fmt.Scan(&code)

// 	// Получаем токен
// 	token, err := config.Exchange(context.Background(), code)
// 	if err != nil {
// 		fmt.Printf("Ошибка при обмене кода на токен: %v\n", err)
// 		return nil
// 	}

// 	return token
// }
