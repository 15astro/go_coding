package main

import "fmt"

func checkValidPalindromes(s string) bool{
	fmt.Println("Using two pointers")
	var startIndex = 0
	var endIndex = len(s) -1  
	for i:=0;i<len(s);i++ {
		if s[startIndex] != s[endIndex]{
			return false
		}
		startIndex += 1
		endIndex -= 1
	}
	return true
}

func main(){
	fmt.Println(checkValidPalindromes("nitin"))
	fmt.Println(checkValidPalindromes("suraj"))
	fmt.Println(checkValidPalindromes("madam"))
	fmt.Println(checkValidPalindromes("madam"))
	
}