package main
import (
	"fmt"
	"go_code/project2020122xianggmu/xiangmu2/service"
	"go_code/project2020122xianggmu/xiangmu2/model"
)
type customerView struct{
	//定义必要字段
	key string //接收用户输入...
	loop bool //表示是否循环的显示主菜单
	//增加一个字段customerService
	customerService *service.CustomerService

}
//显示所有的客户信息
func (this *customerView)list(){
	//首先获取到当前所有的客户信息(在切片中)
	customers:=this.customerService.List()
	//显示
	fmt.Println("-----客户列表-----")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i:=0;i<len(customers);i++{
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n-----客户列表完成-----\n\n")
}
//得到用户的输入，信息构建的客户，并完成添加
func(this *customerView)add(){
	fmt.Println("-----添加客户-----")
	fmt.Println("姓名：")
	name :=""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender :=""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age:=0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone:=""
	fmt.Println(&phone)
	fmt.Scanln("电邮：")
	email:=""
	fmt.Scanln(&email)
	//构建一个新的Customer实例
	//注意：id号，没有让用户输入，id是唯一的，需要系统进行分配
	customer:=model.NewCustomer2(name,gender,age,phone,email)
	if this.customerService.Add(customer){
		fmt.Println("-----添加完成-----")
	}else{
		fmt.Println("-----添加失败-----")
	}

}
//得到用户的输入id,删除该id对应的客户
func(this *customerView)delete(){
	fmt.Println("-----删除客户-----")
	fmt.Println("请选择删除客户编号")
	id:=-1
	fmt.Scanln(&id)
	if id==-1{
		return //放弃删除操作
	
	}
	fmt.Println("确认是否删除（y/N）:")
	choice :=""
	fmt.Scanln(&choice)
	if choice=="y"||choice=="Y"{
		//调用customerService的Delete方法
		if this.customerService.Delete(id){
			fmt.Println("-----删除完成-----")
		}else{
			fmt.Println("-----删除失败-----")
		}
	}
}
//退出软件
func(this *customerView)exit(){
	fmt.Println("是否确认退出软件（y/N）")
	for{
		fmt.Scanln(&this.key)
		if this.key=="Y"||this.key=="y"||this.key=="N"||this.key=="n"{
			break
		}
		fmt.Println("你的输入有错误")
	}
	if this.key=="Y"||this.key=="y"{
		this.loop=false
	}
}


//显示主菜单
func (this *customerView)mainMenu(){
	for{
		fmt.Println("-----客户信息管理软件-----")
		fmt.Println("     1 添 加 客 户")
		fmt.Println("     2 修 改 客 户")
		fmt.Println("     3 删 除 客 户")
		fmt.Println("     4 客 户 列 表")
		fmt.Println("     5 退       出")
		fmt.Print("请选择（1-5）:")
		fmt.Scanln(&this.key)
		switch this.key{
		  case"1":
			this.add()
			 //fmt.Println("添 加 客 户")
		  case"2":
			 fmt.Println("修 改 客 户")
		  case"3":
			//fmt.Println("删 除 客 户")
			this.delete()
		  case"4":
			this.list()
			//fmt.Println("客 户 列 表")
		  case"5":
			this.exit()
			//this.loop=false
		  default:
			fmt.Println("你的输入有错误，请重新输入")
		}
		if !this.loop{
			break
		}
	}
	fmt.Println("你退出了客户关系管理系统...")
}
func main(){
	//在main函数中，创建一个customerView,并运行显示主菜单中
	customerView:=customerView{
		key:"",
		loop:true,
	}
//这里完成对customerView结构体的customerService字段的初始化
    customerView.customerService=service.NewCustomerService()

	//显示主菜单
	customerView.mainMenu()
}