package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Creates the file and builds the initial header
func buildHeader() {
	//NESASM 3.1
	//INESPRG - Specifies the number of 16k prg banks.
	//INESCHR - Specifies the number of 8k chr banks.
	//INESMAP - Specifies the NES mapper used.
	//INESMIR - Specifies VRAM mirroring of the banks. Refer to iNES header document (neshdr20.txt).

	//NESASM 3.6 CE

	//Header
	//The first 0-7 bytes are the same between iNES and NES2.0
	h1 := []byte("NES")               //Bytes 0-2 spell NES $4E $45 $53
	h2, err := hex.DecodeString("1A") // Byte 4 is the MSDOS end of file
	check(err)

	f, err := os.Create("test.nes")
	check(err)
	defer f.Close()
	f.Write(h1)
	f.Write(h2)
	f.Sync()
}

func main() {
	//More infomration on flags here https://gobyexample.com/command-line-flags
	sizePtr := flag.Bool("S", false, "Outputs the size of the ROM data")
	fmt.Println("Hello Owrld:", sizePtr)

}
