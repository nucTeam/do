package do

func (s *RootScope) ServiceForEach(cb func(string, *Scope, any) bool) { s.self.serviceForEach(cb) }
func (s *RootScope) ServiceGet(name string) (any, bool)               { return s.self.serviceGet(name) }
