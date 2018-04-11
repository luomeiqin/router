package main
import (
	"github.com/luomeiqin/micro/router/common"
	"fmt"
)

func main(){
	str := common.Read("abc")
	fmt.Println(str)
}