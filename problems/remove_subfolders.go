package problems

// https://leetcode.com/problems/remove-sub-folders-from-the-filesystem/

import "strings"

type folder struct {
	Name     string
	Children map[string]*folder
	Final    bool
}

func removeSubfolders(folders []string) []string {
	root := &folder{
		Name:     "",
		Children: make(map[string]*folder),
		Final:    false,
	}

	// fill the prefix tree
	for _, f := range folders {
		split := strings.Split(f[1:], "/")

		head := root
		for _, subFolder := range split {
			child, ok := head.Children[subFolder]
			if !ok {
				child = &folder{Name: head.Name + "/" + subFolder, Children: make(map[string]*folder), Final: false}
				head.Children[subFolder] = child
			}
			head = child
		}
		head.Final = true
	}

	visit := []*folder{root} // BFS over directories tree
	ret := []string{}
	for len(visit) > 0 {
		node := visit[0]
		if node.Final {
			ret = append(ret, node.Name)
			visit = visit[1:]
			continue
		}

		for _, ch := range node.Children {
			visit = append(visit, ch)
		}
		visit = visit[1:]
	}
	return ret
}
