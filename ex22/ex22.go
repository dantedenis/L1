package ex22

import (
	"math/big"
)

// BigInt структура для 2х битИнтов, которые внутри основаны на байтовом массиве
type BigInt struct {
	a, b *big.Int
}

// NewBigInt Конструктор от 2х инт64
func NewBigInt(a, b int64) *BigInt {
	return &BigInt{
		a: big.NewInt(a),
		b: big.NewInt(b),
	}
}

// Add Sum (a+b) return new Int
func (b *BigInt) Add() *big.Int {
	return new(big.Int).Add(b.a, b.b)
}

// Sub Differ (a-b) return new Int
func (b *BigInt) Sub() *big.Int {
	return new(big.Int).Sub(b.a, b.b)
}

// Mul Product (a*b) return new Int
func (b *BigInt) Mul() *big.Int {
	return new(big.Int).Mul(b.a, b.b)
}

//Div (a/b) return new Int
func (b *BigInt) Div() *big.Int {
	return new(big.Int).Div(b.a, b.b)
}
