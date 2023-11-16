package dictionary

import (
    "fmt"
    "os"
)

type Entry struct {
}

func (e Entry) String() string {

	return ""
}

type Dictionary struct {
	entries map[string]Entry
}

func New() *Dictionary {

	return nil
}

func (d *Dictionary) Add(word string, definition string) {

}

func (d *Dictionary) Get(word string) (Entry, error) {

	return Entry{}, nil
}

func (d *Dictionary) Remove(word string) {

}

func (d *Dictionary) List() ([]string, map[string]Entry) {

	return []string{}, d.entries
}

func (d *Dictionary) SavuverEnFichier(filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    for word, entry := range d.entries {
        _, err := fmt.Fprintf(file, "%s: %s\n", word, entry.String())
        if err != nil {
            return err
        }
    }

    return nil
}

