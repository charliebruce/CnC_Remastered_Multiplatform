add_executable(RA_CRC_TEST_EXE
    CRC_TEST.CPP
)

target_link_libraries(RA_CRC_TEST_EXE
    RA_CRC
)

add_library(RA_CRC STATIC
    CRC.CPP
    CRC.H
)

add_test(
    NAME RA_CRC_TEST
    COMMAND $<TARGET_FILE:RA_CRC_TEST_EXE>
)

target_include_directories(RA_CRC PUBLIC
    ..
)
