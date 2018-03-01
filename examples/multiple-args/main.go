package main

import (
	"fmt"

	"github.com/samuelngs/go-promise"
)

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
