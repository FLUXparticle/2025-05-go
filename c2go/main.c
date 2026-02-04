#include <stdio.h>
#include "libmylib.h"

// Befehle zum Bauen (macOS):
// go build -buildmode=c-shared -o libmylib.dylib mylib.go
// gcc -o app main.c -L. -lmylib

int main() {
    int sum = Add(2, 40);
    printf("Add: %d\n", sum);

    int p = IsPrime(17);
    printf("IsPrime(17): %d\n", p);

    return 0;
}
