package plate

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	template string
	context  interface{}
}

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
	tpl := template.Must(template.New("t1").Parse(`{{ .Email }}`))
	rec := &Recorder{Template: tpl}
	ctx := struct{ Name string }{"John"}

	buf1 := &bytes.Buffer{}
	buf2 := &bytes.Buffer{}

	err1 := tpl.Execute(buf1, ctx)
	err2 := rec.Execute(buf2, ctx)

	assert.NotNil(t, err1)
	assert.Equal(t, err1, err2)
}

// Tests that recorder accumulates a history of executions
func TestRecorderRecordsExecutions(t *testing.T) {
	commonCtx := struct{ Name string }{"John"}

	var cases = []testCase{
		// A valid template
		{`{{ .Name }}`, commonCtx},
		// A template with a runtime error
		{`{{ .Email }}`, commonCtx},
	}

	var namedCases = []testCase{
		// A valid template
		{`{{ define "t2" }}Hi, {{ .Name }}{{ end }}`, commonCtx},
		// A template with a runtime error
		{`{{ define "t2" }}Hi, {{ .Email }}{{ end }}`, commonCtx},
	}

	for _, c := range cases {
		buf1 := &bytes.Buffer{}
		buf2 := &bytes.Buffer{}

		tpl := template.Must(template.New("t1").Parse(c.template))
		ctx := c.context
		rec := &Recorder{Template: tpl}

		err := tpl.Execute(buf1, ctx)
		_ = rec.Execute(buf2, ctx)

		if !assert.Equal(t, len(rec.execs), 1) {
			t.FailNow()
		}

		assert.Equal(t, rec.execs[0].Error, err)
		assert.Equal(t, rec.execs[0].Output, buf1.Bytes())
		assert.Equal(t, rec.execs[0].Context, ctx)
	}

	for _, c := range namedCases {
		buf1 := &bytes.Buffer{}
		buf2 := &bytes.Buffer{}

		tpl := template.Must(template.New("t1").Parse(c.template))
		ctx := c.context
		rec := &Recorder{Template: tpl}

		err := tpl.ExecuteTemplate(buf1, "t2", ctx)
		_ = rec.ExecuteTemplate(buf2, "t2", ctx)

		if !assert.Equal(t, len(rec.execs), 1) {
			t.FailNow()
		}

		assert.Equal(t, rec.execs[0].Error, err)
		assert.Equal(t, rec.execs[0].Output, buf1.Bytes())
		assert.Equal(t, rec.execs[0].Context, ctx)

		buf1.Reset()
		buf2.Reset()
	}
}
