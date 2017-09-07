package main

import (
    "errors"
    "net/rpc"
    "net/http"
    "fmt"
)

/**
采用http网络协议的rpc  服务端
 */
//
type Args struct {
    A, B int
}

//除法的 商和余数
type Quotient struct {
    Quo, Rem int
}

//要注册成rpc服务的结构
type Arith int

//计算 两个参数的乘积
func (t *Arith) Multiply( args *Args, reply *int ) error {
    *reply = args.A * args.B
    return nil
}

func (t *Arith) Divide( args *Args, quo *Quotient ) error {
    if args.B == 0 {
        return errors.New("除数为0")
    }

    quo.Quo = args.A / args.B
    quo.Rem = args.A % args.B

    return nil
}

func main() {
    arith := new(Arith)
    //注册一个arith的rpc服务
    rpc.Register(arith)
    rpc.HandleHTTP()

    err := http.ListenAndServe(":1234", nil)
    if err != nil {
        fmt.Println(err.Error())
    }
}