# drone-qiniu
上传文件到七牛的drone插件

## Build
```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o qiniu
```

## Params

```
SET ACCESS_KEY=xxx
SET SECRET_KEY=xxx
SET ZONE=huadong
SET BUCKET=nameimtest
SET PREFIX=drone/
SET PATH=e:\drone-qiniu
```