package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Gets the size of the passed file
func getSize(file string) int64 {
	stat, err := os.Stat(file)
	if err != nil {
		log.Fatal(err)
	}
	return stat.Size()
}

// Copies "num" bytes from the "in" file to the "out" file using "buffer"
func copyBytes(in, out *os.File, num int64, buffer []byte) {
	for num > 0 {
		if num < int64(len(buffer)) {
			buffer = buffer[:num]
		}
		n, err := in.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}
		num -= int64(n)
		_, err = out.Write(buffer)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Cuts a file in parts
// Each part has a size of "cutsize"
// The resulting files are named "{file}.{id}"
func CutFile(file string, cutsize int64) {
	in, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()
	size := getSize(file)
	buffer := make([]byte, 1024*1024)
	var i int64
	fi := 0
	for i < size {
		fmt.Printf("%d/%d\r", i, size)
		f, err := os.Create(file + "." + strconv.Itoa(fi))
		if err != nil {
			log.Fatal(err)
		}
		if size-i < cutsize {
			cutsize = size - i
		}
		copyBytes(in, f, cutsize, buffer)
		i += cutsize
		f.Close()
		fi++
	}
	fmt.Printf("%d/%d\n", i, size)
}

// Merges the parts into 1 file
// Pass the file path without ".id"
func MergeFile(file string) {
	out, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	buffer := make([]byte, 1024)
	fi := 0
	for {
		f, err := os.Open(file + "." + strconv.Itoa(fi))
		if err != nil {
			if os.IsNotExist(err) {
				break
			}
			log.Fatal(err)
		}
		fsize := getSize(file + "." + strconv.Itoa(fi))
		copyBytes(f, out, fsize, buffer)
		f.Close()
		fi++
	}
}

// Removes all partial files
func Clean(file string) {
	fi := 0
	for {
		err := os.Remove(file + "." + strconv.Itoa(fi))
		if err != nil {
			if os.IsNotExist(err) {
				break
			}
			log.Fatal(err)
		}
		fi++
	}
}

func main() {
	gb := flag.Int("gb", 0, "Size in Gb")
	mb := flag.Int("mb", 0, "Size in Mb")
	kb := flag.Int("kb", 0, "Size in Kb")
	b := flag.Int("b", 0, "Size in byte")
	file := flag.String("f", "", "Filename (without .number when merging)")
	merge := flag.Bool("m", false, "Merge parts")
	clean := flag.Bool("c", false, "Clean all file parts")
	flag.Parse()
	if *file == "" {
		log.Fatalln("No filename given")
	}
	if *merge {
		MergeFile(*file)
		return
	} else if *clean {
		Clean(*file)
		return
	}
	var cutsize int64
	cutsize += int64(*b)
	cutsize += int64(*kb * 1024)
	cutsize += int64(*mb * 1024 * 1024)
	cutsize += int64(*gb * 1024 * 1024 * 1024)
	if cutsize == 0 {
		log.Fatalln("Enter desired filesize")
	}
	CutFile(*file, cutsize)
}
