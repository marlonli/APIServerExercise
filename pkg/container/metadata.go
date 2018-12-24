package container

import "strings"

// Metadata defines the metadata of applications
type Metadata struct {
	Title       string    `yaml:"title"`
	Version     string    `yaml:"version"`
	Maintainers []*Person `yaml:"maintainers"`
	Company     string    `yaml:"company"`
	Website     string    `yaml:"website"`
	Source      string    `yaml:"source"`
	License     string    `yaml:"license"`
	Description string    `yaml:"description"`
}

type Person struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

// Check whether the name and email format is valid
func (p *Person) IsValid() bool {
	numAt := strings.Count(p.Email, "@")
	if numAt != 1 {
		return false
	}
	atIdx := strings.Index(p.Email, "@")
	dotIdx := strings.LastIndex(p.Email, ".")
	return atIdx < dotIdx && len(p.Name) > 0
}

// Whether the metadata matches a query string
func (m *Metadata) Match(query string) bool {
	// Whether there's a field contains the query string
	match := strings.Contains(m.Title, query) || strings.Contains(m.Version, query) || strings.Contains(m.Company, query) ||
		strings.Contains(m.Website, query) || strings.Contains(m.Source, query) || strings.Contains(m.License, query) ||
		strings.Contains(m.Description, query)
	if match {
		return true
	}

	// Whether the maintainers' information contains the query string
	for _, person := range m.Maintainers {
		if strings.Contains(person.Name, query) || strings.Contains(person.Email, query) {
			match = true
			break
		}
	}
	return match
}

// Check whether the metadata is valid
func (m *Metadata) IsValid() bool {
	pValid := true // pValid indicates whether the maintainers' information is valid
	for _, p := range m.Maintainers {
		if !p.IsValid() {
			pValid = false
			break
		}
	}
	return pValid && len(m.Title) > 0 && len(m.Version) > 0 && len(m.Company) > 0 &&
		len(m.Website) > 0 && len(m.Source) > 0 && len(m.License) > 0 && len(m.Description) > 0
}
