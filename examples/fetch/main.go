package main

import (
	"fmt"

	"github.com/samuelngs/go-promise"
)

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

	fmt.Println("sum: ", d)
}
