#include "day5.h"

typedef struct {
    int P1[2];
    int P2[2];
} Line;

int parse_coor(char* str, int* x, int* y) {
    errno = 0;
    *x = (int)strtol(strtok(str, ","), NULL, 10);
    if(errno != 0){
        return errno;
    }
    errno = 0;
    *y = (int)strtol(strtok(NULL, ","), NULL, 10);
    if(errno != 0){
        return errno;
    }

    return 0;

}

int max(int a , int b) {
    return a > b ? a : b;
}

bool point_on_line(Line* l, const int P[2]) {
    // implementation of https://stackoverflow.com/a/17693189/14790078

    if(!(l->P1[0] == l->P2[0] || l->P1[1] == l->P2[1])) {
        if(
            ((float)(P[1] - l->P1[1])/(float)(l->P2[1] - l->P1[1])) !=
            ((float)(P[0] - l->P1[0])/(float)(l->P2[0] - l->P1[0]))
        ) {
            return false;
        }
    }

    return (P[0] >= l->P1[0] &&
           P[0] <= l->P2[0] &&
           P[1] >= l->P1[1] &&
           P[1] <= l->P2[1])
           ||
            (P[0] <= l->P1[0] &&
            P[0] >= l->P2[0] &&
            P[1] <= l->P1[1] &&
            P[1] >= l->P2[1]);
//    return (l->P1[0] < P[0] < l->P2[0] && l->P1[1] < P[1] < l->P2[1]) ||
//            (l->P1[0] > P[0] > l->P2[0] && l->P1[1] > P[1] > l->P2[1]);
}

void day5() {
    printf("Day 5\n");
    size_t num_lines;

    char** file = load_file_whole("day5.txt", &num_lines);
    if(file == NULL){
        return;
    }
    printf("Number of lines in input: %zu\n", num_lines);

    Line* lines = (Line*)malloc(num_lines*sizeof(Line));
    int maxx = 0, maxy = 0;

    for(int i = 0; i < num_lines; i++) {
        char* from = strtok(file[i], " -> ");
        char* to = strtok(NULL, " -> ");

        int err = parse_coor(from, &lines[i].P1[0], &lines[i].P1[1]);
        if(err != 0) {
            printf("Couldnt parse from coordinates %s at line %i\n", from, i);
            return;
        }
        err = parse_coor(to, &lines[i].P2[0], &lines[i].P2[1]);
        if(err != 0) {
            printf("Couldnt parse to coordinates %s at line %i\n", to, i);
            return;
        }
        maxx = max(maxx, lines[i].P1[0]);
        maxx = max(maxx, lines[i].P2[0]);
        maxy = max(maxy, lines[i].P1[1]);
        maxy = max(maxy, lines[i].P2[1]);
    }
    printf("Board is %i by %i\n", maxx, maxy);

    printf("part 1\n");

    int crossing_points = 0;

    for(int y=0; y<=maxy; y++) {
        for(int x=0; x<=maxx; x++) {
            bool flag = false;
            for(int n=0; n<num_lines; n++) {
                int p[2] = {x, y};
                //for part 1 only consider straight lines
                if(!(lines[n].P1[0] == lines[n].P2[0] || lines[n].P1[1] == lines[n].P2[1])) {
                    continue;
                }
                if(point_on_line(&lines[n], p)) {
                    if(flag){// we already found an earlier line through this point
                        crossing_points++;
                        printf("%i %i\n", x, y);
                        break;
                    }
                    flag = true;
                }
            }
        }
    }

    printf("There are %i crossing points\n", crossing_points);

    free(lines);

}
