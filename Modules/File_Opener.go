package FUCKED

import (
	"bufio"
	"log"
	"os"
)

func FUCKED_LOADER(dstf, dst string) {
	f, x := os.Open(dstf)
	if x != nil {
		log.Fatal(x)
	} else {
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			Filepaths = append(Filepaths, scanner.Text())
		}
	}
	for _, k := range Filepaths {
		if IS_DIR(k) {
			Good_Files = append(Good_Files, k)
		} else {
			Bad_Files = append(Bad_Files, k)
		}
	}
	for _, p := range Good_Files {
		Load(p, dst)
	}
}
