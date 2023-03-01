package datastructure

// 定义一个栈结构
type Stack struct {
	data []interface{} // 存放数据的切片
	top  int           // 栈顶位置
}

// 创建一个新的栈
func NewStack() *Stack {
	return &Stack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

// 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return s.top < 0
}

// 入栈
func (s *Stack) Push(v interface{}) {
	s.top++
	if s.top > len(s.data)-1 { // 栈满了，需要扩容
		s.data = append(s.data, v)
	} else { // 栈未满，直接赋值
		s.data[s.top] = v
	}
}

// 出栈
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() { // 栈空了，返回nil
		return nil
	}
	v := s.data[s.top] // 取出栈顶元素
	s.top--            // 栈顶位置下移一位
	return v           // 返回出栈元素值
}

// 获取栈顶元素（不出栈）
func (s *Stack) Peek() interface{} {
	if s.IsEmpty() { // 栈空了，返回nil
		return nil
	}
	return s.data[s.top] // 返回栈顶元素值（不出栈）
}
