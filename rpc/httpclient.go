package main

import (
    "os"
    "fmt"
    "net/rpc"
    "log"
)

/**
采用http网络协议的rpc  客户端
 */

//
type Args struct {
    A, B int
}

//除法的 商和余数
type Quotient struct {
    Quo, Rem int
}

func main() {
    if len(os.Args)!=2 {
        fmt.Println("使用方法：命令名 域名/IP")
        os.Exit(1)
    }

    serverAddress := os.Args[1]

    client, err := rpc.DialHTTP("tcp", serverAddress + ":1234")
    if err != nil {
        log.Fatal("连接服务端失败:", err)
    }

    //服务端的入参
    args := Args{19, 2}
    //返回值
    var reply int

    err = client.Call("Arith.Multiply", args, &reply) //args因为go中结构体实例化均为指针 所以可以不写&
    if err != nil {
        log.Fatal("服务端计算出错:", err)
    }
    fmt.Printf("计算乘积结果为 %d\n", reply)


    var quot Quotient
    err = client.Call("Arith.Divide", args, &quot)
    if err != nil {
        log.Fatal("除法计算出错:", err)
    }
    fmt.Printf("计算除法 商：%d，余数：%d\n", quot.Quo, quot.Rem)
}