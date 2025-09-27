package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	// HTML template
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
</head>
<body>
    <h1>{{.Title}}</h1>
    <p>Welcome, {{.Name}}!</p>
    <p>Your age is {{.Age}}.</p>
    <p>Your email is {{.Email}}.</p>
    
    {{if .IsAdult}}
    <p>You are an adult.</p>
    {{else}}
    <p>You are not an adult yet.</p>
    {{end}}
    
    {{if .Hobbies}}
    <h2>Your Hobbies:</h2>
    <ul>
    {{range .Hobbies}}
        <li>{{.}}</li>
    {{end}}
    </ul>
    {{else}}
    <p>No hobbies listed.</p>
    {{end}}
</body>
</html>
`

	// Create template
	t, err := template.New("profile").Parse(tmpl)
	if err != nil {
		fmt.Printf("Error parsing template: %v\n", err)
		return
	}

	// Data
	data := struct {
		Title   string
		Name    string
		Age     int
		Email   string
		IsAdult bool
		Hobbies []string
	}{
		Title:   "User Profile",
		Name:    "Frank",
		Age:     28,
		Email:   "frank@example.com",
		IsAdult: true,
		Hobbies: []string{"Reading", "Swimming", "Coding"},
	}

	// Execute template
	err = t.Execute(os.Stdout, data)
	if err != nil {
		fmt.Printf("Error executing template: %v\n", err)
	}
}
