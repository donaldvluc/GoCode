package funding

// Fund : the struct for the fund
type Fund struct {
	// balance is unexported (private), because it's lowercase
	balance int
}

// NewFund : A regular function returning a pointer to a fund
func NewFund(initialBalance int) *Fund {
	// We can returnn a pointer to a new struct without worrying about
	// whether it's on the stack or heap: Go figures that out for us.
	return &Fund{
		balance: initialBalance,
	}
}

// Balance : returns the current fund balance.
// Note : Methods start with a *receiver*, in this case a Fund pointer
func (f *Fund) Balance() int {
	return f.balance
}

// Withdraw : takes out an amount from the current fund.
func (f *Fund) Withdraw(amount int) {
	f.balance -= amount
}
