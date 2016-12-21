package main

import (
  "crypto/rand"
  "crypto/aes"
  "math/big"
  "fmt"
)

// Number of bits in integer
const N int = 64
// size of AES key
const AES_SIZE int = 16

type ServerKeyEq struct {
  s byte
  t byte
  cw []byte // should be length n
}

// Pseudo-random number generator. Alternative 
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

// 0th position is the most significant bit
// True if bit is 1 and False if bit is 0
func getBit(n uint64, pos uint) int {
  val := (n & (1 << (uint(N) - pos)))
  if val > 0 {
    return 1
  } else {
    return 0
  }
}

func generateTreeEq(a []byte, b byte, n int) (*ServerKeyEq, *ServerKeyEq) {
  k0 := new(ServerKeyEq)
  k1 := new(ServerKeyEq)

  cw := make([][]byte, len(a))
  s0 := make([][]byte, n)
  s1 := make([][]byte, n)
  // get random s0, s1, t0
  s0[0] = make([]byte, AES_SIZE)
  s1[0] = make([]byte, AES_SIZE)
  rand.Read(s0[0])
  rand.Read(s1[0])
  ttemp, _ := rand.Int(rand.Reader, big.NewInt(2))
  t0 := ttemp.Int64() & 1
  t1 := t0 ^ 1

  keys := make([][]byte, 3)
  for i := range keys {
    keys[i] = make([]byte, 16)
    rand.Read(keys[i])
  }
  
  for i:=1; i<=n; i++ {
    gs0 := PRF(s0[i-1], keys)
    gs1 := PRF(s1[i-1], keys)

    fmt.Printf("a[i-1]&1: %x\n", a[i-1]&1)
    // if a_i = 1: keep = r, lose = l
    offset := 0
    // if a_i = 0: keep = l, lose = r    
    if (a[i-1]&1) == 0 {
      offset = 17
    } 
    fmt.Printf("offset: %i\n", offset)

    cw[i-1] = make([]byte, 18)
    for k:=0; k<16; k++ {
      // if a[i]== 0: s0r = gs0[17:33], s1r = gs1[17:33]
      // if a[i]== 1: s0l = gs0[0:16], s1l = gs1[0:16]
      cw[i-1][k] = gs0[k+offset] ^ gs1[k+offset]
    }
    // tlcw := t0L ^ t1L ^ a[i] ^ 1
    cw[i-1][17] = gs0[16] ^ gs1[16] ^ a[i-1] ^ 1
    // trcw := t0R ^ t1R ^ a[i]
    cw[i-1][18] = gs0[33] ^ gs1[33] ^ a[i-1] & 1
    fmt.Printf("tlcw: %x\ntrcw: %x\n", cw[i][17], cw[i][18])
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
  return k0, k1
}

func main() {
  generateTreeEq([]byte{0, 3, 2, 5, 2}, 6, 5);
}