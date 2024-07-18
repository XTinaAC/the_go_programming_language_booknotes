/*
	Like libraries/modules in other languages, packages in Go support:
		modularity, encapsulation, separate compilation, and reuse.

	Usually the files of the xxx_xxx pacakge are stored in directory:
		$GOPATH/src/xxx_xxx
	( in my case it's /usr/local/go/src/xxx_xxx )

	By convention, a package's name matches the last segment of its【import path】
*/
package xtina_ac 

import (
	"fmt"
)

/*
	Use "CamelCase" when forming names by combining words;
	Names beginning with upper-case letters are exported 
	( visible & accessible outside of its own package )
*/
const ExportedConst = "导出"
const notExportedConst = "不导出"

func main() {
	// Letters of acronyms & initialisms are always rendered in the SAME CASE
	const escapeHTML = ""
	const htmlEscape = ""

	// byte对应末尾8位的数值
	var tmp1 uint64 = 0b1100111000101010
	fmt.Println(byte(tmp1), byte(tmp1>>(1*8))) // 等值于 0b00101010, 0b11001110
}