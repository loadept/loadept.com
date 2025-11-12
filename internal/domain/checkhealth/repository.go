package checkhealth

import "context"

type CheckHealthRepository interface {
	CheckConnection(ctx context.Context) error
}
