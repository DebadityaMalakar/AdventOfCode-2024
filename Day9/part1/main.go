package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Day9_1 struct {
	label      string
	path       string
	files      map[int][2]int
	fileBlocks map[int][2]int
	freeBlocks map[int]int
	maxBlock   int
}

func NewDay9_1() *Day9_1 {
	return &Day9_1{
		label:      "day9:1",
		path:       "test",
		files:      make(map[int][2]int),
		fileBlocks: make(map[int][2]int),
		freeBlocks: make(map[int]int),
		maxBlock:   0,
	}
}

func (d *Day9_1) parseBlocks(data string) {
	isFile := true

	for i := 0; i < len(data); i++ {
		n, _ := strconv.Atoi(string(data[i]))

		if isFile {
			d.files[len(d.files)] = [2]int{d.maxBlock, n}
			d.fileBlocks[d.maxBlock] = [2]int{len(d.fileBlocks), n}
		} else if n != 0 {
			d.freeBlocks[d.maxBlock] = n
		}

		d.maxBlock += n
		isFile = !isFile
	}
}

func (d *Day9_1) moveBlocks(fileId, start, n int) {
	i := 0
	rn := n

	for i < start {
		if fb, exists := d.freeBlocks[i]; exists {
			s := min(fb, rn)
			d.fileBlocks[i] = [2]int{fileId, s}
			delete(d.freeBlocks, i)

			if fb > s {
				d.freeBlocks[i+s] = fb - s
			}

			if s == rn {
				delete(d.fileBlocks, start)
				break
			}

			rn -= s
			d.fileBlocks[start] = [2]int{fileId, rn}
			i += s
		} else {
			i++
		}
	}
}

func (d *Day9_1) moveFiles() {
	for i := d.maxBlock - 1; i >= 0; i-- {
		if f, exists := d.fileBlocks[i]; exists {
			d.moveBlocks(f[0], i, f[1])
		}
	}
}

func (d *Day9_1) checksum() int {
	i := 0
	result := 0

	for i < d.maxBlock {
		if f, exists := d.fileBlocks[i]; exists {
			for j := 0; j < f[1]; j++ {
				result += i * f[0]
				i++
			}
		} else {
			i++
		}
	}

	return result
}

func (d *Day9_1) pipeline() (int, error) {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}

	data := strings.TrimSpace(string(content))
	d.parseBlocks(data)
	d.moveFiles()
	return d.checksum(), nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	day := NewDay9_1()
	result, err := day.pipeline()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Checksum:", result)
}
