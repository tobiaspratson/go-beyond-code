package main

import (
    "fmt"
    "os"
    "text/template"
)

func main() {
    // Template with advanced variable usage
    tmpl := `
{{$user := .}}
{{$name := .Name}}
{{$age := .Age}}
{{$isAdult := ge .Age 18}}
{{$hobbyCount := len .Hobbies}}

User Information:
================
Name: {{$name}}
Age: {{$age}}
Status: {{if $isAdult}}Adult{{else}}Minor{{end}}
Hobby Count: {{$hobbyCount}}

{{if $hobbies}}
Hobbies:
{{range $index, $hobby := .Hobbies}}
{{$index}}: {{$hobby}}
{{end}}
{{else}}
No hobbies listed.
{{end}}

{{$totalChars := 0}}
{{range .Hobbies}}
{{$totalChars = add $totalChars (len .)}}
{{end}}
Total characters in hobby names: {{$totalChars}}
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
        Hobbies []string
    }{
        Name:    "David",
        Age:     25,
        Hobbies: []string{"Reading", "Swimming", "Coding"},
    }

    // Execute template
    err = t.Execute(os.Stdout, data)
    if err != nil {
        fmt.Printf("Error executing template: %v\n", err)
    }
}