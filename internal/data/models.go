package data

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrRecordNotFound = errors.New("record not found")

type ContactModel struct {
	DB *pgxpool.Pool
}

type Models struct {
	DB       *pgxpool.Pool
	Contacts ContactModel
}

func NewModels(db *pgxpool.Pool) Models {
	return Models{
		DB:       db,
		Contacts: ContactModel{DB: db},
	}
}

type Contact struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"` // Changed from int64 to string
	Email string `json:"email"`
}

func (m ContactModel) Insert(contact *Contact) error {
	query := `
		INSERT INTO contacts (id, name, phone, email)
		VALUES ($1, $2, $3, $4)`

	_, err := m.DB.Exec(context.Background(), query, contact.ID, contact.Name, contact.Phone, contact.Email)
	return err
}

func (m ContactModel) GetByID(id string) (*Contact, error) {
	query := `
		SELECT id, name, phone, email
		FROM contacts
		WHERE id = $1`

	var contact Contact
	err := m.DB.QueryRow(context.Background(), query, id).Scan(&contact.ID, &contact.Name, &contact.Phone, &contact.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return &contact, nil
}

func (m ContactModel) GetAll() ([]*Contact, error) {
	query := `
		SELECT id, name, phone, email
		FROM contacts
		ORDER BY name`

	rows, err := m.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []*Contact
	for rows.Next() {
		var contact Contact
		err := rows.Scan(&contact.ID, &contact.Name, &contact.Phone, &contact.Email)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, &contact)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return contacts, nil
}
