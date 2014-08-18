package plate

import (
	"html/template"
	"text/template"
	"io"
)

type Executable interface{
	Execute(wr io.Writer, data interface{}) error
	ExecuteTemplate(wr io.Writer, name string, data interface{}) error
}

type TemplateMock struct {
	*template.Template
	rendered int32
}

// Satisfies Executable
func (t *TemplateMock) Execute(wr io.Writer, data interface{}) error
func (t *TemplateMock) ExecuteTemplate(wr io.Writer, name string, data interface{}) error

// But has benefits for testing, like:
func (t *TemplateMock) TimesRendered() int
func (t *TemplateMock) Output() []byte
func (t *TemplateMock) ContextReceived() interface{}
func (t *TemplateMock) LastExecution() interface{}

type Recorder struct {
    // Main stuff
    Template Executable

    // Stores exucution info
    execs []Execution
}
