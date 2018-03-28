package JSON编码

import (
	"encoding/json"
	"log"
	"fmt"
)

type Account struct{
	Email     string
	Password  string
	Money     float64
}

type User struct{
	Name    string
	Age     int
	Roles   []string
	Skill   map[string]float64
	Account Account
}

func main() {
	skill := make(map[string]float64)

	skill["python"] = 99.5
	skill["elixir"] = 90
	skill["ruby"] = 80.0

	account := Account{
		"rsj217@gmail.com",
		"123456",
		100.5,
	}

	user := User{
		"rsj217",
		27,
		[]string{"Owner","Master"},
		skill,
		account,
	}

	rs,err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(rs))
}
