//go:build darwin || dragonfly || freebsd || linux || nacl || netbsd || openbsd || solaris
// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package webview2

import (
	"encoding/json"
	"sync"
	"unsafe"

	"github.com/Beastwick18/go-webview2/internal/w32"
)

var (
	windowContext     = map[uintptr]interface{}{}
	windowContextSync sync.RWMutex
)

type browser interface {
	Embed(hwnd uintptr) bool
	Resize()
	Navigate(url string)
	NavigateToString(htmlContent string)
	Init(script string)
	Eval(script string)
	NotifyParentWindowPositionChanged() error
	Focus()
}

type webview struct {
	hwnd       uintptr
	mainthread uintptr
	browser    browser
	autofocus  bool
	maxsz      w32.Point
	minsz      w32.Point
	m          sync.Mutex
	bindings   map[string]interface{}
	dispatchq  []func()
}

type WindowOptions struct {
	Title  string
	Width  uint
	Height uint
	IconId uint
	Center bool
	PosX   int
	PosY   int
}

type WebViewOptions struct {
	Window unsafe.Pointer
	Debug  bool

	// DataPath specifies the datapath for the WebView2 runtime to use for the
	// browser instance.
	DataPath string

	// AutoFocus will try to keep the WebView2 widget focused when the window
	// is focused.
	AutoFocus bool

	// WindowOptions customizes the window that is created to embed the
	// WebView2 widget.
	WindowOptions WindowOptions
}

// New creates a new webview in a new window.
func New(debug bool) WebView { return nil }

// NewWindow creates a new webview using an existing window.
//
// Deprecated: Use NewWithOptions.
func NewWindow(debug bool, window unsafe.Pointer) WebView { return nil }

// NewWithOptions creates a new webview using the provided options.
func NewWithOptions(options WebViewOptions) WebView { return nil }

// NewWithOptions creates a new webview using the provided options.
func NewWithUserAgent(options WebViewOptions, agent string) WebView { return nil }

type rpcMessage struct {
	ID     int               `json:"id"`
	Method string            `json:"method"`
	Params []json.RawMessage `json:"params"`
}

func (w *webview) Create(debug bool, window unsafe.Pointer) bool { return false }

func (w *webview) CreateWithOptions(opts WindowOptions) bool { return false }

func (w *webview) Destroy() {}

func (w *webview) Run() {}

func (w *webview) Terminate() {}

func (w *webview) Window() unsafe.Pointer { return nil }

func (w *webview) Navigate(url string) {}

func (w *webview) SetHtml(html string) {}

func (w *webview) SetTitle(title string) {}

func (w *webview) SetSize(width int, height int, hints Hint) {}

func (w *webview) Init(js string) {}

func (w *webview) Eval(js string) {}

func (w *webview) Dispatch(f func()) {}

func (w *webview) Bind(name string, f interface{}) error { return nil }
