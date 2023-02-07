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
	//TODO

	//Header
	//The first 0-7 bytes are the same between iNES and NES2.0
	h1 := []byte("NES")               //Bytes 0-2 spell NES $4E $45 $53
	h2, err := hex.DecodeString("1A") // Byte 3 is the MSDOS end of file
	check(err)
	h3 := []byte{0b00000000} // Byte 4 is the Size of PRG ROM in 16 KB units
	h4 := []byte{0b00000000} // Byte 5 is the Size of CHR ROM in 8 KB units (value 0 means the board uses CHR RAM)
	h5 := []byte{0b00000000} // Byte 6 Mapper(4b), MultiScreen, trainer, battery, Mirroring (0000,0,0,0,0)

	//For now flags 7-15 unused
	h6 := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	header := append(h1[:], h2[:]...)
	header = append(header[:], h3[:]...)
	header = append(header[:], h4[:]...)
	header = append(header[:], h5[:]...)
	header = append(header[:], h6[:]...)

	f, err := os.Create("test.nes")
	check(err)
	defer f.Close()
	f.Write(header)

	f.Sync()
}

func main() {
	//More infomration on flags here https://gobyexample.com/command-line-flags
	sizePtr := flag.Bool("S", false, "Outputs the size of the ROM data")
	buildHeader()
	fmt.Println("Hello Owrld:", sizePtr)

}
