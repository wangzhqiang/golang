package JSON编码

import (
	"github.com/gin-gonic/gin/json"
	"log"
	"fmt"
)

type User struct{
	Name   string
	Age    int
	Roles  []string
	Skill  map[string]float64
}


func main() {
	skill := make(map[string]float64)

	skill["python"] = 99.5
	skill["elixir"] = 90
	skill["ruby"] = 80.0

	user := User{
		"rsj217",
		27,
		[]string{"Owner","Master"},
		skill,
	}

	rs,err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(rs))
}