package main

import (
    "fmt"
    "os"
    "text/template"
)

func main() {
    // Template with built-in functions
    tmpl := `
Name: {{.Name}}
Age: {{.Age}}
Email: {{.Email}}
Name Length: {{len .Name}}
Age + 10: {{add .Age 10}}
{{if eq .Age 18}}
You are exactly 18!
{{else if lt .Age 18}}
You are under 18.
{{else}}
You are over 18.
{{end}}

{{if .Hobbies}}
Hobbies ({{len .Hobbies}} total):
{{range $index, $hobby := .Hobbies}}
{{$index}}: {{$hobby}}
{{end}}
{{else}}
No hobbies listed.
{{end}}
`

    // Create template with custom functions
    t, err := template.New("profile").Funcs(template.FuncMap{
        "add": func(a, b int) int { return a + b },
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
        Hobbies []string
    }{
        Name:    "David",
        Age:     22,
        Email:   "david@example.com",
        Hobbies: []string{"Reading", "Swimming", "Coding", "Gaming"},
    }

    // Execute template
    err = t.Execute(os.Stdout, data)
    if err != nil {
        fmt.Printf("Error executing template: %v\n", err)
    }
}