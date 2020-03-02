package domain

// Mock Userdatabase
var MockUsersById map[int64]*User = map[int64]*User{
	121: &User{
		"Frank",
		"lkdjalkdsjf",
		121,
	},
	122: &User{
		"Thomas",
		"lkajsdlfkjadsfkl",
		122,
	},
	123: &User{
		"Sabrina",
		"aslfkasdlfkjsdlfkjds",
		123,
	},
}
var MockUsersByEmail map[string]*User = map[string]*User{
	"frank@test.com": &User{
		"Frank",
		"lkdjalkdsjf",
		121,
	},
	"thomas@test.com": &User{
		"Thomas",
		"lkajsdlfkjadsfkl",
		122,
	},
	"sabrina@test.com": &User{
		"Sabrina",
		"aslfkasdlfkjsdlfkjds",
		123,
	},
}
