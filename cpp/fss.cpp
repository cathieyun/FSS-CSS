#include "fss.h"

void FSS::generateTreeEq(ServerKey* k0, ServerKey* k1, uint64_t a, uint64_t b) {
  // n = length of input a - TODO, determine how a is represented
  int n = to_string(a).length();
  int lambda = 128;
  // create arrays size 
  // Question: how long to make s and ts?
  unsigned char s0[n];
  unsigned char s1[n];
  unsigned char t0[n];
  unsigned char t1[n];
  

  // sample random s0 <- {0, 1}^lambda
  if(!RAND_bytes(unsigned char *s0, lambda)) {
    printf("Random byte generation for s0 failed\n");
    exit(1);
  }
  // sample random s1 <- {0, 1}^lambda
  if(!RAND_bytes(unsigned char *s1, lambda)) {
    printf("Random byte generation for s1 failed\n");
    exit(1);
  }
  // sample random t0 <- {0, 1}^lambda
  if(!RAND_bytes(unsigned char *t0, 1)) {
    printf("Random byte generation for t0 failed\n");
    exit(1);
  }
  // take t1 <- t0 XOR 1
  *t1 = *t0 ^ 1

  unsigned char *cw[n];
  for (int i = 1; i < n; i++) {
    if (a[i] == 0) {
      s0k || t0k || s0l || t0l = prf(*(s0 + i - 1));
      s1k || t1k || s1l || t1l = prf(*(s1 + i - 1));
    } else {
      s0l || t0l || s0k || t0k = prf(*(s0 + i - 1));
      s1l || t1l || s1k || t1k = prf(*(s1 + i - 1));
    }
    unsigned char scw = s01 ^ s1l;
    unsigned char tlcw = t0l ^ t1l ^ a[i] ^ 1;
    unsigned char trcw = t0r ^ t1r ^ a[i];
    *(cw + i - 1) = scw || tlcw || trcw;

    // TODO: how to save the keep/lose state? or just copy code...
    *(s0 + i - 1) = s0(keep) ^ *(t0 + i - 1) * scw;
    *(s1 + i - 1) = s1(keep) ^ *(t1 + i - 1) * scw;
    *(t1 + i - 1) = t1(keep) ^ *(t1 + i - 1) * t(keep)cw;
    *(t1 + i - 1) = t1(keep) ^ *(t1 + i - 1) * t(keep)cw;
  }
}