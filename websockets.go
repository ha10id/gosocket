package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func main() {
    http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
        conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

        for {
            // Read message from browser
            msgType, msg, err := conn.ReadMessage()
            if err != nil {
                return
            }
            if string(msg) == "get header" {
                fmt.Printf("%s +++: %s\n", conn.RemoteAddr(), string(msg))
                msg = []byte("{\"body\":\"<div id='root' class='container mx-auto p-0 max-h-screen'></div>\"}")
//                 \"<div id=\"root\" class=\"container mx-auto p-0 max-h-screen\">" +
//                 "<nav class=\"flex items-center justify-between flex-wrap bg-teal-500 p-2\">"+
//                 "<div class=\"flex items-center flex-shrink-0 text-white mr-6\">"+
//                 "<svg class=\"fill-current h-8 w-8 mr-2\" width=\"54\" height=\"54\" viewBox=\"0 0 54 54\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M13.5 22.1c1.8-7.2 6.3-10.8 13.5-10.8 10.8 0 12.15 8.1 17.55 9.45 3.6.9 6.75-.45 9.45-4.05-1.8 7.2-6.3 10.8-13.5 10.8-10.8 0-12.15-8.1-17.55-9.45-3.6-.9-6.75.45-9.45 4.05zM0 38.3c1.8-7.2 6.3-10.8 13.5-10.8 10.8 0 12.15 8.1 17.55 9.45 3.6.9 6.75-.45 9.45-4.05-1.8 7.2-6.3 10.8-13.5 10.8-10.8 0-12.15-8.1-17.55-9.45-3.6-.9-6.75.45-9.45 4.05z\"/></svg>"+
//                 "  <span class=\"font-semibold text-xl tracking-tight\">Приложение</span>"+
//                 "</div>"+
//                 "<div class=\"block lg:hidden\">"+
//                 "<button class=\"flex items-center px-3 py-2 border rounded text-teal-200 border-teal-400 hover:text-white hover:border-white\">"+
//                 "<svg class=\"fill-current h-3 w-3\" viewBox=\"0 0 20 20\" xmlns=\"http://www.w3.org/2000/svg\">"+
//                 "<title>Menu</title>"+
//                 "<path d=\"M0 3h20v2H0V3zm0 6h20v2H0V9zm0 6h20v2H0v-2z\"/>"+
//                 "</svg>"+
//                 "</button>"+
//                 "</div>"+
//                 "<div class=\"w-full block flex-grow lg:flex lg:items-center lg:w-auto\">"+
//                 "<div class=\"text-sm lg:flex-grow\">"+
//                 "<a href=\"#responsive-header\" class=\"block mt-4 lg:inline-block lg:mt-0 text-teal-200 hover:text-white mr-4\">Новости</a>"+
//                 "<a href=\"#responsive-header\" class=\"block mt-4 lg:inline-block lg:mt-0 text-teal-200 hover:text-white mr-4\">Документы</a>"+
//                 "<a href=\"#responsive-header\" class=\"block mt-4 lg:inline-block lg:mt-0 text-teal-200 hover:text-white\">Блог</a>"+
//                 "</div><div>"+
//                 "<a href=\"#\" class=\"inline-block text-sm px-4 py-2 leading-none border rounded text-white border-white hover:border-transparent hover:text-teal-500 hover:bg-white mt-4 lg:mt-0\">Загрузить</a>"+
//                 "</div></div></nav></div>\"}")
            } else if string(msg) == "get script" {
                msg = []byte("<script src=\"/script.js\"></script>")
            } else {
                // Print the message to the console
                fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
            }
            // Write message back to browser
            if err = conn.WriteMessage(msgType, msg); err != nil {
                return
            }
        }
    })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "websockets.html")
    })

    http.ListenAndServe(":8080", nil)
}
