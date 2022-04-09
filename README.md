# re - Simplified Go regular expressions

A wrapper around `rexexp` package

Facts
- Underneath it is using `MustCompile()`. Any errors will result in `panic` at runtime
- Text can be passed in as `string` or `[]byte` but for now `[]byte` is converted to `string` before matching/replacing
- It provides two interfaces: a simple and an object oriented interface
- Because of the above, this package will produce slower code but hopefully provides a nicer interface
- Can be optimised later

## Examples

```go
text1 := `Brown fox is a forest animal. Red fox is a desert animal.`
```

### Match
Using `regexp` package
```go
if regexp.MustCompile(`fox`).MatchString(text1) {
	fmt.Println("Text 1 contains `fox`")
}
```

Using `re` package - simple
```go
if re.Matches(text1, `fox`) {
	fmt.Println("Text 1 contains `fox`")
}
```

Using `re` package - OO way
```go
if re.String(text1).Pattern(`fox`).Matches() {
	fmt.Println("Text 1 contains `fox`")
}
```

### Submatch

Using `regexp` package
```go
if re := regexp.MustCompile(`(\S+)\sfox`); re.MatchString(text1) {
	fmt.Println("Fox found:", re.FindStringSubmatch(text1)[1])
}
```

Using `re` package - simple
```go
if s := re.Submatch(text1, `(\S+)\sfox`, 1); s != "" {
	fmt.Println("Fox found:", s)
}
```

Using `re` package - OO way
```go
if re := re.String(text1).Pattern(`(\S+)\sfox`); re.Matches() {
	fmt.Println("Fox found:", re.Submatch(1))
}
```
### Substitution

Using `regexp` package
```go
text2 := regexp.MustCompile(`(\S+)\s(fox|cat)`).ReplaceAllString(text1, `$2-$1`);
```

Using `re` package - simple
```go
test2 := re.Replace(text1, `(\S+)\s(fox|cat)`, `$2-$1`)
```

Using `re` package - OO way
```go
test2 := re.String(text1).Pattern(`(\S+)\s(fox|cat)`).Replace(`$2-$1`)
```
