#include "lib.h"
#include "malloc.h"
#include "stdio.h"
#include "string.h"
#define CANNED "foocannedbar"

char *just_message(char *input) { return "C Programming"; }

struct Result just_struct() {
  struct Result just = {"foobar"};
  return just;
}

struct Result *just_struct_with_input(char *input) {
  char dest[255];

  strcpy(dest, input);
  printf("input From c: '%s'\n", input);

  struct Result *just = malloc(sizeof(struct Result));
  just->message = malloc(strlen(dest) + 1);
  strcpy(just->message, dest);
  return just;
}

int mutate_struct(struct Result *result) {
  strcpy(result->message, strcat(result->message, CANNED));
  result->status = 1;
  printf("From c: '%s'\n", result->code);
  strcpy(result->code, "ABC");
  result->enabled = 'Y';

  return 1;
}
