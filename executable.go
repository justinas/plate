package plate

import (
    "io"
    )

type Executable interface{
    Execute(wr io.Writer, data interface{}) error
    ExecuteTemplate(wr io.Writer, name string, data interface{}) error
}

// Satisfies Executable
func (t *TemplateMock) Execute(wr io.Writer, data interface{}) error
func (t *TemplateMock) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
