package data

import "time"

// Movie represents a movie in the system
// @Description Movie information
type Movie struct {
	ID        int64     `json:"id"`                // Unique ID for the Movie
	CreatedAt time.Time `json:"-"`                 // Timestamp when movie was added to the Database
	Title     string    `json:"title"`             // Movie Title
	Year      int32     `json:"year,omitempty"`    // Year of Release
	Runtime   int32     `json:"runtime,omitempty"` // 	Movie Runtime (in minutes)
	Genres    []string  `json:"genres,omitempty"`  // Slice of genres for the movie (romance, comedy, etc)
	Version   int32     `json:"version"`           // The version number starts at 1 and will be incremented each time the movie information is updated
}
