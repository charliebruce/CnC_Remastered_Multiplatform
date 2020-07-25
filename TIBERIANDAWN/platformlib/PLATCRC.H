#ifndef PLATCRC_H
#define PLATCRC_H

#include <stdint.h>

#ifdef _MSC_VER
long __cdecl Calculate_CRC(void *buffer, long length);
#elif __cplusplus
extern "C" int32_t Calculate_CRC(void *buffer, long length);
#else
int32_t Calculate_CRC(void *buffer, long length);
#endif

#endif //PLATCRC_H
