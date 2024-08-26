package http

import (
	"encoding/json"
	"github.com/pujilesmana/chat-app/internal/repository"
	"github.com/pujilesmana/chat-app/internal/usecase"
	"gorm.io/gorm"
	"net/http"
)

func RegisterHandler(db *gorm.DB, jwtSecret string) http.HandlerFunc {
	authUsecase := usecase.NewAuthUsecase(repository.NewUserRepositoryPostgres(db), jwtSecret)

	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := authUsecase.Register(req.Username, req.Password); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "user registered"}`))
	}
}

func LoginHandler(db *gorm.DB, jwtSecret string) http.HandlerFunc {
	authUsecase := usecase.NewAuthUsecase(repository.NewUserRepositoryPostgres(db), jwtSecret)

	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, err := authUsecase.Login(req.Username, req.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"token": "` + token + `"}`))
	}
}
