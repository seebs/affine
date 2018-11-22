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

func BenchmarkAffineConcreteInline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for idx := range points {
			_, _ = points[idx].ProjectInline(&affines[idx])
			_, _ = moves[idx].ProjectInline(&affines[idx])
			_, _ = points[idx].AddInline(moves[idx]).ProjectInline(&affines[idx])
		}
	}
}

func BenchmarkAffineConcrete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for idx := range points {
			_, _ = points[idx].Project(&affines[idx])
			_, _ = moves[idx].Project(&affines[idx])
			_, _ = points[idx].AddInline(moves[idx]).Project(&affines[idx])
		}
	}
}


func BenchmarkAffineTypeSwitchHand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for idx := range points {
			_, _ = affines[idx].PVecHand(pVecs[idx])
			_, _ = affines[idx].PVecHand(mVecs[idx])
			// do the computations, then use the precomputed interface
			_ = points[idx].AddInline(moves[idx])
			_, _ = affines[idx].PVecHand(cVecs[idx])
		}
	}
}

func BenchmarkAffineTypeSwitchInlined(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for idx := range points {
			_, _ = affines[idx].PVecInline(pVecs[idx])
			_, _ = affines[idx].PVecInline(mVecs[idx])
			// do the computations, then use the precomputed interface
			_ = points[idx].AddInline(moves[idx])
			_, _ = affines[idx].PVecInline(cVecs[idx])
		}
	}
}

func BenchmarkAffineTypeSwitchCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for idx := range points {
			_, _ = affines[idx].PVecCall(pVecs[idx])
			_, _ = affines[idx].PVecCall(mVecs[idx])
			// do the computations, then use the precomputed interface
			_ = points[idx].Add(moves[idx])
			_, _ = affines[idx].PVecCall(cVecs[idx])
		}
	}
}

func BenchmarkAffineInterface(b *testing.B) {
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
