/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package did

import (
	"github.com/hyperledger/aries-framework-go/pkg/doc/did"
	"github.com/hyperledger/aries-framework-go/pkg/wallet"
)

const method = "local"

// Creator is a DID creator interface
type Creator interface {
	// Creates new local DID document.
	//
	//
	// opts: options to create local DID
	//
	// Returns:
	//
	// did: DID document
	//
	// error: error
	CreateDID(opts ...wallet.DocOpts) (*did.Doc, error)

	// Gets already created DID document by ID.
	//
	// Args:
	//
	// id: DID identifier
	//
	// Returns:
	//
	// did: DID document
	//
	// error: when document is not found or for any other error conditions
	GetDID(id string) (*did.Doc, error)
}

// provider contains dependencies for DID creator and is typically created by using aries.Context()
type provider interface {
	DIDWallet() wallet.DIDCreator
}

// NewLocalDIDCreator returns new Local DID creator
func NewLocalDIDCreator(ctx provider) *LocalDIDCreator {
	return &LocalDIDCreator{ctx.DIDWallet(), method}
}

// LocalDIDCreator creates local DIDs
type LocalDIDCreator struct {
	wallet.DIDCreator
	method string
}

// CreateDID creates new local DID
func (l *LocalDIDCreator) CreateDID(opts ...wallet.DocOpts) (*did.Doc, error) {
	return l.DIDCreator.CreateDID(l.method, opts...)
}

// GetDID gets already created DID document by ID.
func (l *LocalDIDCreator) GetDID(id string) (*did.Doc, error) {
	return l.DIDCreator.GetDID(id)
}
