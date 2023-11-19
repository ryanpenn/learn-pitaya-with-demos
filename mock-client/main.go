package main

func main() {
	RunBot(10, &GameConfig{
		Host:    "localhost:9000",
		Timeout: 10, // 10 seconds
	})
}
