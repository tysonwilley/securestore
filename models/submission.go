package models

type Submission struct{
	ID            string `json:"id"`
	CollectionID  string `json:"collectionId"`
	Title         string `json:"title"`
	Status        string `json:"statusId"`
	Data          string `json:"data"`
	DateCreated   int64  `json:"dateCreated"`
}

func GetSubmission(id string) (*Submission, error) {
	db := connectToDB()
	defer db.Close()

	var submission Submission

	row := db.QueryRow("SELECT * FROM submissions WHERE id=?", id)

	err := row.Scan(
		&submission.ID,
		&submission.Title,
		&submission.CollectionID,
		&submission.Status,
		&submission.DateCreated,
		&submission.Data,
	)

	if err != nil {
		return &submission, err
	}

	return &submission, nil
}