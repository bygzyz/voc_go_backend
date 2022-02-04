#####
```
1、需要结合 gin-web-dev项目整合model迁移和数据校验（gin中不要使用beego或其他框架带的东西）
```

##### GO 生成proto文件命令
```
protoc -I . user.proto --go_out=plugins=grpc:.
```