# simple-clean-server
goを使ってクリーンアーキテクチャを実現する

## 使用ライブラリ

|ライブラリ|用途|備考|
|:---|:---|:---|
|[golang 1.13.3](https://golang.org/)|言語|| 
|[gin](https://github.com/gin-gonic/gin)|ルーティング||
|[wire](https://github.com/google/wire)|Dipendency Injection|現在手動なので変更予定|
|[sqlboiler](https://github.com/volatiletech/sqlboiler)|ORM|現在モックなので変更予定|

## 起動方法
* go modulesの有効化
  ```
  $ export GO111MODULE=on 
  ```
* build
  ```
  $ go build cmd/main.go
  ```
* run
  ```
  $ ./main
  ```
