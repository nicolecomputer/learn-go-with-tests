# Arrays and slices

## Notes on creating arrays

> Arrays have a fixed capacity which you define when you declare the variable.
> We can initialize an array in two ways: [N]type{value1, value2, ..., valueN}
> e.g. numbers := [5]int{1, 2, 3, 4, 5} [...]type{value1, value2, ..., valueN}
> e.g. numbers := [...]int{1, 2, 3, 4, 5}

## Notes on test coverage

Run

> go test -cover

to get coverage

## Variadic functions

> Variadic functions can be called with any number of trailing arguments. For
> example, fmt.Println is a common variadic function.

from [Go by example](https://gobyexample.com/variadic-functions)
