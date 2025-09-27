package main

import (
    "fmt"
    "os"
    "strings"
    "text/template"
    "time"
)

func main() {
    // Template with custom functions
    tmpl := `
Name: {{.Name}}
Age: {{.Age}}
Email: {{.Email}}
Formatted Name: {{title .Name}}
Age in Days: {{ageInDays .Age}}
{{if .IsAdult}}
You are an adult.
{{else}}
You are not an adult yet.
{{end}}

{{if .Hobbies}}
Hobbies:
{{range $index, $hobby := .Hobbies}}
{{$index}}: {{title $hobby}}
{{end}}
{{else}}
No hobbies listed.
{{end}}

Current Time: {{now}}
`

    // Create template with custom functions
    t, err := template.New("profile").Funcs(template.FuncMap{
        "title":      strings.Title,
        "ageInDays":  func(age int) int { return age * 365 },
        "now":        func() string { return time.Now().Format("2006-01-02 15:04:05") },
    }).Parse(tmpl)
    if err != nil {
        fmt.Printf("Error parsing template: %v\n", err)
        return
    }

    // Data
    data := struct {
        Name    string
        Age     int
        Email   string
        IsAdult bool
        Hobbies []string
    }{
        Name:    "eve",
        Age:     25,
        Email:   "eve@example.com",
        IsAdult: true,
        Hobbies: []string{"reading", "swimming", "coding"},
    }

    // Execute template
    err = t.Execute(os.Stdout, data)
    if err != nil {
        fmt.Printf("Error executing template: %v\n", err)
    }
}