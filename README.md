# Plate

[![Build Status](https://travis-ci.org/justinas/plate.svg?branch=master)](https://travis-ci.org/justinas/plate)
[![GoDoc](https://godoc.org/github.com/justinas/plate?status.svg)](https://godoc.org/github.com/justinas/plate)

Plate is a wrapper for Go's `html/template` (or `text/template`)
that helps you test your template execution.

Plate makes it easier to check
if the correct template has been rendered,
whether the correct context has been passed to it
and catch errors that occur.

### Usage
First, rewrite your template variables to have the type `plate.Executor`.
In short, transform this:

```go
var tmpl *template.Template
// or
var templates map[string]*template.Template
```

To this:

```go
var tmpl plate.Executor
// or
var templates map[string]plate.Executor
```

Your templates will implement `plate.Executor` automatically.
Then, in your tests, wrap your template in a `plate.Recorder`:

```go
recorder := plate.New(tmpl)
tmpl = recorder
// or
recorder := plate.New(tmpl)
templates["index.html"] = recorder
```

The template will execute as before, except for one thing:
the recorder will accumulate the result of all executions:
the output that template produced, the context passed to it
and an error returned from an `Execute*()` call.

That information can be checked later
to find out any faults in the execution of your template.

```go
tmpl.Execute(...)
if recorder.LastExecution().Context == nil {
    t.Error("A nil context has been passed to the template")
}
```
