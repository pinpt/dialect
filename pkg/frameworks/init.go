package frameworks

import (
	// language frameworks
	_ "github.com/pinpt/dialect/frameworks/language/nodejs"

	// container frameworks
	_ "github.com/pinpt/dialect/frameworks/container/docker"

	// build environments
	_ "github.com/pinpt/dialect/frameworks/build/circleci"
	_ "github.com/pinpt/dialect/frameworks/build/jenkins"
	_ "github.com/pinpt/dialect/frameworks/build/travis"
)
