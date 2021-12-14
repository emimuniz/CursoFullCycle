package process_transaction

import (
	"github.com/emimuniz/imersao5-gateway/domain/entity"
	mock_repository "github.com/emimuniz/imersao5-gateway/domain/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)


func TestProcessTransaction_ExecuteInvalidCreditCard(t *testing.T){
	input := TransactionDtoInput{
		ID:  "1",
		AccountID: "1",
		CreditCardNumber: "40000000000000",
		CreditCardName: "Wesley Silva",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear: time.Now().Year(),
		CreditCardCVV: 123,
		Amount: 200,
	}

	expectedOutput := TransactionDtoOutPut{
		ID:  "1",
		Status: entity.REJECTED,
		ErrorMessage: "invalid credit card number",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.ID, input.AccountID, input.Amount, expectedOutput.Status, expectedOutput.ErrorMessage).
		Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)
}