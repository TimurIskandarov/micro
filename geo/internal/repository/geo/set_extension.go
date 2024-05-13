package geo

import "context"

func (r *repository) SetExtension(ext string) error {
	_, err := r.pool.Exec(context.Background(), "CREATE EXTENSION IF NOT EXISTS "+ext)
	if err != nil {
		return err
	}

	return nil
}
