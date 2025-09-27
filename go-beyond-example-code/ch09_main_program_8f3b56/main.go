package main

import (
	"fmt"
	"sort"
)

type Contact struct {
	Name  string
	Phone string
	Email string
}

type ContactManager struct {
	contacts map[string]Contact
}

func NewContactManager() *ContactManager {
	return &ContactManager{
		contacts: make(map[string]Contact),
	}
}

func (cm *ContactManager) AddContact(key string, contact Contact) bool {
	if key == "" || contact.Name == "" {
		return false
	}

	cm.contacts[key] = contact
	return true
}

func (cm *ContactManager) GetContact(key string) (Contact, bool) {
	contact, exists := cm.contacts[key]
	return contact, exists
}

func (cm *ContactManager) ListContacts() []Contact {
	var contacts []Contact
	for _, contact := range cm.contacts {
		contacts = append(contacts, contact)
	}

	// Sort by name
	sort.Slice(contacts, func(i, j int) bool {
		return contacts[i].Name < contacts[j].Name
	})

	return contacts
}

func (cm *ContactManager) DeleteContact(key string) bool {
	if _, exists := cm.contacts[key]; exists {
		delete(cm.contacts, key)
		return true
	}
	return false
}

func main() {
	manager := NewContactManager()

	// Add contacts
	manager.AddContact("alice", Contact{
		Name:  "Alice Smith",
		Phone: "555-1234",
		Email: "alice@example.com",
	})

	manager.AddContact("bob", Contact{
		Name:  "Bob Johnson",
		Phone: "555-5678",
		Email: "bob@example.com",
	})

	// List all contacts (sorted)
	fmt.Println("All contacts (sorted by name):")
	for _, contact := range manager.ListContacts() {
		fmt.Printf("- %s (%s)\n", contact.Name, contact.Phone)
	}
}
