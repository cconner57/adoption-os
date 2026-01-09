package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
)

// SEOPet matches the data needed for the preview
type SEOPet struct {
	Name     string
	Bio      string
	ImageURL string
}

var (
	indexHTML      string
	indexHTMLErr   error
	indexHTMLOnce  sync.Once
	indexHTMLMutex sync.RWMutex
)

// getIndexHTML reads and caches the index.html file.
// In development, you might want to skip caching to see changes instantly.
func getIndexHTML() (string, error) {
	indexHTMLOnce.Do(func() {
		// Try multiple paths to find the dist folder
		paths := []string{
			"./frontend/dist/index.html", // Standard relative path
			"./dist/index.html",          // If running from frontend dir?
			"../frontend/dist/index.html",
		}

		var content []byte
		var err error

		found := false
		for _, path := range paths {
			content, err = os.ReadFile(path)
			if err == nil {
				indexHTML = string(content)
				found = true
				break
			}
		}

		if !found {
			// Fallback for dev environment without build
			// Attempt to read source index.html as fallback (though it won't have build assets)
			content, err = os.ReadFile("./frontend/index.html")
			if err == nil {
				indexHTML = string(content)
			} else {
				indexHTMLErr = fmt.Errorf("could not find index.html in expected paths")
			}
		}
	})

	return indexHTML, indexHTMLErr
}

func (app *application) serveDynamicSEO(w http.ResponseWriter, r *http.Request, pet SEOPet) {
	// 1. Get the template (Cached)
	tmpl, err := getIndexHTML()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// 2. Generate Meta Tags
	// Truncate bio if necessary
	description := pet.Bio
	if len(description) > 150 {
		description = description[:147] + "..."
	}

	metaTags := fmt.Sprintf(`
    <meta property="og:title" content="Adopt %s | I Dream of Home Rescue" />
    <meta property="og:description" content="%s" />
    <meta property="og:image" content="%s" />
    <meta name="twitter:card" content="summary_large_image" />
	`, pet.Name, description, pet.ImageURL)

	// 3. Inject into Template
	// We look for the placeholder <!--app-head-->
	modifiedHTML := strings.Replace(tmpl, "<!--app-head-->", metaTags, 1)

	// 4. Serve
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(modifiedHTML))
}
