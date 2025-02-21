package data

import "time"

type Movie struct {
	ID        int64     // Unique ID for the Movie
	CreatedAt time.Time // Timestamp when movie was added to the Database
	Title     string    // Movie Title
	Year      int32     // Year of Release
	Runtime   int32     // 	Movie Runtime (in minutes)
	Genres    []string  // Slice of genres for the movie (romance, comedy, etc)
	Version   int32     // The version number starts at 1 and will be incremented each time the movie information is updated
}
