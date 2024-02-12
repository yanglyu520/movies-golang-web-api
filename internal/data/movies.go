package data

type MovieList struct {
	Movies []Movie `json:"movies"`
}
type Movie struct {
	ID        int64    `json:"id"`
	CreatedAt int64    `json:"-"`
	Title     string   `json:"title"`
	Year      int32    `json:"year,omitempty"`
	Runtime   Runtime  `json:"runtime,omitempty,string"`
	Genres    []string `json:"genres,omitempty"`
	Version   int32    `json:"version"`
}
