package stack

type Stack struct {
	v []int
}

func New() *Stack {
	return &Stack{v: make([]int, 0)}
}

func (s *Stack) Push(num int) {
	s.v = append(s.v, num)
}

func (s *Stack) Pop() int {
	i := len(s.v) - 1
	num := s.v[i]
	s.v = s.v[:i]
	return num
}
