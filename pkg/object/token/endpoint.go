package token

type Endpoint struct {
	Address string `bson:"address"`
	Device  string `bson:"device"`
	Info    string `bson:"info"`
}
