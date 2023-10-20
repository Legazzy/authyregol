package permission

type Permission struct {
	Identifier Identifier `bson:"identifier"`
	Details    Details    `bson:"details"`
}
