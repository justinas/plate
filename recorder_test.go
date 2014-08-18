package plate

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

// This test checks that Recorder executes templates
// as they would be executed without it, that is,
// that it makes no changes to the Execute*() calls,
// context, or output.
func TestRecorderExecutesTemplates(t *testing.T) {
	tpl := template.Must(template.New("t1").Parse(`Hi, {{.}}`))
	rec := &Recorder{Template: tpl}
	ctx := "John"

	buf1 := &bytes.Buffer{}
	buf2 := &bytes.Buffer{}

	err1 := tpl.Execute(buf1, ctx)
	err2 := rec.Execute(buf2, ctx)
	assert.Nil(t, err1)
	assert.Nil(t, err2)

	assert.Equal(t, buf1.String(), "Hi, John")
	assert.Equal(t, buf2.String(), buf1.String())
}

func TestRecorderExecutesNamedTemplates(t *testing.T) {
	tpl := template.Must(template.New("t1").Parse(`{{ define "t2" }}Hi, {{.}}{{ end }}`))
	rec := &Recorder{Template: tpl}
	ctx := "John"

	buf1 := &bytes.Buffer{}
	buf2 := &bytes.Buffer{}

	err1 := tpl.ExecuteTemplate(buf1, "t2", ctx)
	err2 := rec.ExecuteTemplate(buf2, "t2", ctx)
	assert.Nil(t, err1)
	assert.Nil(t, err2)

	assert.Equal(t, buf1.String(), "Hi, John")
	assert.Equal(t, buf2.String(), buf1.String())
}

// Tests that Recorder returns the error
// produced by the inner Executor
func TestRecorderRelaysErrors(t *testing.T) {
	// Lookup a non-existent context member at runtime to produce an error.
	tpl := template.Must(template.New("t1").Parse(`{{ .Name }}`))
	rec := &Recorder{Template: tpl}

	buf1 := &bytes.Buffer{}
	buf2 := &bytes.Buffer{}

	err1 := tpl.ExecuteTemplate(buf1, "t2", nil)
	err2 := rec.ExecuteTemplate(buf2, "t2", nil)

	assert.NotNil(t, err1)
	assert.Equal(t, err1, err2)
}
