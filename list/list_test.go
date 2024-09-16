package list

import (
	"slices"
	"testing"
)

func TestNew(t *testing.T) {
	expected := []int{1, 2, 3, 10, -12, 8}
	ls := New(expected...)
	res := []int{}

	for n := ls.Head(); n != nil; n = n.Next() {
		res = append(res, n.Value)
	}

	if !slices.Equal(res, expected) {
		t.Fatalf("Expected: %v. Got: %v.", expected, res)
	}
}

func TestAdd(t *testing.T) {
	expected := []int{1, 2, 3, 10, -12, 8}
	ls := New[int]()
	res := []int{}

	for _, v := range expected {
		ls.Add(v)
	}

	for n := ls.Head(); n != nil; n = n.Next() {
		res = append(res, n.Value)
	}

	if !slices.Equal(expected, res) {
		t.Fatalf("Expected: %v. Got: %v.", expected, res)
	}
}

func TestAt(t *testing.T) {
	expected := []int{1, 2, 3, 10, -12, 8}
	ls := New(expected...)

	//Positive index
	for i, v := range expected {
		if res := ls.At(i).Value; res != v {
			t.Logf("Index: %d. Expected: %v. Got: %v.", i, v, res)
			t.Fail()
		}
	}

	//Negative index
	for i, v := range expected {
		j := i - ls.Len()
		if res := ls.At(j).Value; res != v {
			t.Logf("Index: %d. Expected: %v. Got: %v.", j, v, res)
			t.Fail()
		}
	}

	//Index > list's length
	for i, v := range expected {
		j := i + ls.Len()
		if res := ls.At(j).Value; res != v {
			t.Logf("Index: %d. Expected: %v. Got: %v.", j, v, res)
			t.Fail()
		}
	}
}

func TestRemove(t *testing.T) {
	sl := []int{1, 2, 3, 4, 5, 6, 7}
	//For deleting whole collection
	ls1 := New(sl...)
	//For deleting a few items
	ls2 := New(sl...)
	expected := New(2, 3, 5, 6)

	for n := ls1.Head(); n != nil; n = ls1.Head() {
		ls1.Remove(n)
	}

	for n := ls1.Head(); n != nil; n = n.Next() {
		t.Fatal("Could not delete list's items")
	}

	//Remove middle item
	ls2.Remove(ls2.At(ls2.Len() / 2))
	//Remove first item
	ls2.Remove(ls2.Head())
	//Remove last item
	ls2.Remove(ls2.Tail())

	for i := range ls2.Len() {
		if expected.At(i).Value != ls2.At(i).Value {
			t.Logf("Index: %d. Expected: %v. Got: %v.", i, expected.At(i).Value, ls2.At(i).Value)
			t.Fail()
		}
	}

}
