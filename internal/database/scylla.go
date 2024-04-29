package database

import (
	"context"
	"fmt"
	"time"

	"github.com/gocql/gocql"
)

type ScyllaStore struct {
	Session *gocql.Session
}

func NewScyllaStore(addr string, ks string) (Store, error) {
	cluster := gocql.NewCluster(addr)
	cluster.Keyspace = ks

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	return &ScyllaStore{Session: session}, nil
}

func (s *ScyllaStore) InitTables() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	query := `
  CREATE TABLE IF NOT EXISTS go_api.users (
    email VARCHAR,
    name VARCHAR,
    hashed_pass text,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY(email)
  )
  `
	err := s.Session.Query(query).WithContext(ctx).Exec()
	return err
}

func (s *ScyllaStore) CreateUser(email string, name string, hashed_pass string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var exists int
	checkUserQuery := `
  SELECT count(1) FROM go_api.users 
  WHERE email = ?
  `
	err := s.Session.Query(checkUserQuery, email).WithContext(ctx).Scan(&exists)
	if err != nil {
		return err
	}
	if exists > 0 {
		return fmt.Errorf("user already exists")
	}

	query := `
  INSERT INTO go_api.users (
    email,
    name,
    hashed_pass,
    created_at,
    updated_at
  ) VALUES (?, ?, ?, ?, ?)
  `
	err = s.Session.Query(
		query,
		email,
		name,
		hashed_pass,
		time.Now(),
		time.Now(),
	).WithContext(ctx).Exec()
	return err
}

func (s *ScyllaStore) GetUser(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	query := `SELECT * FROM go_api.users WHERE email = ?`

	user := new(User)
	err := s.Session.Query(query, email).WithContext(ctx).Scan(
		&user.Email,
		&user.CreatedAt,
		&user.HashedPass,
		&user.Name,
		&user.UpdatedAt,
	)
	return user, err
}

func (s *ScyllaStore) Close() {
	s.Session.Close()
}
