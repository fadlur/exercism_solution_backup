#include "armstrong_numbers.h"
#include <math.h>
bool is_armstrong_number(int candidate) {
    int num, remainder, result = 0, total_digit = 0;
    num = candidate;
    while (num != 0) {
        num/=10;
        total_digit++;
    }

    num = candidate;
    while (num != 0) {
        remainder = num % 10;
        result += pow(remainder, total_digit);
        num /= 10;
    }
    return (result == candidate);
}