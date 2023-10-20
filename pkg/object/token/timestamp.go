package token

type Timestamp struct {
	Creation   int64 `bson:"Creation"`
	Expiration int64 `bson:"expiration"`
}
