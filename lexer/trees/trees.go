package trees

type Node struct{
	child map[rune]*Node
	final bool	
}

func (root *Node) Add(sequence []rune)  {
	if root.child == nil{
		root.child = make(map[rune]*Node)
	}
	switch{
	case len(sequence) > 0:
		if root.child[sequence[0]] == nil{
			root.child[sequence[0]] = new(Node)
		}
		root.child[sequence[0]].Add(sequence[1:])
	default:
		root.final = true
		return
	}
}

func (root *Node) Contains(sequence []rune) bool {
	if root != nil {
		switch {
		case len(sequence) > 1:
			switch {
			case root.child[sequence[0]] != nil:
				return root.child[sequence[0]].Contains(sequence[1:])
			default:
				return false
			}
		case len(sequence) == 1:
			switch {
			case root.child[sequence[0]] != nil:
				return true
			default:
				return false
			}
		default:
			return false
		}
	}
	return false
}

func (root *Node) Is_complete(sequence []rune) bool {
	if root != nil {
		switch {
		case len(sequence) > 1:
			switch {
			case root.child[sequence[0]] != nil:
				return root.child[sequence[0]].Is_complete(sequence[1:])
			default:
				return false
			}
		case len(sequence) == 1:
			switch {
			case root.child[sequence[0]] != nil:
				return root.child[sequence[0]].final
			default:
				return false
			}
		default:
			return false
		}
	} else {
		return false
	}
}

func Get_numbers_tree()*Node{
	root :=new(Node)
	root.child=make(map[rune]*Node,0)
	root.child['0']=new(Node)
	//root.child['.']=root.child['0']
	for i:='1'; i <= '9'; i++ {
		root.child[i] = root.child['0']
	}
	root.child['0'].final = true
	root.child['0'].child = make(map[rune]*Node)
	root.child['0'].child['0'] = root.child['0']
	for i := '1'; i <= '9'; i++ {
		root.child['0'].child[i] = root.child['0']
	}
	root.child['0'].child['E'] = new(Node)
	root.child['0'].child['E'].final = true
	root.child['0'].child['E'].child = make(map[rune]*Node)
	root.child['0'].child['.'] = new(Node)
	root.child['0'].child['.'].final = true
	root.child['0'].child['.'].child = make(map[rune]*Node)

	root.child['0'].child['E'].child['0'] = root.child['0'].child['E']
	root.child['0'].child['.'].child['0'] = root.child['0'].child['.']
	for i := '1'; i <= '9'; i++ {
		root.child['0'].child['E'].child[i] = root.child['0'].child['E']
		root.child['0'].child['.'].child[i] = root.child['0'].child['.']
	}
	root.child['0'].child['.'].child['E'] = root.child['0'].child['E']

	root.child['0'].child['E'].child['.'] = new(Node)
	root.child['0'].child['E'].child['.'].final = true
	root.child['0'].child['E'].child['.'].child = make(map[rune]*Node)
	root.child['0'].child['E'].child['.'].child['0'] = root.child['0'].child['E'].child['.']
	for i := '1'; i <= '9'; i++ {
		root.child['0'].child['E'].child['.'].child[i] = root.child['0'].child['E'].child['.']
	}
	return root
}

func Get_white_spaces() *Node {
	root := new(Node)
	root.child = make(map[rune]*Node)
	root.child[' '] = new(Node)
	root.child[' '].final = true
	root.child['\n'] = new(Node)
	root.child['\n'].final = true
	root.child['\t'] = new(Node)
	root.child['\t'].final = true
	root.child[13] = new(Node)
	root.child[13].final = true
	return root
}
