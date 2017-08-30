package monitor

import (
	"io/ioutil"
	"strings"
	"strconv"
	"time"
	"runtime"
)

type data struct {
	User    uint64            `json:"user"`
	Nice    uint64            `json:"nice"`
	System  uint64        `json:"system"`
	Idle    uint64            `json:"idle"`
	IOWait  uint64        `json:"iowait"`
	IRQ     uint64            `json:"irq"`
	SoftIRQ uint64        `json:"softirq"`
	Total   uint64        `json:"total"`
}

type stat struct {
	User float32	`json:"user"`
	Nice float32	`json:"nice"`
	System float32	`json:"system"`
	Idle float32	`json:"idle"`
	IOWait float32	`json:"iowait"`
	IRQ float32	`json:"irq"`
	SoftIRQ float32	`json:"softirq"`
}

var (
	dataStore data
	statStore stat
)

func init() {
	go startMonitor()
}

func startMonitor() {
	dataStore = getData()
	for true {
		newData := getData()
		statStore = getStat(dataStore, newData)
		dataStore = newData
		time.Sleep(time.Second)
	}
}

func Stat()(stat) {
	return statStore
}

func getData() (output data) {
	if runtime.GOOS == "darwin" {
		output.User = dataStore.User + 1
		output.Nice = dataStore.Nice + 1
		output.System = dataStore.System + 1
		output.Idle = dataStore.Idle + 1
		output.IOWait = dataStore.IOWait + 1
		output.IRQ = dataStore.IRQ + 1
		output.SoftIRQ = dataStore.SoftIRQ + 1
		output.Total = dataStore.Total + 7
		return
	}
	contents, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range(lines) {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, _ := strconv.ParseUint(fields[i], 10, 64)
				switch i {
				default:
				case 1: output.User = val
				case 2: output.Nice = val
				case 3: output.System = val
				case 4: output.Idle = val
				case 5: output.IOWait = val
				case 6: output.IRQ = val
				case 7: output.SoftIRQ = val
				}
				output.Total += val
			}
			return
		}
	}
	return
}

func getStat(old, new data) (output stat) {
	output.User = 100. * (float32)(new.User - old.User) / (float32)(new.Total - old.Total)
	output.Nice = 100. * (float32)(new.Nice - old.Nice) / (float32)(new.Total - old.Total)
	output.System = 100. * (float32)(new.System - old.System) / (float32)(new.Total - old.Total)
	output.Idle = 100. * (float32)(new.Idle - old.Idle) / (float32)(new.Total - old.Total)
	output.IOWait = 100. * (float32)(new.IOWait - old.IOWait) / (float32)(new.Total - old.Total)
	output.IRQ = 100. * (float32)(new.IRQ - old.IRQ) / (float32)(new.Total - old.Total)
	output.SoftIRQ = 100. * (float32)(new.SoftIRQ - old.SoftIRQ) / (float32)(new.Total - old.Total)
	return
}
