package cairo_platinum_test

import (
	"fmt"
	"os"
	"testing"

	"cairo_platinum"
)

func TestFibonacci5ProofVerifies(t *testing.T) {
	fmt.Println(os.Getwd())
	f, err := os.Open("../../tests/testing_data/fibo_5.proof")
	if err != nil {
		t.Errorf("could not open proof file")
	}

	proofBytes := make([]byte, cairo_platinum.MAX_PROOF_SIZE)
	nReadBytes, err := f.Read(proofBytes)
	if err != nil {
		t.Errorf("could not read bytes from file")
	}

	if !cairo_platinum.VerifyCairoProof100Bits(([cairo_platinum.MAX_PROOF_SIZE]byte)(proofBytes), uint(nReadBytes)) {
		t.Errorf("proof did not verify")
	}
}
