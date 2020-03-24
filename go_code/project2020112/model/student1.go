package model
//定义一个结构体变量首字母小写的
type student1 struct{
	Name string
	Score float64
}
//因为student1结构体首字母为小写，因此只能在model包中使用，要想在其他包中也能使用，需要使用工厂模式来解决
func NewStudent(n string,s float64) *student1{
	return &student1{
		Name:n,
		Score:s,
	}
}