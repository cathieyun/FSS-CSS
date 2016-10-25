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
  
  
  fmt.Printf("t1: %x\ngs0: %x\ngs1: %x\n", t1, gs0, gs1)
  return ServerKey {5}, ServerKey {6}
}

func main() {
  Gen(5, 6, 7);
}