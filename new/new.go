package new

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var mux = make(map[string]func(w http.ResponseWriter, r *http.Request))

// StartServer 启动web服务 Server.
func StartServer(Addr, staic string) {
	server := http.Server{
		Addr:        Addr,
		ReadTimeout: 5 * time.Second,
		Handler:     Staic(staic), // 启动文件服务器
	}

	// server.Handler = Staic(staic)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

// HandlerFunc 自定义方法.
func HandlerFunc(router string, fn func(w http.ResponseWriter, r *http.Request)) {

	mux[router] = fn
	return
}

// Staic 自定义的文件服务，返回 ServeMux.
func Staic(s string) *http.ServeMux {
	NewServeMux := http.NewServeMux()
	NewServeMux.Handle("/", &myHandler{})

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	NewServeMux.Handle(s, http.StripPrefix(s, http.FileServer(http.Dir(wd))))
	return NewServeMux
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u := r.URL.String()

	for k := range mux {
		if h, ok := mux[k]; ok && strings.Contains(u, k) && !strings.Contains(u, "/favicon.ico") {
			h(w, r)
			return
		}
	}

	io.WriteString(w, "URL:"+u+"-------------- id:")
}

// ParseQueryString --- 对象序列化
func ParseQueryString(s string) map[string][]string {
	u, err := url.Parse(s)
	if err != nil {
		log.Fatal(err)
	}

	return u.Query()
}
