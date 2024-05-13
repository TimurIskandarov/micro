package geo

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (r *repository) SaveQuery(ctx context.Context, query string) (int, error) {
	q, args, _ := sq.Insert("search_history").
		Columns("history").
		Values(query).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	var id int
	row := r.pool.QueryRow(ctx, q, args...)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
