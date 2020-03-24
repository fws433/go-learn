package main
import(
	"fmt"
	"time"
)
func main(){
	//获取当前时间
	now :=time.Now()
	fmt.Printf("now=%v now type=%T\n",now,now)

	//2.通过Now可以获取到年月日，时分秒
	fmt.Printf("年=%v\n",now.Year())
	fmt.Printf("月=%v\n",int(now.Month()))
	fmt.Printf("日=%v\n",now.Day())

	//1格式化日期时间1
	fmt.Printf("当前年月日%d-%d-%d %d:%d:%d\n",now.Year(),
	now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
	dateStr:=fmt.Sprintf("当前年月日%d-%d-%d %d:%d:%d\n",now.Year(),now.Month(),
	now.Day(),now.Hour(),now.Minute(),now.Second())
	fmt.Printf("dateStr=%v\n",dateStr)

	//2格式化日期时间2
	fmt.Printf(now.Format("2006-01-02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(now.Format("2006"))
	fmt.Println()
	//Unix表示时间戳从1970年1月1日到现在所经过的时间（单位秒），如果经过的时间是纳秒
	fmt.Printf("unix时间戳=%v,unixnano时间戳=%v\n",now.Unix(),now.UnixNano())
}
