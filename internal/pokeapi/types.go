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
