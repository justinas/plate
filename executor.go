package plate

import (
	"io"
)

// Executor is an interface comprised of the metods
// that html/template and text/template use to render themselves.
// Thus any *template.Template implements Executor automatically.
type Executor interface {
	Execute(wr io.Writer, data interface{}) error
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
}
