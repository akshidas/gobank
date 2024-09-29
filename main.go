package main

//
// type book struct {
// 	Name   string `json:"name"`
// 	Author string `json:"author"`
// }
//
// type library struct {
// 	Books []book `json:"books"`
// }

func main() {
	server := NewApiServer(":3333")
	server.Run()

	// lib := library{}
	// lib.readBooks()
	// for _, v := range lib.Books {
	// 	fmt.Println(v.Author, v.Name)
	// }
	// lib.writeBooks()
	// books := []book{
	// 	{Name: "An Introduction to programming in Go", Author: "Caleb Doxsey"},
	// 	{Name: "Go in Action", Author: "William Kennedy"},
	// 	{Name: "The way to Go", Author: "IVO BALBAERT"},
	// }
	//
	// // create the library with books
	// lib := library{Books: books}
	// libJSON, _ := json.Marshal(lib)
	// // create a byte array of a string
	//
	// err := os.WriteFile("hello.json", libJSON, 0777)
	//
	// if err != nil {
	// 	log.Fatalf("%v", err)
	// }
}

//
// func (l *library) writeBooks() {
// 	book := book{
// 		Author: "My fucked up life",
// 		Name:   "Akshay",
// 	}
// 	l.Books = append(l.Books, book)
// 	libJSON, _ := json.Marshal(l.Books)
// 	err := os.WriteFile("hello.json", libJSON, 0777)
//
// 	if err != nil {
// 		log.Fatalf("%v", err)
// 	}
// }
//
// func (l *library) readBooks() {
// 	content, err := os.ReadFile("hello.json")
//
// 	if err != nil {
// 		log.Fatalf("Error while reading a file %v", err)
// 	}
//
// 	err = json.Unmarshal(content, l)
//
// 	if err != nil {
// 		log.Fatalf("Error while unmarshal the content  %v", err)
// 	}
//
// }
