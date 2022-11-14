<!--
 * @Author: tj
 * @Date: 2022-11-10 16:30:29
 * @LastEditors: tj
 * @LastEditTime: 2022-11-14 10:13:41
 * @FilePath: \book\cmd\wasm\readme.md
-->
```
SET GOOS=js
SET GOARCH=wasm
go build -ldflags="-s -w" -o ../../static/main.wasm
```