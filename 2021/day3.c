#include <stdbool.h>
#include "day3.h"

bool check_blocks(char* b1, char* b2, size_t n) {
    if(n == 0) {
        return true;
    }
    for(int i = 0; i<n; i++) {
        if(b1[i] != b2[i]) {
            return false;
        }
    }
    return true;
}

unsigned int msb_word(char* bytes, int word_len, int num_chars, bool filter, bool inverted) {
    char* word = (char*)malloc((word_len + 1) * sizeof(char));
    word[word_len] = '\0';
    int num_lines = num_chars/(word_len+1);

    // iterate through all positions
    for(int z = 0; z<word_len; z++)
    {
        // count the number of ones
        int ones = 0;
        int skipped = 0;
        char* current = NULL;
        for(int i = z; i < num_chars; i+=word_len+1) {
            if(filter) {
                if(!check_blocks(&bytes[i-z], word, z*sizeof(char))) {
                    skipped++;
                    continue;
                }
            }
            current = &bytes[i-z];
            if(bytes[i] == '1'){
                ones++;
            }
        }
        int num_left = num_lines - skipped;
        if(num_left == 1) {
            // there was only one left so copy this into word and be done
            memcpy(word, current, word_len);
            break;
        }
        if(ones > num_left >> 1 || (num_left % 2 == 0 && ones == num_left >> 1)) {
            if(inverted) {
                word[z] = '0';
            } else {
                word[z] = '1';
            }
        } else {
            if(inverted) {
                word[z] = '1';
            } else {
                word[z] = '0';
            }
        }
        printf("Word is: %s\n", word);
    }

    errno = 0;
    unsigned int number = (unsigned int)strtol(word, NULL, 2);
    if(errno != 0){
        printf("Could not convert word\n");
        return -1;
    }
    free(word);
    return number;
}

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
    int num_chars = (int)ftell(fp);
    // set points back to beginning
    rewind(fp);

    printf("Input has %i bytes\n", num_chars);

    char* file_bytes = (char*)malloc(num_chars*sizeof(char));
    fread(file_bytes, num_chars, 1, fp);
    fclose(fp);
    //find the length of the word by looking at length of first line
    int word_len = strcspn(file_bytes, "\n");

    printf("Word length is %i\n", word_len);

    unsigned int gamma_number = msb_word(file_bytes, word_len, num_chars, false, false);

    unsigned int epsilon_number = (~ gamma_number) & ((1 << word_len)-1);

    printf("Gamma number is %u, epsilon number is %u, power consumption is %u\n", gamma_number, epsilon_number,
           gamma_number*epsilon_number);

    printf("part 2\n");

    unsigned int oxygen_number = msb_word(file_bytes, word_len, num_chars, true, false);
    unsigned int co_number = msb_word(file_bytes, word_len, num_chars, true, true);
    printf("Oxygen number is %u, co2 scrubber rate is %u, life support rating is %u\n", oxygen_number, co_number,
           oxygen_number*co_number);


    free(file_bytes);
}