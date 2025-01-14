package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pratikju/go-chat/middleware"

	"golang.org/x/net/websocket"
)

// ListenHTTP starts http server at given hostname and port
func ListenHTTP(hostname string, port int, handler http.Handler) {
	host := fmt.Sprintf("%s:%d", hostname, port)
	log.Println("starting http server at", host)
	if err := http.ListenAndServe(host, handler); err != nil {
		log.Fatal(err)
	}
}

// AttachHandlers attaches all http handler
func AttachHandlers() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/authorize_github", githubAuthorizationHandler)
	http.HandleFunc("/authorize_google", googleAuthorizationHandler)
	http.HandleFunc("/git_home", githubCallbackHandler)
	http.HandleFunc("/google_home", googleCallbackHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/uploads/", uploadViewHandler)
	http.Handle("/assets/", http.FileServer(http.Dir(".")))
	http.Handle("/websocket", websocket.Handler(socketHandler))
	http.HandleFunc("/user", middleware.IsAuthenticated(userHandler))
}
