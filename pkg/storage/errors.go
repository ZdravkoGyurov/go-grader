package storage

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ZdravkoGyurov/go-grader/pkg/errors"
)

const duplicateKeyCode = 11000

func storageError(err error) error {
	if errNoDocuments(err) {
		return errors.Wrap(errors.ErrNoDocuments, err.Error())
	}
	if errIsDuplicate(err) {
		return errors.Wrap(errors.ErrIsDuplicate, err.Error())
	}
	return err
}

func errIsDuplicate(err error) bool {
	var e mongo.WriteException
	if errors.As(err, &e) {
		for _, we := range e.WriteErrors {
			if we.Code == duplicateKeyCode {
				return true
			}
		}
	}
	return false
}

func errNoDocuments(err error) bool {
	return errors.Is(err, mongo.ErrNoDocuments)
}
