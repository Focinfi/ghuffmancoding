huffmancoding package is an implementation of Huffman Coding in Golang

#### Install 

`go get github.com/Focinfi/huffmancoding`

#### Usage

Basiclly, it recieves a string and return a map[string]string, the key is the distinct character in the given string and the value is binary encoding, like "0", "01"

```go
  result := huffmancoding.Encode("111223") // result: map[3:00 2:01 1:1]
```
