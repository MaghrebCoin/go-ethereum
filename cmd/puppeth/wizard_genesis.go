// Copyright 2017 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

// makeGenesis creates a new genesis struct based on some user input.
func (w *wizard) makeGenesis() {
	// Construct a default genesis block
	genesis := &core.Genesis{
		Timestamp:  uint64(time.Now().Unix()),
		GasLimit:   4700000,
		Difficulty: big.NewInt(524288),
		Alloc:      make(core.GenesisAlloc),
		Config: &params.ChainConfig{
			HomesteadBlock:      big.NewInt(0),
			EIP150Block:         big.NewInt(1),
			EIP155Block:         big.NewInt(2),
			EIP158Block:         big.NewInt(3),
			ByzantiumBlock:      big.NewInt(4),
			ConstantinopleBlock: big.NewInt(5),
			PetersburgBlock:     big.NewInt(6),
			IstanbulBlock:       big.NewInt(7),
		},
	}
	// Figure out which consensus engine to choose
	fmt.Println()
	fmt.Println("Which consensus engine to use? (default = clique)")
	fmt.Println(" 1. Ethash - proof-of-work")
	fmt.Println(" 2. Clique - proof-of-authority")
	fmt.Println(" 3. Alien  - delegated-proof-of-stake")

	choice := w.read()
	switch {
	case choice == "1":
		// In case of ethash, we're pretty much done
		genesis.Config.Ethash = new(params.EthashConfig)
		genesis.ExtraData = make([]byte, 32)

	case choice == "" || choice == "2":
		// In the case of clique, configure the consensus parameters
		genesis.Difficulty = big.NewInt(1)
		genesis.Config.Clique = &params.CliqueConfig{
			Period: 15,
			Epoch:  30000,
		}
		fmt.Println()
		fmt.Println("How many seconds should blocks take? (default = 15)")
		genesis.Config.Clique.Period = uint64(w.readDefaultInt(15))

		// We also need the initial list of signers
		fmt.Println()
		fmt.Println("Which accounts are allowed to seal? (mandatory at least one)")

		var signers []common.Address
		for {
			if address := w.readAddress(); address != nil {
				signers = append(signers, *address)
				continue
			}
			if len(signers) > 0 {
				break
			}
		}
		// Sort the signers and embed into the extra-data section
		for i := 0; i < len(signers); i++ {
			for j := i + 1; j < len(signers); j++ {
				if bytes.Compare(signers[i][:], signers[j][:]) > 0 {
					signers[i], signers[j] = signers[j], signers[i]
				}
			}
		}
		genesis.ExtraData = make([]byte, 32+len(signers)*common.AddressLength+65)
		for i, signer := range signers {
			copy(genesis.ExtraData[32+i*common.AddressLength:], signer[:])
		}
	case choice == "" || choice == "3":
		// In the case of alien, configure the consensus parameters
		genesis.Difficulty = big.NewInt(1)
		genesis.Config.Alien = &params.AlienConfig{
			Period:           3,
			Epoch:            201600,
			MaxSignerCount:   21,
			MinVoterBalance:  new(big.Int).Mul(big.NewInt(1000), big.NewInt(1e+18)),
			GenesisTimestamp: uint64(time.Now().Unix()) + (60 * 5), // Add five minutes
			SelfVoteSigners:  []common.UnprefixedAddress{},
		}
		fmt.Println()
		fmt.Println("How many seconds should blocks take? (default = 3)")
		genesis.Config.Alien.Period = uint64(w.readDefaultInt(3))

		fmt.Println()
		fmt.Println("How many blocks create for one epoch? (default = 201600)")
		genesis.Config.Alien.Epoch = uint64(w.readDefaultInt(201600))

		fmt.Println()
		fmt.Println("What is the max number of signers? (default = 21)")
		genesis.Config.Alien.MaxSignerCount = uint64(w.readDefaultInt(21))

		fmt.Println()
		fmt.Println("What is the minimize balance for valid voter ? (default = 1000 ETH)")
		genesis.Config.Alien.MinVoterBalance = new(big.Int).Mul(big.NewInt(int64(w.readDefaultInt(1000))),
			big.NewInt(1e+18))

		fmt.Println()
		fmt.Println("How many block reward one block generate ? (default = 10 ETH)")
		genesis.Config.Alien.BlockReward = new(big.Int).Mul(big.NewInt(int64(w.readDefaultInt(10))),
			big.NewInt(1e+18))

		fmt.Println()
		fmt.Println("How many minutes delay to create first block ? (default = 5 minutes)")
		genesis.Config.Alien.GenesisTimestamp = uint64(time.Now().Unix()) + uint64(w.readDefaultInt(5)*60)

		// We also need the initial list of signers
		fmt.Println()
		fmt.Println("Which accounts are vote by themselves to seal the block?(least one, those accounts will be auto pre-funded)")
		for {
			if address := w.readAddress(); address != nil {

				genesis.Config.Alien.SelfVoteSigners = append(genesis.Config.Alien.SelfVoteSigners, common.UnprefixedAddress(*address))
				genesis.Alloc[*address] = core.GenesisAccount{
					Balance: genesis.Config.Alien.MinVoterBalance, // 2^256 / 128 (allow many pre-funds without balance overflows)
				}
				continue
			}
			if len(genesis.Config.Alien.SelfVoteSigners) > 0 {
				break
			}
		}

		genesis.ExtraData = make([]byte, 32+65)
	default:
		log.Crit("Invalid consensus engine choice", "choice", choice)
	}
	// Consensus all set, just ask for initial funds and go
	fmt.Println()
	fmt.Println("Which accounts should be pre-funded? (advisable at least one)")
	for {
		// Read the address of the account to fund
		if address := w.readAddress(); address != nil {
			genesis.Alloc[*address] = core.GenesisAccount{
				Balance: new(big.Int).Lsh(big.NewInt(1), 256-7), // 2^256 / 128 (allow many pre-funds without balance overflows)
			}
			continue
		}
		break
	}
	fmt.Println()
	fmt.Println("Should the precompile-addresses (0x1 .. 0xff) be pre-funded with 1 wei? (advisable yes)")
	if w.readDefaultYesNo(true) {
		// Add a batch of precompile balances to avoid them getting deleted
		for i := int64(0); i < 256; i++ {
			genesis.Alloc[common.BigToAddress(big.NewInt(i))] = core.GenesisAccount{Balance: big.NewInt(1)}
		}
	}
	// Query the user for some custom extras
	fmt.Println()
	fmt.Println("Specify your chain/network ID if you want an explicit one (default = random)")
	genesis.Config.ChainID = new(big.Int).SetUint64(uint64(w.readDefaultInt(rand.Intn(65536))))

	// All done, store the genesis and flush to disk
	log.Info("Configured new genesis block")

	w.conf.Genesis = genesis
	w.conf.flush()
}

