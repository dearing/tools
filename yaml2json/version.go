package main

import (
	"fmt"
	"os"
)

//Version represents a versioning according to Semantic Versioning 2.0.0 => http://semver.org/spec/v2.0.0.html
type Version struct {
	Major    int    // Major version X (X.y.z | X > 0) MUST be incremented if any backwards incompatible changes are introduced to the public API. It MAY include minor and patch level changes. Patch and minor version MUST be reset to 0 when major version is incremented.
	Minor    int    // Minor version Y (x.Y.z | x > 0) MUST be incremented if new, backwards compatible functionality is introduced to the public API. It MUST be incremented if any public API functionality is marked as deprecated. It MAY be incremented if substantial new functionality or improvements are introduced within the private code. It MAY include patch level changes. Patch version MUST be reset to 0 when minor version is incremented.
	Patch    int    // Patch version Z (x.y.Z | x > 0) MUST be incremented if only backwards compatible bug fixes are introduced. A bug fix is defined as an internal change that fixes incorrect behavior.
	Build    string // Build metadata MAY be denoted by appending a plus sign and a series of dot separated identifiers immediately following the patch or pre-release version. Identifiers MUST comprise only ASCII alphanumerics and hyphen [0-9A-Za-z-]. Identifiers MUST NOT be empty. Build metadata SHOULD be ignored when determining version precedence. Thus two versions that differ only in the build metadata, have the same precedence. Examples: 1.0.0-alpha+001, 1.0.0+20130313144700, 1.0.0-beta+exp.sha.5114f85.
	Codeword string // cute nick name cuz I love them
}

var v = Version{
	Major:    1,
	Minor:    0,
	Patch:    0,
	Build:    "0x55AB0C53",
	Codeword: "drafty esmenet",
}

// Print the details of the version struct, neatly.
func (v *Version) String() string {
	return fmt.Sprintf("%s : %d.%d.%d+%s '%s'", os.Args[0], v.Major, v.Minor, v.Patch, v.Build, v.Codeword)
}
