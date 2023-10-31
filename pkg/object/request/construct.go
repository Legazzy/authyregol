package request

import (
	"time"
)

func NewRequest() *Request {
	req := Request{}

	req.Amount = 0
	req.Cooldown = 3
	req.Last = time.Now().Unix()
	req.Limit = 10

	return &req
}
