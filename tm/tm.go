package tm

import "net/http"

type data struct {
	UnconfirmedTXs	uint64	`json:"unconfirmedTXs"`
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


func Status(tendermintAddress string)(resp *http.Response) {
	resp, _ = http.Get(tendermintAddress + "/status")
	return resp
}
