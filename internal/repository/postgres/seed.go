package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Seed(ctx context.Context, db *pgxpool.Pool) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if err := seedMovies(ctx, tx); err != nil {
		return err
	}

	if err := seedShows(ctx, tx); err != nil {
		return err
	}

	if err := seedShowSeats(ctx, tx); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func seedMovies(ctx context.Context, tx pgx.Tx) error {
	movies := []struct {
		title       string
		description string
		duration    int
	}{
		{
			title:       "Interstellar",
			description: "Sci-fi adventure",
			duration:    169,
		},
		{
			title:       "The Dark Knight",
			description: "Batman vs Joker",
			duration:    152,
		},
	}

	for _, movie := range movies {
		_, err := tx.Exec(
			ctx,
			`
			INSERT INTO movies (title, description, duration_mins)
			VALUES ($1, $2, $3)
			ON CONFLICT DO NOTHING
			`,
			movie.title,
			movie.description,
			movie.duration,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func seedShows(ctx context.Context, tx pgx.Tx) error {
	rows, err := tx.Query(
		ctx,
		`SELECT id, title FROM movies`,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	now := time.Now()

	for rows.Next() {
		var movieID string
		var title string

		if err := rows.Scan(&movieID, &title); err != nil {
			return err
		}

		var start1 time.Time
		var start2 time.Time

		switch title {
		case "Interstellar":
			start1 = time.Date(
				now.Year(), now.Month(), now.Day(),
				10, 0, 0, 0, now.Location(),
			)

			start2 = time.Date(
				now.Year(), now.Month(), now.Day(),
				18, 0, 0, 0, now.Location(),
			)

		case "The Dark Knight":
			start1 = time.Date(
				now.Year(), now.Month(), now.Day(),
				14, 0, 0, 0, now.Location(),
			)

			start2 = time.Date(
				now.Year(), now.Month(), now.Day(),
				21, 0, 0, 0, now.Location(),
			)

		default:
			continue
		}

		shows := []time.Time{
			start1,
			start2,
		}

		for _, start := range shows {
			end := start.Add(3 * time.Hour)

			_, err := tx.Exec(
				ctx,
				`
				INSERT INTO shows (
					movie_id,
					start_time,
					end_time
				)
				VALUES ($1, $2, $3)
				ON CONFLICT DO NOTHING
				`,
				movieID,
				start,
				end,
			)
			if err != nil {
				return err
			}
		}
	}

	return rows.Err()
}

func seedShowSeats(ctx context.Context, tx pgx.Tx) error {
	rows, err := tx.Query(
		ctx,
		`SELECT id FROM shows`,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	var showIDs []string

	for rows.Next() {
		var showID string

		if err := rows.Scan(&showID); err != nil {
			return err
		}

		showIDs = append(showIDs, showID)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	rowLabels := []string{"A", "B", "C", "D", "E"}

	for _, showID := range showIDs {
		for _, rowLabel := range rowLabels {
			for seatNumber := 1; seatNumber <= 8; seatNumber++ {
				seatLabel := fmt.Sprintf("%s%d", rowLabel, seatNumber)

				_, err := tx.Exec(
					ctx,
					`
					INSERT INTO show_seats (
						show_id,
						row_label,
						seat_number,
						seat_label
					)
					VALUES ($1, $2, $3, $4)
					ON CONFLICT DO NOTHING
					`,
					showID,
					rowLabel,
					seatNumber,
					seatLabel,
				)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
