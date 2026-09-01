package memory

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

type Store struct {
	Destination i.IRAddress
	Value       i.IRValue
}

func (s *Store) Exec(context *i.IRContext) {}

func (s *Store) Validate(context *i.IRContext) error {
	return nil
}

func (s *Store) String() string {
	return fmt.Sprintf("store %v %v", s.Destination, s.Value)
}
