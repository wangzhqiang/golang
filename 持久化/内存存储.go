package main

import "fmt"

type Post struct{
	Id      int
	Content string
	Author  string
}

var PostById map[int]*Post
var PostByAuthor map[string][]*Post

func store(post Post) {
	PostById[post.Id] = &post
	PostByAuthor[post.Author] = append(PostByAuthor[post.Author],&post)

}

func main(){
	PostById = make(map[int]*Post)
	PostByAuthor = make(map[string][]*Post)


	post1 := Post{Id:1,Content:"Hello world!",Author:"Sau Sheong"}
	post2 := Post{Id:2,Content:"Bonjour Monde",Author:"Pierre"}
	post3 := Post{Id:3,Content:"Hola Mundo!",Author:"Pedro"}
	post4 := Post{Id:4,Content:"Greetings Earthling",Author:"Sau Sheong"}

	PostByAuthor[post1.Author] = PostByAuthor[post1.Author]  //map[Sau Sheong:[]]
	fmt.Println(PostByAuthor)

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById)   //打印结果 map[2:0xc4200181b0 3:0xc4200181e0 4:0xc420018210 1:0xc420018180]
	fmt.Println(PostByAuthor)  //打印结果 map[Pedro:[0xc4200181e0] Sau Sheong:[0xc420018180 0xc420018210] Pierre:[0xc4200181b0]]

	for _,post := range PostByAuthor["Sau Sheong"] {
		fmt.Println(post)
	}

	for _,post := range PostByAuthor["Pedro"] {
		fmt.Println(post)
	}
}
