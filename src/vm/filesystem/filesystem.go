package filesystem

import (
	"bufio"
	"os"
)

type Reader struct {
	path   string
	out    chan string
	errors chan error
}

func NewReader(path string, out chan string, errors chan error) *Reader {
	return &Reader{path, out, errors}
}

func (r *Reader) Lines() {
	file, err := os.Open(r.path)
	if err != nil {
		r.errors <- err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		r.out <- scanner.Text()
	}

	close(r.out)
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

func (w *Writer) Lines() {
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

func ReadLines2(path string, out chan string, errors chan error) {
	file, err := os.Open(path)
	if err != nil {
		errors <- err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		out <- scanner.Text()
	}

	close(out)
	errors <- scanner.Err()
}

func WriteLines2(path string, in <-chan string, errors chan error, done chan bool) {
	file, err := os.Create(path)
	if err != nil {
		errors <- err
	}
	defer file.Close()

	for line := range in {
		_, err = file.WriteString(line + "\n")
		if err != nil {
			errors <- err
		}
	}
	done <- true
	close(errors)
	close(done)
}
