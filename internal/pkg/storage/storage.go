package storage

import (
	"strconv"

	"go.uber.org/zap"
)

type Storage struct {
	inner  map[string]string
	logger *zap.Logger
}

func NewStorage() (Storage, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return Storage{}, err
	}

	defer logger.Sync()
	logger.Info("created new storage")

	return Storage{
		inner:  make(map[string]string),
		logger: logger,
	}, nil
}

func (r Storage) Set(key, value string) {
	r.logger.Info("key set")
	r.logger.Sync()
	r.inner[key] = value
}

func (r Storage) Get(key string) *string {
	res, ok := r.inner[key]
	if !ok {
		return nil
	}
	r.logger.Info("key obtained", zap.String("key", key), zap.String("value", res))
	r.logger.Sync()
	return &res
}

func (r Storage) GetKind(key string) string {
	value, ok := r.inner[key]

	if !ok {
		return "not found key"
	}
	r.logger.Info("key obtained", zap.String("key", key), zap.String("value", value))
	_, err := strconv.ParseFloat(value, 64)

	if err == nil {
		return "D"
	} else {
		return "S"
	}
}
