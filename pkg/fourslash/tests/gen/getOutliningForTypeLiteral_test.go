package fourslash_test

import (
	"testing"

	"github.com/buke/typescript-go-internal/pkg/fourslash"
	"github.com/buke/typescript-go-internal/pkg/testutil"
)

func TestGetOutliningForTypeLiteral(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `type A =[| {
    a: number;
}|]

type B =[| {
   a:[| {
       a1:[| {
           a2:[| {
               x: number;
               y: number;
           }|]
       }|]
   }|],
   b:[| {
       x: number;
   }|],
   c:[| {
       x: number;
   }|]
}|]`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyOutliningSpans(t)
}
