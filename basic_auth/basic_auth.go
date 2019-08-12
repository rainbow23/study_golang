package main

import (
	"fmt"
	"net/http"
)

const (
	// Basic認証のUserID
	basicAuthUser = "user"
	// Basic認証のPassword
	basicAuthPassword = "password"
)

// 認証が必要なハンドラ用の型
type basicAuthHandler func(w http.ResponseWriter, r *http.Request, u string)

// curl -v -X GET -u user:password localhost:8080/need

func main() {
	// ラッパーかます
	http.HandleFunc("/need", needBasicAuth(needAuthHandler))
	http.HandleFunc("/unneed", unneedBasicAuth(unneedAuthHandler))

	http.ListenAndServe(":8080", nil)
}

// 認証メソッド
// 必要ならDBとかにアクセス
func basicAuthenticate(user, pass string) bool {
	if user == basicAuthUser && pass == basicAuthPassword {
		return true
	}
	return false
}

// 認証が必要な場合のラッパ
func needBasicAuth(fn basicAuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if ok == false || basicAuthenticate(user, pass) == false {
			w.Header().Set("WWW-Authenticate", `Basic realm="auth area"`)
			http.Error(w, "needs authenticate", http.StatusUnauthorized)
			return
		}
		fn(w, r, user)
	}
}

// 認証が不要な場合のラッパ
func unneedBasicAuth(fn basicAuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, "")
	}
}

// 認証が必要なハンドラ
func needAuthHandler(w http.ResponseWriter, r *http.Request, u string) {
	fmt.Fprintln(w, "authed")
}

// 認証が不要なハンドラ
func unneedAuthHandler(w http.ResponseWriter, r *http.Request, u string) {
	fmt.Fprintln(w, "Hello")
}
