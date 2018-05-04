package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

// SuperCompute ...
func SuperCompute() {
	var amt, i uint64
	amt = 1
	// 10 000 000 000
	for i = 1; i < 10000000000; i++ {
		amt += i // amt = amt + i
	}
	fmt.Println(amt)
}

// SuperHeap ...
func SuperHeap() {

	var strArray [10000]string

	for i := 0; i < 10000; i++ {
		for j := 0; j < 1000; j++ {
			strArray[i] += "eee eee eee eee eee eee eee eee"
		}
	}

	fmt.Println(len(strArray[0]))
}

// SuperEfficientHeap ...
func SuperEfficientHeap() {

	var strArray [10000]bytes.Buffer

	for i := 0; i < 10000; i++ {
		for j := 0; j < 1000; j++ {
			strArray[i].WriteString("eee eee eee eee eee eee eee eee")
		}
	}

	fmt.Println(len(strArray[0].String()))
}

func main() {
	fmt.Println("Hello, world!")

	// Current alloc of heap
	var mem runtime.MemStats
	fmt.Println("Heap At Start")
	runtime.ReadMemStats(&mem)
	fmt.Println(mem.Alloc)
	fmt.Println(mem.TotalAlloc)
	fmt.Println(mem.HeapAlloc)
	fmt.Println(mem.HeapSys)

	// CPU Profile
	cpuprofile := "cpu.prof"
	f, err := os.Create(cpuprofile)
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// Run High Compute
	SuperCompute()

	fmt.Println("Heap After SuperCompute")
	runtime.ReadMemStats(&mem)
	fmt.Println(mem.Alloc)
	fmt.Println(mem.TotalAlloc)
	fmt.Println(mem.HeapAlloc)
	fmt.Println(mem.HeapSys)

	// Run High Heap Use
	SuperEfficientHeap()

	fmt.Println("Heap After SuperEfficientHeap")
	runtime.ReadMemStats(&mem)
	fmt.Println(mem.Alloc)
	fmt.Println(mem.TotalAlloc)
	fmt.Println(mem.HeapAlloc)
	fmt.Println(mem.HeapSys)

	// Run High Heap Use
	SuperHeap()

	fmt.Println("Heap After SuperHeap")
	runtime.ReadMemStats(&mem)
	fmt.Println(mem.Alloc)
	fmt.Println(mem.TotalAlloc)
	fmt.Println(mem.HeapAlloc)
	fmt.Println(mem.HeapSys)

	// Memory (Heap) Profiling
	memprofile := "mem.prof"
	mf, err := os.Create(memprofile)
	if err != nil {
		log.Fatal(err)
	}
	pprof.WriteHeapProfile(mf)
	mf.Close()

}
