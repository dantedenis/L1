package ex22

import (
	"math/big"
)

type BigInt struct {
	a, b *big.Int
}

func NewBigInt(a, b int64) *BigInt {
	return &BigInt{
		a: big.NewInt(a),
		b: big.NewInt(b),
	}
}

//Sum (a+b) return new Int
func (b *BigInt) Add() *big.Int {
	return new(big.Int).Add(b.a, b.b)
}

//Differ (a-b) return new Int
func (b *BigInt) Sub() *big.Int {
	return new(big.Int).Add(b.a, b.b)
}

//Product (a*b) return new Int
func (b *BigInt) Mul() *big.Int {
	return new(big.Int).Mul(b.a, b.b)
}

//Div (a/b) return new Int
func (b *BigInt) Div() *big.Int {
	return new(big.Int).Div(b.a, b.b)
}