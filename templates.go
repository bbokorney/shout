package main

// Templates handles rendering message templates
type Templates interface {
	Render(templateName string, data map[string]string) (string, error)
}

// NewTemplates creates a new Templates
func NewTemplates() Templates {
	return templates{}
}

type templates struct {
}

func (t templates) Render(templateName string, data map[string]string) (string, error) {
	return "", nil
}
