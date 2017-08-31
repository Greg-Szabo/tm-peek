package io

import (
	"io/ioutil"
	"strings"
	"strconv"
	"time"
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
	contents, err := ioutil.ReadFile("/proc/diskstats")
	if err != nil {
		contents = []byte("0 0 xvda 0 0 0 0 0 0 0 0 0 0 0")
	}
	lines := strings.Split(string(contents), "\n")
	line := lines[0]
	fields := strings.Fields(line)
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
