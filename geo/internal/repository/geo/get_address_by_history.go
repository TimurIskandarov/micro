package geo

import (
	"context"
	"geo/internal/model"

	sq "github.com/Masterminds/squirrel"
)

func (r *repository) GetAddressByHistory(ctx context.Context, historyId int) ([]*model.Address, error) {
	q, args, _ := sq.Select("address", "latitude", "longitude").
		From("address").
		Where("history_id = ?", historyId).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	rows, err := r.pool.Query(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	addresses := make([]*model.Address, 0)

	// считывание записей
	for rows.Next() {
		address := new(model.Address)

		err := rows.Scan(
			&address.Value,
			&address.Lat,
			&address.Lng,
		)
		if err != nil {
			return nil, err
		}

		addresses = append(addresses, address)
	}

	// убеждаемся, что прошлись по всему набору строк без ошибок
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return addresses, nil
}
