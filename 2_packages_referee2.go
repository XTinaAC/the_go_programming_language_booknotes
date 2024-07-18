package xtina_ac 

import (
	"fmt"
)

func PrintExportedConst() {
	fmt.Println("---4---", ExportedConst)
	fmt.Println("---5---", notExportedConst)
}