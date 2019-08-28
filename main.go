package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
)

//modbus over tcp/ip demo specifically the direct Input channels, should not take to much to include all other functionality

//RegisterAddr map for bool values, for other functionality map will need some form of byte or binary recording
var RegisterAddr = map[uint16]bool{
	1: true,
	2: true,
	3: true,
	4: true,
	5: true,
	6: true,
}

func main() {
	addr := "192.168.11.22:502"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("err listening on %s :%s", addr, err)
	}
	defer listener.Close()

	connectionChannel := make(chan *net.Conn)

	go connctionListener(listener, connectionChannel)

	for {
		//retrive connection from listener go routine
		conn := <-connectionChannel
		buf := bufio.NewReaderSize(conn, 12)
		packet := make(bytes, 12)
		_, err := buf.Read(packet)
		if err != nil {
			fmt.Printf("error reading from connection interface into memory")
		}
		//command head removes transmission ID, and Protocol ID
		commandHead := packet[0:3]
		commandMsg := packet[6:12]

		//length of commandMsg in bytes repersented in uint16, should not exceed 6 bytes
		msgLength := binary.BigEndian.Uint16(packet[3:4])
		// if msgLength => 6 {     /////TOOL for trouble shooting the message lengths
		// 	fmt.Println("command message is larger than 6 bytes, make buffer larger")
		// }

		// this is a default unit idendtifier for all ADAM-6000 series
		unitNum := []byte{1}

		//low and high address to be read from or written to
		lowAddress := binary.BigEndian.Uint16(commandMsg[2:3])
		highAddress := binary.BigEndian.Uint16(commandMsg[4:5])

		switch funcCode := commandMsg[7:7]; funcCode {
		case 1:
			readAddrState()(lowAddress, highAddress)([]byte)
		case 15:
			writeAddrState(lowAddress, highAddress)([]byte)
		default:
			commandResponse := []byte{}
		}

		//assemble message

	}

}

func connctionListener(net.Listener, chan *net.Conn) (*net.Conn chan, err error) {
	for {
		conn, err := listener.Accept()
		if err == nil {
			outputConn <- conn
		}
	}
}

func writeAddrState(lowAddress, highAddress uint16) (commandResponse []byte) {

}

func readAddrState(lowAddress, highAddress uint16) (commandResponse []byte) {

}
