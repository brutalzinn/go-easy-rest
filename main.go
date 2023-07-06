package main

type ResourceData struct {
	Object string
	Data   []Resource
}
type Resource struct {
	Object    string
	Attribute any
	Link      []Link
}

type Link struct {
	Rel  string
	Href string
}

func NewResource(key string, attribute any, links []Link) Resource {
	return Resource{
		Object:    key,
		Attribute: attribute,
		Link:      links,
	}
}
func NewResourceData() *ResourceData {
	return &ResourceData{Data: make([]Resource, 0)}
}
func (r *ResourceData) AddResource(resource Resource) {
	r.Data = append(r.Data, resource)
}

func (r *ResourceData) RemoveResouce(resource Resource) {
	toKeep := make([]Resource, 0)
	for _, item := range r.Data {
		if item.Object != resource.Object {
			toKeep = append(toKeep, item)
		}
	}
	r.Data = toKeep
}

func (r *ResourceData) Size() int {
	return len(r.Data)
}
