package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Day9_2 struct {
	label      string
	path       string
	files      map[int][2]int
	fileBlocks map[int][2]int
	freeBlocks map[int]int
	maxBlock   int
}

func NewDay9_2() *Day9_2 {
	return &Day9_2{
		label:      "day9:2",
		path:       "test",
		files:      make(map[int][2]int),
		fileBlocks: make(map[int][2]int),
		freeBlocks: make(map[int]int),
		maxBlock:   0,
	}
}

func (d *Day9_2) parseBlocks(data string) {
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

func (d *Day9_2) moveFile(fileId, start, n int) {
	for i := 0; i < start; i++ {
		if fb, exists := d.freeBlocks[i]; exists && fb >= n {
			d.freeBlocks[start] = n
			delete(d.freeBlocks, i)
			if fb > n {
				d.freeBlocks[i+n] = fb - n
			}
			d.files[fileId] = [2]int{i, n}
			delete(d.fileBlocks, start)
			d.fileBlocks[i] = [2]int{fileId, n}
			break
		}
	}
}

func (d *Day9_2) moveFiles() {
	sortedFiles := make([][2]int, 0, len(d.files))
	for id, file := range d.files {
		sortedFiles = append(sortedFiles, [2]int{id, file[0]})
	}

	// Sort in descending order based on the start block
	for i := 0; i < len(sortedFiles)-1; i++ {
		for j := 0; j < len(sortedFiles)-i-1; j++ {
			if sortedFiles[j][1] < sortedFiles[j+1][1] {
				sortedFiles[j], sortedFiles[j+1] = sortedFiles[j+1], sortedFiles[j]
			}
		}
	}

	for _, fileInfo := range sortedFiles {
		id := fileInfo[0]
		d.moveFile(id, d.files[id][0], d.files[id][1])
	}
}

func (d *Day9_2) checksum() int {
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

func (d *Day9_2) pipeline() (int, error) {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return 0, err
	}
	data := strings.TrimSpace(string(content))
	d.parseBlocks(data)
	d.moveFiles()
	return d.checksum(), nil
}

func main() {
	day := NewDay9_2()
	result, err := day.pipeline()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Checksum:", result)
}
