package services

import (
	"errors"
	"personcrud/internal/models"
	"personcrud/internal/repository"
)

// Кастомные ошибки для валидации
var (
	ErrInvalidEmail     = errors.New("email is required and must be valid")
	ErrInvalidPhone     = errors.New("phone is required")
	ErrInvalidFirstName = errors.New("first name is required")
	ErrInvalidLastName  = errors.New("last name is required")
	ErrPersonNotFound   = errors.New("person not found")
)

type PersonService struct {
	repo *repository.PersonRepository
}

func NewPersonService(repo *repository.PersonRepository) *PersonService {
	return &PersonService{repo: repo}
}

func (s *PersonService) GetAll() ([]models.Person, error) {
	return s.repo.GetAll()
}

func (s *PersonService) GetByID(id int) (*models.Person, error) {
	if id <= 0 {
		return nil, ErrPersonNotFound
	}
	return s.repo.GetByID(id)
}

func (s *PersonService) Create(person *models.CreatePersonRequest) (int, error) {
	// Валидация с возвратом ошибок
	if err := validatePersonRequest(person); err != nil {
		return 0, err
	}

	return s.repo.Create(person)
}

func (s *PersonService) Update(id int, person *models.CreatePersonRequest) error {
	if id <= 0 {
		return ErrPersonNotFound
	}

	// Валидация с возвратом ошибок
	if err := validatePersonRequest(person); err != nil {
		return err
	}

	return s.repo.Update(id, person)
}

func (s *PersonService) Delete(id int) error {
	if id <= 0 {
		return ErrPersonNotFound
	}
	return s.repo.Delete(id)
}

// Вспомогательная функция для валидации
func validatePersonRequest(person *models.CreatePersonRequest) error {
	if person.Email == "" {
		return ErrInvalidEmail
	}

	// Простая проверка формата email
	if !isValidEmail(person.Email) {
		return ErrInvalidEmail
	}

	if person.Phone == "" {
		return ErrInvalidPhone
	}

	if person.FirstName == "" {
		return ErrInvalidFirstName
	}

	if person.LastName == "" {
		return ErrInvalidLastName
	}

	return nil
}

// Простая проверка email
func isValidEmail(email string) bool {
	return len(email) >= 3 && contains(email, "@") && contains(email, ".")
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && (s[0:len(substr)] == substr || contains(s[1:], substr)))
}
