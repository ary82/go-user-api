package database

import (
	"github.com/gocql/gocql"
)

type Database interface{}

type ScyllaDB struct {
	Session *gocql.Session
}

func NewScyllaDB(addr string, ks string) (*ScyllaDB, error) {
	cluster := gocql.NewCluster(addr)
	cluster.Keyspace = ks

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	return &ScyllaDB{Session: session}, nil
}
