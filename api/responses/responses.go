package responses

// Language representa informações de idioma
type Language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Description representa uma descrição com seu idioma
type Description struct {
	Description string   `json:"description"`
	Language    Language `json:"language"`
}

// Name representa um nome com seu idioma
type Name struct {
	Language Language `json:"language"`
	Name     string   `json:"name"`
}

// PokemonSpecies representa informações da espécie Pokémon
type PokemonSpecies struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// PokemonEntry representa uma entrada Pokémon no Dex
type PokemonEntry struct {
	EntryNumber    int            `json:"entry_number"`
	PokemonSpecies PokemonSpecies `json:"pokemon_species"`
}

// PokedexResponse representa a resposta completa da Pokédex
type PokedexResponse struct {
	ID             int            `json:"id"`
	IsMainSeries   bool           `json:"is_main_series"`
	Name           string         `json:"name"`
	Descriptions   []Description  `json:"descriptions"`
	Names          []Name         `json:"names"`
	PokemonEntries []PokemonEntry `json:"pokemon_entries"`
}
