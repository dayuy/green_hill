package main

func main1() {
	server := NewServer("127.0.0.1", 8888)

	server.Start()
}
