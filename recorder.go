package plate

import (
	"bytes"
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

// Recorder wraps an Executor and
// records results of executions for later checks.
type Recorder struct {
	// The original template to wrap.
	Template Executor

	// Go's templates are already safe to be used in parallel,
	// this mutex only protects our own fields, like `execs`.
	mu sync.RWMutex
	// Stores exucution info
	execs []Execution
}

func (r *Recorder) save(exec Execution) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.execs = append(r.execs, exec)
}

func (r *Recorder) Execute(wr io.Writer, data interface{}) error {
	exec := Execution{Context: data}

	// Substitute the reader
	buf := &bytes.Buffer{}
	writer := io.MultiWriter(buf, wr)

	// Execute and fill out the results
	err := r.Template.Execute(writer, data)
	exec.Output = buf.Bytes()
	exec.Error = err

	r.save(exec)
	return err
}

func (r *Recorder) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	exec := Execution{Context: data}

	// Substitute the reader
	buf := &bytes.Buffer{}
	writer := io.MultiWriter(buf, wr)

	// Execute and fill out the results
	err := r.Template.ExecuteTemplate(writer, name, data)
	exec.Output = buf.Bytes()
	exec.Error = err

	// Save the execution

	r.save(exec)
	return err
}


// Executions() return all executions that have occured
// since the construction of a Recorder (or since Reset()).
func (r *Recorder) Executions() []Execution {
	tmpExecs := make([]Execution, len(r.execs))
	// We do a copy, because callee may mess around with internal []Execution
	// and we do not want this.
	r.mu.RLock()
	defer r.mu.RUnlock()
	copy(tmpExecs, r.execs)
	return tmpExecs
}

// LastExecution() returns the last execution.
// It panics if no executions occured yet.
func (r *Recorder) LastExecution() Execution {
	if len(r.execs) < 1 {
		panic("No executions are available yet.")
	}
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.execs[len(r.execs)-1]
}

// TimesRendered() returns times template was rendered in int
// This is since construction or Reset()
func (r *Recorder) TimesRendered() int {
	return len(r.execs)
}

// FailedExecutions() returns all executions that have Error != nil
func (r *Recorder) FailedExecutions() []Execution {
	failedExecs := make([]Execution, 0)
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _,exec := range r.execs {
		if exec.Error != nil {
			failedExecs = append(failedExecs, exec)
		}
	}

	return failedExecs
}

// Reset() clears all executions. Recorder is thus restored to its initial state.
func (r *Recorder) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.execs = make([]Execution, 0)
}

// Ensure interface compliance
var _ Executor = &Recorder{}
