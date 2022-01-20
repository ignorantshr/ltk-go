package entity

/*
角色：主公，忠臣，反贼，内奸，地主，农民
Lord, loyalist, Rebel, traitor, landlord, peasant
 */

type Role int

const (
	Lord Role = iota
	Loyalist
	Rebel
	Traitor
	Landlord
	Peasant
)
