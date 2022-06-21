package prototypefactory

import "math"

const (
	Simple = iota
	Compound
)

type simpleInterest struct {
	Principal      float64
	RateOfInterest float64
	Time           int
}

type compoundInterest struct {
	Principal      float64
	RateOfInterest float64
	Time           int
}

// Interface
type InterestCalculator interface {
	CalculateInerest() float64
}

func (si *simpleInterest) CalculateInerest() float64 {
	return float64(si.Principal) * float64(si.Time) / 365 * float64(si.RateOfInterest) / 100
}

func (si *compoundInterest) CalculateInerest() float64 {
	return si.Principal * float64(math.Pow(float64(1+si.RateOfInterest/1200), float64(12*si.Time/365)))
}

func NewCalculator(kind int, principal float64, time int, rate float64) InterestCalculator {
	switch kind {
	case Simple:
		return &simpleInterest{Principal: principal, Time: time, RateOfInterest: rate}
	case Compound:
		return &compoundInterest{Principal: principal, Time: time, RateOfInterest: rate}
	default:
		panic("")
	}

}
