#include "fss.h"

void FSS::generateTreeEq(int lambda, int alpha, int beta) {
  // create arrays size 
  unsigned char s0[32];
  unsigned char s1[32];
  unsigned char t0[2];
  unsigned char t1[2];

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

  // for i = n-1 do
  //   let g(s_(a_i)^b[i] = s_0^b || s_1^b || t_0^b || t_1^b)
  //   where s_0^b, s_1^b, t_0^b, t_1^b \in {0, 1}
  
  // NOTE: Where does n come from?
  for (int i = 1; i < n; i++) {
    prf(*(s0 + i - 1));
    prf(*(s1 + i - 1));
  }

}