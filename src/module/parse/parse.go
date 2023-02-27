package parse

import (
	"net/http"
	"log"
	"golang.org/x/time/rate"

	"net/url"
	"strings"
	
	"pdbg.work/shiba/module/setting"
)

var (
	param []string
	
	limiter = rate.NewLimiter(rate.Limit(setting.ServerSecondLimit), setting.MaxClients) // 検討 1sec1000, 30acsess.
)

func Parse(w http.ResponseWriter, r *http.Request) {
	//
	if !limiter.Allow() {
		http.Error(w, "503 Service Temporarily Unavailable.", 503)
		return
	}
	
	// URLの取得
	u, err := url.Parse(r.URL.Path)
	if err != nil {
		log.Fatal(err)
	}
	
	// URLのパース
	url := strings.Split(u.Path, "?")[0] // GETの削除
	param := strings.Split(strings.Trim(url, "/"), "/")
	
	// paramのトリム
	for key, value := range param {
		param[key] = strings.TrimSpace(value)
	}
	
	//　実行
	if (param[0] == "exec") {
		// URLの長さをチェック
		if (len(u.Path) > 1000) {
			http.Error(w, "500 Internal Server Error.", 500)
			return
		}
		
		// 実行する？
		f := param[1]
		param = param[2:]
		
		if (setting.Functions[f] != nil) {
			// POSTのパース
			r.ParseForm()
			
			// 実行
			setting.Functions[f](w, r, param)
			return
		}
	}
	
	// ServeHTTPでhttp.FileServerの実行
	http.FileServer(http.Dir("public")).ServeHTTP(w, r)
}

