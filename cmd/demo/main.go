package main

import (
	"log"

	"github.com/Beastwick18/go-webview2"
)

func main() {
	w := webview2.NewWithUserAgent(webview2.WebViewOptions{
		Debug:     true,
		AutoFocus: true,
		WindowOptions: webview2.WindowOptions{
			Title:  "Minimal webview example",
			PosX:   -404,
			PosY:   -745,
			Width:  800,
			Height: 600,
			IconId: 2, // icon resource id
			Center: false,
		},
	}, "Mozilla/5.0 (Linux; Android 11; SAMSUNG SM-G973U) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/14.2 Chrome/87.0.4280.141 Mobile Safari/537.36")
	if w == nil {
		log.Fatalln("Failed to load webview.")
	}
	w.SetSize(384, 654, webview2.HintNone)
	w.Navigate("https://act.hoyolab.com/app/community-game-records-sea/m.html#/ys/realtime?role_id=611976523&server=os_usa")
	w.Run()
	w.Destroy()
}
