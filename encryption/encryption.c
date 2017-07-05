#include <math.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>
#include <limits.h>
#include <stdbool.h>

int main(){
        char* s = (char *)malloc(10240 * sizeof(char));
        scanf("%s",s);

        int len = strlen(s);

        double sq = sqrt(len);
        int rows = floor(sq);
        int cols = ceil(sq);

        if(rows * cols < len){
                rows++;
        }

        int index;

        for(int i = 0; i < cols; i++){
                for(int j = 0; j < rows; j++){
                        index = (j * cols) + i;
                        if(index < len) {
                                printf("%c", s[index]);
                        }
                }
                printf(" ");
        }

        return 0;
}
