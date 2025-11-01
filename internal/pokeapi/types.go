package pokeapi

type Location struct {
	Count    int
	Next     string
	Previous string
	Results  []Results
}

type Results struct {
	Name string
	Url  string
}
