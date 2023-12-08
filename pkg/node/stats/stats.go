package stats

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

// Get gets the stats.
func Get() (Stats, error) {
	numCPUs, err := cpu.Counts(true)
	if err != nil {
		return Stats{}, err
	}
	mem, err := mem.VirtualMemory()
	if err != nil {
		return Stats{}, err
	}
	disk, err := disk.Usage("/")
	if err != nil {
		return Stats{}, err
	}
	return Stats{
		CPUs: numCPUs,
		MEM:  int64(mem.Total),
		DISK: int64(disk.Total),
	}, nil
}
