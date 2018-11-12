package main

import (
	"io"
	"os"
	"strconv"
)

/*
// Version 1
type Counter struct{}

func (f *Counter) Count(n uint64) uint64 {
	if n == 0 {
		println(strconv.Itoa(0))
		return 0
	}

	cur := n
	println(strconv.FormatUint(cur, 10))
	return f.Count(n - 1)
}*/

// Version 2
type Counter struct {
	Writer io.Writer
}

func (f *Counter) Count(n uint64) uint64 {
	if n == 0 {
		f.Writer.Write([]byte(strconv.Itoa(0) + "\n"))
		return 0
	}

	cur := n
	f.Writer.Write([]byte(strconv.FormatUint(cur, 10) + "\n"))
	return f.Count(n - 1)
}

func main() {
	// Version 1
	//counter := Counter{}
	//counter.Count(3)

	// Version 2
	//c := Counter{os.Stdout}
	//c.Count(3)

	pipeReader, pipeWriter := io.Pipe()
	defer pipeWriter.Close()
	defer pipeReader.Close()

	counter := Counter{
		Writer: pipeWriter,
	}

	tee := io.TeeReader(pipeReader, file)

	go func() {
		io.Copy(os.Stdout, tee)
	}()

	counter.Count(5)
}
