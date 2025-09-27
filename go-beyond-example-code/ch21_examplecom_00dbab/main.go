package main

import (
    "fmt"
    "os"
    "strings"
    "text/template"
    "time"
)

func main() {
    // Template with advanced custom functions
    tmpl := `
User Profile Report
===================
Generated: {{formatTime .GeneratedAt "2006-01-02 15:04:05"}}

Personal Information:
- Name: {{title .Name}} ({{len .Name}} characters)
- Age: {{.Age}} years ({{ageInDays .Age}} days)
- Email: {{.Email}}
- Status: {{statusText .IsActive .IsVerified}}

{{if .Hobbies}}
Hobbies ({{len .Hobbies}} total):
{{range $index, $hobby := .Hobbies}}
{{$index}}: {{title $hobby}} ({{len $hobby}} chars)
{{end}}
{{else}}
No hobbies listed.
{{end}}

{{if .Skills}}
Skills Summary:
{{range $skill, $level := .Skills}}
- {{title $skill}}: {{$level}} ({{skillScore $level}}/10)
{{end}}
Average Skill Level: {{avgSkillLevel .Skills}}
{{end}}

{{if .LastLogin}}
Last Login: {{formatTime .LastLogin "Jan 2, 2006 at 3:04 PM"}}
Days Since Login: {{daysSince .LastLogin}}
{{end}}

{{if .Notes}}
Notes: {{truncate .Notes 50}}
{{end}}
`

    // Create template with comprehensive custom functions
    t, err := template.New("profile").Funcs(template.FuncMap{
        "title":         strings.Title,
        "ageInDays":     func(age int) int { return age * 365 },
        "formatTime":    func(t time.Time, layout string) string { return t.Format(layout) },
        "statusText":    func(isActive, isVerified bool) string {
            if isActive && isVerified { return "Active & Verified ✓" }
            if isActive { return "Active but not verified ⚠" }
            return "Inactive ✗"
        },
        "skillScore":    func(level string) int {
            switch strings.ToLower(level) {
            case "expert": return 10
            case "advanced": return 8
            case "intermediate": return 6
            case "beginner": return 3
            default: return 1
            }
        },
        "avgSkillLevel": func(skills map[string]string) float64 {
            if len(skills) == 0 { return 0 }
            total := 0
            for _, level := range skills {
                switch strings.ToLower(level) {
                case "expert": total += 10
                case "advanced": total += 8
                case "intermediate": total += 6
                case "beginner": total += 3
                default: total += 1
                }
            }
            return float64(total) / float64(len(skills))
        },
        "daysSince":     func(t time.Time) int { return int(time.Since(t).Hours() / 24) },
        "truncate":      func(s string, maxLen int) string {
            if len(s) <= maxLen { return s }
            return s[:maxLen] + "..."
        },
    }).Parse(tmpl)
    if err != nil {
        fmt.Printf("Error parsing template: %v\n", err)
        return
    }

    // Data
    data := struct {
        GeneratedAt time.Time
        Name        string
        Age         int
        Email       string
        IsActive    bool
        IsVerified  bool
        Hobbies     []string
        Skills      map[string]string
        LastLogin   time.Time
        Notes       string
    }{
        GeneratedAt: time.Now(),
        Name:        "frank",
        Age:         30,
        Email:       "frank@example.com",
        IsActive:    true,
        IsVerified:  true,
        Hobbies:     []string{"reading", "swimming", "coding", "gaming"},
        Skills:      map[string]string{"Go": "Expert", "Python": "Advanced", "JavaScript": "Intermediate"},
        LastLogin:   time.Now().AddDate(0, 0, -5), // 5 days ago
        Notes:       "This is a very long note that should be truncated when displayed in the template to keep the output clean and readable.",
    }

    // Execute template
    err = t.Execute(os.Stdout, data)
    if err != nil {
        fmt.Printf("Error executing template: %v\n", err)
    }
}