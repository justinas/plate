package plate

import (
	"html/template"
	"text/template"
	"io"
	"byte"
)

type Execution struct {
	Err bool
	Output []byte
}

type Recorder struct {
    // Main stuff
    Template Executable

    // Stores exucution info
    execs []Execution
}
