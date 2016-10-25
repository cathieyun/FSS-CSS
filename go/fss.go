package main

import (
  "crypto/rand"
  // "crypto/aes"
  "math/big"
  "fmt"
)

type ServerKey struct {
  s byte
}

func PRF(x []byte, k [][]byte) []byte {
  out := make([]byte, 48)
  for i := range keys {
    // get AES_k[i](x) 

    block, err := aes.NewCipher(keys[i])
    if err != nil {
      panic(err.Error())
    }   
    AESx :=  

    // get AES_k[i](x) ^ xw
    out[i*16:(i+1)*16-1] = AESx ^ x
  }
  return out
}

func Gen(a byte, b byte, n byte) (ServerKey, ServerKey) {
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
  gs0 := PRF(s0, keys)
  gs1 := PRF(s1, keys)

  return ServerKey {5}, ServerKey {6}
}

func main() {
  Gen(5, 6, 7);
}