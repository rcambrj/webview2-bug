//go:build windows

package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net"
	"net/http"

	"github.com/jchv/go-webview2"
)

//go:embed html
var html embed.FS

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	defer lis.Close()

	embeddedfs, _ := fs.Sub(html, "html")
	httpfs := http.FS(embeddedfs)
	go http.Serve(lis, http.FileServer(httpfs))

	w := webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:     true,
		AutoFocus: true,
		WindowOptions: webview2.WindowOptions{
			Title: "foo",
		},
	})
	w.SetSize(800, 600, webview2.HintNone)
	w.Navigate("http://" + lis.Addr().String())

	w.Bind("ready", func() { // nolint: errcheck
		fmt.Println("this is called only once")
	})

	// this window.ready() does not execute on the first page load
	// see index.html for the workaround
	w.Init(`window.addEventListener('DOMContentLoaded', function() {
		window.ready();
	});`)

	w.Run()
}