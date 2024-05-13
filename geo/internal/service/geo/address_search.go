package geo

import (
	"context"
	"encoding/json"

	"geo/internal/model"

	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
)

func (s *service) AddressSearch(ctx context.Context, in model.SearchIn) (model.SearchOut, error) {
	var addresses []*model.Address
	var err error

	conn := s.cache.Get()
	defer conn.Close()

	// Проверка наличия значения в кэше
	cachedAddresses, err := redis.String(conn.Do("GET", in.Query))
	if err == redis.ErrNil {
		s.logger.Info("Ключа in.Query в кэше нет", zap.String("key", in.Query))
	} else if err == nil {
		err = json.Unmarshal([]byte(cachedAddresses), &addresses)
		if err == nil {
			// Значение найдено в кэше
			s.logger.Info("адреса получены из кэша", zap.Error(err))
			return model.SearchOut{Addresses: addresses}, nil
		}
	}

	// Значение отсутствует в кэше, делегируем запрос базе данных
	addresses, err = s.repo.GetAddress(ctx, in.Query)
	if err != nil {
		s.logger.Error("адреса не были получены", zap.Error(err))
		return model.SearchOut{}, err
	}

	// Сохранение адресов в кэше
	if jsonAddresses, err := json.Marshal(addresses); err == nil {
		_, err = conn.Do("SET", in.Query, jsonAddresses)
		if err != nil {
			s.logger.Error("адреса в кэше не сохранились", zap.Error(err))
			return model.SearchOut{}, err
		}
	}

	return model.SearchOut{Addresses: addresses}, nil
}
