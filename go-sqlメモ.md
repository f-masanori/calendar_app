## 単一行クエリ

クエリが多くとも1行しか返却しない場合、冗長な定型コードをいくらか省略することができる。

```go
var name string
err = db.QueryRow("SELECT name FROM users WHERE id = ?", 1).Scan(&name)
if err != nil {
	log.Fatal(err)
}
fmt.Println(name)
```

クエリのエラーは`Scan()`が呼び出されるまで延期され、`Scan()`によって返却されます。また、プリペアドステートメントから`QueryRow()`を呼び出すこともできます。

```go
stmt, err := db.Prepare("SELECT name FROM users WHERE id = ?")
if err != nil {
	log.Fatal(err)
}
var name string
err = stmt.QueryRow(1).Scan(&name)
if err != nil {
	log.Fatal(err)
}
fmt.Println(name)
```

