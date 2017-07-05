#include <math.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>
#include <limits.h>
#include <stdbool.h>

int main(){
    int n;
    int k;
    scanf("%d %d",&n,&k);
    int *height = malloc(sizeof(int) * n);
    for(int height_i = 0; height_i < n; height_i++){
       scanf("%d",&height[height_i]);
    }
    int max = INT_MIN;
    for(int i = 0; i < n; i++) {
      if(height[i] > max) {
        max = height[i];
      }
    }

    int num_beverages = 0;

    if(max > k){
      num_beverages = max - k;
    }

    printf("%d\n", num_beverages);

    return num_beverages;
}

