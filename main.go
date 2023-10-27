package main

import (
	"github.com/RaphAlmeida/GoRottenTomato/module"
	"os"
)

func main()  {
	module.Parse(os.Args[1:])
}


