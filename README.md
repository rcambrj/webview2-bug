## Description

This demonstrates the bug in [webview2](github.com/jchv/go-webview2) where the first page load does
not execute any of the `w.Init()` Javascript.

It includes a workaround, where the first page load is detected and reloaded.

If the binding worked, `fmt.Println("this is called only once")` would be called twice, once on the first page load and again on the second page load (the workaround).

## How to use

1. Clone this repository
2. `GOOS=windows go build .`
3. Put `webview2-bug.exe` in your Windows machine
4. Open command line in Windows, run `webview2-bug.exe`

Observe in the console:
```
this is called only once
```