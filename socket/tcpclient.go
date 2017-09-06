package main

import (
    "io/ioutil"
    "net"
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "参数：%s ip/域名+端口", os.Args[0])
        os.Exit(1)
    }

    service := os.Args[1]
    //生成tcp连接句柄
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)

    //连通远程服务器 dial 拨号的意思
    conn, err := net.DialTCP("tcp4", nil, tcpAddr)
    checkError(err)

    //发送请求
    _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
    checkError(err)

    result, err := ioutil.ReadAll(conn)
    checkError(err)

    fmt.Printf(string(result))
    os.Exit(0)
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "致命错误：%s", err.Error())
        os.Exit(1)
    }
}