/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package peer

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/hyperledger/aries-framework-go/pkg/doc/did"
	"github.com/hyperledger/aries-framework-go/pkg/framework/didresolver"
	"github.com/hyperledger/aries-framework-go/pkg/internal/mock/storage"
)

const peerDID = "did:peer:1234"

//nolint:lll
const peerDIDDoc = `{
  "@context": ["https://w3id.org/did/v1","https://w3id.org/did/v2"],
  "id": "did:peer:1234",
  "publicKey": [
    {
      "id": "did:peer:1234#keys-1",
      "type": "Secp256k1VerificationKey2018",
      "controller": "did:peer:1234",
      "publicKeyBase58": "H3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV"
    },
    {
      "id": "did:example:123456789abcdefghw#key2",
      "type": "RsaVerificationKey2018",
      "controller": "did:example:123456789abcdefghw",
      "publicKeyPem": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAryQICCl6NZ5gDKrnSztO\n3Hy8PEUcuyvg/ikC+VcIo2SFFSf18a3IMYldIugqqqZCs4/4uVW3sbdLs/6PfgdX\n7O9D22ZiFWHPYA2k2N744MNiCD1UE+tJyllUhSblK48bn+v1oZHCM0nYQ2NqUkvS\nj+hwUU3RiWl7x3D2s9wSdNt7XUtW05a/FXehsPSiJfKvHJJnGOX0BgTvkLnkAOTd\nOrUZ/wK69Dzu4IvrN4vs9Nes8vbwPa/ddZEzGR0cQMt0JBkhk9kU/qwqUseP1QRJ\n5I1jR4g8aYPL/ke9K35PxZWuDp3U0UPAZ3PjFAh+5T+fc7gzCs9dPzSHloruU+gl\nFQIDAQAB\n-----END PUBLIC KEY-----"
    }
  ],
  "authentication": [
    "did:peer:1234#keys-1",
    {
      "id": "did:example:123456789abcdefghs#key3",
      "type": "RsaVerificationKey2018",
      "controller": "did:example:123456789abcdefghs",
      "publicKeyHex": "02b97c30de767f084ce3080168ee293053ba33b235d7116a3263d29f1450936b71"
    }
  ],
  "service": [
    {
      "id": "did:peer:1234#inbox",
      "type": "SocialWebInboxService",
      "serviceEndpoint": "https://social.example.com/83hfh37dj",
      "spamCost": {
        "amount": "0.50",
        "currency": "USD"
      }
    }
  ],
  "created": "2002-10-10T17:00:00Z",
  "proof": [{
    "type": "LinkedDataSignature2015",
    "created": "2016-02-08T16:02:20Z",
    "creator": "did:example:8uQhQMGzWxR8vw5P3UWH1ja#keys-1",
    "proofValue": "6mdES87erjP5r1qCSRW__otj-A_Rj0YgRO7XU_0Amhwdfa7AAmtGUSFGflR_fZqPYrY9ceLRVQCJ49s0q7-LBA"
  }]
}`

func TestPeerDIDResolver(t *testing.T) {
	prov := storage.NewMockStoreProvider()
	dbstore, err := prov.OpenStore(StoreNamespace)
	require.NoError(t, err)

	context := []string{"https://w3id.org/did/v1"}

	// save did document
	store := NewDIDStore(dbstore)
	err = store.Put(&did.Doc{Context: context, ID: peerDID}, nil)
	require.NoError(t, err)

	resl := NewDIDResolver(store)
	doc, err := resl.Read(peerDID)
	require.NoError(t, err)

	document := &did.Doc{Context: context}
	err = json.Unmarshal(doc, document)
	require.NoError(t, err)
	require.Equal(t, peerDID, document.ID)

	// empty DID
	_, err = resl.Read("")
	require.Error(t, err)

	// missing DID
	// TODO this test should assert that didresolver.ErrNotFound is returned.
	//      that is currently impossible since the underlying store returns a
	//      generic error when the object is not found.
	_, err = resl.Read("did:peer:789")
	require.Error(t, err)
}

func TestWithDIDResolveAPI(t *testing.T) {
	prov := storage.NewMockStoreProvider()
	dbstore, err := prov.OpenStore(StoreNamespace)
	require.NoError(t, err)

	// save did document
	store := NewDIDStore(dbstore)
	peerDoc, err := did.ParseDocument([]byte(peerDIDDoc))
	require.NoError(t, err)
	require.NotNil(t, peerDoc)
	err = store.Put(peerDoc, nil)
	require.NoError(t, err)

	r := didresolver.New(didresolver.WithDidMethod(NewDIDResolver(store)))
	_, err = r.Resolve(peerDID)
	require.NoError(t, err)

	_, err = r.Resolve("did:peer:789")
	require.Error(t, err)
}
