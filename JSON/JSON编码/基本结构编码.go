package JSON编码

import (
	"encoding/json"
	"log"
	"fmt"
)

type Account struct{
	Email string
	password string
	Money float64
}

func main() {
	account := Account{
		"rsj217@gmail.com",
		"123456",
		100.05,
	}

	rs, err := json.Marshal(account)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rs)
	fmt.Println(string(rs))
}
