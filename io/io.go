package io

import (
	"io/ioutil"
	"strings"
	"strconv"
	"time"
	"runtime"
)

type data struct {
	ReadIOs			uint64	`json:"readios"`
	ReadMerges		uint64	`json:"readmerges"`
	ReadSectors		uint64	`json:"readsectors"`
	ReadTicks		uint64	`json:"readticks"`
	WriteIOs		uint64	`json:"writeios"`
	WriteMerges		uint64	`json:"writemerges"`
	WriteSectors	uint64	`json:"writesectors"`
	WriteTicks		uint64	`json:"writeticks"`
	InFlight		uint64	`json:"inflight"`
	IOTicks			uint64	`json:"ioticks"`
	TimeInQueue		uint64	`json:"timeinqueue"`
}

var (
	dataStore data
	statStore data
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

func Stat()(data) {
	return statStore
}

func getData() (output data) {
	if runtime.GOOS == "darwin" {
		return data{
			dataStore.ReadIOs + 1,
			dataStore.ReadMerges + 1,
			dataStore.ReadSectors + 1,
			dataStore.ReadTicks + 1,
			dataStore.WriteIOs + 1,
			dataStore.WriteMerges + 1,
			dataStore.WriteSectors + 1,
			dataStore.WriteTicks + 1,
			dataStore.InFlight + 1,
			dataStore.IOTicks + 1,
			dataStore.TimeInQueue + 1,
		}
	}
	contents, err := ioutil.ReadFile("/proc/diskstats")
	if err != nil {
		return
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range(lines) {
		fields := strings.Fields(line)
		if fields[2] == "vda" {
			fielddata := new([11]uint64)
			for i := 3; i < len(fields); i++ {
				val, _ := strconv.ParseUint(fields[i], 10, 64)
				fielddata[i-3] = val
			}
			return data{
				fielddata[0],
				fielddata[1],
				fielddata[2],
				fielddata[3],
				fielddata[4],
				fielddata[5],
				fielddata[6],
				fielddata[7],
				fielddata[8],
				fielddata[9],
				fielddata[10],
			}
		}
	}
	return
}

func getStat(old, new data) (output data) {
	return data{
		new.ReadIOs - old.ReadIOs,
		new.ReadMerges - old.ReadMerges,
		new.ReadSectors - old.ReadSectors,
		new.ReadTicks - old.ReadTicks,
		new.WriteIOs - old.WriteIOs,
		new.WriteMerges - old.WriteMerges,
		new.WriteSectors - old.WriteSectors,
		new.WriteTicks - old.WriteTicks,
		new.InFlight - old.InFlight,
		new.IOTicks - old.IOTicks,
		new.TimeInQueue - old.TimeInQueue,
	}
}
