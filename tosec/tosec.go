package tosec

import "tosec-manager/tree"

// ShowFileTree displays the file tree structure for the given path
func ShowFileTree(path string) error {
	return tree.Display(path, []string{".go"})
}

