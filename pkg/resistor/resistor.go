package resistor

type Marking int

const (
	UnknownMarking Marking = iota
	AxialStd1Band
	AxialStd3Band = iota + 2
	AxialStd4Band
	AxialStd5Band
	AxialStd6Band
	AxialStd7Band
	SMDStandard
	SMDPrecision
	SMDEIA96
	SMDIndustrial
	SMDMilR11
	SMDMilR39008
)

type Format int

const (
	UnknownFormat Format = iota
	Axial
	SMD
)

type Resistor interface {
	Value() (float64, error)
	Tolerance() float64
	TCR() int
	Type() Marking
}
