package main

import (
  "crypto/rand"
  "crypto/aes"
  "math/big"
  "fmt"
)

type ServerKey struct {
  s byte
}

func PRF(x []byte, keys [][]byte) []byte {
  out := make([]byte, 48)
  for i := range keys {
    // get AES_k[i]
    block, err := aes.NewCipher(keys[i])
    if err != nil {
      panic(err.Error())
    }
    // get AES_k[i](x) 
    temp := make([]byte, 16)
    block.Encrypt(temp, x)
    // get AES_k[i](x) ^ x
    for j := range temp {
      out[i*16+j] = temp[j] ^ x[j]
    }
  }
  return out
}

func Gen(a []byte, b byte, n byte) (ServerKey, ServerKey) {
  // get random s0, s1, t0
  s0 := make([]byte, 16)
  s1 := make([]byte, 16)
  t0, _ := rand.Int(rand.Reader, big.NewInt(2))
  rand.Read(s0)
  rand.Read(s1)
  t1 := t0.Int64() ^ 1

  keys := make([][]byte, 3)
  for i := range keys {
    keys[i] = make([]byte, 16)
    rand.Read(keys[i])
  }

  cw := make([][]byte, len(a))
  for i:=1; i<len(a); i++ {
    gs0 := PRF(s0, keys)
    gs1 := PRF(s1, keys)
    fmt.Printf("a[i]&1: %x\n", a[i]&1)
    offset := 0
    if (a[i]&1) == 0 {
      offset = 17
    } 
    fmt.Printf("offset: %i\n", offset)
    cw[i] = make([]byte, 18)
    for k:=0; k<16; k++ {
      // if a[i]== 0: s0r = gs0[17:33], s1r = gs1[17:33]
      // if a[i]== 1: s0l = gs0[0:16], s1l = gs1[0:16]
      cw[i][k] = gs0[k+offset] ^ gs1[k+offset]
    }
    // tlcw := t0L ^ t1L ^ a[i] ^ 1
    tlcw := gs0[16] ^ gs1[16] ^ a[i] ^ 1
    // trcw := t0R ^ t1R ^ a[i]
    trcw := gs0[33] ^ gs1[33] ^ a[i] & 1
    fmt.Printf("tlcw: %x\ntrcw: %x\n", tlcw, trcw)
  }
/*
    unsigned char scw = s01 ^ s1l;
    unsigned char tlcw = t0l ^ t1l ^ a[i] ^ 1;
    unsigned char trcw = t0r ^ t1r ^ a[i];
    *(cw + i - 1) = scw || tlcw || trcw;

    // TODO: how to save the keep/lose state? or just copy code...
    *(s0 + i - 1) = s0(keep) ^ *(t0 + i - 1) * scw;
    *(s1 + i - 1) = s1(keep) ^ *(t1 + i - 1) * scw;
    *(t1 + i - 1) = t1(keep) ^ *(t1 + i - 1) * t(keep)cw;
    *(t1 + i - 1) = t1(keep) ^ *(t1 + i - 1) * t(keep)cw;
*/  

  fmt.Printf("t1: %x\n", t1)
  return ServerKey {5}, ServerKey {6}
}

func main() {
  Gen([]byte{0, 3, 2, 5, 2}, 6, 7);
}