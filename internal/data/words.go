package data

import (
	"context"
	"database/sql"
	"time"
)

type Word struct {
	ID         string `json:"id"`
	Word       string `json:"word"`
	WordType   string `json:"word_type"`
	Definition string `json:"definition"`
}

type WordModel struct {
	DB *sql.DB
}

func (m WordModel) Insert(word *Word) error {
	query := `
	insert into words (word, word_type, definition)
	values ($1, $2, $3)
	returning id`

	args := []interface{}{word.Word, word.WordType, word.Definition}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return m.DB.QueryRowContext(ctx, query, args...).Scan(&word.ID)
}

func (m WordModel) GetByWord(word string) ([]*Word, error) {
	query := `
		SELECT id, word, word_type, definition
		FROM words
		WHERE LOWER(word) = LOWER($1)`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, word)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var words []*Word

	for rows.Next() {
		var word Word

		err := rows.Scan(
			&word.ID,
			&word.Word,
			&word.WordType,
			&word.Definition,
		)

		if err != nil {
			return nil, err
		}

		words = append(words, &word)
	}

	return words, nil
}
