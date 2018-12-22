package state

import (
	"fmt"

	"github.com/Factom-Asset-Tokens/fatd/factom"
	"github.com/Factom-Asset-Tokens/fatd/fat0"
)

func ProcessEBlock(eb factom.EBlock) error {
	// Get the saved data for this chain.
	chain := chains.Get(eb.ChainID)

	// Skip ignored chains or EBlocks for heights earlier than this chain's
	// state.
	if chain.IsIgnored() || eb.Height <= chain.metadata.Height {
		return nil
	}

	defer func() {
		chains.Set(chain)
	}()

	// Load this Entry Block.
	if err := eb.Get(); err != nil {
		return fmt.Errorf("%#v.Get(): %v", eb, err)
	}
	if !eb.IsPopulated() {
		return fmt.Errorf("%#v.IsPopulated(): false", eb)
	}

	// Check for new chains.
	if eb.IsFirst() {
		// Load first entry of new chain.
		if err := eb.Entries[0].Get(); err != nil {
			return fmt.Errorf("%#v.Get: %v", eb.Entries[0], err)
		}
		if !eb.Entries[0].IsPopulated() {
			return fmt.Errorf("%#v.IsPopulated(): false", eb.Entries[0])
		}
		nameIDs := eb.Entries[0].ExtIDs

		// Ignore chains with NameIDs that don't match the fat0
		// pattern.
		if !fat0.ValidTokenNameIDs(nameIDs) {
			chain.Ignore()
			return nil
		}

		// Track this chain going forward.
		if err := chain.Track(nameIDs); err != nil {
			return err
		}
		if len(eb.Entries) == 1 {
			return nil
		}
		// The first entry cannot be a valid Issuance entry, so discard
		// it and process the rest.
		eb.Entries = eb.Entries[1:]
	} else if !chain.IsTracked() {
		// Ignore chains that are not already tracked.
		chain.Ignore()
		return nil
	}

	if err := chain.ProcessEBlock(eb); err != nil {
		return err
	}
	return nil
}