/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package peer

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/hyperledger/aries-framework-go/pkg/doc/did"
	"github.com/hyperledger/aries-framework-go/pkg/internal/mock/storage"
)

func TestPeerDIDStore(t *testing.T) {
	prov := storage.NewMockStoreProvider()
	dbstore, err := prov.OpenStore(StoreNamespace)
	require.NoError(t, err)

	context := []string{"https://w3id.org/did/v1"}

	did1 := "did:peer:1234"
	did2 := "did:peer:4567"

	store := NewDIDStore(dbstore)

	// put
	err = store.Put(&did.Doc{Context: context, ID: did1}, nil)
	require.NoError(t, err)

	// put
	err = store.Put(&did.Doc{Context: context, ID: did2}, nil)
	require.NoError(t, err)

	// get
	doc, err := store.Get(did1)
	require.NoError(t, err)
	require.Equal(t, did1, doc.ID)

	// get - empty id
	_, err = store.Get("")
	require.Error(t, err)

	// get - invalid id
	_, err = store.Get("did:peer:789")
	require.Error(t, err)

	// put - empty id
	err = store.Put(&did.Doc{ID: ""}, nil)
	require.Error(t, err)

	// put - missing doc
	err = store.Put(nil, nil)
	require.Error(t, err)

	// get - not json document
	err = dbstore.Put("not-json", []byte("not json"))
	require.NoError(t, err)
	v, err := store.Get("not-json")
	require.NotNil(t, err)
	require.Nil(t, v)
	require.Contains(t, err.Error(), "delta data fetch from store failed")
}
