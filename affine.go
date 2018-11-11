// Package affine is a little sketch of some affine math
// stuff, broken out from the miracle modus because I wanted
// to make a benchmark of a thing I ran into while messing
// with it.
package affine

import (
	"math"
)

// Affine is a trivial affine matrix
// { a, c, e }
// { b, d, f }
type Affine struct {
	A, B, C, D, E, F float32
}

type Vec interface {
	XY() (float32, float32)
	Project(Affine) (float32, float32)
}

type Point struct {
	X, Y float32
}

type Move struct {
	X, Y float32
}

func (p Point) Project(a Affine) (float32, float32) {
	return a.A*p.X + a.C*p.Y + a.E, a.B*p.X + a.D*p.Y + a.F
}

func (m Move) Project(a Affine) (float32, float32) {
	return a.A*m.X + a.C*m.Y, a.B*m.X + a.D*m.Y
}

func (p Point) Add(m Move) Point {
	return Point{X: p.X + m.X, Y: p.Y + m.Y}
}

func (p Point) XY() (float32, float32) { return p.X, p.Y }
func (m Move) XY() (float32, float32) { return m.X, m.Y }

func (a Affine) PVec(v Vec) (float32, float32) {
	switch v := v.(type) {
	case Point:
		return a.A*v.X + a.C*v.Y + a.E, a.B*v.X + a.D*v.Y + a.F
	case Move:
		return a.A*v.X + a.C*v.Y, a.B*v.X + a.D*v.Y
	}
	return 0, 0
}

func (a Affine) PVecInterface(v Vec) (float32, float32) {
	return v.Project(a)
}

// Project applies the affine matrix.
func (a Affine) Project(x0, y0 float32) (x1, y1 float32) {
	return a.A*x0 + a.C*y0 + a.E, a.B*x0 + a.D*y0 + a.F
}

// Unproject reverses projection.
func (a Affine) Unproject(x1, y1 float32) (x0, y0 float32) {
	// subtract translation, multiply by inverse of upper left 2x2
	d := (a.A * a.D) - (a.B * a.C)
	x1, y1 = (x1-a.E)/d, (y1-a.F)/d
	return x1*a.D - y1*a.B, y1*a.A - x1*a.C
}

// Scale scales by X and Y.
func (a *Affine) Scale(xs, ys float32) *Affine {
	a.A, a.C, a.E = a.A*xs, a.C*xs, a.E*xs
	a.B, a.D, a.F = a.B*ys, a.D*ys, a.F*ys
	return a
}

// Rotate rotates by an angle.
func (a *Affine) Rotate(theta float32) *Affine {
	s64, c64 := math.Sincos(float64(theta))
	s, c := float32(s64), float32(c64)
	a.A, a.B, a.C, a.D = a.A*c+a.C*s, a.B*c+a.D*s, a.C*c-a.A*s, a.D*c-a.B*s
	return a
}

// IdentityAffine yields the identity matrix.
func IdentityAffine() Affine {
	return Affine{A: 1, D: 1}
}
