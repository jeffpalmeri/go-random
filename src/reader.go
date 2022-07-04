package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main2() {
	// sentence := []byte("This is my sentence!")
	// fmt.Println(string(sentence[:10]))

	// data, err := os.ReadFile("sample.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))

	// fileReader()
	// fileReaderReadAt()
	stringReader()
	// connectionReader()

  c, err :=	os.Create("created.txt")
	if err != nil {
		panic(err)
	}
	str := []byte("This is a string!")
	c.Write(str)

}

// Reads the file starting at some byte offset
func fileReaderReadAt() {
	f, err := os.Open("sample.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := make([]byte, 10)
	n, err := f.ReadAt(buf, 8)
	if err == io.EOF {
		fmt.Println("EOF Reached")
	}
	if err != nil {
		fmt.Println("An error here")
		panic(err)
	}

	fmt.Printf("Number of bytes read: %v", n)
	fmt.Println(string(buf))
}

// Reads the whole file in chunks
func fileReader() {
	f, err := os.Open("sample.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	readerToStandardOutput(f, 5)
	// buf := make([]byte, 10)
	// for {
	// 	n, err := f.Read(buf)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		break
	// 	}
	// 	if n > 0 {
	// 		fmt.Println(string(buf[:n]))
	// 	} 
	// }
}

func stringReader() {
	s := strings.NewReader("A short but very interesting string")
	readerToStandardOutput(s, 25)
	// buf := make([]byte, 2)

	// for {
	// 	n, err := s.Read(buf)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	if n > 0 {
	// 		fmt.Println(string(buf[:n]))
	// 	}
	// }
}

func connectionReader() {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Fprint(conn, "GET HTTP 1.0\r\n\r\n")
	readerToStandardOutput(conn, 25)

	// buf := make([]byte, 10)
	
	// for {
	// 	n, err := conn.Read(buf)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	if n > 0 {
	// 		fmt.Println(string(buf[:n]))
	// 	}
	// }
}

func readerToStandardOutput(r io.Reader, bufSize int) {
	buf := make([]byte, bufSize)
	
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if n > 0 {
			fmt.Println(string(buf[:n]))
		}
	}
}