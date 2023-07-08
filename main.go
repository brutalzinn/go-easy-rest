package goeasyrest

var Version string = "1.0.4"

type ResourceData struct {
	Data []Resource `json:"data"`
}
type Resource struct {
	Object    string `json:"object"`
	Attribute any    `json:"attribute"`
	Link      []Link `json:"link"`
}

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

func NewRes(key string, attribute any, links []Link) Resource {
	return Resource{
		Object:    key,
		Attribute: attribute,
		Link:      links,
	}
}
func NewResData() *ResourceData {
	return &ResourceData{Data: make([]Resource, 0)}
}
func (r *ResourceData) Add(resource Resource) {
	r.Data = append(r.Data, resource)
}

func (r *ResourceData) Remove(resource Resource) {
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
