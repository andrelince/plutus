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
	GetAccountTransactions(ctx context.Context, accountID uint) ([]model.Transaction, error)
	DeleteAccount(ctx context.Context, userID, accountID uint) error
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
	foundUser := model.User{ID: userID}
	err := r.g.
		WithContext(ctx).
		First(&foundUser).
		Error

	// user must not be deleted
	if err != nil {
		return model.Account{}, err
	}

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
		Unscoped().
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

	foundUser := model.User{ID: account.UserID}
	err = r.g.
		WithContext(ctx).
		First(&foundUser).
		Error

	// user must not be deleted
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

func (r AccountRepo) GetAccountTransactions(ctx context.Context, accountID uint) ([]model.Transaction, error) {
	var transactions []model.Transaction
	err := r.g.
		Where("account_id = ?", accountID).
		Order("created_at desc").
		Find(&transactions).
		Error
	return transactions, err
}

func (r AccountRepo) DeleteAccount(ctx context.Context, userID, accountID uint) error {
	account := model.Account{ID: accountID, UserID: userID}
	err := r.g.
		WithContext(ctx).
		First(&account).
		Error

	if err != nil {
		return err
	}

	res := r.g.
		WithContext(ctx).
		Delete(&model.Account{}, accountID)
	return res.Error
}
