package week08

import (
	"fmt"
	"sync"
)

// Product to be built by an employee
type Product struct {
	Materials Materials

	builtBy Employee
	mu      sync.RWMutex
}

func (p *Product) String() string {
	return p.Materials.String()
}

// BuiltBy returns the employee that built the product.
// A return value of "0" means no employee has built the product yet.
func (p *Product) BuiltBy() Employee {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.builtBy
}

// Build builds the product by the given employee.
// Returns an error if the product is not valid
// Returns an error if the product is already built
// Returns an error if the employee is not valid
func (p *Product) Build(e Employee, w *Warehouse) error {
	// error check

	if err := p.IsValid(); err != nil {
		return err
	}

	if err := e.IsValid(); err != nil {
		return err
	}

	p.mu.RLock()
	// retrieve materials from warehouse
	for k, v := range p.Materials {
		w.Retrieve(k, v)
	}
	p.mu.RUnlock()

	// mark the product as built
	p.mu.RLock()
	p.builtBy = e
	p.mu.RUnlock()

	return nil
}

// IsValid returns an error if the product is invalid.
// A valid product has a quantity > 0.
func (p *Product) IsValid() error {

	p.mu.RLock()
	defer p.mu.RUnlock()

	if len(p.Materials) == 0 {
		return ErrInvalidMaterials(len(p.Materials))
	}

	return nil
}

// IsBuilt returns an error if the product is not built,
// or if the product is invalid.
func (p *Product) IsBuilt() error {
	if err := p.IsValid(); err != nil {
		return err
	}

	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.builtBy == 0 {
		return ErrProductNotBuilt(fmt.Sprintf("product is not built: %v", p))
	}

	return nil
}
