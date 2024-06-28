package main

import "children-songs/routers"

func main() {
	r := routers.Router()

	r.Run(":8081") // 监听并在 0.0.0.0:8080 上启动服务
}
