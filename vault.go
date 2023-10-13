package apisecrets

import (
	"errors"
	"example/apisecrets/encrypt"
)

type Vault struct {
	encodingKey string
	keyValues map[string]string
}

func Memory(encodingKey string) Vault {
	return Vault{
		encodingKey: encodingKey,
		keyValues: make(map[string]string),
	}
}

func (v *Vault) Get(key string) (string, error) {
	cipherHex, ok := v.keyValues[key]

	if !ok {
		return "", errors.New("secret: no value for choosen key")
	}

	ret, err := encrypt.Decrypt(v.encodingKey, cipherHex)

	if err != nil {
		return "", err
	}

	return ret, nil
}

func (v *Vault) Set(key, value string) error {
	encryptedValue, err := encrypt.Encrypt(v.encodingKey, key)
	if err != nil {
		return err
	}

	v.keyValues[key] = encryptedValue
	return nil
}