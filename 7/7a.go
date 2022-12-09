package seven

import (
	"fmt"
	"log"
	"strings"

	"github.com/Jaytpa01/advent-of-code-2022/utils"
)

func A() int {

	// input := utils.OpenInputFile("./7/sample.txt")
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

	total := 0
	for _, val := range sizes {
		if val <= 100000 {
			total += val
		}
	}

	return total
}

func parseCommand(inputCmd string) (string, string) {

	splitCmd := strings.Split(inputCmd, " ")

	cmd := splitCmd[1]

	// if list cmd, return that and no arg
	if cmd == CMD_LIST {
		return CMD_LIST, ""
	}

	if cmd == CMD_CHANGE_DIR {
		return cmd, splitCmd[2]
	}

	return "", ""
}
