package affine

import (
	"testing"
)

var (
	points  []Point
	moves   []Move
	pVecs   []Vec
	mVecs   []Vec
	cVecs   []Vec
	affines []Affine
)

func init() {
	for i := 0; i < 10000; i++ {
		p := Point{X: float32(i), Y: float32(i)}
		m := Move{X: float32(i), Y: float32(i)}
		points = append(points, p)
		moves = append(moves, m)
		pVecs = append(pVecs, p)
		mVecs = append(mVecs, m)
		cVecs = append(cVecs, p.Add(m))
		affines = append(affines, Affine{A: 2, B: 2, E: float32(i), F: float32(i)})
	}
}

func BenchmarkAffineTypeSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for idx := range points {
			_, _ = affines[idx].PVec(points[idx])
			_, _ = affines[idx].PVec(moves[idx])
			_, _ = affines[idx].PVec(points[idx].Add(moves[idx]))
		}
	}
}

func BenchmarkAffineInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for idx := range points {
			_, _ = affines[idx].PVecInterface(points[idx])
			_, _ = affines[idx].PVecInterface(moves[idx])
			_, _ = affines[idx].PVecInterface(points[idx].Add(moves[idx]))
		}
	}
}

func BenchmarkAffineTypeSwitchPre(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for idx := range points {
			_, _ = affines[idx].PVec(pVecs[idx])
			_, _ = affines[idx].PVec(mVecs[idx])
			// do the computations, then use the precomputed interface
			_ = points[idx].Add(moves[idx])
			_, _ = affines[idx].PVec(cVecs[idx])
		}
	}
}

func BenchmarkAffineInterfacePre(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for idx := range points {
			_, _ = affines[idx].PVecInterface(pVecs[idx])
			_, _ = affines[idx].PVecInterface(mVecs[idx])
			// do the computations, then use the precomputed interface
			_ = points[idx].Add(moves[idx])
			_, _ = affines[idx].PVecInterface(cVecs[idx])
		}
	}
}

func BenchmarkAffineConcrete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for idx := range points {
			_, _ = points[idx].Project(affines[idx])
			_, _ = moves[idx].Project(affines[idx])
			_, _ = points[idx].Add(moves[idx]).Project(affines[idx])
		}
	}
}
