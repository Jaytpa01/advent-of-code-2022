package seven

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/Jaytpa01/advent-of-code-2022/utils"
)

func B() int {
	// fs := utils.NewFileScanner("./7/sample.txt")
	fs := utils.NewFileScanner("./7/input.txt")
	defer fs.Close()

	root := newRootDir()
	currentDir := root

	for fs.Scan() {
		line := fs.Text()

		if isCommand(line) {
			cmd, arg := parseCommand(line)

			if cmd == CMD_CHANGE_DIR {
				currentDir = currentDir.changeDir(arg)
				continue
			}
			continue
		}

		if strings.HasPrefix(line, "dir") {
			dirName := strings.Split(line, " ")[1]
			currentDir.createSubDirectory(dirName)
			continue
		}

		var fileName string
		var fileSize int
		_, err := fmt.Sscanf(line, "%d %s", &fileSize, &fileName)
		if err != nil {
			log.Fatalf("error parsing line (%s): %v", line, err)
		}

		currentDir.Contents[fileName] = fileSize
	}

	// root.DFS()
	sizes := []int{}
	_, sizes = root.DirSize(sizes)
	sort.Ints(sizes)

	unused := 70000000 - sizes[len(sizes)-1]

	for _, val := range sizes {
		if unused+val >= 30000000 {
			return val
		}
	}

	return -1
}
