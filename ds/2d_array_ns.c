#include <math.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>
#include <limits.h>
#include <stdbool.h>

int main(){
    int *arr[6];
    for (int i=0; i<6; i++)
         arr[i] = (int *)malloc(6 * sizeof(int));

    for(int arr_i = 0; arr_i < 6; arr_i++){
       for(int arr_j = 0; arr_j < 6; arr_j++){
          scanf("%d",&arr[arr_i][arr_j]);
       }
    }

    int sum = INT_MIN;
    for(int arr_i = 0; arr_i < 4; arr_i++){
       for(int arr_j = 0; arr_j < 4; arr_j++){
          int h_sum = 0;

          for(int h_j = 0; h_j < 3; h_j++){
            h_sum += arr[arr_i][arr_j + h_j];
          }

          h_sum += arr[arr_i+1][arr_j+1];

          for(int h_j = 0; h_j < 3; h_j++){
            h_sum += arr[arr_i + 2][arr_j + h_j];
          }

          if(h_sum > sum) {
            sum = h_sum;
          }
       }
    }

    printf("%d", sum);
}


