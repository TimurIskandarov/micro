package geo

import (
	"context"

	sq "github.com/Masterminds/squirrel"
)

// Если совпадение поискового запроса с сохраненным составляет более 70% символов, то возвращаем адреса, привязанные к истории поиска.
// Если совпадение поискового запроса с сохраненным составляет менее 70% символов, то возвращаем адреса, полученные от сервиса dadata.ru.
func (r *repository) FindQuery(ctx context.Context, query string) (int, error) {
	q, args, _ := sq.Select("id", "history").
		From("search_history").
		Where("levenshtein(history, ?) <= LENGTH(?) * 0.3", query, query).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	var id int
	row := r.pool.QueryRow(ctx, q, args...)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
