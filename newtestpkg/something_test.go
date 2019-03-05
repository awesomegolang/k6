package newtestpkg

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func init() {
	fmt.Printf("newtestpkg init environment: %#q", os.Environ())
}

func setEnv(t *testing.T, newEnv []string) (restoreEnv func()) {
	actuallSetEnv := func(env []string) {
		os.Clearenv()
		t.Logf("Cleared env, setting it to %#q", env)
		for _, e := range env {
			val := ""
			pair := strings.SplitN(e, "=", 2)
			if len(pair) > 1 {
				val = pair[1]
			}
			t.Logf("\tSetting '%s' to be '%s'", pair[0], val)
			require.NoError(t, os.Setenv(pair[0], val))
		}
	}
	t.Logf("setEnv(%q) called", newEnv)
	oldEnv := os.Environ()
	actuallSetEnv(newEnv)

	return func() {
		t.Log("restoreEnv reached")
		actuallSetEnv(oldEnv)
	}
}

func TestConfigConsolidation(t *testing.T) {
	restoreEnv := setEnv(t, []string{"test=mest", "asdfg=ashasdhasd"})
	defer restoreEnv()
	t.Error("did not expect this to happen")
}
