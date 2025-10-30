package main

import (
	"errors"
	"fmt"
	"strings"
)

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}
type ContactManager struct {
	contacts []Contact
	nextID   int
}

func (cm *ContactManager) AddContact(name, email, phone string) Contact {
	cm.nextID++
	contact := Contact{
		ID:    cm.nextID,
		Name:  name,
		Email: email,
		Phone: phone,
	}
	cm.contacts = append(cm.contacts, contact)
	return contact
}
func (cm *ContactManager) GetContact(id int) (Contact, error) {
	for _, c := range cm.contacts {
		if c.ID == id {
			return c, nil
		}
	}
	return Contact{}, errors.New("contact introuvable")
}
func (cm *ContactManager) ListContacts() []Contact {
	return cm.contacts
}
func (cm *ContactManager) DeleteContact(id int) error {
	for i, c := range cm.contacts {
		if c.ID == id {
			cm.contacts = append(cm.contacts[:i], cm.contacts[i+1:]...)
			return nil
		}
	}
	return errors.New("contact non trouvé")
}
func (cm *ContactManager) SearchByName(query string) []Contact {
	var results []Contact
	for _, c := range cm.contacts {
		if strings.Contains(strings.ToLower(c.Name), strings.ToLower(query)) {
			results = append(results, c)
		}
	}
	return results
}

func (cm *ContactManager) SearchByEmail(email string) (Contact, error) {
	for _, c := range cm.contacts {
		if strings.EqualFold(c.Email, email) {
			return c, nil
		}
	}
	return Contact{}, errors.New("aucun contact trouvé avec cet email")
}
func main() {
	cm := &ContactManager{}

	for {
		fmt.Println("\n=== Gestionnaire de Contacts ===")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Lister les contacts")
		fmt.Println("3. Rechercher un contact (par nom ou email)")
		fmt.Println("4. Supprimer un contact")
		fmt.Println("5. Quitter")

		var choice int
		fmt.Print("Choix : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name, email, phone string
			fmt.Print("Nom : ")
			fmt.Scan(&name)
			fmt.Print("Email : ")
			fmt.Scan(&email)
			fmt.Print("Téléphone : ")
			fmt.Scan(&phone)
			c := cm.AddContact(name, email, phone)
			fmt.Println("Contact ajouté :", c)

		case 2:
			contacts := cm.ListContacts()
			if len(contacts) == 0 {
				fmt.Println("Aucun contact.")
			} else {
				fmt.Println("Liste des contacts :")
				for _, c := range contacts {
					fmt.Printf("%d - %s | %s | %s\n", c.ID, c.Name, c.Email, c.Phone)
				}
			}

		case 3:
			var option int
			fmt.Println("1. Recherche par nom")
			fmt.Println("2. Recherche par email")
			fmt.Print("Choix : ")
			fmt.Scan(&option)

			switch option {
			case 1:
				var query string
				fmt.Print("Nom à rechercher : ")
				fmt.Scan(&query)
				results := cm.SearchByName(query)
				if len(results) == 0 {
					fmt.Println("Aucun contact trouvé.")
				} else {
					for _, c := range results {
						fmt.Printf("%d - %s | %s | %s\n", c.ID, c.Name, c.Email, c.Phone)
					}
				}
			case 2:
				var email string
				fmt.Print("Email à rechercher : ")
				fmt.Scan(&email)
				contact, err := cm.SearchByEmail(email)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("%d - %s | %s | %s\n", contact.ID, contact.Name, contact.Email, contact.Phone)
				}
			}

		case 4:
			var id int
			fmt.Print("ID du contact à supprimer : ")
			fmt.Scan(&id)
			err := cm.DeleteContact(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Contact supprimé avec succès.")
			}

		case 5:
			fmt.Println("Au revoir 👋")
			return

		default:
			fmt.Println("Choix invalide, réessayez.")
		}
	}
}
