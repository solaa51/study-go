package main

import (
    "strings"
    "time"
    "net"
    "os"
    "fmt"
    "strconv"
)

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "致命错误：%s", err.Error())
        os.Exit(1)
    }
}

func main() {

    service := ":7777"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)

    listener, err := net.ListenTCP("tcp4", tcpAddr)
    checkError(err)
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }

        //多并发处理链接请求
        go handleClient(conn)
    }
}

//处理单个链接请求 精简版本
func handleClient(conn net.Conn) {
    defer conn.Close()

    request := make([]byte, 128)
    readLen, err := conn.Read(request)
    if err != nil { //读取内容出错 断开连接
        fmt.Println(err)
        os.Exit(0)
    }

    daytime := time.Now().String()
    var data = fmt.Sprintf("接收到的长度是：%d，时间是：%s", readLen, daytime)
    conn.Write([]byte(data))
}

//长连接版本 且根据请求内容 返回不同信息
func handleClient2(conn net.Conn) {
    //设置超时时间 为2分钟 断开连接
    conn.SetReadDeadline(time.Now().Add(2*time.Minute))

    //分配空间 128 一次最大读取数量
    request := make([]byte, 128)

    defer conn.Close()

    //持续读取内容
    for {
        //读取内容 返回的是 内容长度
        readLen, err := conn.Read(request)
        if err != nil { //读取内容出错 断开连接
            fmt.Println(err)
            break
        }

        if readLen == 0 {
            break
        } else if strings.TrimSpace( string(request[:readLen]) )=="timestamp" {
            daytime :=  strconv.FormatInt( time.Now().Unix(), 10 )
            conn.Write([]byte(daytime))
        } else {
            daytime := time.Now().String()
            conn.Write([]byte(daytime))
        }

        request = make([]byte, 128) //清空内容
    }
}