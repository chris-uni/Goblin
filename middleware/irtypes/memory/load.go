package memory

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Load struct {
	Destination i.IRTemporary
	Source      i.IRAddress
}

func (l *Load) Exec(state *i.IRExecutionState) (i.IRValue, error) {

	index := l.Destination.Index

	val, err := state.Resolve(l.Source)
	if err != nil {
		return nil, err
	}

	state.PushTemporaries(index, val)

	// Increment PC.
	state.PC++

	return nil, nil
}

func (l *Load) Validate(context *i.IRContext) error {

	_, err := i.ResolveIRType(l.Source, context)
	if err != nil {
		return err
	}

	_, err = i.ResolveIRType(l.Destination, context)
	if err != nil {
		return err
	}

	return nil
}

func (l *Load) String() string {
	return fmt.Sprintf("load %v %v", l.Destination, l.Source)
}
