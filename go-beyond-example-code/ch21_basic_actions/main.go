package main

import (
    "fmt"
    "os"
    "text/template"
)

func main() {
    // Template with various actions
    tmpl := `
Name: {{.Name}}
Age: {{.Age}}
Email: {{.Email}}
Is Adult: {{if .IsAdult}}Yes{{else}}No{{end}}
{{if .IsAdult}}
You can vote!
{{end}}
{{range .Hobbies}}
- {{.}}
{{end}}
`

    // Create template
    t, err := template.New("profile").Parse(tmpl)
    if err != nil {
        fmt.Printf("Error parsing template: %v\n", err)
        return
    }

    // Data
    data := struct {
        Name     string
        Age      int
        Email    string
        IsAdult  bool
        Hobbies  []string
    }{
        Name:    "Bob",
        Age:     30,
        Email:   "bob@example.com",
        IsAdult: true,
        Hobbies: []string{"Reading", "Swimming", "Coding"},
    }

    // Execute template
    err = t.Execute(os.Stdout, data)
    if err != nil {
        fmt.Printf("Error executing template: %v\n", err)
    }
}