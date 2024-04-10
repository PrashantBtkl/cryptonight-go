========================================================================
>               CryptonightGPU hashing algorithm                       <
========================================================================

The CryptonightGPU hash verification algorithm is the block verification
algorithm for the Ryo blockchain. As you may know, when a new block is 
verified in blockchain technology, that event is known as "mining".  All
mining can be summerized to the follow:

`
while (true) {
	nonce := rand(0, 2^64)
	if hashsum(block, nonce) < block.Target {
		return nonce
	}
} 
`
It's the hashsum() < target part that we're interested in here. I'd like
to have that in a Go package.

-----------
- Hashes
-----------

A "hashsum" typically takes a range of derived large numbers (big.Ints), 
and returns the sum of these values in hexdecimal encoding.  Expect to 
use math/big for this.

As such, the package we will be developing can be summarized to two func 
prototypes and one struct.

type Block struct {
    // Properties may vary
    Header	string
    SeedHash	string
    Height	uint64
    Nonce	uint64
    Difficulty	string
}

(h* Hasher) func Sum(Block) (*big.Int, error)  // Calculates the sum of the hash
(h* Hasher) func Verify(Block) (bool, error)


Ultimately, the comparison of sum() back to the block.Target is what 
we're after:  

`
H(f) < target = pass
`

========================================================================
>                  CryptonightGPU Algorithm Specifics                  <
========================================================================

From what I've read, CryptonightGPU is an algorithm that's forked pretty
far from it's original Cryptonight parent.  In reading the source code,
it seems to use a term known as "keccak" often.

This algorithm is algo used in the Conceal, and Equilibria block chains;
though it seems there maybe variations between the flavors.


https://github.com/ConcealNetwork/conceal-core/blob/master/src/crypto/hash.c
https://github.com/EquilibriaCC/Equilibria/blob/master/src/crypto/hash.c
https://github.com/fireice-uk/xmr-stak/blob/master/xmrstak/backend/cryptonight.hpp
