package geo

import (
	"context"
	"geo/internal/model"
	"strconv"

	sq "github.com/Masterminds/squirrel"
)

func (r *repository) SaveAddress(ctx context.Context, historyId int, address *model.Address) (int, error) {
	latitude, _ := strconv.ParseFloat(address.Lat, 64)
	longitude, _ := strconv.ParseFloat(address.Lng, 64)

	q, args, _ := sq.Insert("address").
		Columns(
			"history_id",
			"address",
			"latitude",
			"longitude",
		).
		Values(
			historyId,
			address.Value,
			latitude,
			longitude,
		).
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
