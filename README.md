# pubfncomment

pubfncomment is analyzer that detects a function which has no comment.

## Instruction

```sh
go install github.com/sivchari/pubfncomment/cmd/pubfncomment
```

## Usage

```go
package a

func ok() {}

// OK is a function
func OK() {}

func NG() {}
```

```console
go vet -vettool=(which pubfncomment) ./...

# a
./main_test.go:11:2: public function (NG) should have comment
```

## CI

### GitHub Actions

```yaml
- name: install pubfncomment
  run: go install github.com/sivchari/pubfncomment/cmd/pubfncomment

- name: run tenv
  run: go vet -vettool=`which pubfncomment` ./...
```
