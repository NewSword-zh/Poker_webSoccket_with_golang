package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	//需要如何配置可以不用绝对路径
	http.ServeFile(w, r, "D:\\awesomeProject\\src\\WebSocketT_T\\home.html")//使用html作为返回值
}

func serveJs(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	//需要如何配置可以不用绝对路径
	http.ServeFile(w, r, "D:\\awesomeProject\\src\\WebSocketT_T\\js\\poker.min.js")//使用html作为返回值
}
func main() {
	flag.Parse()//解析命令行传入的参数
	hub := newHub()
	go hub.run()//广播站开始工作
	go testPush(hub)//定时推送消息
	fmt.Println("服务器推送服务已开启")

	http.HandleFunc("/js", serveJs)
	http.HandleFunc("/index", serveHome)
	http.HandleFunc("/poker", servePoker)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {

		serveTest(hub, w, r)
	})


	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}


}

func servePoker(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	poker := Poker{Color:"heart",Num:"5",X:"300px",Y:"400px"}

	pokerJson,_ := json.Marshal(poker)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")//配置运行跨域
	io.WriteString(w,string(pokerJson))
	//fmt.Fprint(w,poker_json)

}