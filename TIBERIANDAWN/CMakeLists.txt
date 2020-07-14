
add_library(TD_MIXFILE
    CCFILE.H
    LINK.CPP
    LINK.H
    MIXFILE.CPP
    MIXFILE.H
)

target_include_directories(TD_MIXFILE PRIVATE
    platformlib
)
