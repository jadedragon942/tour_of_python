// main.go
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func executePythonCode(code string) (string, error) {
    tmpfile, err := os.CreateTemp("/tmp", "code-*.py")
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

    tmpDir := filepath.Dir(tmpfile.Name())
    cmd := exec.Command("podman", "run", "--rm", "-v", fmt.Sprintf("%s:/scripts", tmpDir), "-w", "/scripts", "python:3.11", "python", filepath.Base(tmpfile.Name()))

    var out, stderr bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &stderr
    err = cmd.Run()
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

func main() {
	http.HandleFunc("/execute", codeHandler)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
