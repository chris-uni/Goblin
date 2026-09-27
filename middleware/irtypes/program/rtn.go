package program

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Return struct {
	Value i.IRValue
}

func (r *Return) Exec(state *i.IRExecutionState) (i.IRValue, error) {

	val, err := state.Resolve(r.Value)
	if err != nil {
		return nil, err
	}

	result := &i.IRExecutionResult{
		Value: val,
	}

	state.Result = result

	return nil, nil
}

func (r *Return) Validate(state *i.IRContext) error {

	return nil
}

func (r *Return) String() string {
	return fmt.Sprintf("rtn %v", r.Value)
}
