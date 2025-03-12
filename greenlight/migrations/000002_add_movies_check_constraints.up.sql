-- To check if runtime is greater than 0 at database level.
ALTER TABLE movies ADD CONSTRAINT movies_runtime_check CHECK (runtime >= 0);

-- To check if the year is between 1888 and current year at database level
ALTER TABLE movies ADD CONSTRAINT movies_year_check CHECK(year BETWEEN 1888 AND date_part('year', now()));

-- To check if values in genres are more than 1 and less than 5 at database level
ALTER TABLE movies ADD CONSTRAINT movies_genres_check CHECK(array_length(genres, 1) BETWEEN 1 AND 5);