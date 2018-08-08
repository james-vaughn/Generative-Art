package seir

type SeirModel struct {
	Config SeirConfig
	S      []int
	E      []int
	I      []int
	R      []int
}

type SeirModelPoint struct {
	SeirConfig
	S float64
	E float64
	I float64
	R float64
}

func NewModel(config SeirConfig) SeirModel {
	return SeirModel{
		Config: config,
		S:      make([]int, config.Steps),
		E:      make([]int, config.Steps),
		I:      make([]int, config.Steps),
		R:      make([]int, config.Steps),
	}
}

//Uses Euler's method to compute discrete points on each of the curves
func (sm SeirModel) CalculatePoints() {
	stepSize := float64(sm.Config.Days) / float64(sm.Config.Steps)

	pt := SeirModelPoint{
		sm.Config,
		sm.Config.S,
		sm.Config.E,
		sm.Config.I,
		sm.Config.R,
	}

	for x := 0; x < sm.Config.Steps; x++ {
		pt.S = pt.S + stepSize*dS(pt)
		pt.E = pt.E + stepSize*dE(pt)
		pt.I = pt.I + stepSize*dI(pt)
		pt.R = pt.R + stepSize*dR(pt)

		sm.S[x] = int(pt.S)
		sm.E[x] = int(pt.E)
		sm.I[x] = int(pt.I)
		sm.R[x] = int(pt.R)
	}

}

func dS(pt SeirModelPoint) float64 {
	N := pt.N()
	return pt.Mu*(N-pt.S) - pt.Beta*(pt.S*pt.I/N) - pt.Nu*pt.S
}

func dE(pt SeirModelPoint) float64 {
	N := pt.N()
	return pt.Beta*(pt.S*pt.I/N) - (pt.Mu+pt.Sigma)*pt.E
}

func dI(pt SeirModelPoint) float64 {
	return pt.Sigma*pt.E - (pt.Mu+pt.Gamma)*pt.I
}

func dR(pt SeirModelPoint) float64 {
	return pt.Gamma*pt.I - pt.Mu*pt.R + pt.Nu*pt.S
}
