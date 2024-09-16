package main

import (
	"flag"
	"fmt"
	"list/m/v2/list"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {
	var cpuprofile, memprofile string
	flag.StringVar(&cpuprofile, "cpuprofile", "", "write cpu profile to file")
	flag.StringVar(&memprofile, "memprofile", "", "write mem profile to file")

	flag.Parse()

	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			log.Fatal("Could not create cpu profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("Could not start cpu profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	var ls = list.New(1, 2, 3, 4, 5)

	flag.Parse()

	ls.ForEach(func(el *int, i int) {
		fmt.Printf("%d ", *el)
	})
	fmt.Println()

	ls2 := ls.Map(func(el, i int) int {
		el = el + i + 1
		return el
	})

	for n := ls2.Head(); n != nil; n = n.Next() {
		fmt.Printf("%d ", n.Value)
	}
	fmt.Println()

	ls.ForEach(func(el *int, i int) {
		*el = *el + i + 2
		fmt.Printf("%d ", *el)
	})

	for n := ls.Tail(); n != nil; n = ls.Tail() {
		ls.Remove(n)
	}

	ls.ForEach(func(el *int, i int) {
		fmt.Print(*el)
	})

	res1 := ls.Reduce(func(acc, el int) int {
		return acc + el
	})
	fmt.Printf("\n%d\n", res1)

	res2 := ls2.Reduce(func(acc, el int) int {
		return acc + el
	})
	fmt.Printf("%d\n", res2)

	if memprofile != "" {
		f, err := os.Create(memprofile)
		if err != nil {
			log.Fatal("Could not create mem profile: ", err)
		}
		defer f.Close()
		runtime.GC()
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("Could not start mem profile: ", err)
		}
	}
}
