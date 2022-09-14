package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	fmt.Println("Building...")
	os.Setenv("GOOS", "js")
	os.Setenv("GOARCH", "wasm")
	cmd := exec.Command("go", "build", "-o", "src/main.wasm", "./src/")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully Built!")

	id := "dist"
	if _, err := os.Stat(id); os.IsNotExist(err) {
		err = os.Mkdir(id, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	fa, err := os.Create(filepath.Join(id, "index.html"))
	if err != nil {
		panic(err)
	}
	defer fa.Close()
	fb, err := os.Create(filepath.Join(id, "main.wasm"))
	if err != nil {
		panic(err)
	}
	defer fb.Close()
	fc, err := os.Create(filepath.Join(id, "wasm_exec.js"))
	if err != nil {
		panic(err)
	}
	defer fc.Close()

	_fa, err := os.Open("index.html")
	if err != nil {
		panic(err)
	}
	defer _fa.Close()
	_fb, err := os.Open("src/main.wasm")
	if err != nil {
		panic(err)
	}
	defer _fb.Close()
	_fc, err := os.Open("src/wasm_exec.js")
	if err != nil {
		panic(err)
	}
	defer _fc.Close()

	io.Copy(fa, _fa)
	fmt.Println("Successfully Copied index.html")
	io.Copy(fb, _fb)
	fmt.Println("Successfully Copied main.wasm")
	io.Copy(fc, _fc)
	fmt.Println("Successfully Copied wasm_exec.js")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		buildPath := id
		f, err := os.Open(filepath.Join(buildPath, r.URL.Path))
		if os.IsNotExist(err) {
			index, err := os.ReadFile(filepath.Join(buildPath, "index.html"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusAccepted)
			w.Write(index)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		http.FileServer(http.Dir(buildPath)).ServeHTTP(w, r)
	})
	fmt.Println("Serving...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}
