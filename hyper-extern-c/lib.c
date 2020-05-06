#include "lib.h"
#include "malloc.h"
#include "stdio.h"
#include "string.h"

char *just_message(char *input) { return "C Programming"; }

struct Result just_struct() {
  struct Result just = {"foobar"};
  return just;
}

struct Result *just_struct_with_input(char *input) {
  char dest[255];

  strcpy(dest, input);

  struct Result *just = malloc(sizeof(struct Result));
  just->message = malloc(strlen(dest) + 1);
  strcpy(just->message, dest);
  return just;
}
