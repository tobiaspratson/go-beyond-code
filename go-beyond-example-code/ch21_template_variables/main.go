package main

import (
    "fmt"
    "os"
    "text/template"
)

func main() {
    // Template with variables
    tmpl := `
{{$name := .Name}}
{{$age := .Age}}
{{$isAdult := ge .Age 18}}

Hello, {{$name}}!
You are {{$age}} years old.
{{if $isAdult}}
You are an adult.
{{else}}
You are not an adult yet.
{{end}}

{{$hobbies := .Hobbies}}
{{if $hobbies}}
Your hobbies are:
{{range $index, $hobby := $hobbies}}
{{$index}}: {{$hobby}}
{{end}}
{{else}}
You have no hobbies.
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
        Name    string
        Age     int
        Hobbies []string
    }{
        Name:    "Charlie",
        Age:     16,
        Hobbies: []string{"Gaming", "Music"},
    }

    // Execute template
    err = t.Execute(os.Stdout, data)
    if err != nil {
        fmt.Printf("Error executing template: %v\n", err)
    }
}