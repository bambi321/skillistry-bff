package handlers

import (
	"net/http"
	"skillistry-bff-app/services"
)

func CreateSkill(services *services.Container) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		return
	}
}
