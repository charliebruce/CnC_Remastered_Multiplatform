
add_library(gbuffer_win32 STATIC
    DRAWBUFF.H
    GBUFFER.CPP
    GBUFFER.H
    ICONCACH.H
    TILE.H
    WWSTD.H
)

add_executable(gbuffer_win32_test_exe
    GBUFFER_TEST.CPP
)

add_dependencies(gbuffer_win32_test_exe
    gbuffer_win32
)

add_test(NAME gbuffer_win32_test COMMAND gbuffer_win32_test_exe)
