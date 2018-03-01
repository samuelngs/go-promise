# go-promise
Promise implementation in Golang

## Example

```go
func main() {
	t := promise.
		New(func(resolve promise.Accept, reject promise.Decline) {

			r, err := http.Get("https://google.com")
			if err != nil {
				reject(err)
				return
			}
			defer r.Body.Close()

			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				reject(err)
				return
			}

			resolve(b)
		}).
		Then(func(b []byte) string {
			return string(b)
		}).
		Catch(func(err error) string {
			return err.Error()
		})

	d := t.Result()

	fmt.Println("fetch response: ", d) // fetch response: <!doctype html><html itemscope="" itemtype="...
}
```

## Nested promises

```go
func main() {
	t := promise.
		New(func(resolve promise.Accept, reject promise.Decline) {
			resolve(2)
		}).
		Then(func(n int) promise.Promise {
			return promise.
				Resolve(n).
				Then(func(n int) int {
					return n + 3
				})
		})

	d := t.Result()

	fmt.Println("sum: ", d) // sum: 5
}
```

## Multiple arguments and returns
```go
func main() {

	t := promise.
		New(func(resolve promise.Accept, reject promise.Decline) {
			resolve(1)
		}).
		Then(func(n int) (int, int, float64, promise.Promise) {
			return n, 2, 3.5, promise.Resolve(4.5)
		}).
		Then(func(a, b int, c, d float64) int {
			return int(float64(a) + float64(b) + c + d)
		})

	d := t.Result()

	fmt.Println("total: ", d)
}
```
