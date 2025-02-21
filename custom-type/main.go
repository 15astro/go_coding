package main

import (
 "fmt"
 "os"
 "flag"
)

type language string

var greetings = map[language]string{
"english": "This is english",
"hindi": "This is hindi",
"marathi": "This is marathi",
}

func greet(l language) (string, error){
   msg, ok := greetings[l]
   if !ok {
      return "", fmt.Errorf("Unsupported language: %v", l)
   }
   return msg, nil
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
   msg, err := greet(language(lang))
   if err != nil {
   fmt.Printf("Something went wrong: %v", err)
   os.Exit(1)
  }
  fmt.Println(msg)
}
