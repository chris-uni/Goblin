package irtypes

type IRType int

const (
	IRTypeNumber IRType = iota
	IRTypeString
	IRTypeBoolean
	IRTypeLabel
	IRTypeUndefined
)

func (t IRType) String() string {
	switch t {
	case IRTypeNumber:
		return "number"
	case IRTypeString:
		return "string"
	case IRTypeBoolean:
		return "boolean"
	default:
		return "unknown"
	}
}

type IRContext struct {
	Commands []IRCommand

	Storage     []IRValue
	Temporaries []IRValue
	Labels      []IRLabel

	Symbols map[string]IRAddress

	PC int
}

type IRResult struct {
	Commands []IRCommand
	Value    IRValue
}

/*
Pushes a new command into the context.
*/
func (context *IRContext) Push(com IRCommand) {
	context.Commands = append(context.Commands, com)
	context.PC++
}

func (c *IRContext) AllocateAddress() IRAddress {

	storage := IRAddress{
		Index: len(c.Storage),
	}

	c.Storage = append(c.Storage, nil)
	return storage
}

func (c *IRContext) AllocateTemporary() IRTemporary {

	temporary := IRTemporary{
		Index: len(c.Temporaries),
	}

	c.Temporaries = append(c.Temporaries, nil)
	return temporary
}

func (c *IRContext) AllocateLabel() IRLabel {

	label := IRLabel{
		Value: len(c.Labels),
	}

	c.Labels = append(c.Labels, label)
	return label
}

func (c *IRContext) StoreSymbol(name string, address IRAddress) {

	c.Symbols[name] = address
}
