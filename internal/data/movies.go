package data

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"github.com/yanglyu520/movies-golang-web-api/internal/validator"
	"log"
	"time"
)

type MovieList struct {
	Movies   []*Movie `json:"movies"`
	Metadata Metadata `json:"metadata"`
}

// MovieModel struct wraps a sql.DB connection pool and allows us to work with Movie struct type
// and the movies table in our database.
type MovieModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}
type Movie struct {
	ID        int64    `json:"id"`
	CreatedAt int64    `json:"-"`
	Title     string   `json:"title"`
	Year      int32    `json:"year,omitempty"`
	Runtime   int32    `json:"runtime,omitempty"`
	Genres    []string `json:"genres,omitempty"`
	Version   int32    `json:"version"`
}

func (m MovieModel) GetAll(title string, genres []string, filters Filters) ([]*Movie, Metadata, error) {
	// Update the SQL query to include the filter conditions.
	query := fmt.Sprintf(
		`SELECT count(*) OVER(), id, created_at, title, year, runtime, genres, version 
				FROM movies
				WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1='')
				AND (genres @> $2 OR $2 = '{}')
				ORDER BY %s %s, id ASC
				LIMIT $3 OFFSET $4`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{title, pq.Array(genres), filters.limit(), filters.offset()}
	// Pass the title and genres as the placeholder parameter values.
	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}
	defer rows.Close()

	totalRecords := 0
	movies := []*Movie{}
	for rows.Next() {
		var movie Movie
		err := rows.Scan(
			&totalRecords,
			&movie.ID,
			&movie.CreatedAt,
			&movie.Title,
			&movie.Year,
			&movie.Runtime,
			pq.Array(&movie.Genres),
			&movie.Version,
		)
		if err != nil {
			return nil, Metadata{}, err
		}
		movies = append(movies, &movie)
	}
	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)
	return movies, metadata, nil
}

func (m MovieModel) Insert(movie *Movie) error {
	query := `
		INSERT INTO movies (title, year, runtime, genres, created_at) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id, created_at, version
		`

	// Create a context with a 3-second timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Create an args slice containing the values for the placeholder parameters from the movie
	// struct. Declaring this slice immediately next to our SQL query helps to make it nice and
	// clear *what values are being user where* in the query
	args := []interface{}{movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres), movie.CreatedAt}

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&movie.ID, &movie.CreatedAt, &movie.Version)
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year > 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must contain at most 5 genres")
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
}
