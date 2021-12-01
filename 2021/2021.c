#include "2021.h"

void day1() {
    printf("Day 1\n");
    size_t num_lines;
    int num_increases=0;

    char** file = load_file_whole("day1.txt", &num_lines);
    if(file == NULL){
        return;

    }
    printf("Number of lines in input: %zu\n", num_lines);
    int* file_i = convert_array_to_int(file, num_lines);

    printf("part 1\n");
    for(int i = 1; i<num_lines; i++) {
        if(file_i[i] > file_i[i-1]) {
            num_increases++;
        }
    }
    printf("Number of increases: %i\n", num_increases);

    printf("part 2\n");
    num_increases = 0;
    int buffer[3] = {file_i[0], file_i[1], file_i[2]};
    int buffer_i = 0;
    for(int i = 3; i<num_lines; i++) {
        int A = sum(buffer, 3);
        buffer[buffer_i] = file_i[i];
        buffer_i = (buffer_i + 1) % 3;
        int B = sum(buffer, 3);
        if(B > A) {
            num_increases++;
        }
    }

    printf("Number of increases for window: %i\n", num_increases);
}
