package _package

// Package defines a struct holding information regarding package
type Package struct {
	Name        string
	Version     string
	Maintainer  string
	BuildConfig *BuildConfig
}

// Product defines build products output. Name stands for human readable name
// of the product. Source and Destination field stand for source and destination
// paths, respectively.
type Product struct {
	Name        string
	Source      string
	Destination string
}

// BuildConfig defines basic build instructions. IsPlaceholder defines additional
// functionalities - i.e. build system may be used in the future to create OCI
// images or ISO images.
type BuildConfig struct {
	IsPlaceholder bool
	Prerequisites *Prerequisites
	Steps         *BuildSteps
	Products      []*Product
}

// Prerequisites holds list of requirements for the target package. Tools need to
// be installed on the build machine. Packages are other packages that the product
// may depend on. DevelopmentLibraries are libraries needed for build. Libraries
// are requirements also carried on to the product package and will be required
// on the target machine.
type Prerequisites struct {
	Tools                []*Requirement
	Packages             []*Requirement
	DevelopmentLibraries []*Requirement
	Libraries            []*Requirement
}

// Requirement is a package either for target or build host that are needed.
// Name stands for package name, MinVersion and MaxVersion stand for package
// versions (inclusively - <= or >= equivalent in other package systems)
type Requirement struct {
	Name       string
	MinVersion string
	MaxVersion string
}

// BuildSteps is a container holding information how to build a package for
// specified architecture. SameArch stands for requirement that the same main
// architecture should be used for building the package (i.e. amd64 for amd64
// package. This flag may be set to "false" i.e. on Python Packages or Golang
// binaries where the host system supports building for different architectures.
type BuildSteps struct {
	SameArch      bool
	Architectures map[string]*Architecture
}

// Architecture is a struct containing subtypes of an architecture and their
// respective build instructions. I.e. AMD64 supports many of subtypes, as well
// as a generic subtype. I.e. bdver2 could be a subtype for AMD Bulldozer V2
// architecture.
type Architecture struct {
	Name     string
	Subtypes []*SubtypeRecipe
}

// SubtypeRecipe contains information on how to build a package for specified
// architecture subtype. "generic" should be used as a generic container.
// Commands field holds list of commands that will be executed and watched.
// EnvironmentVariables field holds map of environment variables that will be set
// before the build system will be executed.
type SubtypeRecipe struct {
	EnvironmentVariables map[string]string
	Commands             []string
}
