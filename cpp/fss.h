#ifndef FSS_H
#define FSS_H

#include "openssl-aes.h"

#include <openssl/rand.h>

class FSS {
  private:



  public:
    FSS();
    void generateTreeEq(ServerKeyEq* k0, ServerKeyEq* k1);
    void evaluateEq();
}

struct ServerKey {
  // TODO: flesh out more
  unsigned char s;
};