package main

import "testing"

func TestServe(t *testing.T) {
    want := "Syava pidor"
    if got := Serve(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
