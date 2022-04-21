package locality

import "github.com/pkg/errors"

type Code uint
type StatisticalCode uint
type Name string
type Status uint8

const (
	Raion Status = iota + 1
	Oras
	Sector
	Municipiu
	Localitate
	Comuna
	Sat
)

func (t Status) ToString() string {
	switch t {
	case Raion:
		return "Raion"
	case Oras:
		return "Oras"
	case Sector:
		return "Sector"
	case Municipiu:
		return "Municipiu"
	case Localitate:
		return "Localitate"
	case Comuna:
		return "Comuna"
	case Sat:
		return "Sat"
	}
	return ""
}

type Locality struct {
	Code
	StatisticalCode
	Name
	Status
}

func NewLocality(code Code, statisticalCode StatisticalCode, name Name, status Status) (*Locality, error) {
	if code == 0 {
		return nil, errors.New("empty locality-service code")
	}
	if name == "" {
		return nil, errors.New("empty locality-service name")
	}

	return &Locality{
		Code:            code,
		StatisticalCode: statisticalCode,
		Name:            name,
		Status:          status,
	}, nil
}
