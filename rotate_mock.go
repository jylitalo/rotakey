package rotakey

type MockExec struct{}

func NewMockExec() ExecIface {
	return MockExec{}
}

func (client MockExec) Execute(params ExecuteInput) error {
	return nil
}
