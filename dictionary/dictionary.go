package dictionary

import (
	"errors"
	"fmt"
	"os"
	"sync"
)

type Entry struct {
	Word       string `json:"word"`
	Definition string `json:"definition"`
}

type Dictionary struct {
	entries map[string]Entry
	lock    sync.RWMutex
}

func NewDictionary() *Dictionary {
	return &Dictionary{
		entries: make(map[string]Entry),
	}
}
func (d *Dictionary) Add(word string, definition string) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	if len(word) < 3 || len(definition) < 5 {
		return errors.New("Les données ne respectent pas les règles de validation")
	}

	d.entries[word] = Entry{Definition: definition}
	return nil
}

func (d *Dictionary) Get(word string) (Entry, error) {
	d.lock.RLock()
	defer d.lock.RUnlock()

	entry, ok := d.entries[word]
	if !ok {
		return Entry{}, errors.New("Mot non trouvé")
	}

	return entry, nil
}

func (d *Dictionary) Remove(word string) {
	d.lock.Lock()
	defer d.lock.Unlock()

	delete(d.entries, word)
}

func (d *Dictionary) List() map[string]Entry {
	d.lock.RLock()
	defer d.lock.RUnlock()

	return d.entries
}

func (d *Dictionary) SavuverEnFichier(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	/*for word, entry := range d.entries {
	    _, err := fmt.Fprintf(file, "%s: %s\n", word, entry.string())
	    if err != nil {
	        return err
	    }
	}*/
	fmt.Println("Succèssssssssssssss.")
	return nil
}
