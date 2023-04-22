# twID
twID package 提供了一個驗證台灣身分證字號的功能。這個功能可以檢查身分證字號的格式和 checksum 是否正確。

## 安裝
使用者可以使用 go get 安裝 twID package：
```bash
go get github.com/r567tw/twID
```

## 使用方法
以下是 twID package 的使用範例：
```go
package main

import (
	"fmt"
	"github.com/r567tw/twID"
)

func main() {
	id := "A123456789"

	if twID.Verify(id) {
		fmt.Printf("%s is a valid Taiwan national identification number.\n", id)
	} else {
		fmt.Printf("%s is not a valid Taiwan national identification number.\n", id)
	}
}
```

## 貢獻方式
本專案歡迎開發者提交 Pull Request。如果有任何更改或問題，請先開一個 issue 來討論。

## 授權
twID package 使用 MIT License。



