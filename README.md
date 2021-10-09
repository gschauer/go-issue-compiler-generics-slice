### Introduction

This repository contains a stripped down version of
[collect.Ordering](https://github.com/abc-inc/goava/tree/feature/ordering/collect/ordering)
from [github.com/abc-inc/goava](https://github.com/abc-inc/goava/), which is heavily inspired by [Guava](https://github.com/google/guava).

This repository is created solely for reproducing a internal compiler error in Go 1.18 as of 2021-10-09.

### What version of Go are you using (`go version`)?

<pre>
$ gotip version
go version devel go1.18-e1c294a 2021-10-09 01:04:29 +0000 linux/amd64
</pre>

### Does this issue reproduce with the latest release?

gotip is required for the code to compile.

go 1.17.1 displays the following error instead
(when lowering the module requirement to 1.17):

<pre>
./issue.go:28:6: internal compiler error: Cannot export a generic function (yet): Lexicographical
</pre>

### What operating system and processor architecture are you using (`go env`)?

<details><summary><code>go env</code> Output</summary><br><pre>
$ go env
GO111MODULE=""
GOARCH="amd64"
GOBIN=""
GOCACHE="/home/gschauer/.cache/go-build"
GOENV="/home/gschauer/.config/go/env"
GOEXE=""
GOEXPERIMENT=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOINSECURE=""
GOMODCACHE="/home/gschauer/go/pkg/mod"
GONOPROXY=""
GONOSUMDB=""
GOOS="linux"
GOPATH="/home/gschauer/go"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/home/gschauer/sdk/gotip"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/home/gschauer/sdk/gotip/pkg/tool/linux_amd64"
GOVCS=""
GOVERSION="devel go1.18-e1c294a 2021-10-09 01:04:29 +0000"
GCCGO="gccgo"
GOAMD64="v1"
AR="ar"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
GOMOD="/media/data/dev/go-issue-compiler-generics-slice/go.mod"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/user/1000/go-build2665377813=/tmp/go-build -gno-record-gcc-switches"
</pre></details>

### What did you do?

<pre>
go install golang.org/dl/gotip@latest
gotip download
git clone https://github.com/gschauer/go-issue-compiler-generics-slice.git
cd go-issue-compiler-generics-slice

gotip run -gcflags=-G=3 .
# Output:
# *main.Comp[[]string]

gotip test -gcflags=-G=3 .
# Compilation fails - see below
</pre>

### What did you expect to see?

The test should compile and pass, or at least display the actual type of the variable `o`.

### What did you see instead?

See also: https://github.com/gschauer/go-issue-compiler-generics-slice/actions

<pre>
# github.com/gschauer/go-issue-compiler-generics-slice_test [github.com/gschauer/go-issue-compiler-generics-slice.test]
./issue_test.go:14:60: internal compiler error: found illegal assignment main.Comp[%2eshape.[]string_0] -> main.Ordering[%2eshape.[]string_0]; :
	main.Comp[%2eshape.[]string_0] does not implement main.Ordering[%2eshape.[]string_0] (wrong type for Compare method)
		have Compare(.shape.[]string_0, .shape.[]string_0) int
		want Compare(main.Comp.Reverse.T, main.Comp.Reverse.T) int

goroutine 1 [running]:
runtime/debug.Stack()
	/home/gschauer/sdk/gotip/src/runtime/debug/stack.go:24 +0x65
cmd/compile/internal/base.FatalfAt({0x58c000, 0xc0}, {0xcfb0cf, 0x27}, {0xc000430998, 0x3, 0x3})
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/base/print.go:227 +0x1ca
cmd/compile/internal/base.Fatalf(...)
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/base/print.go:196
cmd/compile/internal/noder.assignconvfn({0xe3fcf8, 0xc0004a7930}, 0xc000517730)
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/noder/transform.go:436 +0x1b2
cmd/compile/internal/noder.typecheckaste(0xc0, {0xc000518b90, 0xc0004a7930}, 0x0, 0xc000430ab8, {0xc00070f8b0, 0x1, 0x1})
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/noder/transform.go:512 +0x179
cmd/compile/internal/noder.transformReturn(0xc000518b90)
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/noder/transform.go:533 +0xa8
cmd/compile/internal/noder.(*subster).node.func1({0xe40270, 0xc0005189b0})
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/noder/stencil.go:991 +0x94c
cmd/compile/internal/noder.(*subster).node(0xc000517d50, {0xe40270, 0xc0005189b0})
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/noder/stencil.go:1180 +0xa5
cmd/compile/internal/noder.(*subster).list(0xc0006a4f60, {0xc00070f870, 0x1, 0xc00049f450})
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/noder/stencil.go:1438 +0x8e
cmd/compile/internal/noder.(*irgen).genericSubst(0xc0004400f0, 0xc000518a00, 0xc0003ab790, {0xc0006a4f28, 0x1, 0x1}, 0x1, 0xc0003a38c0)
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/noder/stencil.go:757 +0xce6
cmd/compile/internal/noder.(*irgen).getInstantiation(0xc0004400f0, 0xc0003ab790, {0xc0006a4de0, 0x1, 0x1}, 0x1)
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/noder/stencil.go:635 +0x205
cmd/compile/internal/noder.(*irgen).instantiateMethods(0xc0004ff180)
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/noder/stencil.go:543 +0x24e
cmd/compile/internal/noder.(*irgen).getDictionarySym(0xc0004400f0, 0xc00041d520, {0xc0006a4b28, 0xcad2c0, 0x1}, 0x0)
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/noder/stencil.go:1664 +0x50d
cmd/compile/internal/noder.(*irgen).getDictionaryValue(0xc0004312a0, 0x0, {0xc0006a4b28, 0xc000480160, 0xc0003a3278}, 0x80)
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/noder/stencil.go:1748 +0x25
cmd/compile/internal/noder.(*irgen).getDictOrSubdict(0xc0004400f0, 0xc00041d520, {0xe3e718, 0xc0004fc480}, 0xc000115a00, {0xc0006a4b28, 0xc000115800, 0x2}, 0x78)
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/noder/stencil.go:584 +0xbd
cmd/compile/internal/noder.(*irgen).stencil.func1({0xe3e718, 0xc0004fc480})
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/noder/stencil.go:109 +0x32f
cmd/compile/internal/ir.Visit.func1({0xe3e718, 0xc0004fc480})
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/ir/visit.go:105 +0x30
cmd/compile/internal/ir.(*AssignStmt).doChildren(0xc00009e730, 0xc0003a2cf0)
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/ir/node_gen.go:152 +0x82
cmd/compile/internal/ir.DoChildren(...)
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/ir/visit.go:94
cmd/compile/internal/ir.Visit.func1({0xe3e330, 0xc00009e730})
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/ir/visit.go:106 +0x57
cmd/compile/internal/ir.doNodes({0xc0006b3900, 0x3, 0x0}, 0xc0003a2cf0)
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/ir/node_gen.go:1512 +0x67
cmd/compile/internal/ir.(*Func).doChildren(0xe3f078, 0xc000412c60)
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/ir/func.go:152 +0x2e
cmd/compile/internal/ir.DoChildren(...)
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/ir/visit.go:94
cmd/compile/internal/ir.Visit.func1({0xe3f078, 0xc000412c60})
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/ir/visit.go:106 +0x57
cmd/compile/internal/ir.Visit({0xe3f078, 0xc000412c60}, 0xc0006b3940)
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/ir/visit.go:108 +0xb8
cmd/compile/internal/noder.(*irgen).stencil(0xc0004400f0)
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/noder/stencil.go:90 +0x238
cmd/compile/internal/noder.(*irgen).generate(0xc0004400f0, {0xc00000e470, 0x1, 0x2})
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/noder/irgen.go:307 +0x359
cmd/compile/internal/noder.check2({0xc00000e470, 0x1, 0x1})
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/noder/irgen.go:93 +0x175
cmd/compile/internal/noder.LoadPackage({0xc00001e220, 0x1, 0x0})
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/noder/noder.go:90 +0x335
cmd/compile/internal/gc.Main(0xd0e5f0)
	/home/gschauer/sdk/gotip/src/cmd/compile/internal/gc/main.go:190 +0xaf3
main.main()
	/home/gschauer/sdk/gotip/src/cmd/compile/main.go:55 +0xdd

FAIL	github.com/gschauer/go-issue-compiler-generics-slice [build failed]
FAIL
</pre>

### Remark

The root cause seems to be a combination of exporting a type with `[]T` __AND__
using it in another package.
Executing the code in the corresponding "`_test`" package is sufficient to encounter the issue.
As long as the code is used solely inside the very same package, it runs successfully.

The `-G=3` frontend does not match the types `.shape.[]string_0` and `main.Comp.Reverse.T`.

`GOEXPERIMENT=unified` (and `-vet=off`) allows to compile and run the code successfully.

<pre>
$ GOEXPERIMENT=unified gotip test -vet=off .
# Output:
# ok      github.com/gschauer/go-issue-compiler-generics-slice    0.001s
</pre>
