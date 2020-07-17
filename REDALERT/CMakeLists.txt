add_library(RA_CRC STATIC
    CRC.CPP
    CRC.H
)

target_include_directories(RA_CRC PUBLIC
    ..
)
