package bc

import (
	"chain/crypto/hash160"
	"chain/crypto/hash256"
	"chain/fedchain/script"
)

type (
	// AssetID is the Hash160 of the issuance script
	// for the asset and the genesis block of the chain
	// where it appears.
	AssetID [20]byte

	// IssuanceID is the hash160 of the outpoint
	// of the issuing transaction.
	IssuanceID [20]byte
)

// ComputeAssetID computes the asset ID of the asset defined by
// the given issuance script and genesis block hash.
func ComputeAssetID(issuance script.Script, genesis [32]byte) AssetID {
	buf := append([]byte{}, genesis[:]...)
	sh := hash256.Sum(issuance)
	buf = append(buf, sh[:]...)
	return hash160.Sum(buf)
}

// Issuance computes the ID of the asset issuance
// that occurred in p.
func Issuance(p Outpoint) (id IssuanceID) {
	h := hash160.New()
	p.WriteTo(h)
	h.Sum(id[:0])
	return id
}

// AssetDefinitionPointer is a Hash256 value of data associated
// with a specific AssetID.
// This is issuer's authenticated description of their asset.
type AssetDefinitionPointer struct {
	AssetID        AssetID
	DefinitionHash [32]byte
}
