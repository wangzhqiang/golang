package JSON编码

import (
	"github.com/gin-gonic/gin/json"
	"log"
	"fmt"
)

//有时候输出的json希望是数字的字符串，而定义的字段是数字类型，那么就可以使用string选项。

type Account struct {
	Email    string  `json:"email"`
	Password string  `json:"password,omitempty"`
	Money    float64 `json:"money,string"`
}

func main() {
	account := Account{
		Email:    "rsj217@gmail.com",
		Password: "123",
		Money:    100.50,
	}

	rs,err := json.Marshal(account)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(rs))

}