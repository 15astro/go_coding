package main

import (
 "fmt"
 "os"
 "lang"
 "net"
)

fun lookUp(rType string, hostName string) ([]string,error) {
  records, err := net.LookupHost(hostName)
  if err != nil {
   return [], fmt.Errorf("Unable to resolv: %v type %v", hostname, rType)
   }
   return records, nil
 
}

func main(){

  var rType string 
  var hostName string
  
  flag.StringVar(&rType, "type", "A", "Specify record type")
  flag.StringVar(&hostName, "hostname", "", "Specify host")

  flag.Parse()
  
  if flag.NFlag() == 0 {
  flag.Usage()
  os.Exit(1)
}
  fmt.Println("type", rType)
  fmt.Println("hostname", hostname)
  //

}
