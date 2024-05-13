package geo

import (
	"context"
	"geo/internal/model"
)

func (r *repository) GetAddress(ctx context.Context, query string) ([]*model.Address, error) {
	// поиск запроса в базе
	historyId, err := r.FindQuery(ctx, query)
	// если запрос в базе есть, то
	if err == nil {
		addresses, err := r.GetAddressByHistory(ctx, historyId)
		if err != nil {
			return nil, err
		}
		return addresses, nil
	}

	// если запроса в базе нет, то обращаемся к api за адресами
	addresses, err := r.CallExternalApi(ctx, "address", query)
	if err != nil {
		return nil, err
	}

	// сохраняем запрос
	historyId, err = r.SaveQuery(ctx, query)
	if err != nil {
		return nil, err
	}
	// сохраняем связь между запросом и адресами
	if _, err = r.SaveHistoryAddress(ctx, historyId); err != nil {
		return nil, err
	}
	// сохраняем адреса
	for _, address := range addresses {
		if _, err = r.SaveAddress(ctx, historyId, address); err != nil {
			return nil, err
		}
	}

	return addresses, nil
}
