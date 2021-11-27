#include "2020practice.h"

int day1() {
    printf("Day 1\n");
    size_t num_lines;
    int product;

    char** file = load_file_whole("day1.txt", &num_lines);
    if(file == NULL){
        return -1;

    }
    printf("Number of lines in input: %zu\n", num_lines);
    int* file_i = convert_array_to_int(file, num_lines);

    printf("Trying to find entires summing to 2020\n");
    for(int i = 0; i<num_lines; i++) {
        for(int j=i+1;j<num_lines;j++) {
            if(file_i[i]+file_i[j] == 2020){
                product = file_i[i]*file_i[j];
                printf("Lines %i and %i sum to 2020, product is: %i\n", i, j, product);
                return product;
            }
        }
    }
    return -1;
}
