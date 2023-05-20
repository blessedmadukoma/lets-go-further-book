package data

import (
	"database/sql"
	"time"

	"github.com/blessedmadukoma/greenlight/internal/validator"

	"github.com/lib/pq"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`    // release year
	Runtime   Runtime   `json:"runtime,omitempty"` // movie run time (in minutes)
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"` // the version number starts at 1 and will be incremented each time the movie information is updated
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")
	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")
	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
}

// MovieModel defines the methods for interacting with movies data.
type MovieModel struct {
	DB *sql.DB
}

// Insert will insert a new movie into the database.
func (m MovieModel) Insert(movie *Movie) error {
	query := `INSERT INTO movies (title, year, runtime, genres) VALUES ($1, $2, $3, $4) RETURNING id, created_at, version`

	args := []interface{}{movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres)}

	return m.DB.QueryRow(query, args...).Scan(&movie.ID, &movie.CreatedAt, &movie.Version)
}

// Get will return the movie with the provided ID. If no matching movie is
func (m MovieModel) Get(id int64) (*Movie, error) {
	return nil, nil
}

// Update will update the movie with the provided ID and information.
func (m MovieModel) Update(movie *Movie) error {
	return nil
}

// Delete will delete a movie from the database.
func (m MovieModel) Delete(id int64) error {
	return nil
}
