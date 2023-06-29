package notification_usecase

import "context"

func (u *usecase) CountUnreadByUser(ctx context.Context, userId int) (int64, error) {
	return u.repository.CountUnreadByUser(ctx, userId)
}
