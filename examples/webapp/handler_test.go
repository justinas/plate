package webapp

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/justinas/plate"
)

// Some of the test assertions fail on purpose in this package.
// These are marked by a commment "FAILS".

func TestTemplateRendering(t *testing.T) {
	// Replace the template with a plate.Recorder.
	// The `tpl` variable is already set by init(),
	// so we can just wrap that.
	recorder := plate.New(tpl)
	tpl = recorder

	// Set up a test server, just to distance ourselves
	// from the handler itself
	//
	// This could also be done without setting up a
	// httptest Server and just calling the handler
	// with a ResponseRecorder and a Request directly.
	server := httptest.NewServer(http.HandlerFunc(myHandler))
	defer server.Close()

	// Do a request to our server.
	_, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	if recorder.TimesExecuted() != 1 {
		t.Error("The template didn't get executed")
	}

	// Grab the last execution.
	exec := recorder.LastExecution()

	// Having that, we can check for errors...
	if exec.Error != nil {
		t.Error("An error occured while rendering the template")
	}

	// ...context...
	ctx, ok := recorder.LastExecution().Context.(map[string]interface{})
	if !ok {
		t.Error("A wrong context type was used to execute the template")
	} else if ctx["Answer"] != 42 {
		// FAILS
		t.Errorf("Context contained the wrong answer: %d instead of 42", ctx["Answer"])
	}

	// ...and the output (though that usually matches the http response body)
	if !strings.HasPrefix(string(exec.Output), "The answer is") {
		// FAILS
		t.Error("Template starts with a wrong sentence")
		t.Errorf("Template output: %#s", exec.Output)
	}
}
