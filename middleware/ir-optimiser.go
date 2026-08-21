package middleware

type Optimiser struct {
	passes   []Pass
	ir       []IRCommand
	runAgain bool
}

type Pass interface {
	Run([]IRCommand) ([]IRCommand, bool)
}

type DeadCodeEliminator struct{}

func (d *DeadCodeEliminator) Run(in []IRCommand) ([]IRCommand, bool) {
	return in, false
}

func Optimise(validatedIR []IRCommand) ([]IRCommand, error) {

	optimiser := Optimiser{
		passes:   make([]Pass, 0),
		ir:       make([]IRCommand, 0),
		runAgain: true,
	}

	// Dead-code eleminator.
	dec := DeadCodeEliminator{}

	optimiser.ir = validatedIR
	optimiser.passes = append(optimiser.passes, &dec)

	for optimiser.runAgain {

		optimiser.runAgain = false

		for _, pass := range optimiser.passes {

			out, changed := pass.Run(optimiser.ir)
			if changed {
				optimiser.runAgain = true
			}

			optimiser.ir = out
		}
	}

	return optimiser.ir, nil
}
