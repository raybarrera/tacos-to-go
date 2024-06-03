package main

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	http.ListenAndServe(":3000", r)
}

func ParseMentions(text string) ([]string, error) {
	exp, err := regexp.Compile("(?:<@)(\\w+)")
	if err != nil {
		return nil, err
	}
	matched := exp.FindAllString(text, -1)
	return matched, nil
}

func ParseEmoji(emoji, text string) []string {
	next := strings.Index(text, emoji)
	start := 0
	end := 0
	increment := len(emoji)
	if next < 0 {
		return nil
	}
	result := make([]string, 0)
	for next > -1 {
		start = next
		end = start + increment
		word := text[start:end]
		result = append(result, word)
		start = end
		if start >= len(text) {
			break
		}
		text = text[start:]
		next = strings.Index(text, emoji)
	}
	return result
}
