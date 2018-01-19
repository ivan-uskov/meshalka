package actions

import (
	"meshalka/model"
	"meshalka/contexts"
)

type RequestContext struct {
	Data string
	User *model.User
	Ctx  *contexts.Context
}