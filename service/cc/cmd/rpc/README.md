## proto 生成代码
1. 执行指令
```
cd gcc-zero/service/cc/cmd/rpc
goctl rpc proto -src cc.proto -dir .
```

2. 把 gcc-zero/service/cc/cmd/rpc/cc中一个客户端代码移至ccclient文件夹下