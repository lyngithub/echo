package utils

import (
	"github.com/shopspring/decimal"
)

func StringToFloat(s string) float64 {
	temp, _ := decimal.NewFromString(s)
	result, _ := temp.Float64()
	return result
}

func Sub(a string, b string) string {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.Sub(db).StringFixed(10)
}

func SubFix(a string, b string, fix int32) string {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.Sub(db).StringFixed(fix)
}

func SubNotFix(a string, b string) string {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.Sub(db).String()
}

func Div(a string, b string) string {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.Div(db).StringFixed(10)
}

func DivFix(a string, b string, fix int32) string {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.Div(db).StringFixed(fix)
}

func Add(a string, b string) string {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.Add(db).StringFixed(10)
}

func Mul(a string, b string) string {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.Mul(db).StringFixed(10)
}

func MulFix(a string, b string, fix int32) string {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.Mul(db).StringFixed(fix)
}

func CurrencyPrice(a string, currency float32) string {
	da, _ := decimal.NewFromString(a)
	dc := decimal.NewFromFloat(float64(currency))
	dr := dc.Div(decimal.New(100, 0)).Add(decimal.New(1, 0))
	return da.Mul(dr).StringFixed(2)
}

func MulSubRatio(a string, bs ...string) string {
	da, _ := decimal.NewFromString(a)
	dr := decimal.New(1, 0)
	for _, b := range bs {
		db, _ := decimal.NewFromString(b)
		dr = dr.Sub(db)
	}
	return da.Mul(dr).StringFixed(10)
}
func MulAddRatio(a string, bs ...string) string {
	da, _ := decimal.NewFromString(a)
	dr := decimal.New(1, 0)
	for _, b := range bs {
		db, _ := decimal.NewFromString(b)
		dr = dr.Add(db)
	}
	return da.Mul(dr).StringFixed(10)
}

func GE(a string, b string) bool {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.GreaterThanOrEqual(db)
}

func LessThen(a string, b string) bool {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.LessThan(db)
}

func LessThenAndEqual(a string, b string) bool {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.LessThanOrEqual(db)
}

func Greater(a string, b string) bool {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.GreaterThan(db)
}

func Equal(a string, b string) bool {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.Equal(db)
}

func MulUnFixed(a string, b string) string {
	da, _ := decimal.NewFromString(a)
	db, _ := decimal.NewFromString(b)
	return da.Mul(db).String()
}

func Fix(a string, fix int32) string {
	da, _ := decimal.NewFromString(a)
	return da.StringFixed(fix)
}

func RoundDown(a string, fix int32) string {
	da, _ := decimal.NewFromString(a)
	resp := da.RoundDown(fix)

	return resp.StringFixed(fix)
}

func Round(a string, fix int32) string {
	da, _ := decimal.NewFromString(a)
	resp := da.Round(fix)
	return resp.StringFixed(fix)
}
