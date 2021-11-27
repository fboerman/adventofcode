#pragma clang diagnostic push
#pragma ide diagnostic ignored "cert-err34-c"
#include "utils.h"

int count_lines_file(char fname[]) {
    FILE * fp;
    int count = 0;
    char c;

    fp = fopen(fname, "r");
    if (fp == NULL){
        printf("Could not open file!\n");
        return -1;
    }

    for(c = (char)getc(fp); c != EOF; c = (char)getc(fp)){
        if(c == '\n') {
            count++;
        }
    }

    return count;
}

char **load_file_whole(char fname[], size_t* num_lines) {
    FILE* fp;
    char* line = NULL;
    size_t len = 0;
    ssize_t read;
    *num_lines = count_lines_file(fname);
    char** file = (char**)malloc(*num_lines * sizeof(char*));
    if (*num_lines == -1){
        return NULL;
    }

    fp = fopen("day1.txt", "r");

    for(int i = 0; i<*num_lines; i++){
        getline(&line, &len, fp);
        line[strcspn(line, "\n")] = '\0';
        file[i] = (char*)malloc(strlen(line) + 1);
        strcpy(file[i], line);
    }

    fclose(fp);

    return file;
}

int* convert_array_to_int(char **arr, size_t s) {
    int* result = (int*)malloc(s*sizeof(int));

    for(int i = 0; i < s;i++) {
        result[i] = atoi(arr[i]);
    }

    return result;
}

#pragma clang diagnostic pop