## Stringer

[Stringer](https://pkg.go.dev/fmt#Stringer) is the interface that fmt.Print uses
to convert a type into a string

## Finding places where errors should be handled

Install errcheck

> go install github.com/kisielk/errcheck@latest

`errcheck .` will find places where errors shoudl be handled but aren't.
