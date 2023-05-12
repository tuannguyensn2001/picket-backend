package wallet_usecase

import (
	"context"
	"database/sql"
	"github.com/rs/zerolog/log"
	"picket/src/internal/dto"
)

func (u *usecase) Increase(ctx context.Context, input dto.IncreaseBalanceWalletInput) error {
	return u.repository.Transaction(ctx, func(ctx context.Context) error {
		wallet, err := u.repository.FindByUserId(ctx, input.UserId)
		if err != nil {
			log.Error().Err(err).Send()
			return err
		}
		wallet.Balance += input.Amount
		err = u.repository.Save(ctx, wallet)
		if err != nil {
			log.Error().Err(err).Send()
			return err
		}

		return nil
	}, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})

}
