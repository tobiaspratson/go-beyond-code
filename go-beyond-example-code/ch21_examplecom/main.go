package main

import (
    "fmt"
    "os"
    "text/template"
)

func main() {
    // Template with advanced built-in functions
    tmpl := `
User Profile:
=============
Name: {{.Name}} ({{len .Name}} characters)
Age: {{.Age}}
Email: {{.Email}}

Age Analysis:
{{if eq .Age 18}}
- You are exactly 18!
{{else if lt .Age 18}}
- You are under 18 ({{sub 18 .Age}} years to go)
{{else}}
- You are over 18 ({{sub .Age 18}} years past)
{{end}}

Status Check:
{{if and .IsActive .IsVerified}}
- Account is active and verified ✓
{{else if .IsActive}}
- Account is active but not verified ⚠
{{else}}
- Account is inactive ✗
{{end}}

{{if .Hobbies}}
Hobbies ({{len .Hobbies}} total):
{{range $index, $hobby := .Hobbies}}
{{printf "%d. %s (length: %d)" $index $hobby (len $hobby)}}
{{end}}

First hobby: {{index .Hobbies 0}}
Last hobby: {{index .Hobbies (sub (len .Hobbies) 1)}}
{{else}}
No hobbies listed.
{{end}}

{{if .Skills}}
Skills:
{{range $skill, $level := .Skills}}
- {{$skill}}: {{$level}}
{{end}}
{{end}}
`

    // Create template with custom functions
    t, err := template.New("profile").Funcs(template.FuncMap{
        "sub": func(a, b int) int { return a - b },
    }).Parse(tmpl)
    if err != nil {
        fmt.Printf("Error parsing template: %v\n", err)
        return
    }

    // Data
    data := struct {
        Name       string
        Age        int
        Email      string
        IsActive   bool
        IsVerified bool
        Hobbies    []string
        Skills     map[string]string
    }{
        Name:       "Eve",
        Age:        25,
        Email:      "eve@example.com",
        IsActive:   true,
        IsVerified: true,
        Hobbies:    []string{"Reading", "Swimming", "Coding", "Gaming"},
        Skills:     map[string]string{"Go": "Expert", "Python": "Intermediate", "JavaScript": "Beginner"},
    }

    // Execute template
    err = t.Execute(os.Stdout, data)
    if err != nil {
        fmt.Printf("Error executing template: %v\n", err)
    }
}