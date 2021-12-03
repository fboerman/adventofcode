#include "day2.h"

void day2() {
    printf("Day 2\n");
    size_t num_lines;

    char** file = load_file_whole("day2.txt", &num_lines);
    if(file == NULL){
        return;

    }
    printf("Number of lines in input: %zu\n", num_lines);
    int horizontal = 0, depth = 0, depth_2 = 0;
    // for part 2 depth is the aim

    for(int i = 0; i < num_lines; i++) {
        char* dir = strtok(file[i], " ");
        errno = 0;
        int steps = (int)strtol(strtok(NULL, " "), NULL, 10);
        if(errno != 0){
            printf("Could not convert number at line %i, exiting\n", i);
            return;
        }
        if(strcmp(dir, "forward") == 0) {
            horizontal += steps;
            depth_2 += depth*steps;
        } else if(strcmp(dir, "up") == 0) {
            depth -= steps;
        } else if(strcmp(dir, "down") == 0) {
            depth += steps;
        } else {
            printf("Unkown instruction %s at line %i, existing\n", dir, i);
            return;
        }
    }
    printf("Part 1\n");
    printf("Final position: %i, final depth: %i, product: %i\n", horizontal, depth, depth*horizontal);

    printf("Part 2\n");
    printf("Final position: %i, final depth: %i, product: %i\n", horizontal, depth_2, depth_2*horizontal);
}
