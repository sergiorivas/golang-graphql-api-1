package resolvers

type RootResolver struct {
}

func (_ *RootResolver) Hello() string { return "Hello, world!12" }
