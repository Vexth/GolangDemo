package router

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Vexth/GolangDemo/new"
)

// SayHello xxxxxxxx.
func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s", r.Form)
	u := new.ParseQueryString(r.URL.String())
	for k := range u {
		fmt.Println(k, u[k][0])
	}

	// 需要post的数据,以key-value形式提交
	data := make(map[string]string)
	data["appId"] = "10001"
	data["appKey"] = "test"
	data["option"] = "upload"
	data["sign"] = "1111"

	j, _ := json.Marshal(data)

	io.WriteString(w, string(j))
}
