## RustFFI
FFI (Foreign Function Interface) 翻译过来叫做外部函数接口，最早来自于`Common Lisp` 的规范（Wiki）。

FFI的作用简单来说就是允许一种语言去调用另一种语言，在不同的语言中会有不同的实现，比如在Go中的cgo, Python中的ctypes， Haskell中的CAPI（之前还有一个ccall）等。

对于Go和Rust而言，它们的FFI需要与C语言对象进行通信，而这部分其实是由操作系统根据API中的调用约定来完成的。

## Rust安装
```
brew install rustup-init

rustup-init

rm -rf ~/.profile ~/.zshenv

echo '. "$HOME/.cargo/env"' >> .zprofile

rustup self uninstall
```

## 用Cargo创建项目
使用Rust的Cargo工具创建一个名叫rustdemo的项目，这里由于我增加了--lib的选项，使用其内置的library模板。
```
cargo new --lib rustdemo

cd rustdemo
```

## 准备Rust代码
暴露出来一个函数名叫做rustdemo，接收一个外部的参数，并将其打印出来。之后从Rust这边再设置一个字符串。

> CString::new(str_name).unwrap().into_raw() 被转换为原始指针，以便之后由 C 语言处理。

`lib.rs`
```
extern crate libc;
use std::ffi::{CStr, CString};

#[no_mangle] 
pub extern "C" fn rustdemo(name: *const libc::c_char) -> *const libc::c_char {
    let cstr_name = unsafe { CStr::from_ptr(name) };
    let mut str_name = cstr_name.to_str().unwrap().to_string();
    println!("Rust get Input:  \"{}\"", str_name);
    let r_string: &str = " Rust say: Hello Go ";
    str_name.push_str(r_string);
    CString::new(str_name).unwrap().into_raw()
}
```

## 编译Rust代码
`Cargo.toml`
```
[package]
name = "rustdemo"
version = "0.1.0"
edition = "2021"

[lib]
crate-type = ["cdylib"]

[dependencies]
libc = "0.2"
```

然后进行编译，得出dylib动态库文件。（Macos）
```
cargo build --release
```

## 准备Go代码
`main.go`
```
package main

/*
#cgo LDFLAGS: -L./ -lrustdemo
#include <stdlib.h>
#include "./rustdemo.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	s := "Go say: Hello Rust"

	input := C.CString(s)
	defer C.free(input)
	o := C.rustdemo(input)
	output := C.GoString(o)
	fmt.Printf("%s\n", output)
}
```

> 同时，为了能够让Go程序能正常调用Rust函数，这里我们还需要声明其头文件。

`rustdemo.h`
```
char* rustdemo(char *name);
```

## 编译Go代码
在Go编译的时候，我们需要开启CGO（默认都是开启的），同时需要链接到Rust构建出来的`librustdemo.dylib`文件，我们将该文件放到当前目录中。
```
cp ./target/release/librustdemo.dylib .
```

编译：
```
go build -o go-rust  -ldflags="-r ./" main.go

./go-rust
```
output：
``` 
Rust get Input:  "Go say: Hello Rust"
Go say: Hello Rust Rust say: Hello Go
```

> 可以看到，第一行的输出是由 Go 传入了 Rust ， 第二行中则是从 Rust 再传回 Go 的了。符合我们的预期。


