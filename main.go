package main

import "mtdn.io/Kagerou/routers"

func main() {
	router := routers.InitRouter()
	router.Run("127.0.0.1:8080")
}
