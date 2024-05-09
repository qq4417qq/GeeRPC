package main

import (
	"fmt"
	"geerpc"
	"log"
	"net"
	"sync"
	"time"
)

func startServer(addr chan string) {
	// pick a free port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	geerpc.Accept(l)
}

func main() {
	log.SetFlags(0)                         //清除每个日志条目前加上日期、时间和文件的名称及行号
	addr := make(chan string)               //addr是一个string类型的通道
	go startServer(addr)                    //创建服务器
	client, _ := geerpc.Dial("tcp", <-addr) //直到接受到addr再产生conn
	defer func() { _ = client.Close() }()   //关闭client

	time.Sleep(time.Second)
	// send request & receive response
	var wg sync.WaitGroup //加入wg用来go计数
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := fmt.Sprintf("geerpc req %d", i)
			var reply string
			if err := client.Call("Foo.Sum", args, &reply); err != nil { //call，先send，再Write，再
				log.Fatal("call Foo.Sum error:", err)
			}
			log.Println("reply:", reply)
		}(i)
	}
	wg.Wait()
}
