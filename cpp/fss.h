#ifndef FSS_H
#define FSS_H

#include "openssl-aes.h"
#include <openssl/rand.h>
#include <string>

struct ServerKey {
  // TODO: flesh out more
  unsigned char s;
};

class FSS {
  private:

  public:
    FSS();
    void generateTreeEq(ServerKey* k0, ServerKey* k1, uint64_t a, uint64_t b);
    void evaluateEq();
};

#endif