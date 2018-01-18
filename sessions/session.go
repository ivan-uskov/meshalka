package sessions

import (
	"github.com/gorilla/sessions"
	"net/http"
	"meshalka/model"
	"meshalka/database"
)

const cookieName = `mesh`

type session struct {
	store *sessions.CookieStore
	db    database.Database
}

type Session interface {
	Login(w http.ResponseWriter, r *http.Request, u *model.User) bool
	AutoLoginFilter(r *http.Request) (*model.User, bool)
}

func New(secret string, db database.Database) Session {
	return &session{sessions.NewCookieStore([]byte(secret)), db}
}

func (s *session) Login(w http.ResponseWriter, r *http.Request, u *model.User) bool {
	session, err := s.store.Get(r, cookieName)
	if err != nil {
		return false
	}

	session.Values[`user_id`] = u.UserId
	session.Save(r, w)
	return true
}

func (s *session) AutoLoginFilter(r *http.Request) (*model.User, bool) {
	session, err := s.store.Get(r, cookieName)
	if err != nil || session.IsNew {
		return nil, false
	}

	userIdV, ok := session.Values[`user_id`]
	if !ok {
		return nil, false
	}

	userId, ok := userIdV.(uint64)
	if !ok {
		return nil, false
	}

	user, err := model.NewUserRepository(s.db).SelectUserById(userId)
	if err != nil {
		return nil, false
	}

	return user, true
}
