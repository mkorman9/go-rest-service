package main

type Cat struct {
	Id				int64	`json:"id" db:"ID"`
	RoleName 		string	`json:"roleName" db:"ROLE_NAME"`
	Name 			string 	`json:"name" db:"NAME"`
	DuelsWon 		int		`json:"duelsWon" db:"DUELS_WON"`
}
