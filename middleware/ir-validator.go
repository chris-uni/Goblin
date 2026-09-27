/*
Goblin IR Validator v0.1
Author: Chris J.M. Wing
Date: 21/08/2026

Input:
	Raw GoblinIR from the Reducer stage.
Output:
	A set of valid i.IRCommands that represent a validated GoblinIR program.
*/

package middleware

import (
	"fmt"

	i "goblin.org/main/middleware/irtypes"
)

func Validate(commands []i.IRCommand, context *i.IRContext) ([]i.IRCommand, error) {

	context.PC = 0

	for _, command := range commands {

		err := command.Validate(context)
		if err != nil {
			return nil, fmt.Errorf("validation error: %v\n", err)
		}
	}

	return context.Commands, nil
}
