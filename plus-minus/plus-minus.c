#include <math.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>
#include <limits.h>
#include <stdbool.h>

int main(){
        int n;
        scanf("%d",&n);
        int arr[n];
        for(int arr_i = 0; arr_i < n; arr_i++){
                scanf("%d",&arr[arr_i]);
        }

        int positives = 0;
        int negatives = 0;
        int zeroes = 0;

        for(int arr_i = 0; arr_i < n; arr_i++){
                if(arr[arr_i] > 0) {
                        positives++;
                } else if(arr[arr_i] < 0) {
                        negatives++;
                } else {
                        zeroes++;
                }

        }
        float size = (float)n;
        float positive_ratio = positives/size;
        float negative_ratio = negatives/size;
        float zeroes_ratio = zeroes/size;

        printf("%.6f\n", positive_ratio);
        printf("%.6f\n", negative_ratio);
        printf("%.6f\n", zeroes_ratio);

        return 0;
}
