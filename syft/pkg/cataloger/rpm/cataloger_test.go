package rpm

import (
	"testing"

	"github.com/anchore/syft/syft/pkg/cataloger/internal/pkgtest"
)

func Test_DBCataloger_Globs(t *testing.T) {
	tests := []struct {
		name     string
		fixture  string
		expected []string
	}{
		{
			name:    "obtain DB files",
			fixture: "test-fixtures/glob-paths",
			expected: []string{
				"var/lib/rpm/Packages",
				"var/lib/rpm/Packages.db",
				"var/lib/rpm/rpmdb.sqlite",
				"var/lib/rpmmanifest/container-manifest-2",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pkgtest.NewCatalogTester().
				FromDirectory(t, test.fixture).
				ExpectsResolverContentQueries(test.expected).
				TestCataloger(t, NewRpmDBCataloger())
		})
	}
}

func Test_RPMFileCataloger_Globs(t *testing.T) {
	tests := []struct {
		name     string
		fixture  string
		expected []string
	}{
		{
			name:    "obtain rpm files",
			fixture: "test-fixtures/glob-paths",
			expected: []string{
				"dive-0.10.0.rpm",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pkgtest.NewCatalogTester().
				FromDirectory(t, test.fixture).
				ExpectsResolverContentQueries(test.expected).
				TestCataloger(t, NewFileCataloger())
		})
	}
}
