package provision

type Startup interface {
	WriteStartup() string
}
