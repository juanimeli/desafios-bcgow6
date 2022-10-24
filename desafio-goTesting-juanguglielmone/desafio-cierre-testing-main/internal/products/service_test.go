package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockRepo struct {
	DataMock   []Product
	ErrService error
}

func (m MockRepo) GetAllBySeller(seller string) ([]Product, error) {
	if m.ErrService != nil {
		return []Product{}, m.ErrService
	}
	return m.DataMock, nil
}

func TestService(t *testing.T) {

	//arrange
	esperado := []Product{{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	},
	}
	myRepoMock := MockRepo{
		DataMock: esperado,
	}
	service := NewService(myRepoMock)

	//act

	resultado, err := service.GetAllBySeller("1")

	//assert
	assert.Nil(t, err)
	assert.Equal(t, esperado, resultado)

}

func TestServiceFail(t *testing.T) {

	//arrange
	esperadoError := errors.New("Ha ocurrido un error")
	myRepoMock := MockRepo{
		DataMock:   nil,
		ErrService: errors.New("Ha ocurrido un error"),
	}
	service := NewService(myRepoMock)

	//act

	resultado, err := service.GetAllBySeller("1")

	//assert

	assert.EqualError(t, err, esperadoError.Error())
	assert.Empty(t, resultado)
}
