package main

import "snippetbox.suphachok.net/internal/models"

type templateData struct {
	Snippet models.Snippet
	Snippets []models.Snippet
}