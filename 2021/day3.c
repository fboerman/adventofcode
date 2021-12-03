#include "day3.h"



void day3() {
    printf("Day 3\n");
    size_t num_lines;

    FILE* fp = fopen("day3.txt", "r");
    if (fp == NULL){
        printf("Could not open file!\n");
        return;
    }

    //find the end of the file
    fseek(fp, 0, SEEK_END);
    // get the byte number of the end
    long num_chars = ftell(fp);
    // set points back to beginning
    rewind(fp);

    printf("Input has %i bytes\n", num_chars);

    char* file_bytes = (char*)malloc(num_chars*sizeof(char));
    fread(file_bytes, num_chars, 1, fp);
    fclose(fp);
    //find the length of the word by looking at length of first line
    int word_len = strcspn(file_bytes, "\n");

    printf("Word length is %i\n", word_len);

    char* gamma_word = (char*)malloc((word_len+1)*sizeof(char));
    gamma_word[word_len] = '\0';

    // iterate through all positions
    for(int z = 0; z<word_len; z++)
    {
        // count the number of ones
        int ones = 0;
        for(int i = z; i < num_chars; i+=word_len+1) {
            if(file_bytes[i] == '1'){
                ones++;
            }
        }
        if(ones > (num_chars/(word_len+1)) >> 1) {
            gamma_word[z] = '1';
        } else {
            gamma_word[z] = '0';
        }
    }
    printf("Gamma word is: %s\n", gamma_word);
    errno = 0;
    unsigned int gamma_number = (unsigned int)strtol(gamma_word, NULL, 2);
    if(errno != 0){
        printf("Could not convert gamma word\n");
        return;
    }

    unsigned int epsilon_number = (~ gamma_number) & ((1 << word_len)-1);

    printf("Gamma number is %u, epsilon number is %u, power consumption is %u\n", gamma_number, epsilon_number,
           gamma_number*epsilon_number);

}