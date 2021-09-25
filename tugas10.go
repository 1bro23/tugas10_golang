package main

import "fmt"
import "time"

func main(){
  fmt.Println("AKu")
  timer()
  fmt.Println("Memiliki Hobi Typing Keyboard")
  timer()
  fmt.Println("Dan lahir pada 11-11-2011")
}

func timer(){
  for i:=1;i<6;i++{
    time.Sleep(time.Second*1)
    fmt.Print(i)
  }
}
