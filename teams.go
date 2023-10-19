package main

type team struct {
	name string
	// id must be three uppercase letters
	id string
	// country should be three uppercase letters using the [FIFA country
	// codes]
	//
	// [FIFA country codes]: https://en.wikipedia.org/wiki/List_of_FIFA_country_codes
	country string
	// id of team that this team is paired with
	paired string
	// UEFA club coefficient
	coeff int
	// Team is in pot 1
	pot1 bool
}

var teams = []team{
	{"Barcelona", "FCB", "ESP", "MAD", 126_233, true},
	{"Olympic Lyon", "OLY", "FRA", "PSG", 118_166, true},
	{"Bayern Munich", "BAY", "GER", "SGE", 96_333, true},
	{"Chelsea", "CHE", "ENG", "", 81_366, true},
	{"Paris Saint-Germain", "PSG", "FRA", "OLY", 97_166, false},
	{"Slavia Praha", "SLP", "CZE", "", 39_233, false},
	{"Real Madrid", "MAD", "ESP", "FCB", 37_233, false},
	{"Rosengård", "ROS", "SWE", "HAC", 33_399, false},
	{"St. Pölten", "STP", "AUT", "", 30_050, false},
	{"Benfica", "SLB", "POR", "", 22_800, false},
	{"Häcken", "HAC", "SWE", "ROS", 22_399, false},
	{"Roma", "ROM", "ITA", "", 21_000, false},
	{"Ajax Amsterdam", "AJA", "NED", "", 18_400, false},
	{"Paris FC", "PFC", "FRA", "", 18_166, false},
	{"Eintracht Frankfurt", "SGE", "GER", "BAY", 17_333, false},
	{"Brann", "BRA", "NOR", "", 7_100, false},
}
