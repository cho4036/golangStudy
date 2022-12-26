package setting

var singletonInstance *setting

type setting struct {
	color string
}

func (s *setting) Color() string {
	return s.color
}

func GetInstance() *setting {
	if singletonInstance == nil {
		singletonInstance = &setting{}
	}
	return singletonInstance
}

func (s *setting) SetColor(color string) {
	s.color = color
}
