package domain

import "context"

type ValidationRepository interface {
	RecordExists(ctx context.Context, params ExistsParams) (bool, error)
	ValidateExistence(ctx context.Context, params ExistsParams) (bool, error)
}
