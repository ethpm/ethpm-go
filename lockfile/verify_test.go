package lockfile

import (
	"encoding/json"
	"testing"
)

func TestValidateCompilerInfo(t *testing.T) {
	for i, test := range []struct {
		compilerInfo CompilerInfo
		err          error
	}{
		{
			compilerInfo: CompilerInfo{"serpent"},
			err:          "invalid compiler type selected: serpent",
		},
		{
			compilerInfo: CompilerInfo{"LLL"},
			err:          "invalid compiler type selected: LLL",
		},
		{
			compilerInfo: CompilerInfo{"solc"},
			err:          "",
		},
		{
			compilerInfo: CompilerInfo{"solcjs"},
			err:          "",
		},
	} {
		err = test.compilerInfo.validate()
		if err != nil && len(test.err) == 0 {
			t.Errorf("%d failed. Expected no err but got: %v", i, err)
			continue
		}
		if err == nil && len(test.err) != 0 {
			t.Errorf("%d failed. Expected err: %v but got none", i, test.err)
			continue
		}
		if err != nil && len(test.err) != 0 && err.Error() != test.err {
			t.Errorf("%d failed. Expected err: '%v' got err: '%v'", i, test.err, err)
			continue
		}
	}
}
