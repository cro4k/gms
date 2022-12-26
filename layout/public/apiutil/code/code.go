package code

import "github.com/cro4k/ginx"

const (
	OK   = 1
	Fail = 0
)

const (
	ErrIncorrectUsername = "incorrect username or password"
)

func init() {
	ginx.SetCodeMap(map[int]int{
		ginx.CodeOK:   OK,
		ginx.CodeFail: Fail,
	})
}
