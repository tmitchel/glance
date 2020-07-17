package glance

import (
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/lib/pq" // postgres drivers
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// statement builder using postgres style
var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

// Database provides an interface with methods any underlying database
// must provide to be used.
type Database interface {
	// gets
	GetCard(string) (Card, error)
	GetCards() ([]Card, error)
	GetClaimedCards() ([]Card, error)
	GetUnclaimedCards() ([]Card, error)
	GetUser(string) (User, error)
	GetUsers() ([]User, error)
	GetUserCurrentCard(string) (Card, error)

	// creates
	CreateCard(Card) error
	CreateUser(User) error
	ClaimCard(string, string) (Card, error)
	UnclaimCard(string, string) (Card, error)
	UpdateStatus(string, string, int) (Card, error)

	Close()
}

// database wraps the postgres database to provide methods.
type database struct {
	*sql.DB
}

// OpenDatabase connects to the postgres database
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
	err := psql.Select("id", "title", "content", "status", "creator", "claimed", "created_at").
		From("cards").Where(sq.Eq{"id": id}).RunWith(d.DB).QueryRow().
		Scan(&card.ID, &card.Title, &card.Content, &card.Status, &card.Creator, &card.Claimed, &card.CreatedAt)
	return card, err
}

func (d *database) GetCards() ([]Card, error) {
	var cards []Card
	rows, err := psql.Select("id", "title", "content", "status", "creator", "claimed", "created_at").
		From("cards").RunWith(d.DB).Query()
	if err != nil {
		return cards, nil
	}

	for rows.Next() {
		var card Card
		err := rows.Scan(&card.ID, &card.Title, &card.Content, &card.Status, &card.Creator, &card.Claimed, &card.CreatedAt)
		if err != nil {
			continue
		}
		cards = append(cards, card)
	}

	return cards, nil
}

func (d *database) GetClaimedCards() ([]Card, error) {
	var cards []Card
	rows, err := psql.Select("id", "title", "content", "status", "creator", "claimed", "created_at").
		From("cards").Where(sq.Eq{"claimed": true}).RunWith(d.DB).Query()
	if err != nil {
		return cards, nil
	}

	for rows.Next() {
		var card Card
		err := rows.Scan(&card.ID, &card.Title, &card.Content, &card.Status, &card.Creator, &card.Claimed, &card.CreatedAt)
		if err != nil {
			continue
		}
		cards = append(cards, card)
	}

	return cards, nil
}

func (d *database) GetUnclaimedCards() ([]Card, error) {
	var cards []Card
	rows, err := psql.Select("id", "title", "content", "status", "creator", "claimed", "created_at").
		From("cards").Where(sq.Eq{"claimed": false}).RunWith(d.DB).Query()
	if err != nil {
		return cards, nil
	}

	for rows.Next() {
		var card Card
		err := rows.Scan(&card.ID, &card.Title, &card.Content, &card.Status, &card.Creator, &card.Claimed, &card.CreatedAt)
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

func (d *database) GetUsers() ([]User, error) {
	var users []User
	rows, err := psql.Select("id", "name", "email", "password").
		From("users").RunWith(d.DB).Query()
	if err != nil {
		return users, nil
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			continue
		}
		users = append(users, user)
	}

	return users, err
}

func (d *database) GetUserCurrentCard(id string) (Card, error) {
	var card Card
	err := psql.Select("id", "title", "content", "status", "creator", "claimed", "created_at").
		From("cards").Join("cards_users cu ON (cu.card_id = id)").Where(sq.Eq{"cu.user_id": id}).
		RunWith(d.DB).QueryRow().
		Scan(&card.ID, &card.Title, &card.Content, &card.Status, &card.Creator, &card.Claimed, &card.CreatedAt)
	return card, err
}

func (d *database) CreateCard(c Card) error {
	_, err := psql.Insert("cards").
		Columns("id", "title", "content", "status", "creator", "claimed", "created_at").
		Values(c.ID, c.Title, c.Content, c.Status, c.Creator, c.Claimed, c.CreatedAt).
		RunWith(d.DB).Exec()
	return err
}

func (d *database) CreateUser(u User) error {
	_, err := psql.Insert("users").
		Columns("id", "name", "email", "password").
		Values(u.ID, u.Name, u.Email, u.Password).
		RunWith(d.DB).Exec()
	if err != nil {
		return err
	}

	_, err = psql.Insert("cards_users").
		Columns("user_id", "card_id").
		Values(u.ID, nil).
		RunWith(d.DB).Exec()

	return err
}

func (d *database) ClaimCard(uid, cid string) (Card, error) {
	var c Card
	_, err := psql.Update("cards_users").
		Set("card_id", cid).
		Where(sq.Eq{"user_id": uid}).
		RunWith(d.DB).Exec()
	if err != nil {
		return c, err
	}

	_, err = psql.Update("cards").Set("claimed", true).Where(sq.Eq{"id": cid}).RunWith(d.DB).Exec()
	if err != nil {
		return c, err
	}

	err = psql.Select("id", "title", "content", "status", "creator", "claimed", "created_at").
		From("cards").Where(sq.Eq{"id": cid}).RunWith(d.DB).QueryRow().
		Scan(&c.ID, &c.Title, &c.Content, &c.Status, &c.Creator, &c.Claimed, &c.CreatedAt)
	return c, err
}

func (d *database) UnclaimCard(uid, cid string) (Card, error) {
	var c Card
	_, err := psql.Update("cards_users").
		Set("card_id", nil).
		Where(sq.Eq{"user_id": uid}).
		RunWith(d.DB).Exec()
	if err != nil {
		return c, err
	}

	_, err = psql.Update("cards").Set("claimed", false).Where(sq.Eq{"id": cid}).RunWith(d.DB).Exec()
	if err != nil {
		return c, err
	}

	err = psql.Select("id", "title", "content", "status", "creator", "claimed", "created_at").
		From("cards").Where(sq.Eq{"id": cid}).RunWith(d.DB).QueryRow().
		Scan(&c.ID, &c.Title, &c.Content, &c.Status, &c.Creator, &c.Claimed, &c.CreatedAt)
	return c, err
}

func (d *database) UpdateStatus(cid, uid string, status int) (Card, error) {
	var c Card
	_, err := psql.Update("cards").Set("status", status).Where(sq.Eq{"id": cid}).RunWith(d.DB).Exec()
	if err != nil {
		return c, err
	}

	err = psql.Select("id", "title", "content", "status", "creator", "claimed", "created_at").
		From("cards").Where(sq.Eq{"id": cid}).RunWith(d.DB).QueryRow().
		Scan(&c.ID, &c.Title, &c.Content, &c.Status, &c.Creator, &c.Claimed, &c.CreatedAt)
	return c, err
}

// Close closes the database.
func (d *database) Close() {
	d.DB.Close()
}
