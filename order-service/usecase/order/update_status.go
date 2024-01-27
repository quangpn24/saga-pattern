package order

import "context"

func (uc *UseCase) UpdateStatus(ctx context.Context, id string, status string) error {
	return uc.repo.UpdateStatus(ctx, id, status, "")
}
