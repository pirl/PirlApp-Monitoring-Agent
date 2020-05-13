package agent_pirlapp

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)


func GetAllStats() {
	v, _ := mem.VirtualMemory()


}
