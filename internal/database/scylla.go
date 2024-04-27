package database

import (
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

func (s *ScyllaStore) Close() {
	s.Session.Close()
}
