package csv

import (
	"regexp"
	"sort"
	"time"
)

type JSONRecordDictionary map[string]*JSONRecord

type JSONContacts struct {
	Contacts []*JSONRecord `json:"contacts"`
}

type JSONRecord struct {
	FirstName       string    `json:"firstname"`
	LastName        string    `json:"lastname"`
	Email           string    `json:"email"`
	InscriptionDate time.Time `json:"inscription_date"`
}

// add adds the contact into the dictionary
func (dict JSONRecordDictionary) add(newContact *JSONRecord) {
	if newContact == nil {
		return
	}
	if !newContact.isValid() {
		return
	}

	// check if we already have this contact in the dictionary
	if contact, exists := dict[newContact.Email]; exists {
		// if the existing contact is older than the new contact
		if contact.InscriptionDate.Before(newContact.InscriptionDate) {
			dict[contact.Email] = newContact
		}
		return
	}

	dict[newContact.Email] = newContact
}

// toContact translates the record dictionary into a JSON Contact object.
func (dict JSONRecordDictionary) toContact() JSONContacts {
	var records []*JSONRecord

	for _, record := range dict {
		records = append(records, record)
	}

	contact := JSONContacts{
		Contacts: records,
	}

	return contact
}

// sort sorts the contact from the most recent to the oldest.
func (contact JSONContacts) sort() {
	sort.Slice(contact.Contacts, func(i, j int) bool {
		return contact.Contacts[i].InscriptionDate.After(contact.Contacts[j].InscriptionDate)
	})
}

// isValid checks whether the JSONRecord is valid.
func (rec *JSONRecord) isValid() bool {
	if rec == nil {
		return false
	}

	valid := true

	if len(rec.FirstName) >= 50 {
		stats.InvalidFirstname++
		valid = false
	}

	if len(rec.LastName) >= 50 {
		stats.InvalidLastname++
		valid = false
	}

	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression
	regex := regexp.MustCompile(pattern)

	if len(rec.Email) >= 255 || !regex.MatchString(rec.Email) {
		stats.InvalidEmail++
		return false
	}
	return valid
}
