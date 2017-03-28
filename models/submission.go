package models

import (
	"database/sql"
	"errors"
	"strings"
	"time"
	"secureStore/encryption"
	"fmt"
)

type Submission struct{
	ID            string `json:"id"`
	CollectionID  string `json:"collectionId,omitempty"`
	Title         string `json:"title"`
	Status        string `json:"statusId,omitempty"`
	Data          string `json:"data,omitempty"`
	DateCreated   int64  `json:"dateCreated"`
}

type Submissions []Submission

func GetSubmission(id string) (*Submission, error) {
	db := connectToDB()
	defer db.Close()

	var submission Submission

	row := db.QueryRow("SELECT * FROM submissions WHERE id=?", id)

	err := row.Scan(
		&submission.ID,
		&submission.Title,
		&submission.CollectionID,
		&submission.DateCreated,
		&submission.Data,
	)

	if err != nil {
		return &submission, err
	}

	submission.Data = encryption.Decrypt(submission.Data, submission.ID)

	return &submission, nil
}

func InsertSubmission(requestBody []byte, collectionId string) (*Submission, error) {
	var submission Submission

	if ! hasCollectionId(collectionId) {
		return &submission, errors.New("Invalid collectionId.")
	}

	id          := newUUID()
	dateCreated := getTimestamp()
	title       := fmt.Sprintf("Submission: %s", time.Now().String())

	replacer := strings.NewReplacer("\n", "", "\t", "")
	data := string(requestBody)
	data = replacer.Replace(data)
	data = encryption.Encrypt(data, id)

	db := connectToDB()
	defer db.Close()

	stmt, _ := db.Prepare("	INSERT INTO submissions (`id`, `title`, `collectionId`, `dateCreated`, `data`) VALUES (?, ?, ?, ?, ?)")
	_, err  := stmt.Exec(id, title, collectionId, dateCreated, data)

	if err != nil {
		return &submission, err
	}

	submission.ID           = id
	submission.Title        = title
	submission.CollectionID = collectionId
	submission.DateCreated  = dateCreated
	submission.Data         = data

	return &submission, nil
}

func DeleteSubmission(submissionId string) (error) {
	db := connectToDB(); defer db.Close()

	stmt, _ := db.Prepare("DELETE FROM submissions WHERE id = ?")
	result, err  := stmt.Exec(submissionId)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf("%s not found. No records removed.", submissionId))
	}

	return nil
}

func GetSubmissions(collectionId string) (*Submissions, error) {
	db := connectToDB()
	defer db.Close()

	submissions := make(Submissions, 0)

	rows, err := db.Query("SELECT `id`, `title`, `dateCreated` FROM submissions WHERE collectionId =?", collectionId)

	if err != nil {
		return &submissions, err
	}

	for rows.Next() {
		var submission Submission

		rows.Scan(
			&submission.ID,
			&submission.Title,
			&submission.DateCreated,
		)

		submissions = append(submissions, submission)
	}

	return &submissions, nil
}

func hasCollectionId(collectionId string) bool {
	var output interface{}

	db := connectToDB(); defer db.Close()
	err := db.QueryRow("SELECT `id` FROM collections WHERE id=?", collectionId).Scan(&output)

	return sql.ErrNoRows != err
}
