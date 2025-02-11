package real

import "strconv"

const (
	MaxReciprocalIterations = 100
)

func (x *Real) Reciprocal() *Real {
	xscaled := x.Copy()
	xscaled.exponent = 0

	s := xscaled.String()
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic("could not parse float")
	}
	f = 1 / f
	z0 := NewFloat64(f)

	z := z0
	two := NewInt64(2)

	for i := 0; i < MaxReciprocalIterations; i++ {
		zn := z.Mul(two.Sub(xscaled.Mul(z)))
		if zn.Compare(z) == 0 {
			z = zn
			break
		}
		z = zn
	}

	z.exponent += x.exponent * -1
	z.trim()
	return z
}
