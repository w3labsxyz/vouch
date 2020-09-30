// Code generated by fastssz. DO NOT EDIT.
package wallet

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the SigningContainer object
func (s *SigningContainer) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the SigningContainer object to a target array
func (s *SigningContainer) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Root'
	if len(s.Root) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, s.Root...)

	// Field (1) 'Domain'
	if len(s.Domain) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, s.Domain...)

	return
}

// UnmarshalSSZ ssz unmarshals the SigningContainer object
func (s *SigningContainer) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 64 {
		return ssz.ErrSize
	}

	// Field (0) 'Root'
	if cap(s.Root) == 0 {
		s.Root = make([]byte, 0, len(buf[0:32]))
	}
	s.Root = append(s.Root, buf[0:32]...)

	// Field (1) 'Domain'
	if cap(s.Domain) == 0 {
		s.Domain = make([]byte, 0, len(buf[32:64]))
	}
	s.Domain = append(s.Domain, buf[32:64]...)

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the SigningContainer object
func (s *SigningContainer) SizeSSZ() (size int) {
	size = 64
	return
}

// HashTreeRoot ssz hashes the SigningContainer object
func (s *SigningContainer) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the SigningContainer object with a hasher
func (s *SigningContainer) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Root'
	if len(s.Root) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(s.Root)

	// Field (1) 'Domain'
	if len(s.Domain) != 32 {
		err = ssz.ErrBytesLength
		return
	}
	hh.PutBytes(s.Domain)

	hh.Merkleize(indx)
	return
}