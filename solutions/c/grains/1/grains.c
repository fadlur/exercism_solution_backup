#include "grains.h"
uint64_t square(uint8_t index) {
    return (index < 1 || index > 64) ? 0: 1ul << (index - 1);
}
uint64_t total() {
    uint64_t result = 0;
    for (uint8_t i = 0; i <= 64; i++) {
        result += square(i);
    }
    return result;
}