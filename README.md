# Go Practice 


### Reference Links
- [Callicoder.com Golang Methods Tutorial (VERY GOOD!)](https://www.callicoder.com/golang-methods-tutorial/).  I am borrowing and learning from Rajeev Singh's article with examples here.
- 
- 
### The first app

```sh
$ go run hello.go
```
hello world

```sh
go build hello.go
ls
```
hello   hello.go

```sh
./hello
```
hello world

### Getting started with Practice

Go Methods are functions with a special _receiver_ argument.

```go
func (receiver Type) MethodName(parameterList) (returnTypes) {
}
```

Notice the way we called the method IsAbove() on the Point instance p.  It’s exactly like the way you call methods in an object-oriented programming language.
