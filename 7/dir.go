package seven

import (
	"fmt"
	"strings"
)

const (
	CMD_LIST = "ls"

	CMD_CHANGE_DIR      = "cd"
	CMD_CHANGE_DIR_OUT  = ".."
	CMD_CHANGE_DIR_ROOT = "/"
)

type dir struct {
	Name     string
	Contents map[string]int

	Parent  *dir
	SubDirs []*dir
}

func newDir(name string, parent *dir, subDirs []*dir) *dir {
	return &dir{
		Name:     name,
		Contents: make(map[string]int),
		Parent:   parent,
		SubDirs:  subDirs,
	}
}

func newRootDir() *dir {
	return newDir("/", nil, nil)
}

func (d *dir) createSubDirectory(name string) {
	new := newDir(name, d, nil)

	d.SubDirs = append(d.SubDirs, new)
}

// cd command
func (currentDir *dir) changeDir(dirName string) *dir {

	// check if its a specific command (i.e go to root, or go to parent)
	switch dirName {
	case CMD_CHANGE_DIR_OUT:
		return currentDir.Parent

	// this should go up the directory chain until the root is found
	case CMD_CHANGE_DIR_ROOT:
		dir := currentDir

		for dir.Parent != nil {
			dir = dir.Parent
		}

		return dir

	}

	// otherwise search subdirs
	for _, d := range currentDir.SubDirs {
		if d.Name == dirName {
			return d

		}
	}

	return nil

}

func isCommand(input string) bool {
	return strings.HasPrefix(input, "$")
}

func (d *dir) DFS() {
	if d.Parent != nil {
		fmt.Printf("Name: %s, Parent: %s\n", d.Name, d.Parent.Name)
	} else {
		fmt.Println("-", d.Name)
	}

	var b strings.Builder
	b.WriteString("Files:\n")
	for k := range d.Contents {
		b.WriteString(fmt.Sprintf("  - %s\n", k))
	}

	fmt.Println(b.String())

	for _, child := range d.SubDirs {
		child.DFS()
	}
}

func (d *dir) DirSize(sizes []int) (int, []int) {

	mySize := 0
	for _, v := range d.Contents {
		mySize += v
	}

	childSize := 0
	for _, s := range d.SubDirs {
		var t int
		t, sizes = s.DirSize(sizes)
		childSize += t
	}

	total := mySize + childSize
	sizes = append(sizes, total)

	return total, sizes
}
