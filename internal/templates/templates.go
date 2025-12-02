package templates

import (
	"bytes"
	"fmt"
	"text/template"

	_ "embed"
)

//go:embed day/solver.go.tmpl
var daySolverTemplate string

//go:embed day/solver_test.go.tmpl
var dayTestTemplate string

var (
	solverTmpl = template.Must(template.New("daySolver").Parse(daySolverTemplate))
	testTmpl   = template.Must(template.New("dayTest").Parse(dayTestTemplate))
)

type dayData struct {
	Year    int
	Day     int
	Package string
}

// RenderDaySolver returns the scaffolded solver implementation.
func RenderDaySolver(year, day int) ([]byte, error) {
	data := dayData{Year: year, Day: day, Package: fmt.Sprintf("day%02d", day)}
	var buf bytes.Buffer
	if err := solverTmpl.Execute(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// RenderDayTest returns the scaffolded solver tests.
func RenderDayTest(year, day int) ([]byte, error) {
	data := dayData{Year: year, Day: day, Package: fmt.Sprintf("day%02d", day)}
	var buf bytes.Buffer
	if err := testTmpl.Execute(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
