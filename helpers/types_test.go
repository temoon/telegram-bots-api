package helpers

import (
	"fmt"
	"testing"
)

// region Bool

var testCasesBool = []bool{true, false}

func TestBool(t *testing.T) {
	for _, value := range testCasesBool {
		if ref := Bool(value); ref == nil || *ref != value {
			t.Error("Value:", value, "Got:", *ref)
		}
	}
}

func BenchmarkBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bool(testCasesBool[i%len(testCasesBool)])
	}
}

func ExampleBool() {
	fmt.Println(*Bool(true))
	// Output: true
}

// endregion

// region Int32

var testCasesInt32 = []int32{-42, 0, 42}

func TestInt32(t *testing.T) {
	for _, value := range testCasesInt32 {
		if ref := Int32(value); ref == nil || *ref != value {
			t.Error("Value:", value, "Got:", *ref)
		}
	}
}

func BenchmarkInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int32(testCasesInt32[i%len(testCasesInt32)])
	}
}

func ExampleInt32() {
	fmt.Println(*Int32(42))
	// Output: 42
}

// endregion

// region Int64

var testCasesInt64 = []int64{-42, 0, 42}

func TestInt64(t *testing.T) {
	for _, value := range testCasesInt64 {
		if ref := Int64(value); ref == nil || *ref != value {
			t.Error("Value:", value, "Got:", *ref)
		}
	}
}

func BenchmarkInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int64(testCasesInt64[i%len(testCasesInt64)])
	}
}

func ExampleInt64() {
	fmt.Println(*Int64(42))
	// Output: 42
}

// endregion

// region String

var testCasesString = []string{"", " ", "foo"}

func TestString(t *testing.T) {
	for _, value := range testCasesString {
		if ref := String(value); ref == nil || *ref != value {
			t.Error("Value:", value, "Got:", *ref)
		}
	}
}

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(testCasesString[i%len(testCasesString)])
	}
}

func ExampleString() {
	fmt.Println(*String("foo"))
	// Output: foo
}

// endregion

// region BoolOrFalse

type CaseBoolOrFalse struct {
	Ref    *bool
	Return bool
}

var testCasesBoolOrFalse = []CaseBoolOrFalse{
	{Ref: nil, Return: false},
	{Ref: Bool(true), Return: true},
	{Ref: Bool(false), Return: false},
}

func TestBoolOrFalse(t *testing.T) {
	for _, testCase := range testCasesBoolOrFalse {
		if value := BoolOrFalse(testCase.Ref); value != testCase.Return {
			t.Error("Ref:", testCase.Ref, "Expected:", testCase.Return, "Got:", value)
		}
	}
}

func BenchmarkBoolOrFalse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BoolOrFalse(testCasesBoolOrFalse[i%len(testCasesBoolOrFalse)].Ref)
	}
}

func ExampleBoolOrFalse() {
	fmt.Println(BoolOrFalse(Bool(true)))
	fmt.Println(BoolOrFalse(nil))
	// Output:
	// true
	// false
}

// endregion

// region Int32OrZero

type CaseInt32OrZero struct {
	Ref    *int32
	Return int32
}

var testCasesInt32OrZero = []CaseInt32OrZero{
	{Ref: nil, Return: 0},
	{Ref: Int32(-42), Return: -42},
	{Ref: Int32(0), Return: 0},
	{Ref: Int32(42), Return: 42},
}

func TestInt32OrZero(t *testing.T) {
	for _, testCase := range testCasesInt32OrZero {
		if value := Int32OrZero(testCase.Ref); value != testCase.Return {
			t.Error("Ref:", testCase.Ref, "Expected:", testCase.Return, "Got:", value)
		}
	}
}

func BenchmarkInt32OrZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int32OrZero(testCasesInt32OrZero[i%len(testCasesInt32OrZero)].Ref)
	}
}

func ExampleInt32OrZero() {
	fmt.Println(Int32OrZero(Int32(42)))
	fmt.Println(Int32OrZero(nil))
	// Output:
	// 42
	// 0
}

// endregion

// region Int64OrZero

type CaseInt64OrZero struct {
	Ref    *int64
	Return int64
}

var testCasesInt64OrZero = []CaseInt64OrZero{
	{Ref: nil, Return: 0},
	{Ref: Int64(-42), Return: -42},
	{Ref: Int64(0), Return: 0},
	{Ref: Int64(42), Return: 42},
}

func TestInt64OrZero(t *testing.T) {
	for _, testCase := range testCasesInt64OrZero {
		if value := Int64OrZero(testCase.Ref); value != testCase.Return {
			t.Error("Ref:", testCase.Ref, "Expected:", testCase.Return, "Got:", value)
		}
	}
}

func BenchmarkInt64OrZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int64OrZero(testCasesInt64OrZero[i%len(testCasesInt64OrZero)].Ref)
	}
}

func ExampleInt64OrZero() {
	fmt.Println(Int64OrZero(Int64(42)))
	fmt.Println(Int64OrZero(nil))
	// Output:
	// 42
	// 0
}

// endregion

// region StringOrEmpty

type CaseStringOrEmpty struct {
	Ref    *string
	Return string
}

var testCasesStringOrEmpty = []CaseStringOrEmpty{
	{Ref: nil, Return: ""},
	{Ref: String(""), Return: ""},
	{Ref: String(" "), Return: " "},
	{Ref: String("foo"), Return: "foo"},
}

func TestStringOrEmpty(t *testing.T) {
	for _, testCase := range testCasesStringOrEmpty {
		if value := StringOrEmpty(testCase.Ref); value != testCase.Return {
			t.Error("Ref:", testCase.Ref, "Expected:", testCase.Return, "Got:", value)
		}
	}
}

func BenchmarkStringOrEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringOrEmpty(testCasesStringOrEmpty[i%len(testCasesStringOrEmpty)].Ref)
	}
}

func ExampleStringOrEmpty() {
	fmt.Println(StringOrEmpty(String("foo")))
	fmt.Println(StringOrEmpty(nil))
	// Output: foo
}

// endregion
