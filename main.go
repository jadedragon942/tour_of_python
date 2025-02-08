// main.go
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

var templates = template.Must(template.New("index.html").Funcs(template.FuncMap{
	"safeHTML": func(s string) template.HTML {
		return template.HTML(s)
	},
}).ParseFiles("templates/index.html"))

func countSlides() int {
	maxSlide := 0
	entries, err := os.ReadDir("slides")
	if err != nil {
		return 0
	}

	for _, entry := range entries {
		if entry.IsDir() {
			if num, err := strconv.Atoi(entry.Name()); err == nil {
				if num > maxSlide {
					maxSlide = num
				}
			}
		}
	}
	return maxSlide
}

func tourHandler(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile(`/tour/welcome/(\d+)`)
	match := re.FindStringSubmatch(r.URL.Path)
	if len(match) != 2 {
		http.Redirect(w, r, "/tour/welcome/1", http.StatusFound)
		return
	}

	slideNumber, err := strconv.Atoi(match[1])
	if err != nil {
		http.Error(w, "Invalid slide number", http.StatusBadRequest)
		return
	}

	slidePath := filepath.Join("slides", fmt.Sprintf("%d", slideNumber))
	titlePath := filepath.Join(slidePath, "title")
	markdownPath := filepath.Join(slidePath, "content.md")
	pythonPath := filepath.Join(slidePath, "script.py")

	titleContent, err := os.ReadFile(titlePath)
	if err != nil {
		http.Error(w, "Title file not found", http.StatusNotFound)
		return
	}

	markdownContent, err := os.ReadFile(markdownPath)
	if err != nil {
		http.Error(w, "Markdown file not found", http.StatusNotFound)
		return
	}

	// Preprocess markdown content to ensure proper list formatting
	markdownRe := regexp.MustCompile(`(?m)^-([^\s])`)
	processedContent := markdownRe.ReplaceAll(markdownContent, []byte("- $1"))

	pythonContent, err := os.ReadFile(pythonPath)
	if err != nil {
		http.Error(w, "Python script file not found", http.StatusNotFound)
		return
	}

	// Create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extensions)

	// Create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{
		Flags: htmlFlags,
	}
	renderer := html.NewRenderer(opts)

	// Parse and render markdown with custom renderer
	parsedContent := p.Parse(processedContent)
	htmlContent := markdown.Render(parsedContent, renderer)

	data := struct {
		Title        string
		MarkdownHTML template.HTML
		PythonCode   string
		MaxSlides    int
	}{
		Title:        string(titleContent),
		MarkdownHTML: template.HTML(htmlContent),
		PythonCode:   string(pythonContent),
		MaxSlides:    countSlides(),
	}

	templates.ExecuteTemplate(w, "index.html", data)
}

func main() {
	http.HandleFunc("/tour/welcome/", tourHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/tour/welcome/1", http.StatusFound)
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server started at 127.0.0.1:9000")
	log.Fatal(http.ListenAndServe("127.0.0.1:9000", nil))
}
