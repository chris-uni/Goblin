package memory

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Load struct {
	Destination i.IRTemporary
	Source      i.IRAddress
}

func (l *Load) Exec(context *i.IRContext) {}

func (l *Load) Validate(context *i.IRContext) error {

	/*
		For load, we need to check the source exists as we are loading a stored variable into temporary memory.
	*/
	_, err := i.ResolveIRType(l.Source, context)
	if err != nil {
		return err
	}

	return nil
}

func (l *Load) String() string {
	return fmt.Sprintf("load %v %v", l.Destination, l.Source)
}
