package cmd

import (
	"log"
)

type CompositionRoot struct {
	configs Config

	closers []Closer
}

func NewCompositionRoot(configs Config) *CompositionRoot {
	return &CompositionRoot{
		configs: configs,
	}
}

///////////////////////////////////////////////////////////
//////////////////// LIFECYCLE ////////////////////////////
///////////////////////////////////////////////////////////

func (cr *CompositionRoot) RegisterCloser(c Closer) {
	cr.closers = append(cr.closers, c)
}

func (cr *CompositionRoot) CloseAll() {
	for _, c := range cr.closers {
		if err := c.Close(); err != nil {
			log.Printf("close error: %v", err)
		}
	}
}
