package repository

import (
	"context"
	"database/sql"
	"personcrud/internal/models"

	"github.com/uptrace/bun"
)

type PersonRepository struct {
	db *bun.DB
}

func NewPersonRepository(db *bun.DB) *PersonRepository {
	return &PersonRepository{db: db}
}

func (r *PersonRepository) GetAll() ([]models.Person, error) {
	var persons []models.Person

	err := r.db.NewSelect().
		Model(&persons).
		Order("id ASC").
		Scan(context.Background())

	return persons, err
}

func (r *PersonRepository) GetByID(id int) (*models.Person, error) {
	var person models.Person

	err := r.db.NewSelect().
		Model(&person).
		Where("id = ?", id).
		Scan(context.Background())

	if err != nil {
		// Обрабатываем случай, когда запись не найдена
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &person, nil
}

func (r *PersonRepository) Create(person *models.CreatePersonRequest) (int, error) {
	newPerson := &models.Person{
		Email:     person.Email,
		Phone:     person.Phone,
		FirstName: person.FirstName,
		LastName:  person.LastName,
	}

	_, err := r.db.NewInsert().
		Model(newPerson).
		Exec(context.Background())

	if err != nil {
		return 0, err
	}

	return newPerson.ID, nil
}

func (r *PersonRepository) Update(id int, person *models.CreatePersonRequest) error {
	result, err := r.db.NewUpdate().
		Model(&models.Person{
			Email:     person.Email,
			Phone:     person.Phone,
			FirstName: person.FirstName,
			LastName:  person.LastName,
		}).
		Where("id = ?", id).
		Exec(context.Background())

	if err != nil {
		return err
	}

	// Проверяем, была ли обновлена хотя бы одна запись
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *PersonRepository) Delete(id int) error {
	result, err := r.db.NewDelete().
		Model((*models.Person)(nil)).
		Where("id = ?", id).
		Exec(context.Background())

	if err != nil {
		return err
	}

	// Проверяем, была ли удалена хотя бы одна запись
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
