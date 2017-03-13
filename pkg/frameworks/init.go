package frameworks

import (
	// language frameworks
	_ "github.com/pinpt/dialect/pkg/frameworks/language/nodejs"

	// container frameworks
	_ "github.com/pinpt/dialect/pkg/frameworks/container/docker"

	// build environments
	_ "github.com/pinpt/dialect/pkg/frameworks/build/circleci"
	_ "github.com/pinpt/dialect/pkg/frameworks/build/jenkins"
	_ "github.com/pinpt/dialect/pkg/frameworks/build/travis"
)
