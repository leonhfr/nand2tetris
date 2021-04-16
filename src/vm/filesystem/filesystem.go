package filesystem

import (
	"bufio"
	"os"
	"path"
	"path/filepath"
)

type Reader struct {
	path      string
	directory bool
	out       chan string
	errors    chan error
}

func NewReader(path string, directory bool, out chan string, errors chan error) *Reader {
	return &Reader{path, directory, out, errors}
}

func (r *Reader) Read() {
	defer close(r.out)

	if !r.directory {
		r.readFile(r.path)
		return
	}

	var files []string
	err := filepath.Walk(r.path, func(file string, info os.FileInfo, err error) error {
		if path.Ext(file) == ".vm" {
			files = append(files, file)
		}
		return nil
	})
	if err != nil {
		r.errors <- err
	}

	for _, path := range files {
		r.readFile(path)
	}
}

func (r *Reader) readFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		r.errors <- err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		r.out <- scanner.Text()
	}

	r.errors <- scanner.Err()
}

type Writer struct {
	path   string
	in     <-chan string
	errors chan error
	done   chan bool
}

func NewWriter(path string, in <-chan string, errors chan error, done chan bool) *Writer {
	return &Writer{path, in, errors, done}
}

func (w *Writer) Write() {
	file, err := os.Create(w.path)
	if err != nil {
		w.errors <- err
	}
	defer file.Close()

	for line := range w.in {
		_, err = file.WriteString(line + "\n")
		if err != nil {
			w.errors <- err
		}
	}
	w.done <- true
	close(w.errors)
	close(w.done)
}
