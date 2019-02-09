package main

import (
	"reflect"
	"testing"
)

func TestParseResponseReturnsProperData(t *testing.T) {
	response := []byte{28, 3, 3, 237, 0, 0, 3, 128, 0, 0, 15, 174, 193, 106, 216, 30, 224, 9, 165, 224, 185, 135, 206,
		26, 0, 0, 0, 0, 0, 0, 0, 0, 224, 9, 168, 224, 2, 173, 41, 84, 224, 9, 168, 224, 2, 174, 42, 221}

	seconds, fraction := parseResponse(response)

	if !reflect.DeepEqual(seconds, []byte{224, 9, 168, 224}) {
		t.Error("Wrong seconds value")
	}

	if !reflect.DeepEqual(fraction, []byte{2, 174, 42, 221}) {
		t.Error("Wrong fraction value")
	}
}

func TestGetParsedTimestampsReturnsProperData(t *testing.T) {
	timestamp, timestampNano := getParsedTimestamps([]byte{224, 9, 168, 224}, []byte{2, 174, 42, 221})

	if timestamp != 1549740640 {
		t.Error("Wrong timestamp value")
	}

	if timestampNano != 44968669 {
		t.Error("Wrong timestampNano value")
	}
}
