# GoEvdev

reading input events from devices mouse, keyboard, ...

example

```golang
package main

import (
	evdev "github.com/g4291/goevdev"
)

func main() {
	reader, err := evdev.NewReader("/dev/input/event15")
    if err != nil {
        panic(err)
    }
	defer reader.Close()

	for {
		event, _ := reader.ReadEvent([]uint16{evdev.EV_SYN, evdev.EV_MSC})
		fmt.Println("Event:", "\nTime: ", event.Timestamp, "\nType: ", event.Typ, "\nCode: ", event.Code, "\nValue:", event.Value)
        fmt.Println()
	}

}
```
