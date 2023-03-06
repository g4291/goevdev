package evdev

import (
	"encoding/binary"
	"os"
	"time"
)

func inSLice(v uint16, s []uint16) bool {
	if s == nil {
		return false
	}

	for _, e := range s {
		if e == v {
			return true
		}
	}

	return false
}

type Event struct {
	Typ       uint16
	Code      uint16
	Timestamp time.Time
	Value     int32
}

func NewEvent(data []byte) (*Event, error) {
	e := &Event{}
	e.Typ = binary.LittleEndian.Uint16(data[16:18])
	e.Code = binary.LittleEndian.Uint16(data[18:20])
	e.Timestamp = time.Unix(
		int64(binary.LittleEndian.Uint64(data[0:8])),
		int64(binary.LittleEndian.Uint64(data[8:16])))
	e.Value = int32(binary.LittleEndian.Uint32(data[20:]))

	return e, nil
}

type Reader struct {
	fh *os.File
}

func NewReader(filename string) (*Reader, error) {
	fh, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	er := &Reader{}
	er.fh = fh
	return er, nil
}

// read one event, skip sync, misc events
func (r *Reader) ReadEvent(skipTypes []uint16) (*Event, error) {
	data := make([]byte, 24)
	r.fh.Read(data)
	e, err := NewEvent(data)
	if err == nil {
		// skip
		if inSLice(e.Typ, skipTypes) {
			// read next
			return r.ReadEvent(skipTypes)
		}
		return e, nil
	}

	return e, err
}

func (r *Reader) Close() {
	// close file handler
	r.fh.Close()
}
