package apisecrets

import (
	"encoding/json"
	"errors"
	"example/apisecrets/cipher"
	"io"
	"os"
	"sync"
)

type Vault struct {
	encodingKey string
	filepath    string
	mutex       sync.Mutex
	keyValues   map[string]string
}

func File(encodingKey string, filepath string) *Vault {
	return &Vault{
		encodingKey: encodingKey,
		filepath:    filepath,
	}
}

func (v *Vault) load() error {
	file, err := os.Open(v.filepath)
	// if we do not have a file initialize the map
	if err != nil {
		v.keyValues = make(map[string]string)
		return nil
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	r, err := cipher.DecryptReader(v.encodingKey, file)
	if err != nil {
		return err
	}

	return v.readKeyValues(r)
}

func (v *Vault) readKeyValues(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(&v.keyValues)
}

func (v *Vault) save() error {
	file, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	w, err := cipher.EncryptWriter(v.encodingKey, file)
	if err != nil {
		return err
	}

	return v.writeKeyValues(w)
}

func (v *Vault) writeKeyValues(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(v.keyValues)
}

func (v *Vault) Get(key string) (string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	err := v.load()

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

	err := v.load()
	if err != nil {
		return err
	}

	v.keyValues[key] = value
	err = v.save()
	if err != nil {
		return err
	}

	return nil
}
