#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int cmpr(const void *a, const void *b) {
    int* aint = (int*)a;
    int* bint = (int*)b;
    if (*aint > *bint) {
        return 1;
    }
    if (*aint < *bint) {
        return -1;
    }
    return 0;
}


int main() {
    FILE *fp = fopen("input.txt", "r");
    int c;
    int *left_side = malloc(0);
    int *right_side = malloc(0);
    int col = 0;
    char buf[50] = {0};
    int ln = 0;
    int sum = 0;
    while (c = fgetc(fp)) {
        if (c == '\n' || c == EOF) {
            col = 0;
            int i = 0;
            char left[20] = {0};
            char right[20] = {0};
            while (buf[i] != ' ') {
                left[i] = buf[i];
                i++;
            }
            while (buf[i] == ' ') {
                i++;
            }
            int j = 0;
            while (buf[i] != 0) {
                right[j] = buf[i];
                i++;
                j++;
            }
            int left_int = atoi(left);
            int right_int = atoi(right);
            ln++;
            left_side = realloc(left_side, ln * sizeof(int));
            left_side[ln-1] = left_int;
            right_side = realloc(right_side, ln * sizeof(int));
            right_side[ln-1] = right_int;
            memset(buf, 0, 50);
            if (c == EOF) {
                break;
            }
        } else {
            buf[col] = (char)c;
            col++;
        }
    }

    qsort(left_side, ln, sizeof(int), cmpr);
    qsort(right_side, ln, sizeof(int), cmpr);

    /*
    Part 1:
    for (int i = 0; i < ln; i++) {
        sum += abs(right_side[i]-left_side[i]);
    }
    */

    /*
    Part 2:
    */
    for (int i = 0; i < ln; i++) {
        int partial = 0;
        for (int j = 0; j < ln; j++) {
            if (left_side[i] == right_side[j]) {
                partial++;
            }
        }
        sum += partial * left_side[i];
    }

    printf("%d\n", sum);
    fclose(fp);
}