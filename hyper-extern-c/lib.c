#include "lib.h"
#include "string.h"

struct Result verifier(char input[255]) {
  struct Result Verified;
  strcpy(Verified.message, strcat("C Programming", input));
  return Verified;
}
