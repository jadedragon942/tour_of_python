// main.go
package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "os/exec"
)

func executePythonCode(code string) (string, error) {
    log.Printf("Executing Python Code:\n%s", code)
    
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
    
    if err != nil {
        return stderr.String(), err
    }
    return out.String(), nil
}

func codeHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            http.Error(w, "Unable to read request body", http.StatusBadRequest)
            return
        }

        output, err := executePythonCode(string(body))
        if err != nil {
            fmt.Fprint(w, output) // Send only the Python output, suppressing unnecessary errors
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

