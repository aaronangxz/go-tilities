package go_tilities

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func (n *TreeNode) PrintNode() {
	if n != nil {
		fmt.Print(n.Val)
	}
	fmt.Println()
}

func TreeBFS(root *TreeNode) []*TreeNode {
	var (
		//Maintain a queue of nodes to visit
		q Queue
		//Output values of BFS
		out []*TreeNode
	)

	//Root will be the first node to visit
	q.Enqueue(root)

	//Continue to visit until queue is empty
	for !q.IsEmpty() {
		//Visit the first node in queue
		node := q.Front().(*TreeNode)
		out = append(out, node)

		//If left node of current node is not empty, add into the queue to visit next
		if node.Left != nil {
			q.Enqueue(node.Left)
		}

		//If right node of current node is not empty, add into the queue to visit next
		if node.Right != nil {
			q.Enqueue(node.Right)
		}

		//Pop the first node in the queue (current node)
		q.Dequeue()
	}
	return out
}

func RecursiveTreeDFS(root *TreeNode) []*TreeNode {
	var visited []*TreeNode
	return recurse(root, visited)
}

func recurse(root *TreeNode, visited []*TreeNode) []*TreeNode {
	visited = append(visited, root)

	if root.Left != nil {
		visited = recurse(root.Left, visited)
	}

	if root.Right != nil {
		visited = recurse(root.Right, visited)
	}
	return visited
}

func IterativeTreeDFS(root *TreeNode) []*TreeNode {
	//Stack to store elements to visit
	var s Stack

	//Store root node first
	s.Push(root)

	//Keep track of visited nodes
	var visited []*TreeNode

	//Continue as long as stack is not empty
	for !s.IsEmpty() {
		//Visit the top element in stack
		node := s.Top().(*TreeNode)
		visited = append(visited, node)

		//pop the last element
		s.Pop()

		//If right element exists, push onto stack
		if node.Right != nil {
			s.Push(node.Right)
		}

		//If left element exists, push onto stack
		if node.Left != nil {
			s.Push(node.Left)
		}

		//Push right -> left because left node needs to be visited first
	}
	return visited
}
