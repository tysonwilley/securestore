package models

type Submissions []Submission

func GetSubmissions(collectionId string) (*Submissions, error) {
	db := connectToDB()
	defer db.Close()

	submissions := make(Submissions, 0)

	rows, err := db.Query("SELECT * FROM submissions WHERE collectionId =?", collectionId)

	if err != nil {
		return &submissions, err
	}

	for rows.Next() {
		var submission Submission

		rows.Scan(
			&submission.ID,
			&submission.Title,
			&submission.CollectionID,
			&submission.Status,
			&submission.DateCreated,
			&submission.Data,
		)

		submissions = append(submissions, submission)
	}

	return &submissions, nil
}
