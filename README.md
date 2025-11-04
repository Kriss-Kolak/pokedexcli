# PokePoke - A Command Line Pokédex

A interactive command-line Pokédex application built in Go that allows users to explore Pokémon locations and information using the PokeAPI.

## Features

- Browse location areas from the Pokémon world
- Navigate through different locations using 'next' and 'previous' commands
- View Pokémon available in each location area
- Built-in caching system for improved performance
- Interactive command-line interface

## Installation

```bash
# Clone the repository
git clone https://github.com/Kriss-Kolak/PokePoke.git

# Navigate to the project directory
cd PokePoke

# Build the project
go build

# Run the application
./pokedexcli
```

## Commands

- `help` - Display a list of available commands
- `exit` - Exit the Pokédex
- `map` - Display the names of 20 location areas
- `mapb` - Display the previous 20 location areas
- `explore [location-area]` - List Pokémon in a specific location area
- `catch [pokemon-name]` - Catch specific Pokémon
- `inspect [pokemon-name]` - List caught Pokémon stats  
- `pokedex` - List all caught Pokémons 

## Technical Details

- Built in Go
- Uses the [PokeAPI](https://pokeapi.co/) for Pokémon data
- Implements an in-memory cache with configurable TTL
- Clean architecture with separation of concerns:
	- `internal/pokeapi`: API client and data structures
	- `internal/pokecache`: Caching implementation
	- Command handlers for different CLI operations

## Project Structure

```
PokePoke/
├── internal/
│   ├── pokeapi/
│   │   ├── client.go
│   │   ├── config.go
│   │   ├── data.go
│   │   └── types.go
│   └── pokecache/
│       ├── cache_test.go
│       └── cache.go
├── main.go
├── repl.go
├── command_exit.go
├── command_help.go
└── README.md
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- [PokeAPI](https://pokeapi.co/) for providing the Pokémon data
- The Go community for excellent documentation and tools