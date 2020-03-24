package main
import(
	"fmt"
	"sort"
	"math/rand"
)
//1.声明Hero结构体
type Hero struct{
	Name string
	Age int
}
//2声明一个Hero结构体切片类型
type HeroSlice []Hero
//3实现Interface接口
func (hs HeroSlice)Len() int{
	return len(hs)
}