package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lfdelima3/Backend-Go-Bet/src/util"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	util.RegisterCustomValidations(validate)
}

func ValidateRequest(model interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(model); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Erro na validação dos dados",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}

		if err := validate.Struct(model); err != nil {
			errors := make(map[string]string)
			for _, err := range err.(validator.ValidationErrors) {
				field := err.Field()
				tag := err.Tag()
				errors[field] = getValidationErrorMessage(field, tag)
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Erro na validação dos dados",
				"errors":  errors,
			})
			c.Abort()
			return
		}

		c.Set("validatedData", model)
		c.Next()
	}
}

func ValidateQuery(model interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindQuery(model); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Erro na validação dos parâmetros de consulta",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}

		if err := validate.Struct(model); err != nil {
			errors := make(map[string]string)
			for _, err := range err.(validator.ValidationErrors) {
				field := err.Field()
				tag := err.Tag()
				errors[field] = getValidationErrorMessage(field, tag)
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Erro na validação dos parâmetros de consulta",
				"errors":  errors,
			})
			c.Abort()
			return
		}

		c.Set("validatedQuery", model)
		c.Next()
	}
}

// Retorna mensagens de erro amigáveis para cada tipo de validação
func getValidationErrorMessage(field, tag string) string {
	switch tag {
	case "required":
		return "Campo obrigatório"
	case "email":
		return "Email inválido"
	case "valid_email":
		return "Email inválido"
	case "min":
		return "Valor muito pequeno"
	case "max":
		return "Valor muito grande"
	case "future_date":
		return "Data deve ser futura"
	case "past_date":
		return "Data deve ser passada"
	case "valid_score":
		return "Placar inválido"
	case "valid_odds":
		return "Probabilidade inválida"
	case "valid_amount":
		return "Valor monetário inválido"
	case "valid_team_name":
		return "Nome do time inválido"
	case "strong_password":
		return "A senha deve conter pelo menos 8 caracteres, incluindo letras maiúsculas, minúsculas, números e caracteres especiais"
	default:
		return "Campo inválido"
	}
}
