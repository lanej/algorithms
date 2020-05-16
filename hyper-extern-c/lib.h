struct Result {
  char *message;
  int status;
  char code[3];
  char enabled;
};

char *just_message(char *input);
struct Result just_struct();
struct Result *just_struct_with_input(char *input);
int mutate_struct(struct Result *result);
