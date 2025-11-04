package pokeapi

type Location struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Results `json:"results"`
}

type Results struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationAreaID struct {
	Id                int                `json:"id"`
	Name              string             `json:"name"`
	GameIndex         int                `json:"game_index"`
	Location          Location           `json:"location"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon NamedAPIResource `json:"pokemon"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
}

type Pokemon struct {
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
	Weight         int           `json:"weight"`
	Height         int           `json:"height"`
}

type PokemonStat struct {
	BaseStat int              `json:"base_stat"`
	Stat     NamedAPIResource `json:"stat"`
}

type PokemonType struct {
	Type NamedAPIResource `json:"type"`
}
