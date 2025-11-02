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

type Config struct {

	Next     string
	Previous string
}

func get_config(){
	
}