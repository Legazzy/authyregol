package request

type Request struct {
	Amount   int64 `bson:"amount"`
	Cooldown int64 `bson:"amound"`
	Limit    int64 `bson:"limit"`
	Last     int64 `bson:"last"`
}
