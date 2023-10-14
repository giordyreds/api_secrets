package apisecrets

import (
	"encoding/json"
	"errors"
	"example/apisecrets/encrypt"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type Vault struct {
	encodingKey string
	filepath string
	mutex sync.Mutex
	keyValues map[string]string
}

func File(encodingKey string, filepath string) (*Vault) {
	return &Vault{
		encodingKey: encodingKey,
		filepath: filepath,
	}
}

func (v *Vault) loadKeyValues() error {
	file, err := os.Open(v.filepath)
	// if we do not have a file initialize the map
	if err != nil {
		v.keyValues = make(map[string]string)
		return nil
	}

	defer file.Close()

	var sb strings.Builder
	_, err = io.Copy(&sb, file)
	if err != nil {
		return err
	}

	decryptedJSON, err := encrypt.Decrypt(v.encodingKey, sb.String())
	if err != nil {
		return err
	}

	fmt.Println(decryptedJSON)
	
	r := strings.NewReader(decryptedJSON)

	dec := json.NewDecoder(r)
	err = dec.Decode(&v.keyValues)
	if err != nil {
		return err
	}

	return nil
}

func (v *Vault) saveKeyValues() error {
	var sb strings.Builder
	enc := json.NewEncoder(&sb)

	err := enc.Encode(v.keyValues)
	if err != nil {
		return err
	}

	enryptedJSON, err := encrypt.Encrypt(v.encodingKey, sb.String())
	if err != nil {
		return err
	}

	file, err := os.OpenFile(v.filepath, os.O_RDWR | os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = fmt.Fprint(file, enryptedJSON)
	if err != nil {
		return err
	}

	return nil
}

func (v *Vault) Get(key string) (string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	
	err := v.loadKeyValues()

	if err != nil {
		return "", err
	}

	value, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("secret: no value for choosen key")
	}

	return value, nil
}

func (v *Vault) Set(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	err := v.loadKeyValues()
	if err != nil {
		return err
	}

	v.keyValues[key] = value
	err = v.saveKeyValues()
	if err != nil {
		return err
	}

	return nil
}