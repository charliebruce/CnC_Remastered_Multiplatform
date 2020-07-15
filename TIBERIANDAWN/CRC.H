#ifndef TD_CRC_H
#define TD_CRC_H

#ifdef _MSC_VER
long __cdecl Calculate_CRC(void *buffer, long length);
#else
extern "C" long Calculate_CRC(void *buffer, long length);
#endif

#endif
