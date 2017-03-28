package models

import (
	"errors"
	"fmt"
)

type Collection struct{
	ID            string `json:"id"`
	Title         string `json:"title"`
	OwnerID       string `json:"ownerId"`
	Recipients    string `json:"recipients"`
	DateCreated   int64  `json:"dateCreated"`
}

func GetCollection(id string) (*Collection, error) {
	db := connectToDB()
	defer db.Close()

	var collection Collection

	row := db.QueryRow("SELECT * FROM collections WHERE id=?", id)

	err := row.Scan(
		&collection.ID,
		&collection.Title,
		&collection.OwnerID,
		&collection.DateCreated,
		&collection.Recipients,
	)

	if err != nil {
		return &collection, err
	}

	return &collection, nil
}

func InsertCollection(title, ownerId, recipients string) (*Collection, error) {
	var collection Collection

	id          := newUUID()
	dateCreated := getTimestamp()

	db := connectToDB()
	defer db.Close()

	stmt, _ := db.Prepare("	INSERT INTO collections (`id`, `title`, `ownerId`, `recipients`, `dateCreated`) VALUES (?, ?, ?, ?, ?)")
	_, err  := stmt.Exec(id, title, ownerId, recipients, dateCreated)

	if err != nil {
		return &collection, err
	}

	collection.ID           = id
	collection.Title        = title
	collection.OwnerID      = ownerId
	collection.DateCreated  = dateCreated
	collection.Recipients   = recipients

	return &collection, nil
}

func DeleteCollection(collectionId string) (error) {
	db := connectToDB(); defer db.Close()

	stmt, _ := db.Prepare("DELETE FROM collections WHERE id = ?")
	result, err  := stmt.Exec(collectionId)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf("%s not found. No records removed.", collectionId))
	}

	return nil
}