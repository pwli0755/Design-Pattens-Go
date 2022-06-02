package main

type folder struct {
	children []inode
	name     string
}

func (f *folder) print(indentation string) {
	// fmt.Println(indentation + f.name)
	for _, sub := range f.children {
		sub.print(indentation + f.name + "/")
	}
}

func (f *folder) clone() inode {
	// new_folder := folder{
	// 	children: make([]inode, len(f.children)),
	// 	name:     f.name + "_clone",
	// }

	// copy(new_folder.children, f.children)

	cloneFolder := folder{name: f.name + "_clone"}
	var tempChildren []inode
	for _, i := range f.children {
		copy_children := i.clone()
		tempChildren = append(tempChildren, copy_children)
	}
	cloneFolder.children = tempChildren
	return &cloneFolder
}
