package cannon

import (
	"encoding/json"
	"fmt"

	"github.com/BLASTchain/blast/cannon/mipsevm"
	"github.com/BLASTchain/blast/bl-service/ioutil"
)

func parseState(path string) (*mipsevm.State, error) {
	file, err := ioutil.OpenDecompressed(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open state file (%v): %w", path, err)
	}
	defer file.Close()
	var state mipsevm.State
	err = json.NewDecoder(file).Decode(&state)
	if err != nil {
		return nil, fmt.Errorf("invalid mipsevm state (%v): %w", path, err)
	}
	return &state, nil
}
