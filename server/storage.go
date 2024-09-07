package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	//To avoid conflicts with dockerized postgres, make sure to create it with:
	//docker run --name postgres -e POSTGRES_PASSWORD=gobank -p 5433:5432 -d postgres
	//Map port 5432 of the container to 5433 of the host
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable port=5433"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{db: db}, nil
}

func (s *PostgresStore) Init() error {
	return s.createWordCombiTable()
}

func (s *PostgresStore) createWordCountTable() error {
	query := `CREATE TABLE IF NOT EXISTS word_count (
			word VARCHAR(255) PRIMARY KEY,
			count INT
			)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) createWordCombiTable() error {
	query := `CREATE TABLE IF NOT EXISTS word_combi (
			word1 VARCHAR(255),
			word2 VARCHAR(255),
			res VARCHAR(255)
			)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateWordCount(wc *WordCount) error {
	query := `INSERT INTO word_count (word, count) VALUES ($1, $2)`
	_, err := s.db.Query(query, wc.Word, wc.Count)

	if err != nil {
		return err
	}

	fmt.Println("CREATED")
	return nil
}

func (s *PostgresStore) GetWordCount(word string) (*WordCount, bool, error) {
	query := `SELECT * FROM word_count WHERE word = $1`
	row := s.db.QueryRow(query, word)

	wc := &WordCount{}
	if err := row.Scan(&wc.Word, &wc.Count); err != nil {
		if err == sql.ErrNoRows {
			// Word not found, return nil and false without an error
			return nil, false, nil
		}
		// An actual error occurred, return the error
		return nil, false, err
	}

	// Word found, return the WordCount and true
	return wc, true, nil
}

func (s *PostgresStore) CreateWordCombi(wc *WordCombination) error {
	query := `INSERT INTO word_combi (word1, word2, res) VALUES ($1, $2, $3)`
	_, err := s.db.Query(query, wc.Word1, wc.Word2, wc.Result)

	if err != nil {
		return err
	}

	fmt.Println("CREATED")
	return nil
}
