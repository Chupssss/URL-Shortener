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

func (s *UrlServ) Create(originalUrl string) string {
	shortCode := generateCode()
	s.repo.Save(shortCode, originalUrl)
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

func (s *UrlServ) Resolve(shortUrl string) (string, bool) {
	url, ok := s.repo.Get(shortUrl)
	if !ok {
		return "", false
	}

	s.repo.IncClicks(shortUrl)

	return url, true
}

func (s *UrlServ) Stats(shortUrl string) (string, int, time.Time, bool) {
	return s.repo.GetUrlStats(shortUrl)
}
