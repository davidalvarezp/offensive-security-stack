package reporting

import (
	"fmt"
	"html/template"
	"os"
)

// ExportHTML exports all reports to an HTML file
func (rm *ReportManager) ExportHTML(path string) error {
	const tpl = `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>OSS Reports</title>
<style>
body { font-family: Arial, sans-serif; margin: 20px; }
h1 { color: #2c3e50; }
h2 { color: #34495e; }
pre { background: #ecf0f1; padding: 10px; border-radius: 5px; }
</style>
</head>
<body>
<h1>OSS Reports</h1>
<p>Generated: {{index . 0.Timestamp}}</p>
{{range .}}
<h2>{{.ModuleName}} - {{.Title}}</h2>
<p>Timestamp: {{.Timestamp}}</p>
<pre>{{.Content}}</pre>
{{end}}
</body>
</html>
`

	t, err := template.New("report").Parse(tpl)
	if err != nil {
		return fmt.Errorf("failed to parse HTML template: %w", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create HTML file: %w", err)
	}
	defer f.Close()

	if err := t.Execute(f, rm.Reports); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}
