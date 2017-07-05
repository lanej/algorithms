#include <math.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>
#include <limits.h>
#include <stdbool.h>

void solve(int n, int k) {
        if(k == 0) {
                for(int i = 1; i <= n; i++) {
                        printf("%d ", i);
                }
        } else if(k == 1 && n != 3) {
                for(int i = n; i > 0 && i <= n; i--) {
                        printf("%d ", i);
                }
        } else if(k * 2 == n){
                for(int i = k+1; i <= n; i++) {
                        printf("%d ", i);
                }
                for(int i = 1; i <= k; i++) {
                        printf("%d ", i);
                }
        } else {
                puts("-1");
                return;
        }

        puts("");
}

int main(){
        int t;
        scanf("%d",&t);
        for(int i = 0; i < t; i++){
                int n;
                int k;
                scanf("%d %d",&n,&k);
                solve(n, k);
        }
        return 0;
}

