package middleware

import i "goblin.org/main/middleware/irtypes"

type Optimiser struct {
	passes   []Pass
	ir       []i.IRCommand
	runAgain bool
}

func (o Optimiser) registerPass(p Pass) {
	o.passes = append(o.passes, p)
}

type Pass interface {
	Run([]i.IRCommand) ([]i.IRCommand, bool)
}

type LabelResolution struct{}

func (l *LabelResolution) Run(in []i.IRCommand) ([]i.IRCommand, bool) {
	return in, false
}

func Optimise(validatedIR []i.IRCommand, context *i.IRContext) ([]i.IRCommand, error) {

	optimiser := Optimiser{
		passes:   make([]Pass, 0),
		ir:       make([]i.IRCommand, 0),
		runAgain: true,
	}

	// Dead-code eleminator.
	lr := LabelResolution{}

	optimiser.ir = validatedIR
	optimiser.registerPass(&lr)

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
