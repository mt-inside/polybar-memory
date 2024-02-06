package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/shirou/gopsutil/mem"
)

var width = float64(16) // TODO: arg

func main() {

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		vm, err := mem.VirtualMemoryWithContext(ctx)
		cancel()
		if err != nil {
			continue
		}
		ram := float64(vm.Total)
		fmt.Print(strings.Repeat("█", int((float64(vm.Used)/ram)*width+0.5))) // total - free - buffers - cached
		fmt.Print(strings.Repeat("▓", int((float64(vm.Buffers)/ram)*width+0.5)))
		fmt.Print(strings.Repeat("░", int((float64(vm.Cached)/ram)*width+0.5)))
		fmt.Print(strings.Repeat("▁", int((float64(vm.Free)/ram)*width+0.5)))
		fmt.Println()

		time.Sleep(1 * time.Second)
	}
}
