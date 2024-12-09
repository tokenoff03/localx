package docs

import "github.com/swaggo/swag"

// Шаблон для Swagger-документации
const docTemplate = `{
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "version": "{{.Version}}",
        "contact": {}
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "schemes": {{ marshal .Schemes }},
    "paths": {{ marshal .Paths }}
}`

// SwaggerInfo содержит информацию о спецификации API
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",                     // Версия API
	Host:             "localhost:8080",            // Хост API
	BasePath:         "/",                         // Базовый путь
	Schemes:          []string{"http"},            // Схема (http/https)
	Title:            "My API Documentation",      // Заголовок документации
	Description:      "API Documentation Example", // Описание API
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
