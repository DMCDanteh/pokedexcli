package pokeapi

type LocationAreaDetailResponse struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon PokemonReference `json:"pokemon"`
}

type PokemonReference struct {
	Name string `json:"name"`
}
