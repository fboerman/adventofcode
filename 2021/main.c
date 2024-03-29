#include <stdio.h>
#include "day1.h"
#include "day2.h"
#include "day3.h"
#include "day5.h"

int main() {
    char choice;

    printf("Advent of code 2021 Frank Boerman (c)\n");

    printf("Choose day: ");
    scanf("%c", &choice); getchar();

    switch(choice) {
        case '1':
            day1();
            break;
        case '2':
            day2();
        break;
        case '3':
            day3();
            break;
        case '5':
            day5();
            break;
        default:
            printf("Invalid choice!\n");
            return -1;
    }

    return 0;
}
