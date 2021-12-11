package proccess_transaction

import (
	mock_brokers "FullCycle/adapter/brokers/mock"
	"FullCycle/domain/entity"
	mock_repository "FullCycle/domain/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	TOPIC = "transactions_result"
)

func TestProccessTransaction_ExecuteInvalidCreditCard(t *testing.T) {
	input := TransactionDtoInput{
		"1",
		"1",
		"4000000000",
		"Jonas Silveira",
		05,
		time.Now().Year(),
		123,
		200,
	}

	expectedOutput := TransactionDtoOutput{
		"1",
		entity.REJECT,
		"Numero de cartao invalido",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().Insert(input.ID,
		input.AccountID,
		expectedOutput.Status,
		expectedOutput.ErrorMessage,
		input.Amount).Return(nil)
	producerMock := mock_brokers.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.ID), TOPIC)

	usecase := NewProcessTransaction(repositoryMock, producerMock, TOPIC)
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)

}

func TestProccessTransaction_ExecuteRejectedTransaction(t *testing.T) {
	input := TransactionDtoInput{
		"1",
		"1",
		"4193523830170205",
		"Jonas Silveira",
		05,
		time.Now().Year(),
		123,
		10000,
	}

	expectedOutput := TransactionDtoOutput{
		"1",
		entity.REJECT,
		"you dont have limit enough",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().Insert(input.ID,
		input.AccountID,
		expectedOutput.Status,
		expectedOutput.ErrorMessage,
		input.Amount).Return(nil)

	producerMock := mock_brokers.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.ID), TOPIC)

	usecase := NewProcessTransaction(repositoryMock, producerMock, TOPIC)
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)

}

func TestProccessTransaction_ExecuteAprovedTransaction(t *testing.T) {
	input := TransactionDtoInput{
		"1",
		"1",
		"4193523830170205",
		"Jonas Silveira",
		05,
		time.Now().Year(),
		123,
		100,
	}

	expectedOutput := TransactionDtoOutput{
		"1",
		entity.APPROVED,
		"",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().Insert(input.ID,
		input.AccountID,
		expectedOutput.Status,
		expectedOutput.ErrorMessage,
		input.Amount).Return(nil)

	producerMock := mock_brokers.NewMockProducerInterface(ctrl)
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.ID), TOPIC)

	usecase := NewProcessTransaction(repositoryMock, producerMock, TOPIC)
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)

}
