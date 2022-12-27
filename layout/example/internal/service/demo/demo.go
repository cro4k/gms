package demo

type DemoService struct {
}

func (s *DemoService) SayHello(name string) (string, error) {
	return "Hello " + name + "!", nil
}
