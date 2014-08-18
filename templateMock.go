package plate



type TemplateMock struct {
    *template.Template
    rendered int32
}

func (t *TemplateMock) TimesRendered() int
func (t *TemplateMock) Output() []byte
func (t *TemplateMock) ContextReceived() interface{}
func (t *TemplateMock) LastExecution() interface{}
