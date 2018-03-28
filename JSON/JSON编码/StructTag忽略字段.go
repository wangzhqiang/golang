package JSON编码

import (
	"github.com/gin-gonic/gin/json"
	"log"
	"fmt"
)

type Account struct {
	Email     string  `json:"email"`
	Password  string  `json:"pass_word,omitempty"`
	Money     float64 `json:"money"`
}

func main() {
	account := Account{
		Email: "rsj217@gmail.com",
		Password:"",
		Money: 100.5,
	}

	rs,err := json.Marshal(account)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(rs))
}
