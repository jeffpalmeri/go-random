package main

import (
	"encoding/binary"
	"fmt"
	"os"

	singlylinkedlists "github.com/jeffpalmeri/go-files/src/DataStructures/SinglyLinkedLists"
	dynamicprogramming "github.com/jeffpalmeri/go-files/src/algorithms/DynamicProgramming"
)

type Store struct {
	file *os.File
	name []byte
	size uint64
}

func (s *Store) Append(p []byte) (n uint64, pos uint64, err error) {
	pos = s.size
	ssl := singlylinkedlists.NewList(1)
	fmt.Println(ssl)
	// buf := make([]byte, 10)
	// buf2 := bufio.NewWriter(s.file)

	// Writing the length of the record in bytes
	// if err := binary.Write(s.file, binary.BigEndian, uint64(len(p))); err != nil {
	if err := binary.Write(s.file, binary.BigEndian, p); err != nil {
		fmt.Println("I will find you")
		return 0, 0, err
	}

	// w, err := buf2.Write(p)
	// if err != nil {
	// 	fmt.Println("This write error?")
	// 	return 0, 0, err
	// }

	// fmt.Println("Bytes written: ", w)

	// _, err = s.file.Write([]byte(buf2.Flush().Error()))
	// _, err = s.file.Write([]byte(buf2.Flush().Error()))
	// s.file.Write((buf2))
	// if err != nil {
	// 	fmt.Println("Problem with this write")
		// panic(err)
	// }


	// w += 8
	// s.size += uint64(w)
	// return uint64(w), pos, nil
	return 1,1,nil
}

func main() {
	// thing1 := []int{1,2,3}
	thing1 := make([]int, 3)
	thing2 := thing1
	
	thing2[0] = 333
	thing1[0] = 444
	fmt.Println("thing1: ", thing1)
	fmt.Println("thing2: ", thing2)


	dynamicprogramming.LongestCommonSubsequence("hi", "world")

	// f, err := os.Open("storeFile.txt")
	// f, err := os.OpenFile("storeFile.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, os.ModeAppend)
	f, err := os.OpenFile("storeFile.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("openFile error: ")
		panic(err)
	}
	defer f.Close()
	s := newStore(f, "myStore")

	_, _, err = s.Append([]byte("Hi"))

	if err != nil {
		fmt.Println("Error here?")
		panic(err)
	}

	fmt.Println("Store: \n", s)
}

func newStore(f *os.File, storeName string) *Store {
	fi, err := os.Stat(f.Name())
	if err != nil {
		panic(err)
	}
	size := fi.Size()

	return &Store{
		file: f,
		name: []byte(storeName),
		size: uint64(size),
	}
}

// func main() {
// 	fmt.Println("Another one")

// 	file, err := os.Open("sample.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
	
// 	data := make([]byte, 100)
// 	count, err := file.Read(data)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("read %d bytes: %q\n", count, data[:count])

// }