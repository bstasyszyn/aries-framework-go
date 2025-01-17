/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package didcomm

// MockAuthCrypt mock auth crypt
type MockAuthCrypt struct {
	EncryptValue func(payload, senderPubKey []byte, recipients [][]byte) ([]byte, error)
	DecryptValue func(envelope []byte) ([]byte, error)
}

// Encrypt mock encrypt
func (m *MockAuthCrypt) Encrypt(payload, senderPubKey []byte,
	recipients [][]byte) ([]byte, error) {
	return m.EncryptValue(payload, senderPubKey, recipients)
}

// Decrypt mock decrypt
func (m *MockAuthCrypt) Decrypt(envelope []byte) ([]byte, error) {
	return m.DecryptValue(envelope)
}
