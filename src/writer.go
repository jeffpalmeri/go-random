package main

import "os"

func main3() {
	f, err := os.OpenFile("created.txt",  os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Write([]byte("Hello"))
}

