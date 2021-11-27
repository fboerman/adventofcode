#ifndef INC_2021_UTILS_H
#define INC_2021_UTILS_H
#include <stdio.h>
#include <stdlib.h>
#include <malloc.h>
#include <string.h>

int count_lines_file(char fname[]);
char** load_file_whole(char fname[], size_t* lines);
int* convert_array_to_int(char** arr, size_t s);

#endif //INC_2021_UTILS_H
