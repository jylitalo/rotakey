package rotakey

type MockExec struct{}

func NewMockExec() ExecIface {
	return MockExec{}
}

func (me MockExec) Execute(params ExecuteInput) error {
	return nil
}