// importGenesis imports a Geth genesis spec into puppeth.
func (w *wizard) importGenesis() {
	// Request the genesis JSON spec URL from the user
	fmt.Println()
	fmt.Println("Where's the genesis file? (local file or http/https url)")
	url := w.readURL()

	// Convert the various allowed URLs to a reader stream
	var reader io.Reader

	switch url.Scheme {
	case "http", "https":
		// Remote web URL, retrieve it via an HTTP client
		res, err := http.Get(url.String())
		if err != nil {
			log.Error("Failed to retrieve remote genesis", "err", err)
			return
		}
		defer res.Body.Close()
		reader = res.Body

	case "":
		// Schemaless URL, interpret as a local file
		file, err := os.Open(url.String())
		if err != nil {
			log.Error("Failed to open local genesis", "err", err)
			return
		}
		defer file.Close()
		reader = file

	default:
		log.Error("Unsupported genesis URL scheme", "scheme", url.Scheme)
		return
	}
	// Parse the genesis file and inject it successful
	var genesis core.Genesis
	if err := json.NewDecoder(reader).Decode(&genesis); err != nil {
		log.Error("Invalid genesis spec: %v", err)
		return
	}
	log.Info("Imported genesis block")

	w.conf.Genesis = &genesis
	w.conf.flush()
}

