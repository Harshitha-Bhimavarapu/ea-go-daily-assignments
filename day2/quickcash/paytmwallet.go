package quickcash

type PaytmWallet struct {
	quickcash QuickCash
}

func (paytmWallet *PaytmWallet) WithDrawCash(amount float64) (float64, string) {
	return paytmWallet.quickcash.getCash(amount)
}
