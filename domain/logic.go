package domain

import "io"

type GlobalState struct {
	Accounts   []BankAccount
	Operations []Operation
	Categories []Category
}

type Dumper interface {
	Dump(state GlobalState, writer io.Writer) error
}
type Loader interface {
	Load(reader io.Reader) GlobalState
}
