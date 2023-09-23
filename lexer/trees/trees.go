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
