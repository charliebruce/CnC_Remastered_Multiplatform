#include <iostream>
#include "CRC.H"

int main() {
    CRCEngine engine;
    engine("big", 3);
    engine("check.shp", 9);
    auto value = static_cast<std::int32_t>(engine);
    constexpr std::int32_t expected = -724769914;
    if (value == expected) {
        return 0;
    }

    std::cerr << "Calculated CRC wrong:" << std::endl
              << "Expected:" << std::endl
              << " Dec: " << expected << std::endl
              << " Hex: " << std::uppercase << std::hex << expected << std::endl
              << "Actual:" << std::endl
              << " Dec:" << std::dec << value << std::endl
              << " Hex:" << std::uppercase << std::hex << value << std::endl;
    return 1;
}
