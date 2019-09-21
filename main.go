package main

import ( 
	"fmt"
	profile "github/biplabpokhrel/helloworld/profile"
	name "github/biplabpokhrel/helloworld/profile/name"
)

func main() {
	fmt.Println(profile.Display())
	fmt.Println(name.FullName())
}