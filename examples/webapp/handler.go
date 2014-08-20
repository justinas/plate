// Package webapp demonstrates the use of plate in detail.
package webapp

import (
	"html/template"
	"net/http"

	"github.com/justinas/plate"
)

// Use plate.Executor instead of *template.Template
// This way, you can reuse the variable for both
// production and tests.
var tpl plate.Executor

func myHandler(w http.ResponseWriter, r *http.Request) {
	// Is that really the right answer?
	ctx := map[string]interface{}{"Answer": 41}

	// Render the template as always.
	err := tpl.Execute(w, ctx)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

func init() {
	// Oops, a misspelling.
	tplStr := `The anzwer is {{ .Answer }}`
	tpl = template.Must(template.New("").Parse(tplStr))
}
