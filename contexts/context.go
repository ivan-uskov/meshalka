package contexts

import (
	"meshalka/database"
	"meshalka/config"
	"meshalka/sessions"
)

type Context struct {
	Session  sessions.Session
	Database database.Database
	Config   *config.Config
}

func NewContext(path string) (*Context, error) {
	c, err := config.Load(path)
	if err != nil {
		return nil, err
	}

	db, err := database.New(&c.Database)
	if err != nil {
		return nil, err
	}

	return &Context{sessions.New(c.SessionSecret, db), db, c}, nil
}