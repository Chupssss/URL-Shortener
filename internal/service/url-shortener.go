package service

import (
	"URL-Shortener/internal/repos"
	"math/rand"
	"time"
)

type UrlServ struct {
	repo *repos.UrlRepo
}

func NewUrlService(repo *repos.UrlRepo) *UrlServ {
	return &UrlServ{repo: repo}
}

func (service *UrlServ) Create(originalUrl string) string {
	shortCode := generateCode()
	service.repo.Save(shortCode, originalUrl)
	return shortCode
}

func generateCode() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().UnixNano())

	code := make([]byte, 5)
	for i := 0; i < len(code); i++ {
		code[i] = letters[rand.Intn(len(letters))]
	}

	return string(code)
}
