package list

import (
	"math/rand"
	"testing"
	"time"
)

const size = 512

func BenchmarkAdd(b *testing.B) {
	rand.New(rand.NewSource(time.Now().Unix()))

	b.ResetTimer()
	for range b.N {
		b.StopTimer()
		ls := New[int]()
		source := make([]int, size)

		for i := range source {
			source[i] = rand.Int()
		}

		b.StartTimer()
		for i := 0; i < size; i++ {
			ls.Add(source[i])
		}
	}
}

func BenchmarkAt(b *testing.B) {
	rand.New(rand.NewSource(time.Now().Unix()))

	b.ResetTimer()
	for range b.N {
		b.StopTimer()
		ls := New[int]()

		for range size {
			ls.Add(rand.Int())
		}

		b.StartTimer()
		for i := 0; i < size; i++ {
			ls.At(i)
		}
	}
}

func BenchmarkRemove(b *testing.B) {
	rand.New(rand.NewSource(time.Now().Unix()))

	b.ResetTimer()
	for range b.N {
		b.StopTimer()
		ls := New[int]()

		for range size {
			ls.Add(rand.Int())
		}

		b.StartTimer()
		for n := ls.Head(); n != nil; n = ls.Head() {
			ls.Remove(n)
		}
	}
}
