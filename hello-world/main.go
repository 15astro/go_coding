package main

import (
 "fmt"
 "os"
 "flag"
)

func greet(l string) string{
   switch l{
   case "english":
     return "This is english"
   case "hindi":
     return "This is hindi"
   case "marathi":
     return "This is Marathi"
}
    return "Hello World!"

}

func main(){
/*   msgEn := greet("english")
   fmt.Println(msgEn)
   msgHindi := greet("hindi")
   fmt.Println(msgHindi)
   msgMarathi := greet("marathi")
   fmt.Println(msgMarathi)
*/
    var lang string
    flag.StringVar(&lang, "lang", "", "Specify language")
    flag.Parse()

   if flag.NFlag() == 0 {
   flag.Usage()
   os.Exit(1)
  }
   fmt.Println(greet(lang))
}
