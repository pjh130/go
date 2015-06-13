package databaselib

import (
	"database/sql"
	"fmt"
)

// Provider contains global session methods and saved SessionStores.
// it can operate a SessionStore by its id.
type Provider interface {
	DbInit(config string) error
	GetDb() *sql.DB
	Exec(query string, args ...interface{}) (sql.Result, error)
}

var provides = make(map[string]Provider)

// Register makes a session provide available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, provide Provider) {
	if provide == nil {
		panic("session: Register provide is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provider " + name)
	}
	provides[name] = provide
}

type Manager struct {
	provider Provider
	config   string
}

func (this *Manager) GetProvider() Provider {
	return this.provider
}

func NewManager(provideName, config string) (*Manager, error) {
	fmt.Println("length: ", len(provides))
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("unknown provide %q (forgotten import?)", provideName)
	}

	err := provider.DbInit(config)
	if err != nil {
		return nil, err
	}

	return &Manager{
		provider,
		config,
	}, nil
}
