# Linter for Go


各种流行的 Go Linter 介绍，例如 Gocyclo、bodyclose、sqlrows、GolangCI-Lint、reviewdog等

<!--more-->

## 一、 Gocyclo

**Gocyclo**，函数圈复杂度(cyclomatic complexities)计算工具，计算方法：

- `1` is the base complexity of a function
- `+1` for each `if`, `for`, `case`, `&&` or `||`

> **NOTE**
>
> 1. 在 Go 语言中，由于 `if err != nil` 的特殊情况存在，因此，当圈复杂度超过 15 时，则表明函数较复杂；
> 其他语言中圈复杂度阈值一般设置为 10；
> 2. `switch` 中的 `default` 并不会增加函数的圈复杂度；

### Install

`gocyclo` commond tool

```Bash
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
```

### Usage

```Bash
Calculate cyclomatic complexities of Go functions.
Usage:
    gocyclo [flags] <Go file or directory> ...

Flags:
    -over N               show functions with complexity > N only and
                          return exit code 1 if the set is non-empty
    -top N                show the top N most complex functions only
    -avg, -avg-short      show the average complexity over all functions;
                          the short option prints the value without a label
    -ignore REGEX         exclude files matching the given regular expression

The output fields for each line are:
<complexity> <package> <function> <file:line:column>
```

### Example

```Go
package main

import (
 "fmt"
 "strconv"
)

func main() {
    var a = 10
    if a == 10 {
        f()
    } else {
        fmt.Printf("%s", strconv.Itoa(a))
    }

    switch a{
    case 10:
        fmt.Println(a)
    default:
        fmt.Println("default")
    }
}

func f() {
    a := 10
    b := 12

    if a != b {
        // do something
        fmt.Println("a != b")
    }
}
```

```Bash
$ gocyclo gocyclo-test/main.go 
3 main main gocyclo-test/main.go:8:1
2 main f gocyclo-test/main.go:24:1
```

## 二、bodyclose

`bodyclose` 是检查 `res.Body` 是否被正确关闭的静态检查工具；

### Install

```sh
go install github.com/timakin/bodyclose
```

### Usage

这里借用 `GolangCI-Linter` 的方式使用 `bodyclose`

```Go
package Kyten

import (
 "fmt"
 "io"
 "net/http"
)

func f() error{
    resp, err := http.Get("http://example.com/")
    if err != nil {
        return err
    }
    // defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    fmt.Println(body)
    return nil
}
```

```sh
$ golangci-lint run --disable-all -E bodyclose main.go
main.go:10:26: response body must be closed (bodyclose)
    resp, err := http.Get("http://example.com/")
```

> Tip:
> 避免使用 `http` 库中 `body` 忘记 `close` 的更优方案是：**对 Go 官方提供的 `http` 进行封装，使 caller 不用显示调用 `close` 函数.**

## 三、sqlrows

`sqlrows` is a static code analyzer which helps uncover bugs by reporting a diagnostic for mistakes of sql.Rows usage.

### Install

```sh
go install github.com/gostaticanalysis/sqlrows
```

### Usage

```Go
package sqlrowstest_test

import (
 "context"
 "database/sql"
)

func f(ctx context.Context, db *sql.DB) (interface{}, error) {
    rows, err := db.QueryContext(ctx, "SELECT * FROM users")
    defer rows.Close()

    if err != nil {
        return nil, err
    }
    // defer rows.Close()

    for rows.Next() {
        err = rows.Scan()
        if err != nil {
            return nil, err
        }
    }
    return nil, nil
}
```

```sh
go vet -vettool=$(which sqlrows) main.go                  (base) 
# command-line-arguments
./main.go:10:11: using rows before checking for errors
```

```sh
go vet -vettool=$(which sqlrows) main.go                  (base) 
# command-line-arguments
./main.go:9:33: rows.Close must be called
```

## 四、GolangCI-Lint

`golangci-lint` is a fast Go linters runner. It runs linters in parallel, uses caching, supports YAML configuration, integrates with all major IDEs, and includes over a hundred linters. **生产级静态分析工具**

### 本地安装

注意版本：`v1.57.2`

```sh
# binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.57.2

golangci-lint --version
```

### Usage

在不进行任何配置的情况下，GolangCI-Lint 将默认采用启动以下 Linters:
`errcheck`, `gosimple`, `govet`, `ineffassign`, `staticcheck`, `unused`

![GolangCI-Linter Default config](/img/golangci-linter-default.png)

还可以通过传递 `-E`/`--enable` 参数来启动 Linter，传递 `-D`/`--disable` 来禁用 Linter

```sh
golangci-lint run --disable-all -E errcheck
```

### Visual Studio Code 集成

VSCode配置 `settings.json`

```json
"go.lintTool": "golangci-lint",
"go.lintFlags": [
  "--fast" // Using it in an editor without --fast can freeze your editor.
]
```

> Golangci-lint automatically discovers `.golangci.yml` config for edited file: Don't need to configure it in VS Code settings.

更多详细信息，请参考[官方文档](https://golangci-lint.run/)

## 五、reviewdog

A code review dog who keeps your codebase healthy.

> 将常用的 Linter 集成在 CI 上线流程中，可以保证项目代码质量的下限

## 六、Reference

- [Cyclomatic complexity](https://en.wikipedia.org/wiki/Cyclomatic_complexity)
- [Gocyclo](https://github.com/fzipp/gocyclo)
- [bodyclose](https://github.com/timakin/bodyclose)
- [sqlrows](https://github.com/gostaticanalysis/sqlrows)
- [GolangCI-Lint](https://github.com/golangci/golangci-lint)
- [static analysis](https://github.com/analysis-tools-dev/static-analysis)
- [reviewdog](https://github.com/reviewdog/reviewdog)

