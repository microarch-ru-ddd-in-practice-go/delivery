package cmd

type Closer interface {
	Close() error
}
