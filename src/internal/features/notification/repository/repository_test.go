package notification_repository

import (
	"context"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"picket/src/base"
	"picket/src/internal/entities"
	"testing"
)

func TestCreate(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	base.TestWithDatabase(func(db *gorm.DB) {
		repo := New(db)

		t.Run("create notification without payload", func(t *testing.T) {
			n := entities.Notification{
				From:     1,
				To:       2,
				Type:     1,
				Template: "test",
			}
			err := repo.Create(context.Background(), &n)
			require.Nil(t, err)
		})

		t.Run("create notification with payload", func(t *testing.T) {
			n := entities.Notification{
				From:     1,
				To:       2,
				Type:     1,
				Template: "test",
				Payload: map[string]string{
					"user_id": "1",
				},
			}
			err := repo.Create(context.Background(), &n)
			require.Nil(t, err)
		})

		t.Run("get notification with payload", func(t *testing.T) {
			n := entities.Notification{
				From:     1,
				To:       2,
				Type:     1,
				Template: "test",
				Payload: map[string]string{
					"user_id": "1",
				},
			}
			_ = repo.Create(context.Background(), &n)

			var n2 entities.Notification
			_ = db.Model(&entities.Notification{}).Last(&n2).Error
			require.Equal(t, "1", n2.Payload["user_id"])

		})
	})
}
