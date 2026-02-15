package mapping

import (
	"github.com/jinzhu/copier"
	"github.com/shopspring/decimal"
)

var defaultOption = copier.Option{
	Converters: []copier.TypeConverter{
		{
			SrcType: decimal.Decimal{},
			DstType: float64(0),
			Fn: func(src interface{}) (interface{}, error) {
				return src.(decimal.Decimal).InexactFloat64(), nil
			},
		},
	},
}

func Copy(dst any, src any) error {
	return copier.CopyWithOption(dst, src, defaultOption)
}
