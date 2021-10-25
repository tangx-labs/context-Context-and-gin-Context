package main

import (
	"context"
	"fmt"
	"testing"
)

type StringAsKey string

func Test_Main(t *testing.T) {

	ctx := context.TODO()

	key1 := StringAsKey("1")
	ctx = WithContext(ctx, key1, "1")

	FromContext(ctx, key1, "StringAsKey: ")
	FromContext(ctx, "1", "string: ")

}

func WithContext(ctx context.Context, key interface{}, val string) context.Context {
	return context.WithValue(ctx, key, val)
}

func FromContext(ctx context.Context, key interface{}, msg string) {
	v := ctx.Value(key)

	val, ok := v.(string)
	if ok {
		fmt.Println(msg, val)
		return
	}

	fmt.Println(msg, "invalid key")

}
