package cpu

import (
	"io/ioutil"
	"strings"
	"strconv"
	"time"
	"fmt"
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
		fmt.Println(statStore)
		dataStore = newData
		time.Sleep(time.Second)
	}
}

func Stat()(stat) {
	return statStore
}

func getData() (output data) {
	contents, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		contents = []byte("cpu 0 0 0 0 0 0 0 0")
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range(lines) {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			fielddata := make([]uint64,8)
			for i := 1; i < 8; i++ {
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
	div := new.Total - old.Total
	if div == 0 {
		div = 1
	}
	return stat{
		100. * (float32)(new.User - old.User) / (float32)(div),
		100. * (float32)(new.Nice - old.Nice) / (float32)(div),
		100. * (float32)(new.System - old.System) / (float32)(div),
		100. * (float32)(new.Idle - old.Idle) / (float32)(div),
		100. * (float32)(new.IOWait - old.IOWait) / (float32)(div),
		100. * (float32)(new.IRQ - old.IRQ) / (float32)(div),
		100. * (float32)(new.SoftIRQ - old.SoftIRQ) / (float32)(div),
	}
}
