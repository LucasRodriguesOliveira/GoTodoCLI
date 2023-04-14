package todomodel

import "strconv"

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

func (i *Item) PrettyPriority() string {
	switch i.Priority {
	case 1:
		return "[HIGH]"
	case 2:
		return "[MEDIUM]"
	case 3:
		return "[LOW]"
	default:
		return ""
	}
}

func (i *Item) SetPosition(p int) {
	i.position = p
}

func (i *Item) Label() string {
	return strconv.Itoa(i.position) + "."
}

func (i *Item) Pos() int {
	return i.position
}

func (i *Item) PrettyDone() string {
	if i.Done {
		return "[DONE]"
	}

	return ""
}
