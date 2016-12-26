package main

import (
  "crypto/rand"
  "crypto/aes"
  "math/big"
  "fmt"
  "log"
)

// Number of bits in integer
const N int = 64
// size of AES key
const AES_SIZE int = 16

type CWEq struct {
  scw []byte
  tlcw byte
  trcw byte
}

type ServerKeyEq struct {
  s []byte
  t byte 
  cw []CWEq // should be length n
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

func generateTreeEq(a, b uint64) (*ServerKeyEq, *ServerKeyEq, *big.Int) {
  k0 := new(ServerKeyEq)
  k1 := new(ServerKeyEq)

  k0.cw = make([]CWEq, N)
  k1.cw = make([]CWEq, N)

  // get random s0, s1, t0
  s0 := make([]byte, AES_SIZE)
  s1 := make([]byte, AES_SIZE)
  rand.Read(s0)
  rand.Read(s1)
  ttemp, _ := rand.Int(rand.Reader, big.NewInt(2))
  t0 := byte(ttemp.Int64()) & 1
  t1 := t0 ^ 1

  k0.s = s0
  k1.s = s1
  k0.t = t0
  k1.t = t1

  keys := make([][]byte, 3)
  for i := range keys {
    keys[i] = make([]byte, 16)
    rand.Read(keys[i])
  }
  
  for i:=0; i<N; i++ {
    gs0 := PRF(s0, keys)
    gs1 := PRF(s1, keys)

    a_i := getBit(a, uint(i))
    fmt.Printf("a_i: %x\n", a_i)
    // if a_i = 0: keep = l, lose = r
    keep := 0
    lose := 17
    // else: keep = r, lose = l   
    if (a_i == 1) {
      keep = 17
      lose = 0
    } 
    k0.cw[i].scw = make([]byte, AES_SIZE)
    k1.cw[i].scw = make([]byte, AES_SIZE)
    for j:=0; j<16; j++ {
      // if a[i]== 0: s0r = gs0[17:33], s1r = gs1[17:33]
      // if a[i]== 1: s0l = gs0[0:16], s1l = gs1[0:16]
      k0.cw[i].scw[j] = gs0[j+lose] ^ gs1[j+lose]
      k1.cw[i].scw[j] = gs0[j+lose] ^ gs1[j+lose]
    }
    fmt.Printf("scw: %x\n", k0.cw[i].scw)
    fmt.Printf("gs0: %x\ngs1: %x\na[i]: %x\n", gs0[16], gs1[16], a_i)
    
    // tlcw := t0L ^ t1L ^ a[i] ^ 1
    k0.cw[i].tlcw = gs0[16] ^ gs1[16] ^ byte(a_i) ^ 1
    k1.cw[i].tlcw = gs0[16] ^ gs1[16] ^ byte(a_i) ^ 1
    // trcw := t0R ^ t1R ^ a[i]
    k0.cw[i].trcw = gs0[33] ^ gs1[33] ^ byte(a_i)
    k1.cw[i].trcw = gs0[33] ^ gs1[33] ^ byte(a_i)
    fmt.Printf("tlcw: %x\ntrcw: %x\n", k0.cw[i].tlcw, k0.cw[i].trcw)
  
    // s_b^i = s_b^keep xor t_b^(i-1) for b = 0, 1
    for k:=0; k<16; k++ {
      s0[k] = gs0[k+keep] * k0.cw[i].scw[k]
      s1[k] = gs1[k+keep] * k1.cw[i].scw[k]
    }
    // TODO: which bit get xor'ed with t? least or most? Implement.
    tcw_keep := k0.cw[i].tlcw
    if (a_i == 1) {
      tcw_keep = k0.cw[i].trcw
    }
    t0 = gs0[16+keep] ^ t0 * tcw_keep
    t1 = gs1[16+keep] ^ t1 * tcw_keep 
  }
  // CW^n = (-1)^t_1^n[beta - convert(s0) + convert(s1)]
  // Generate 64 bit prime for abelien group
  p, err := rand.Prime(rand.Reader, N)
  if err != nil {
    log.Fatal(err)
  }
  k0.cw[AES_SIZE].scw = make([]byte, AES_SIZE)




  // TODO: get CW^n
  return k0, k1, p
}

func main() {
  generateTreeEq(27, 6);
}