package memory

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Store struct {
	Destination i.IRAddress
	Value       i.IRValue
}

func (s *Store) Exec(state *i.IRExecutionState) (i.IRValue, error) {

	val, err := state.Resolve(s.Value)
	if err != nil {
		return nil, err
	}

	state.PushStorage(s.Destination.Index, val)

	// Increment PC.
	state.PC++

	return nil, nil
}

func (s *Store) Validate(context *i.IRContext) error {

	_, err := i.ResolveIRType(s.Destination, context)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) String() string {
	return fmt.Sprintf("store %v %v", s.Destination, s.Value)
}
