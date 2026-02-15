#include "leap.h"
bool leap_year(int year) {
    // int current_year_by_four = year;
    // int current_year_by_four_hundred = year;
    // bool leap_by_four = (current_year_by_four % 4 == 0);
    // bool leap_by_four_hundred = (current_year_by_four_hundred % 400 == 0);
    // return (leap_by_four && leap_by_four_houndred);
    return (year % 4 == 0) && (year % 100 != 0 || year % 400 == 0);
}