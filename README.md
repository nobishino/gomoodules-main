## Example1

```
go get github.com/nobishino/gomodules-a@v1.2.0
```

を実行すると次のように表示されます。

```
go get github.com/nobishino/gomodules-a@v1.2.0
go: downloading github.com/nobishino/gomodules-a v1.2.0
go: added github.com/nobishino/gomodules-a v1.2.0
go: added github.com/nobishino/gomodules-c v1.3.0
go: added github.com/nobishino/gomodules-d v1.2.0
```

この状態で`go list -m all`を実行すると次のようになるでしょう。
```
github.com/nobishino/gomoodules-main
github.com/nobishino/gomodules-a v1.2.0
github.com/nobishino/gomodules-c v1.3.0
github.com/nobishino/gomodules-d v1.2.0
```

`m.go`に次の`init`関数を追加して、`gomodule-a`を実際に使うようにしましょう。そうしないと、`go mod tidy`を実行したときに`gomodules-a`への依存が削除されてしまうため、便宜的に依存させておきます。

```go
func init() {
	var a a.A
	_ = a.Number
	_ = a.Print
}
```

`go mod tidy`を実行してから`go list -m all`を実行します。

```
go get github.com/nobishino/gomodules-b@v1.2.0
```