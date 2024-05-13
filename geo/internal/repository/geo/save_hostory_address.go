package geo

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

func (r *repository) SaveHistoryAddress(ctx context.Context, history_id int) (int, error) {
	q, args, _ := sq.Insert("history_search_address").
		Columns("history_id").
		Values(history_id).
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
