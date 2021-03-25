package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ZdravkoGyurov/go-grader/pkg/errors"
	"github.com/ZdravkoGyurov/go-grader/pkg/log"
)

func (s *Storage) Transaction(ctx context.Context, f func(ctx context.Context) error) error {
	err := s.mongoClient.UseSession(ctx, func(sessionCtx mongo.SessionContext) error {
		err := sessionCtx.StartTransaction()
		if err != nil {
			err = errors.Wrap(err, "failed to start transaction")
			return err
		}

		if err = f(sessionCtx); err != nil {
			if abortTransactionErr := sessionCtx.AbortTransaction(sessionCtx); abortTransactionErr != nil {
				abortTransactionErr = errors.Wrap(abortTransactionErr, "failed to abort transaction")
				return err
			}
			err = errors.Wrap(err, "failed to execute transaction operations")
			return err
		}

		if commitTransactionErr := sessionCtx.CommitTransaction(sessionCtx); commitTransactionErr != nil {
			err = errors.Wrap(err, "failed to commit transaction")
			return err
		}
		return nil
	})
	if err != nil {
		err = errors.Wrap(err, "failed to execute transaction")
		log.Error().Println(err)
		return err
	}

	return nil
}
