package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string


var (
 errNotFound = errors.New("Not Found")
 errWordExists = errors.New("That word already exists")
 errCantUpdate =errors.New("Cant update non-existing word")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
  value, exists := d[word];
  if exists {
    return value, nil
  }
  return "", errNotFound 
}

// Add a word to the Dictionary
func (d Dictionary) Add(word, def string) error {
  _, err := d.Search(word)

  switch err {
    case errNotFound:
      d[word] = def
    case nil:
      return errWordExists
  }
  return nil
}

// Update the word
func (d Dictionary) Update(word, def string) error {
  _, err := d.Search(word)
  switch err {
    case nil: 
      d[word] = def
    case errNotFound:
      return errCantUpdate
  }
  return nil
}

// Delete the word
func (d Dictionary) Delete(word string){
  delete(d, word)
}