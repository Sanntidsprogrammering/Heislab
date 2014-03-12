package main
//heis

import (
	"net"
	."fmt"
	"time"
	"strings"
	"strconv"
	"math/rand"
	//"os/exec"
)

var myIP string
var lastDigitsMyIP int
var myBroadIP string
var lastRegisterdFloor, nextFloor, direction int
var UP,DOWN int

func write(conn net.Conn, msg string){
    data := []byte(msg)
	conn.Write(data)
}

func getMyIP() string{
	con1, _ := net.Dial("tcp", "google.com:http")
	myIP := con1.LocalAddr().String()
	str := strings.Split(myIP,":")
	myIP = str[0]
	con1.Close()
	return myIP
}

func getLastDigitsInIP(IP string) int {
	splitIP := strings.Split(IP,".")
	lastDigitsInIP,_ := strconv.Atoi(splitIP[len(splitIP)-1])
	return lastDigitsInIP
}

func executeOrder(key int){
    var nextFloor,lastRegisterdFloor, stopped int
    rand.Seed( time.Now().UnixNano())
	
	go updateLastRegisterdFloor(key, &lastRegisterdFloor,&stopped)
	go sendNextFloor(key,&nextFloor)
	
	for{
    	if nextFloor == lastRegisterdFloor{
        	time.Sleep(2000*time.Millisecond)
        	nextFloor = rand.Intn(4)
        	Println("master says nextFloor :", nextFloor)
        }
        time.Sleep(1*time.Millisecond)
    }
}

func updateLastRegisterdFloor(key int, lastRegisterdFloor *int, stopped *int){
	
	lstnAddr, _ := net.ResolveUDPAddr("udp",":20"+strconv.Itoa(key+100))
	lstnConn, _ := net.ListenUDP("udp",lstnAddr)
	
	for{
		data := make([]byte, 1024)
		t := time.Now().Add(time.Duration(10000*time.Millisecond))
		lstnConn.SetReadDeadline(t)
		_, _, err := lstnConn.ReadFromUDP(data)
		if err != nil{
			Println("elevator 'key' is down")
			//registerElevatorDown()
			break
		}
		Msg := strings.Split(string(data),"\x00")
		Msg = strings.Split(Msg[0],"#")
		response,_ := strconv.Atoi(Msg[0])
		*stopped,_ = strconv.Atoi(Msg[1])
		if response != *lastRegisterdFloor{
			*lastRegisterdFloor = response
			Println("master hears lastRegisterdFloor :", *lastRegisterdFloor, " stopped : ", *stopped)
		}else{
			*lastRegisterdFloor = response
		}
    }
    
    lstnConn.Close()
}

func sendNextFloor(key int, nextFloor *int){

    broadIP := myBroadIP+":20"+strconv.Itoa(key)
	broadConn,_ := net.Dial("udp",broadIP)	
	
	for{
	    write(broadConn, strconv.Itoa(*nextFloor))
		time.Sleep(1000*time.Millisecond)
	}
	
	broadConn.Close()
}

func main(){
    myIP = getMyIP()
    lastDigitsMyIP = getLastDigitsInIP(myIP)
    key := lastDigitsMyIP
    Println(key)
    executeOrder(key)

}
