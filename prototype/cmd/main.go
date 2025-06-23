package main

import "github.com/ThalesMonteir0/gof_design_partners/prototype"

func main() {
	contact := prototype.NewContact("thales", "88888")
	contact.Print("\t")

	contactCloned := contact.Clone()
	contactCloned.Print("\t")
}
