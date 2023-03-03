package main

import (
	"flag"

	evdev "github.com/g4291/goevdev"
	l "github.com/g4291/gologger"
)

var (
	params struct {
		device string
		log    string
	}
)

func init() {
	flag.StringVar(&params.device, "device", "", "input device path")
	flag.StringVar(&params.log, "log", "", "log file path")
	flag.Parse()

	if params.device == "" {
		l.Fatal("bad input device, use -device device_path")
	}

	if params.log != "" {
		l.SetOutput(
			l.FileOutput(params.log),
			l.StdOutput(),
		)
	}
}

func main() {

	reader, err := evdev.NewReader(params.device)
	l.ErrorCheck(err, true, nil)
	defer reader.Close()

	for {
		event, _ := reader.ReadEvent()
		l.Info("Event:", "\nTime: ", event.Timestamp, "\nType: ", event.Typ, "\nCode: ", event.Code, "\nValue:", event.Value)
	}

}
