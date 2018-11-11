package affine

import (
	"testing"
)

var (
	points  []Point
	moves   []Move
	affines []Affine
)

func init() {
	for i := 0; i < 10000; i++ {
		points = append(points, Point{X: float32(i), Y: float32(i)})
		moves = append(moves, Move{X: float32(i), Y: float32(i)})
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

func BenchmarkAffineConcrete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for idx := range points {
			_, _ = points[idx].Project(affines[idx])
			_, _ = moves[idx].Project(affines[idx])
			_, _ = points[idx].Add(moves[idx]).Project(affines[idx])
		}
	}
}
