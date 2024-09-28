package main

func main() {
	server := NewApiServer(":3333")
	server.Run()
}
