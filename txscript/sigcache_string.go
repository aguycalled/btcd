// Copyright (c) 2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package txscript

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"sync"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/wire"
)

// newSigKey creates a new sigcache lookup key using the passed paramters. This
// lookup key is is the result of: SHA-256(nonce || sigHash || signature || pubkey).
func newSigKeyString(nonce [wire.HashSize]byte, sigHash wire.ShaHash,
	sig *btcec.Signature, pubKey *btcec.PublicKey) string {

	hasher := sha256.New()
	hasher.Write(nonce[:])
	hasher.Write(sigHash[:])
	hasher.Write(sig.Serialize())
	hasher.Write(pubKey.SerializeCompressed())

	return string(hasher.Sum(nil))
}

// StringSigCache implements an ECDSA signature verification cache with a randomized
// entry eviction policy. Only valid signatures will be added to the cache. The
// benefits of StringSigCache are two fold. Firstly, usage of StringSigCache mitigates a DoS
// attack wherein an attack causes a victim's client to hang due to worst-case
// behavior triggered while processing attacker crafted invalid transactions. A
// detailed description of the mitigated DoS attack can be found here:
// https://bitslog.wordpress.com/2013/01/23/fixed-bitcoin-vulnerability-explanation-why-the-signature-cache-is-a-dos-protection/.
// Secondly, usage of the StringSigCache introduces a signature verification
// optimization which speeds up the validation of transactions within a block,
// if they've already been seen and verified within the mempool.
type StringSigCache struct {
	sync.RWMutex
	validSigs  map[string]struct{}
	maxEntries uint
	cacheNonce [wire.HashSize]byte
}

// NewStringSigCache creates and initializes a new instance of StringSigCache. Its sole
// parameter 'maxEntries' represents the maximum number of entries allowed to
// exist in the StringSigCache at any particular moment. Random entries are evicted
// to make room for new entries that would cause the number of entries in the
// cache to exceed the max.
func NewStringSigCache(maxEntries uint) (*StringSigCache, error) {
	cache := &StringSigCache{
		validSigs:  make(map[string]struct{}),
		maxEntries: maxEntries,
	}

	// Read a 32 byte nonce to use as a salt the SHA-256 invocations for
	// each entry.
	if _, err := rand.Read(cache.cacheNonce[:]); err != nil {
		return nil, err
	}

	return cache, nil
}

// Exists returns true if an existing entry of 'sig' over 'sigHash' for public
// key 'pubKey' is found within the StringSigCache. Otherwise, false is returned.
//
// NOTE: This function is safe for concurrent access. Readers won't be blocked
// unless there exists a writer, adding an entry to the StringSigCache.
func (s *StringSigCache) Exists(sigHash wire.ShaHash, sig *btcec.Signature, pubKey *btcec.PublicKey) bool {
	s.RLock()
	_, ok := s.validSigs[newSigKeyString(s.cacheNonce, sigHash, sig, pubKey)]
	s.RUnlock()
	return ok
}

// Add adds an entry for a signature over 'sigHash' under public key 'pubKey'
// to the signature cache. In the event that the StringSigCache is 'full', an
// existing entry is randomly chosen to be evicted in order to make space for
// the new entry.
//
// NOTE: This function is safe for concurrent access. Writers will block
// simultaneous readers until function execution has concluded.
func (s *StringSigCache) Add(sigHash wire.ShaHash, sig *btcec.Signature, pubKey *btcec.PublicKey) {
	s.Lock()
	defer s.Unlock()

	if s.maxEntries <= 0 {
		return
	}

	// If adding this new entry will put us over the max number of allowed
	// entries, then evict an entry.
	if uint(len(s.validSigs)+1) > s.maxEntries {
		// Generate a cryptographically random hash.
		randHashBytes := make([]byte, wire.HashSize)
		_, err := rand.Read(randHashBytes)
		if err != nil {
			// Failure to read a random hash results in the proposed
			// entry not being added to the cache since we are
			// unable to evict any existing entries.
			return
		}

		// Try to find the first entry that is greater than the random
		// hash. Use the first entry (which is already pseudo random due
		// to Go's range statement over maps) as a fall back if none of
		// the hashes in the rejected transactions pool are larger than
		// the random hash.
		var foundEntry string
		for sigEntry := range s.validSigs {
			if foundEntry == "" {
				foundEntry = sigEntry
			}
			if bytes.Compare([]byte(sigEntry), randHashBytes) > 0 {
				foundEntry = sigEntry
				break
			}
		}
		delete(s.validSigs, foundEntry)
	}

	s.validSigs[newSigKeyString(s.cacheNonce, sigHash, sig, pubKey)] = struct{}{}
}