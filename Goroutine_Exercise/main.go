package main

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

type CPUMonitor struct {
}

func (m *CPUMonitor) CPUMonitor(ctx context.Context) string {
	percent, err := cpu.PercentWithContext(ctx, 1*time.Second, false)

	if err != nil {
		return "N/A"
	}

	value := fmt.Sprintf("%2.f%%", percent[0])
	return value
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	cpuMornitor := CPUMonitor{}
	cpuPercent := cpuMornitor.CPUMonitor(ctx)

	fmt.Println(cpuPercent)
}
