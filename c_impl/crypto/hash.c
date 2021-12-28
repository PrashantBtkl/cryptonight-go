#include <stddef.h>
#include <stdint.h>
#include <string.h>

#include "hash-ops.h"
#include "keccak.h"

void cn_fast_hash(const void* data, size_t length, char* hash)
{
	keccak((const uint8_t*)data, length, (uint8_t*)hash, 32);
}
