package main
//heis

import (
	"net"
	."fmt"
	"time"
	"strings"
	"strconv"
	//"math/rand"
	//"os/exec"
)

var myIP string
var lastDigitsMyIP int
var myBroadIP string
var lastRegisterdFloor, nextFloor, direction, stopped int
var UP,DOWN,YES,NO int

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

func sendLastRegisterdFloorToMaster(broadConn net.Conn){
	for{
    	msg := strconv.Itoa(lastRegisterdFloor)+"#"+strconv.Itoa(stopped)
    	data := []byte(msg)
		broadConn.Write(data)
		time.Sleep(20*time.Millisecond)
	}
}

func updateNextFloor(lstnConn *net.UDPConn){
	for{
		data := make([]byte, 1024)
		lstnConn.ReadFromUDP(data)
		Msg := strings.Split(string(data),"\x00")
		response,_ := strconv.Atoi(Msg[0])
		if response != nextFloor{
			nextFloor = response
			Println("elevator hears nextFloor : ",nextFloor)
		}
    }
}

func goToNextFloor(){
	if lastRegisterdFloor == nextFloor {	//kjør heisen til nextFloor
    	stopped = YES //åpne dører og skru av lamper
        Println("stopped : ",stopped)
        time.Sleep(1000*time.Millisecond)
    }else{
       	time.Sleep(1000*time.Millisecond) //kjører til neste etasje
       	lastRegisterdFloor = (lastRegisterdFloor+1)%4
       	stopped = NO
       	Println("elevator says lastRegisterdFloor : ",lastRegisterdFloor, " stopped : ",stopped)
   }
}

func main(){
    UP,DOWN = 1,0
    YES,NO = 1,0
    lastRegisterdFloor,nextFloor = 0,0
    direction = UP
    myBroadIP = "129.241.187.255"
    myIP = getMyIP()
    lastDigitsMyIP = getLastDigitsInIP(myIP)
    time.Sleep(500*time.Millisecond)
    
    lstnAddr, _ := net.ResolveUDPAddr("udp",":20"+strconv.Itoa(lastDigitsMyIP))
	lstnConn, _ := net.ListenUDP("udp",lstnAddr)
	
    broadIP := myBroadIP+":20"+strconv.Itoa(lastDigitsMyIP+100)
	broadConn,_ := net.Dial("udp",broadIP)
	
	go sendLastRegisterdFloorToMaster(broadConn)
	go updateNextFloor(lstnConn)
	
	for{
		goToNextFloor()
	}

}
