package plate

import (
	"io"
	"sync"
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

// Recorder wraps an Executable and
// records results of executions for later checks.
type Recorder struct {
	// The original template to wrap.
	Template Executable

	mu sync.RWMutex
	// Stores exucution info
	execs []Execution
}

func (r *Recorder) Execute(wr io.Writer, data interface{}) error {
	return nil
}

func (r *Recorder) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	return nil
}

// Ensure interface compliance
var _ Executable = &Recorder{}
