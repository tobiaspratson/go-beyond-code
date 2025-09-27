package main

import (
    "fmt"
    "os"
    "text/template"
)

func main() {
    // Define a simple template
    tmpl := `Hello, {{.Name}}!
You are {{.Age}} years old.
Your email is {{.Email}}.`

    // Create template
    t, err := template.New("greeting").Parse(tmpl)
    if err != nil {
        fmt.Printf("Error parsing template: %v\n", err)
        return
    }

    // Data to fill the template
    data := struct {
        Name  string
        Age   int
        Email string
    }{
        Name:  "Alice",
        Age:   25,
        Email: "alice@example.com",
    }

    // Execute template
    err = t.Execute(os.Stdout, data)
    if err != nil {
        fmt.Printf("Error executing template: %v\n", err)
    }
}