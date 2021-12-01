#include <stdio.h>
#include "2021.h"

int main() {
    char choice;

    printf("Advent of code 2021 Frank Boerman (c)\n");

    printf("Choose day: ");
    scanf("%c", &choice);

    switch(choice) {
        case '1':
            day1();
            break;
        default:
            printf("Invalid choice!\n");
            return -1;
    }

    return 0;
}
