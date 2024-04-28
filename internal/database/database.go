package database

type Store interface {
	Close()
	InitTables() error

	CreateUser(email string, name string, hashed_pass string) error
	GetUser(email string) (*User, error)
}
