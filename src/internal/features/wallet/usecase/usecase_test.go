package wallet_usecase

//func TestIncrease(t *testing.T) {
//	config, err := config2.GetConfig()
//	if err != nil {
//		log.Fatal().Err(err).Send()
//	}
//	dbTest := config.DbTest
//	base.HookReset(context.TODO(), dbTest)
//	defer func() {
//		base.HookReset(context.TODO(), dbTest)
//	}()
//
//	dbTest.Create(&entities.Wallet{
//		UserId:  1,
//		Balance: 0,
//	})
//
//	repo := wallet_repository.New(dbTest)
//	usecase := New(repo)
//
//	t.Run("increase success", func(t *testing.T) {
//		input := dto.IncreaseBalanceWalletInput{
//			UserId: 1,
//			Amount: 100,
//		}
//		err := usecase.Increase(context.TODO(), input)
//
//		var wallet entities.Wallet
//		dbTest.Where("user_id = ?", input.UserId).First(&wallet)
//
//		require.Nil(t, err)
//		require.Equal(t, input.UserId, wallet.UserId)
//		require.GreaterOrEqual(t, wallet.Balance, input.Amount)
//
//	})
//
//	t.Run("increase success when have multiple increase", func(t *testing.T) {
//		input := dto.IncreaseBalanceWalletInput{
//			UserId: 1,
//			Amount: 100,
//		}
//		var wg sync.WaitGroup
//		for i := 0; i < 10; i++ {
//			wg.Add(1)
//			go func() {
//				defer wg.Done()
//				usecase.Increase(context.TODO(), input)
//			}()
//		}
//
//		wg.Wait()
//		var wallet entities.Wallet
//		dbTest.Where("user_id = ?", input.UserId).First(&wallet)
//
//		require.Nil(t, err)
//		require.Equal(t, input.Amount, wallet.Balance)
//	})
//}
