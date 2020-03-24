package service
import(
	"go_code/project2020122xianggmu/xiangmu2/model"
)
//该CustomerService完成对Customer的操作，
//包括增删改查
type CustomerService struct{
	customers []model.Customer
	//声明一个字段，表示当前切片含有多少个客户
	//该字段后面，还可以作为新客户的id+1
	customerNum int
}
//编写一个方法，可以返回*CustomerService
func NewCustomerService() *CustomerService{
	//为了能够看到有客户在切片中，我们初识话一个客户
	customerService:=&CustomerService{}
	customerService.customerNum=1
	customer:=model.NewCustomer(1,"张三","男",20,"112","fws@qq.com")
	customerService.customers=append(customerService.customers,customer)
	return customerService
}
//返回客户切片
func (this *CustomerService)List() []model.Customer{
	return this.customers
}
//添加客户到customers切片
func (this *CustomerService)Add(customer model.Customer)bool{
	this.customerNum++//3
	customer.Id=this.customerNum//4
	this.customers=append(this.customers,customer)//1
	return true//2
}
func (this *CustomerService)Delete(id int)bool{
	index :=this.FindById(id)
	//如果index==-1,说明没有这个客户
	if index==-1{
		return false
	}
	//如何从切片中删除一个元素
	this.customers=append(this.customers[:index],this.customers[index+1:]...)
	return true
}
//根据id查找客户在切片中对应的下标，如果没有客户，返回-1
func(this *CustomerService)FindById(id int)int{
	index :=-1
	//遍历this.customers切片
	for i:=0;i<len(this.customers);i++{
		if this.customers[i].Id==id{
			//找到
			index=i
		}
	}
	return index
}