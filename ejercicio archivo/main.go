package main

import (
	"fmt"
	"os"

)
/*
func createFile(p string) *os.File{
	fmt.Println("creando")
	f,err :=os.Create("ola.txt")
	if err!= nil {
		panic(err)

	}
return f
}
func writeFile(f *os.File){
fmt.Println("escribiendo")
fmt.Fprintf(f,"fdnbosdn sdfosfbknofb jsd")

}
func CloseFile(f *os.File){
	fmt.Println("cerrando")
	err:=f.Close()
	if err!=nil{
		fmt.Fprintf(os.Stderr,"error:%v\n",err)
		os.Exit(1)
	}
}*/
func main() {
	f,_ :=os.Create("copypaste.txt")

	defer f.Close()
	
	fmt.Fprintf(f,"hola")
	

}
