package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

var msgsToSendCh = make(chan string)
var msgsToRecvCh = make(chan string)
var stop = false
var waitingToSendMsg = false
var waitingToRecvMsg = false

// Helper method that will panic on an error, or print a successful message and continue
func check(err error, message string) {
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", message)
}

func connectToServer() (c net.Conn) {
	CONNECT := "127.0.0.1:" + "5555"
	c, err := net.Dial("tcp", CONNECT)
	check(err, "Connected to server.")
	return c
}

//Will change timings later to synchronize better to take into account sending and receiving at the same time
func recvMsgs(conn net.Conn) {
	fmt.Println("Starting recvMsgs goroutine...")

	waitingToRecvMsg = true
	message, _ := bufio.NewReader(conn).ReadString('\n')
	msgsToRecvCh <- message
	time.Sleep(time.Second * 5)
	waitingToRecvMsg = false

	fmt.Println("End of recvMsgs goroutine")
}

func sendMsgs() {
	fmt.Println("Starting sendMsgs goroutine...")

	reader := bufio.NewReader(os.Stdin)
	waitingToSendMsg = true
	fmt.Print(">> ")
	userInput, _ := reader.ReadString('\n')
	msgsToSendCh <- userInput
	time.Sleep(time.Second * 3)
	waitingToSendMsg = false

	fmt.Println("End of recvMsgs goroutine")
}

func checkForMsgs(conn net.Conn) {
	for {
		if waitingToSendMsg {
			fmt.Println("Waiting to send msg...")
			time.Sleep(time.Second * 2)
		} else {
			go sendMsgs()
		}
		if waitingToRecvMsg {
			fmt.Println("Waiting to recv msg...")
			time.Sleep(time.Second * 2)
		} else {
			go recvMsgs(conn)
		}
		fmt.Println("Sleeping for 6s...")
		time.Sleep(time.Second * 6)
	}
}

func InitializeClient() {
	conn := connectToServer()

	go sendMsgs()
	go recvMsgs(conn)
	time.Sleep(time.Second * 2) // Have to wait because the timing of when boolean changes in sendMsgs and recv and timing of checkForMsgs
	go checkForMsgs(conn)
	i := 1
	for {
		fmt.Println("Loop counter: ", i)

		select {
		case msgToSend := <-msgsToSendCh:
			fmt.Println("msgToSend case")
			fmt.Fprintf(conn, msgToSend+"\n")
			if msgToSend == "STOP" {
				stop = true
			}
		case msgToRecv := <-msgsToRecvCh:
			fmt.Println("msgToRecv case")
			fmt.Println(msgToRecv)
		}

		time.Sleep(time.Second * 2)
		if stop == true {
			break
		}
		fmt.Println("end of for loop. counter = ", i)
		i++
	}

}
