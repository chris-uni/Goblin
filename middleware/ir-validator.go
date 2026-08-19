package middleware

/*
Stages of validator:

1. Invalid address (storage)
2. Invalid temporary
3. Type errors
4. Invalid instruction operands
5. Invalid control flow
*/

func Validate(commands []IRCommand) ([]IRCommand, error) {

	validatedCommands := make([]IRCommand, 0)

	for _, command := range commands {

		validatedCommands = append(validatedCommands, command)
	}

	return validatedCommands, nil
}
