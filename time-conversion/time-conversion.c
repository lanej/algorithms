#include <math.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>
#include <limits.h>
#include <stdbool.h>

int main() {
        int hours, mins, seconds;
        char* offset = (char *)malloc(sizeof(char) * 2);

        scanf("%d:%d:%d%s", &hours, &mins, &seconds, offset);

        if(strncmp(offset, "PM", 2) == 0 && hours < 12){
                hours += 12;
        }

        if(strncmp(offset, "AM", 2) == 0 && hours >= 12){
                hours -= 12;
        }

        printf("%.2d:%.2d:%.2d\n", hours, mins, seconds);

        return 0;
}

