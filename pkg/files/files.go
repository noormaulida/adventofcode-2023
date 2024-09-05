package files

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
)

func ReadFiles(filepath string) []string {
	fmt.Println("Opening input file ", filepath)
	_, callingFile, _, ok := runtime.Caller(1)
	if !ok {
		panic("unable to find caller so cannot build path to read file")
	}

	fileInput, err := os.Open(path.Join(path.Dir(callingFile), filepath))
	if err != nil {
		panic(err)
	}

	defer func(fileInput *os.File) {
		err := fileInput.Close()
		if err != nil {
			panic(err)
		}
	}(fileInput)

	scanner := bufio.NewScanner(fileInput)
	scanner.Split(bufio.ScanLines)

	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

	}
	return lines
}