// manageGenesis permits the modification of chain configuration parameters in
// a genesis config and the export of the entire genesis spec.
func (w *wizard) manageGenesis() {
	// Figure out whether to modify or export the genesis
	fmt.Println()
	fmt.Println(" 1. Modify existing configurations")
	fmt.Println(" 2. Export genesis configurations")
	fmt.Println(" 3. Remove genesis configuration")

	choice := w.read()
	switch choice {
	case "1":
		// Fork rule updating requested, iterate over each fork
		fmt.Println()
		fmt.Printf("Which block should Homestead come into effect? (default = %v)\n", w.conf.Genesis.Config.HomesteadBlock)
		w.conf.Genesis.Config.HomesteadBlock = w.readDefaultBigInt(w.conf.Genesis.Config.HomesteadBlock)

		fmt.Println()
		fmt.Printf("Which block should EIP150 (Tangerine Whistle) come into effect? (default = %v)\n", w.conf.Genesis.Config.EIP150Block)
		w.conf.Genesis.Config.EIP150Block = w.readDefaultBigInt(w.conf.Genesis.Config.EIP150Block)

		fmt.Println()
		fmt.Printf("Which block should EIP155 (Spurious Dragon) come into effect? (default = %v)\n", w.conf.Genesis.Config.EIP155Block)
		w.conf.Genesis.Config.EIP155Block = w.readDefaultBigInt(w.conf.Genesis.Config.EIP155Block)

		fmt.Println()
		fmt.Printf("Which block should EIP158/161 (also Spurious Dragon) come into effect? (default = %v)\n", w.conf.Genesis.Config.EIP158Block)
		w.conf.Genesis.Config.EIP158Block = w.readDefaultBigInt(w.conf.Genesis.Config.EIP158Block)

		fmt.Println()
		fmt.Printf("Which block should Byzantium come into effect? (default = %v)\n", w.conf.Genesis.Config.ByzantiumBlock)
		w.conf.Genesis.Config.ByzantiumBlock = w.readDefaultBigInt(w.conf.Genesis.Config.ByzantiumBlock)

		fmt.Println()
		fmt.Printf("Which block should Constantinople come into effect? (default = %v)\n", w.conf.Genesis.Config.ConstantinopleBlock)
		w.conf.Genesis.Config.ConstantinopleBlock = w.readDefaultBigInt(w.conf.Genesis.Config.ConstantinopleBlock)
		if w.conf.Genesis.Config.PetersburgBlock == nil {
			w.conf.Genesis.Config.PetersburgBlock = w.conf.Genesis.Config.ConstantinopleBlock
		}
		fmt.Println()
		fmt.Printf("Which block should Petersburg come into effect? (default = %v)\n", w.conf.Genesis.Config.PetersburgBlock)
		w.conf.Genesis.Config.PetersburgBlock = w.readDefaultBigInt(w.conf.Genesis.Config.PetersburgBlock)

		fmt.Println()
		fmt.Printf("Which block should Istanbul come into effect? (default = %v)\n", w.conf.Genesis.Config.IstanbulBlock)
		w.conf.Genesis.Config.IstanbulBlock = w.readDefaultBigInt(w.conf.Genesis.Config.IstanbulBlock)

		out, _ := json.MarshalIndent(w.conf.Genesis.Config, "", "  ")
		fmt.Printf("Chain configuration updated:\n\n%s\n", out)

		w.conf.flush()

	case "2":
		// Save whatever genesis configuration we currently have
		fmt.Println()
		fmt.Printf("Which file to save the genesis into? (default = %s.json)\n", w.network)
		out, _ := json.MarshalIndent(w.conf.Genesis, "", "  ")
		if err := ioutil.WriteFile(w.readDefaultString(fmt.Sprintf("%s.json", w.network)), out, 0644); err != nil {
			log.Error("Failed to save genesis file", "err", err)
		}
		log.Info("Exported existing genesis block")

	case "3":
		// Make sure we don't have any services running
		if len(w.conf.servers()) > 0 {
			log.Error("Genesis reset requires all services and servers torn down")
			return
		}
		log.Info("Genesis block destroyed")

		w.conf.Genesis = nil
		w.conf.flush()
	default:
		log.Error("That's not something I can do")
		return
	}
}

// saveGenesis JSON encodes an arbitrary genesis spec into a pre-defined file.
func saveGenesis(folder, network, client string, spec interface{}) {
	path := filepath.Join(folder, fmt.Sprintf("%s-%s.json", network, client))

	out, _ := json.MarshalIndent(spec, "", "  ")
	if err := ioutil.WriteFile(path, out, 0644); err != nil {
		log.Error("Failed to save genesis file", "client", client, "err", err)
		return
	}
	log.Info("Saved genesis chain spec", "client", client, "path", path)
}
