#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <memory.h>

int main() {
    FILE *fp = fopen("input.txt", "r");
    if (fp == NULL) {
        return EXIT_FAILURE;
    }
    int c = 0;
    int ln = 0;
    int col = 0;
    char line[100];
    int safe = 0;
    while (1) {
        c = fgetc(fp);
        if (c == '\n' || c == EOF) {
            col = 0;
            char *tok = strtok(line, " ");
            char *ntok = strtok(NULL, " ");
            int *diffs = malloc(0);
            int i = 0;
            while (ntok != NULL) {
                int a = atoi(tok);
                int b = atoi(ntok);
                int dif = b - a;
                diffs = realloc(diffs, sizeof(int) * (i+1));
                diffs[i] = dif;
                strcpy(tok, ntok);
                ntok = strtok(NULL, " ");
                i++;
            }
            int ok = 1;
            for (int j = 0; j < i; j++) {
                if (abs(diffs[j]) < 1 || abs(diffs[j]) > 3) {
                    ok = 0;
                    break;
                }
            }

            int dec = 1;
            for (int j = 0; j < i; j++) {
                if (diffs[j] > 0) {
                    dec = 0;
                    break;
                }
            }

            int inc = 1;
            for (int j = 0; j < i; j++) {
                if (diffs[j] < 0) {
                    inc = 0;
                    break;
                }
            }

            if (!dec && !inc) {
                ok = 0;
            }
            if (dec && inc) {
                ok = 0;
            }

            safe += ok;
            ln++;
            if (c == EOF) {
                break;
            }
        } else {
            line[col] = c;
            col++;
        }
    }
    printf("%d\n", safe);
    fclose(fp);
}