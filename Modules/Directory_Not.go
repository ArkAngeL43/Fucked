package FUCKED

import (
	"fmt"
	"os"
)

func IS_DIR(filename string) bool {
	f, x := os.Stat(filename)
	if x != nil {
		return false // not a file
	} else {
		fmt.Println("--------------------------------------------------")
		fmt.Println("[+] FILE DATA  - IS DIR    |  ", f.IsDir())
		fmt.Println("[+] FILE DATA  - MOD TIME  |  ", f.ModTime())
		fmt.Println("[+] FILE DATA  - FILE MODE |  ", f.Mode())
		fmt.Println("[+] FILE DATA  - FILE NAME |  ", f.Name())
		fmt.Println("[+] FILE DATA  - FILE SIZE |  ", f.Size())
		fmt.Println("[+] FILE DATA  - FILE SYS  |  ", f.Sys())
		fmt.Println("---------------------------------------------------")
		return true // is a file
	}
}
