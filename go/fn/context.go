package fn


type Context struct {
	results *Results
	ResMap map[Identifier][]*KubeObject
}

type ErrContextEnd struct {}
func (e *ErrContextEnd) Error() string {
	return ""
}

func (c *Context) AddGeneralResult(message string, severity Severity) {
	*c.results = append(*c.results, &Result{Message: message, Severity: severity})
}

func (c *Context) AddErrResultAndDie(message string, o *KubeObject) {
	c.AddErrResult(message, o)
	panic(ErrContextEnd{})
}

func (c *Context) AddErrResult(message string, o *KubeObject) {
	var r *Result
	if o != nil {
		r = ConfigObjectResult(message, o, Error)
	} else {
		r = GeneralResult(message, Error)
	}
	*c.results = append(*c.results, r)
}
