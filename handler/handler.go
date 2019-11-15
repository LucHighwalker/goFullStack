package handler

import (
	"golang-starter-pack/article"
	"golang-starter-pack/magic"
	"golang-starter-pack/user"
)

type Handler struct {
	userStore    user.Store
	articleStore article.Store
	magicStore   magic.Store
}

func NewHandler(us user.Store, as article.Store, ms magic.Store) *Handler {
	return &Handler{
		userStore:    us,
		articleStore: as,
		magicStore:   ms,
	}
}
