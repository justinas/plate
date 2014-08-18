package plate

import (
	"html/template"
	"text/template"
	"io"
	"byte"
)

// Execution represents one occurence of template being executed.
// It provides access to the output produced,
// the context that was passed to the template
// and the error returned from the Execute*() function, if any.
type Execution struct {
	Output  []byte
	Context interface{}

	Error error
}

type Recorder struct {
    // Main stuff
    Template Executable

    // Stores exucution info
    execs []Execution
}
