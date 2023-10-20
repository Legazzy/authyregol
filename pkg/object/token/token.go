package token

type Token struct {
	Credential Credential `bson:"credential"`
	Endpoint   Endpoint   `bson:"endpoint"`
	Identifier Identifier `bson:"identifier"`
	Resource   Resource   `bson:"resource"`
	Timestamp  Timestamp  `bson:"timestamp"`
}
