// main.go
package main
import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/gomarkdown/markdown"
)

var templates = template.Must(template.New("index.html").Funcs(template.FuncMap{
	"safeHTML": func(s string) template.HTML {
		return template.HTML(s)
	},
}).ParseFiles("templates/index.html"))

func sendCodeToAxiom(code, stdout, stderr string) error {
	axiomKey := os.Getenv("AXIOM_KEY")
	if axiomKey == "" {
		panic("please specify AXIOM_KEY")
	}
	axiomDataset := os.Getenv("AXIOM_DATASET")
	if axiomDataset == "" {
		panic("please specify AXIOM_DATASET")
	}

	payload := []map[string]interface{}{
		{
			"data": map[string]string{
				"code":   code,
				"stdout": stdout,
				"stderr": stderr,
			},
		},
	}

	// Marshal the payload into JSON.
	body, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		return fmt.Errorf("Error marshaling JSON: %v", err)
	}

	// Construct the ingest URL.
	url := fmt.Sprintf("https://api.axiom.co/v1/datasets/%s/ingest", axiomDataset)

	// Create a new HTTP request.
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return fmt.Errorf("error creating request: %v", err)
	}

	// Set the required headers: Authorization and Content-Type.
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", axiomKey))
	req.Header.Set("Content-Type", "application/json")

	// Send the request using the default HTTP client (you could configure your own).
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error performing request: %v", err)
		return fmt.Errorf("error performing request: %v", err)
	}
	defer resp.Body.Close()

	// Optionally, handle non-OK status codes or read response body here.
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Printf("Non-2xx status code received: %d\n", resp.StatusCode)
		// You could read and log resp.Body for details.
		return fmt.Errorf("non-2xx status code received: %d\n", resp.StatusCode)
	}

	return nil
}

func executePythonCode(code string) (string, error) {
	log.Printf("Executing Python Code:\n%s", code)

	err := sendCodeToAxiom(code, "", "")
	if err != nil {
		return "", err
	}

	tmpfile, err := os.CreateTemp("", "code-*.py")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(code)); err != nil {
		return "", err
	}
	if err := tmpfile.Close(); err != nil {
		return "", err
	}

	cmd := exec.Command("python3", "./sandbox.py", tmpfile.Name())
	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()

	log.Printf("Stdout:\n%s", out.String())
	log.Printf("Stderr:\n%s", stderr.String())

	err = sendCodeToAxiom(code, out.String(), stderr.String())
	if err != nil {
		return "", err
	}

	if err != nil {
		return stderr.String(), err
	}
	return out.String(), nil
}

func codeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusBadRequest)
			return
		}

		output, err := executePythonCode(string(body))
		if err != nil {
			http.Error(w, output, http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, output)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
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

	pythonContent, err := os.ReadFile(pythonPath)
	if err != nil {
		http.Error(w, "Python script file not found", http.StatusNotFound)
		return
	}

	htmlContent := markdown.ToHTML(markdownContent, nil, nil)

	data := struct {
		Title        string
		MarkdownHTML template.HTML
		PythonCode   string
	}{
		Title:        string(titleContent),
		MarkdownHTML: template.HTML(htmlContent),
		PythonCode:   string(pythonContent),
	}

	templates.ExecuteTemplate(w, "index.html", data)
}

func main() {
	http.HandleFunc("/execute", codeHandler)
	http.HandleFunc("/tour/welcome/", tourHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/tour/welcome/1", http.StatusFound)
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server started at :9000")
	log.Fatal(http.ListenAndServe("127.0.0.1:9000", nil))
}
