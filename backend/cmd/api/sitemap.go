package main

import (
	"fmt"
	"net/http"
	"time"
)

func (app *application) sitemapHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Static URLs (Priority 1.0/0.8)
	// Base URL should ideally be from config, but we'll assume a standard deployment or environment var
	// For now, hardcoding or using a placeholder.
	baseURL := "https://happy-tails.com" // Or derived from app config if available

	type urlEntry struct {
		Loc        string
		LastMod    string
		ChangeFreq string
		Priority   string
	}

	staticPages := []urlEntry{
		{Loc: baseURL + "/", LastMod: time.Now().Format("2006-01-02"), ChangeFreq: "daily", Priority: "1.0"},
		{Loc: baseURL + "/adopt", LastMod: time.Now().Format("2006-01-02"), ChangeFreq: "daily", Priority: "0.9"},
		{Loc: baseURL + "/volunteer", LastMod: time.Now().Format("2006-01-02"), ChangeFreq: "weekly", Priority: "0.8"},
		{Loc: baseURL + "/about", LastMod: time.Now().Format("2006-01-02"), ChangeFreq: "monthly", Priority: "0.7"},
	}

	// 2. Dynamic URLs (Pets)
	pets, err := app.models.Pets.GetAllAvailablePets()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// 3. Build XML String
	// Standard sitemap protocol: https://www.sitemaps.org/protocol.html
	xml := `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
	xml += `<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">` + "\n"

	// Add Static
	for _, page := range staticPages {
		xml += fmt.Sprintf(`  <url>
    <loc>%s</loc>
    <lastmod>%s</lastmod>
    <changefreq>%s</changefreq>
    <priority>%s</priority>
  </url>`+"\n", page.Loc, page.LastMod, page.ChangeFreq, page.Priority)
	}

	// Add Dynamic (Pets)
	for _, pet := range pets {
		// Assuming frontend route is /adopt/pet/:id
		loc := fmt.Sprintf("%s/adopt/pet/%s", baseURL, pet.ID)
		lastMod := pet.UpdatedAt.Format("2006-01-02")
		xml += fmt.Sprintf(`  <url>
    <loc>%s</loc>
    <lastmod>%s</lastmod>
    <changefreq>daily</changefreq>
    <priority>0.8</priority>
  </url>`+"\n", loc, lastMod)
	}

	xml += `</urlset>`

	// 4. Return Response
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(xml))
}
