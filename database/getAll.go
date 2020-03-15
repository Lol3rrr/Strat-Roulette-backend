package database

import (
	"context"
	"time"
)

// GetAll loads all the entries matching the given filter
func (s *session) GetAll(query map[string]interface{}, results interface{}) error {
	if !s.isConnectionAlive() {
		err := s.connect()
		if err != nil {
			return err
		}
	}

	ctx, cancelCtx := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelCtx()

	filter := convertQueryToPrimitive(query)
	cur, err := s.MongoCollection.Find(ctx, filter)
	if err != nil {
		return err
	}

	return cur.All(context.TODO(), results)
}
