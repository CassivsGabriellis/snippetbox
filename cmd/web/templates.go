package main

import (
	"snippetbox.cassius.github/internal/models"
)

type templateData struct {
	Snippet models.Snippet
	Snippets []models.Snippet
}
