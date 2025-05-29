package util

import (
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

// Registra validações personalizadas
func RegisterCustomValidations(v *validator.Validate) {
	// Validação de data futura
	v.RegisterValidation("future_date", validateFutureDate)

	// Validação de data passada
	v.RegisterValidation("past_date", validatePastDate)

	// Validação de placar válido
	v.RegisterValidation("valid_score", validateScore)

	// Validação de odds (probabilidades)
	v.RegisterValidation("valid_odds", validateOdds)

	// Validação de valor monetário
	v.RegisterValidation("valid_amount", validateAmount)

	// Validação de nome de time
	v.RegisterValidation("valid_team_name", validateTeamName)

	// Validação de email
	v.RegisterValidation("valid_email", validateEmail)

	// Validação de senha forte
	v.RegisterValidation("strong_password", validateStrongPassword)
}

// Valida se a data é futura
func validateFutureDate(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if !ok {
		return false
	}
	return date.After(time.Now())
}

// Valida se a data é passada
func validatePastDate(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if !ok {
		return false
	}
	return date.Before(time.Now())
}

// Valida se o placar é válido (números não negativos)
func validateScore(fl validator.FieldLevel) bool {
	score, ok := fl.Field().Interface().(int)
	if !ok {
		return false
	}
	return score >= 0
}

// Valida se as odds são válidas (maior que 1.0)
func validateOdds(fl validator.FieldLevel) bool {
	odds, ok := fl.Field().Interface().(float64)
	if !ok {
		return false
	}
	return odds > 1.0
}

// Valida se o valor monetário é válido (positivo)
func validateAmount(fl validator.FieldLevel) bool {
	amount, ok := fl.Field().Interface().(float64)
	if !ok {
		return false
	}
	return amount > 0
}

// Valida se o nome do time é válido (apenas letras, números e espaços)
func validateTeamName(fl validator.FieldLevel) bool {
	name, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9\s]+$`, name)
	return matched
}

// Valida se o email é válido
func validateEmail(fl validator.FieldLevel) bool {
	email, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, email)
	return matched
}

// Valida se a senha é forte (mínimo 8 caracteres, letra maiúscula, minúscula, número e caractere especial)
func validateStrongPassword(fl validator.FieldLevel) bool {
	password, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	matched, _ := regexp.MatchString(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`, password)
	return matched
}
