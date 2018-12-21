package mongo

import (
	"database/sql"

	"github.com/mattrax/Mattrax/internal/mattrax"
)

type policyRepository struct {
	db *sql.DB
}

func (r policyRepository) Create(policy *mattrax.Policy) error {
	return nil
}

func (r policyRepository) Remove(policy *mattrax.Policy) error {
	return nil

}

func (r policyRepository) Find(id mattrax.PolicyID) (*mattrax.Policy, error) {
	return &mattrax.Policy{}, nil
}

func NewPolicyRepository(db *sql.DB) mattrax.PolicyRepository {
	return policyRepository{db}
}
