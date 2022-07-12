package go_tilities

type Stack []interface{}

//IsEmpty Helper function to check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

//Push the new element onto the stack
func (s *Stack) Push(e interface{}) {
	*s = append(*s, e)
}

//Pop Removes the element on the top of the stack
func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	index := len(*s) - 1
	*s = (*s)[:index]
}

//Top Returns the element on top of the stack
func (s *Stack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	index := len(*s) - 1
	return (*s)[index]
}
