package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/plutus/banking-api/repositories/model"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=./mocks/mock_account_connector.go -package=mocks github.com/plutus/banking-api/repositories AccountConnector
type AccountConnector interface {
	CreateAccount(ctx context.Context, userID uint) (model.Account, error)
	GetAccountByUserIDAndID(ctx context.Context, userID, accountID uint) (model.Account, error)
	CreateTransaction(ctx context.Context, accountID uint, transaction model.Transaction) (model.Transaction, error)
}

type TransactionSettings struct {
	BaseCurrency   string
	TransactionFee float64
}

type AccountRepo struct {
	g        *gorm.DB
	settings TransactionSettings
}

func NewAccountRepo(g *gorm.DB, settings TransactionSettings) AccountRepo {
	return AccountRepo{
		g:        g,
		settings: settings,
	}
}

func (r AccountRepo) CreateAccount(ctx context.Context, userID uint) (model.Account, error) {
	account := model.Account{UserID: userID}
	res := r.g.
		WithContext(ctx).
		Create(&account)
	return account, res.Error
}

func (r AccountRepo) GetAccountByUserIDAndID(ctx context.Context, userID, accountID uint) (model.Account, error) {
	found := model.Account{ID: accountID, UserID: userID}
	foundRes := r.g.
		WithContext(ctx).
		First(&found)
	return found, foundRes.Error
}

func (r AccountRepo) CreateTransaction(ctx context.Context, accountID uint, transaction model.Transaction) (model.Transaction, error) {
	account := model.Account{ID: accountID}
	err := r.g.
		WithContext(ctx).
		First(&account).
		Error
	if err != nil {
		return model.Transaction{}, err
	}

	tx := r.g.
		Begin(&sql.TxOptions{Isolation: sql.LevelReadCommitted}).
		WithContext(ctx)
	defer tx.Rollback()

	netAmount := transaction.Amount - r.settings.TransactionFee
	isCredit := transaction.Type == "credit"

	conversionRate := 1.0
	if transaction.CurrencyCode != r.settings.BaseCurrency {
		conversionRate, err = r.getConversionRate(transaction.CurrencyCode, r.settings.BaseCurrency, time.Now())
		if err != nil {
			return model.Transaction{}, err
		}
	}

	// convert between currencies
	convertedAmount := netAmount * conversionRate

	if isCredit {
		account.Balance += convertedAmount
	} else {
		if account.Balance < convertedAmount {
			return model.Transaction{}, errors.New("insufficient funds")
		}
		account.Balance -= convertedAmount
	}

	// update account balance
	if res := tx.Save(&account); res.Error != nil {
		return model.Transaction{}, res.Error
	}

	// update transaction
	transaction.AccountID = accountID
	transaction.TransactionFee = r.settings.TransactionFee
	transaction.ConvertedAmount = convertedAmount
	transaction.ConversionRate = conversionRate
	transaction.Status = "completed"

	if res := tx.Create(&transaction); res.Error != nil {
		return model.Transaction{}, res.Error
	}

	return transaction, tx.Commit().Error
}

func (r AccountRepo) getConversionRate(fromCurrencyCode, toCurrencyCode string, transactionDate time.Time) (float64, error) {
	var conversionRate model.CurrencyConversionRate
	err := r.g.
		Where("from_currency_code = ? AND to_currency_code = ? AND effective_date <= ?", fromCurrencyCode, toCurrencyCode, transactionDate).
		Order("effective_date DESC").
		First(&conversionRate).
		Error
	return conversionRate.ConversionRate, err
}
