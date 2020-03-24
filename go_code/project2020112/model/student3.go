package model
//定义一个结构体，首字母小写，并且字段score也小写
type student struct{
	Name string
	score float64
}
//因为结构体首字母小写，所以使用工厂模式来解决
func NewStudent1(n string,s float64)*student{
	return &student{
		Name:n,
		score:s,
	}
}
//如果score字段首字母小写，则其他包中不可以直接使用，我们可以提供一个方法
func (s *student)GetScore()float64{
	return s.score
}