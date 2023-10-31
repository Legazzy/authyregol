package service

type Service struct {
	Details    Details    `bson:"details"`
	Identifier Identifier `bson:"identifier"`
}
