package data

import (
	"context"
	"database/sql"
	"time"
)

type Word struct {
	ID         int64 `json:"id"`
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
		select id, word, word_type, definition
		from words
		where lower(word) = lower($1)`

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

func (m WordModel) GetAllForUser(userId int64) ([]Word, error) {
	query := `
	SELECT words.id, words.word, words.word_type, word.definition
	FROM words
	INNER JOIN users_words 
	ON users_words.word_id = words.id
	INNER JOIN users 
	ON users_words.user_id = users.id
	WHERE users.id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var words []Word

	for rows.Next() {
		var word Word

		err := rows.Scan(&word.ID, &word.Word, &word.WordType, &word.Definition)
		if err != nil {
			return nil, err
		}

		words = append(words, word)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return words, nil

}

func (m WordModel) AddForUser(userID int64, wordId int64) error {
	query := `
	insert into users_words (user_id, word_id)
	values ($1, $2)`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := m.DB.ExecContext(ctx, query, userID, wordId)
	return err
}
