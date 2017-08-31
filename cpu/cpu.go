package cpu

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
	for {
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
		return data{
			dataStore.User + 1,
			dataStore.Nice + 1,
			dataStore.System + 1,
			dataStore.Idle + 1,
			dataStore.IOWait + 1,
			dataStore.IRQ + 1,
			dataStore.SoftIRQ + 1,
			dataStore.Total + 7,
		}
	}
	contents, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range(lines) {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			fielddata := new([8]uint64)
			for i := 1; i < len(fields); i++ {
				val, _ := strconv.ParseUint(fields[i], 10, 64)
				fielddata[0] += val
				fielddata[i] = val
			}
			return data{
				fielddata[1],
				fielddata[2],
				fielddata[3],
				fielddata[4],
				fielddata[5],
				fielddata[6],
				fielddata[7],
				fielddata[0],
			}
		}
	}
	return
}

func getStat(old, new data) (output stat) {
	return stat{
		100. * (float32)(new.User - old.User) / (float32)(new.Total - old.Total),
		100. * (float32)(new.Nice - old.Nice) / (float32)(new.Total - old.Total),
		100. * (float32)(new.System - old.System) / (float32)(new.Total - old.Total),
		100. * (float32)(new.Idle - old.Idle) / (float32)(new.Total - old.Total),
		100. * (float32)(new.IOWait - old.IOWait) / (float32)(new.Total - old.Total),
		100. * (float32)(new.IRQ - old.IRQ) / (float32)(new.Total - old.Total),
		100. * (float32)(new.SoftIRQ - old.SoftIRQ) / (float32)(new.Total - old.Total),
	}
}
