
package main

import (
	. "net"
	//. "os"
	. "fmt"
)

var err error

func read(conn Conn){
	data := make([]byte, 1024)
	_, err = conn.Read(data)
	checkError(err,"read")
	msg := string(data[:])
	Println(msg)
}

func write(conn Conn, msg string){
	data := []byte(msg)
	_, err = conn.Write(data)
	checkError(err,"write")
}

func checkError(err error,str string){
		if err != nil{
			Println("Error in "+str)
	}
}

func main(){
	conn, err := Dial("tcp", "129.241.187.161:33546")
	checkError(err,"dial")
	read(conn)
	write(conn, "hei\x00")
	read(conn)

	Println("connect")
	//adr:= conn.LocalAddr()
	lstn, err := Listen("tcp", ":10485")
	checkError(err,"Listen")
	Println(1)
	write(conn,"Connect to: "+"129.421.187.158"+":10485\x00")
	Println(2)
	for{
		con2,err := lstn.Accept()
		checkError(err,"for")
		Println(con2)
		Printf("%d",3)
		if (err != nil){
			break
		}
	}
	Println(4)

	//write(con2,"hei igjen")
	//read(con2)
	
}
