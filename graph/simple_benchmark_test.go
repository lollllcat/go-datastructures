package graph

import (
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkAddEdge(b *testing.B) {
	sgraph := NewSimpleGraph()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			a1 := rand.Int()
			a2 := rand.Int()
			b.ResetTimer()
			sgraph.AddEdge("A"+strconv.Itoa(a1), "B"+strconv.Itoa(a2))
		}
	})
}
