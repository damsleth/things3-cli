package db

import (
	"database/sql"
	"fmt"
	"net/url"
	"path/filepath"

	_ "modernc.org/sqlite"
)

// Store wraps a Things database connection.
type Store struct {
	conn *sql.DB
	path string
}

// Open opens a Things database at the provided path in read-only mode.
func Open(path string) (*Store, error) {
	if path == "" {
		return nil, fmt.Errorf("empty database path")
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("resolve database path: %w", err)
	}
	dsn := sqliteDSN(abs)
	conn, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}
	if err := conn.Ping(); err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("open database: %w", err)
	}
	return &Store{conn: conn, path: abs}, nil
}

// OpenDefault resolves the Things database path and opens it.
func OpenDefault(override string) (*Store, string, error) {
	path, err := ResolveDatabasePath(override)
	if err != nil {
		return nil, "", err
	}
	store, err := Open(path)
	if err != nil {
		return nil, path, err
	}
	return store, path, nil
}

// Close closes the underlying database connection.
func (s *Store) Close() error {
	if s == nil || s.conn == nil {
		return nil
	}
	return s.conn.Close()
}

// Path returns the resolved database path.
func (s *Store) Path() string {
	if s == nil {
		return ""
	}
	return s.path
}

func sqliteDSN(path string) string {
	u := url.URL{Scheme: "file", Path: path}
	q := u.Query()
	q.Set("mode", "ro")
	u.RawQuery = q.Encode()
	return u.String()
}
