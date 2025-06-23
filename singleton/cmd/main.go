package main

import "github.com/ThalesMonteir0/gof_design_partners/singleton"

func main() {
	for i := 0; i < 10; i++ {
		singleton.InitDataBase()
	}
}
