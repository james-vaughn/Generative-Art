package seir

type SeirConfig struct {
	Beta  float64
	Gamma float64
	Sigma float64
	Mu    float64
	Nu    float64

	S float64
	E float64
	I float64
	R float64

	Days  int
	Steps int
}

func NewConfig() SeirConfig {
	return SeirConfig{}
}

func (sc SeirConfig) N() float64 {
	return sc.S + sc.E + sc.I + sc.R
}
