#include "CRC.H"

/***********************************************************************************************
 * Calculate_CRC -- Calculates a one-way hash from a data block.                               *
 *                                                                                             *
 *    This routine is used to create a hash value from a data block. The algorithm is similar  *
 *    to a CRC, but is faster.                                                                 *
 *                                                                                             *
 * INPUT:   buffer   -- Pointer to a buffer of data to be 'hashed'.                            *
 *                                                                                             *
 *          len      -- The length of the buffer to compute the hash upon.                     *
 *                                                                                             *
 * OUTPUT:  Returns with a 32bit hash value calculated from the specified buffer.              *
 *                                                                                             *
 * WARNINGS:   none                                                                            *
 *                                                                                             *
 * HISTORY:                                                                                    *
 *   03/02/1996 JLB : Created.                                                                 *
 *=============================================================================================*/
extern "C"  {
long Calculate_CRC(void * buffer, long len)
{
	return(CRCEngine()(buffer, len));
}
}
