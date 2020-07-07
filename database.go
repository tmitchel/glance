package glance

import (
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// statement builder using postgres style
var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type Database interface {
	// gets
	GetCard(string) (Card, error)
	GetCards() ([]Card, error)
	GetUser(string) (User, error)

	// creates
	CreateCard(Card) error
	CreateUser(User) error

	Close()
}

// database wraps the postgres database to provide methods.
type database struct {
	*sql.DB
}

// New connects to the postgres database
// and returns that connection.
func OpenDatabase(psqlInfo string) (Database, error) {
	for {
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			return nil, errors.Wrap(err, "Error opening database")
		}

		// make sure we have a good connection
		err = db.Ping()
		if err != nil {
			time.Sleep(time.Second)
			logrus.Errorf("Error pinging database %v", err)
		} else {
			logrus.Info("Connected to database.")
			return &database{db}, nil
		}
	}
}

func (d *database) GetCard(id string) (Card, error) {
	var card Card
	err := psql.Select("id", "title", "content", "status", "creator", "volunter", "created_at").
		From("cards").Where(sq.Eq{"id": id}).RunWith(d.DB).QueryRow().
		Scan(&card.ID, &card.Title, &card.Content, &card.Status, &card.Creator, &card.Volunteer, &card.CreatedAt)
	return card, err
}

func (d *database) GetCards() ([]Card, error) {
	var cards []Card
	rows, err := psql.Select("id", "title", "content", "status", "creator", "volunteer", "created_at").
		From("cards").RunWith(d.DB).Query()
	if err != nil {
		return cards, nil
	}

	for rows.Next() {
		var card Card
		err := rows.Scan(&card.ID, &card.Title, &card.Content, &card.Status, &card.Creator, &card.Volunteer, &card.CreatedAt)
		if err != nil {
			continue
		}
		cards = append(cards, card)
	}

	return cards, nil
}

func (d *database) GetUser(email string) (User, error) {
	var user User
	err := psql.Select("id", "name", "email", "password").
		From("users").Where(sq.Eq{"email": email}).RunWith(d.DB).QueryRow().
		Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	return user, err
}

func (d *database) CreateCard(c Card) error {
	_, err := psql.Insert("cards").
		Columns("id", "title", "content", "status", "creator", "volunteer", "created_at").
		Values(c.ID, c.Title, c.Content, c.Status, c.Creator, c.Volunteer, c.CreatedAt).
		RunWith(d.DB).Exec()
	return err
}

func (d *database) CreateUser(u User) error {
	_, err := psql.Insert("users").
		Columns("id", "name", "email", "password").
		Values(u.ID, u.Name, u.Email, u.Password).
		RunWith(d.DB).Exec()
	return err
}

// Close closes the database.
func (d *database) Close() {
	d.DB.Close()
}
