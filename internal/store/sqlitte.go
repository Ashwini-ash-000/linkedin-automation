package store

import (
	"database/sql"
	"time"

	_ "modernc.org/sqlite"
)

type Store struct {
	db *sql.DB
}

func New(path string) (*Store, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	s := &Store{db: db}
	if err := s.init(); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Store) init() error {
	_, err := s.db.Exec(`
	CREATE TABLE IF NOT EXISTS connections (
		profile_id TEXT PRIMARY KEY,
		sent_at DATETIME
	);

	CREATE TABLE IF NOT EXISTS messages (
		profile_id TEXT PRIMARY KEY,
		sent_at DATETIME
	);
	`)
	return err
}

func (s *Store) HasSent(profileID string) (bool, error) {
	row := s.db.QueryRow("SELECT 1 FROM connections WHERE profile_id = ?", profileID)
	var x int
	err := row.Scan(&x)
	if err == sql.ErrNoRows {
		return false, nil
	}
	return err == nil, err
}

func (s *Store) MarkSent(profileID string) error {
	_, err := s.db.Exec(
		"INSERT OR IGNORE INTO connections VALUES (?, ?)",
		profileID,
		time.Now(),
	)
	return err
}

func (s *Store) CountSentToday() (int, error) {
	row := s.db.QueryRow(
		"SELECT COUNT(*) FROM connections WHERE date(sent_at) = date('now')",
	)
	var c int
	return c, row.Scan(&c)
}

func (s *Store) HasMessaged(profileID string) (bool, error) {
	row := s.db.QueryRow("SELECT 1 FROM messages WHERE profile_id = ?", profileID)
	var x int
	err := row.Scan(&x)
	if err == sql.ErrNoRows {
		return false, nil
	}
	return err == nil, err
}

func (s *Store) MarkMessaged(profileID string) error {
	_, err := s.db.Exec(
		"INSERT OR IGNORE INTO messages VALUES (?, ?)",
		profileID,
		time.Now(),
	)
	return err
}
