// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
)

// Genesis hashes to enforce below configs on.
var (
	MainnetGenesisHash = common.HexToHash("0xd4e56740f876aef8c010b86a40d5f56745a118d0906a34e69aec8c0db1cb8fa3")
	RopstenGenesisHash = common.HexToHash("0x41941023680923e0fe4d74a34bdac8141f2540e3ae90623718e47d66d1ca4a2d")
	RinkebyGenesisHash = common.HexToHash("0x6341fd3daf94b748c72ced5a5b26028f2474f5f00d824504e4fa37a75767e177")
	GoerliGenesisHash  = common.HexToHash("0xbf7e331f7f7c1dd2e05159666b3bf8bc7a8a3a9eb1d518969eab529dd9b88c1a")
)

// TrustedCheckpoints associates each known checkpoint with the genesis hash of
// the chain it belongs to.
var TrustedCheckpoints = map[common.Hash]*TrustedCheckpoint{
	MainnetGenesisHash: MainnetTrustedCheckpoint,
	RopstenGenesisHash: RopstenTrustedCheckpoint,
	RinkebyGenesisHash: RinkebyTrustedCheckpoint,
	GoerliGenesisHash:  GoerliTrustedCheckpoint,
}

// CheckpointOracles associates each known checkpoint oracles with the genesis hash of
// the chain it belongs to.
var CheckpointOracles = map[common.Hash]*CheckpointOracleConfig{
	MainnetGenesisHash: MainnetCheckpointOracle,
	RopstenGenesisHash: RopstenCheckpointOracle,
	RinkebyGenesisHash: RinkebyCheckpointOracle,
	GoerliGenesisHash:  GoerliCheckpointOracle,
}

var (
	// MainnetChainConfig is the chain parameters to run a node on the main network.
	MainnetChainConfig = &ChainConfig{
		ChainID:             big.NewInt(8848),
		HomesteadBlock:      big.NewInt(1),
		EIP150Block:         big.NewInt(2),
		EIP150Hash:          common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		EIP155Block:         big.NewInt(3),
		EIP158Block:         big.NewInt(3),
		ByzantiumBlock:      big.NewInt(4),
		ConstantinopleBlock: nil,
		Alien: &AlienConfig{
			Period:           3,
			Epoch:            201600,
			MaxSignerCount:   21,
			TrantorBlock:     new(big.Int).SetUint64(2968888),
			MinVoterBalance:  new(big.Int).Mul(big.NewInt(100), big.NewInt(1e+18)),
			GenesisTimestamp: 1554004800,
			SelfVoteSigners: []common.UnprefixedAddress{
				common.UnprefixedAddress(common.HexToAddress("0x6e83430ca56ee33a26e5ce87239cb251981ccc2b")),
				common.UnprefixedAddress(common.HexToAddress("0x1807efcb4dc252ff6958eaab770c8b3936a5378f")),
				common.UnprefixedAddress(common.HexToAddress("0x350fccf36124cecd26318e9931414ce872bdb68c")),
				common.UnprefixedAddress(common.HexToAddress("0x09cbad80e089754f610cb8771d9eca05e4e22bdb")),
				common.UnprefixedAddress(common.HexToAddress("0x9d507c10960531c9adc0ffdc9d9c735167275caf")),
				common.UnprefixedAddress(common.HexToAddress("0xc252c0f4d460554c679532072c8dbecd8d9ee89b")),
				common.UnprefixedAddress(common.HexToAddress("0xd8f68e2af8a061f0ea5e57ab7aca1b7fa96dab8a")),
				common.UnprefixedAddress(common.HexToAddress("0x0c58019b9c8e293e3be8d3fd50f77af5f2e84bb7")),
				common.UnprefixedAddress(common.HexToAddress("0x65dd958f7433cbe8353401d131c925e0424330b6")),
				common.UnprefixedAddress(common.HexToAddress("0x076c15f06f36b15544f1e97b4aacbd358d60cdf0")),
				common.UnprefixedAddress(common.HexToAddress("0xecfd032885b4b9e69ab732e800c72296733165d7")),
				common.UnprefixedAddress(common.HexToAddress("0x1a7910fe43b49b8bc33c04cb138cb2a8e1842f32")),
				common.UnprefixedAddress(common.HexToAddress("0xf52fe2e8decbbb3b00ebec7a1a50a41055d784ea")),
				common.UnprefixedAddress(common.HexToAddress("0x35ef874a0f12581fd01fd2b178da7472475e253c")),
				common.UnprefixedAddress(common.HexToAddress("0x90d4a9e77bf64b58f7c07d3bc19f8bb5e9d49031")),
				common.UnprefixedAddress(common.HexToAddress("0x7bd38c427c685fbecbbe0daf49cda466b6475cc6")),
				common.UnprefixedAddress(common.HexToAddress("0xdb1f586092917033e15298663594abb01eb98e39")),
				common.UnprefixedAddress(common.HexToAddress("0x49574ad7832ff9a9214eb462cce2accf35f9118c")),
				common.UnprefixedAddress(common.HexToAddress("0xc8a7ca612be71d84c82c2c1fefbd035517df6745")),
				common.UnprefixedAddress(common.HexToAddress("0x7e13706bab4bfae1f856d75e96676ab27eeea083")),
				common.UnprefixedAddress(common.HexToAddress("0xc5981e7fb6726be96345a732de6206bb1d66b963")),
				common.UnprefixedAddress(common.HexToAddress("0xba99e0bb3fb9537db76a8ac1e76ebca5177954c9")),
				common.UnprefixedAddress(common.HexToAddress("0xd039d1feb6b13c3abe5089da9157fd41104c1aee")),
				common.UnprefixedAddress(common.HexToAddress("0x532c8772925e4b55a6bc99e954aa4cacc7d152b3")),
				common.UnprefixedAddress(common.HexToAddress("0xb464963fcb52b4666577987538a45e68876dc4e7")),
				common.UnprefixedAddress(common.HexToAddress("0x8967f6d04ce36683ebe08c55caa15a177447f983")),
				common.UnprefixedAddress(common.HexToAddress("0x05f39bfe9588f9297b8f3b019a3ee336efe47c47")),
				common.UnprefixedAddress(common.HexToAddress("0x0c59dd1a15c3d5db4b4297cd79bfe72b60affc3e")),
				common.UnprefixedAddress(common.HexToAddress("0x8f05387c4d637288dd197e26d5bdd3cb7087793c")),
				common.UnprefixedAddress(common.HexToAddress("0x02289f35b60c97e27141c6aeb2691d25b531c755")),
				common.UnprefixedAddress(common.HexToAddress("0x39e18521278e5121fdb0b691e84869bd4c645241")),
				common.UnprefixedAddress(common.HexToAddress("0x4b4a0c8cb17b50d8d22610b307c349b63560ca4b")),
				common.UnprefixedAddress(common.HexToAddress("0x9b25f97fa4e3892d9a86ac035a338b36dace5c4b")),
				common.UnprefixedAddress(common.HexToAddress("0xc520c15d943603dc333ebf6b5e39eb4d509fc1f8")),
				common.UnprefixedAddress(common.HexToAddress("0xda863ba260a36a11e3ea953b61de4a0eeffaa6f5")),
				common.UnprefixedAddress(common.HexToAddress("0x3692048ef49479294bcfe9ee7e97508633756f3f")),
				common.UnprefixedAddress(common.HexToAddress("0x27f7fcf7938618dfb0fc3668cd6fe7c1f7315870")),
				common.UnprefixedAddress(common.HexToAddress("0xf34961e5654a76335e0480bd7c7d370ad41ac74f")),
				common.UnprefixedAddress(common.HexToAddress("0x92918ee96f529fdabab1a1ffda627c3d6b442ad9")),
				common.UnprefixedAddress(common.HexToAddress("0x2024cc8d89f7cbd09a4085fbc729e3b9ee92c1be")),
				common.UnprefixedAddress(common.HexToAddress("0x1b5887157beff2e2eff9ea9b8409f3ca1b6a052f")),
				common.UnprefixedAddress(common.HexToAddress("0xad11612be2d9811ffe80f9e9ec1bbdc0ff34067c")),
				common.UnprefixedAddress(common.HexToAddress("0x4e3011ab5b261cff133f4e8fb597dd0980814a94")),
				common.UnprefixedAddress(common.HexToAddress("0x2834dc6b4b054fcf9cb206df4cce17fa0044826b")),
				common.UnprefixedAddress(common.HexToAddress("0x7a2da45fd12d9bd44227ec58a5f0c3085ef18bf1")),
				common.UnprefixedAddress(common.HexToAddress("0xe8ae4d470fb87381f34a77c992a1de53fc2d2a3c")),
				common.UnprefixedAddress(common.HexToAddress("0x777689118d95751e1d709d7134adddd387226ac3")),
				common.UnprefixedAddress(common.HexToAddress("0x0d6556b96b2b7cd095bf42aa2c287df99f22fc87")),
				common.UnprefixedAddress(common.HexToAddress("0xb97b279af3aa97655e6592b320e94505b41631ec")),
				common.UnprefixedAddress(common.HexToAddress("0xbce13d77339971d1f5f00c38f523ba7ee44c95ed")),
			},
		},
	}

	// TestnetChainConfig contains the chain parameters to run a node on the Ropsten test network.
	TestnetChainConfig = &ChainConfig{
		ChainID:             big.NewInt(8341),
		HomesteadBlock:      big.NewInt(1),
		EIP150Block:         big.NewInt(2),
		EIP150Hash:          common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		EIP155Block:         big.NewInt(3),
		EIP158Block:         big.NewInt(3),
		ByzantiumBlock:      big.NewInt(4),
		ConstantinopleBlock: nil,
		Alien: &AlienConfig{
			Period:           3,
			Epoch:            201600,
			MaxSignerCount:   21,
			MinVoterBalance:  new(big.Int).Mul(big.NewInt(100), big.NewInt(1e+18)),
			TrantorBlock:     big.NewInt(695000),
			GenesisTimestamp: 1554004800,
			SelfVoteSigners: []common.UnprefixedAddress{
				common.UnprefixedAddress(common.HexToAddress("0xbe6865ffcbbe5f9746bef5c84b912f2ad9e52075")),
				common.UnprefixedAddress(common.HexToAddress("0x4909b4e54395de9e313ad8a2254fe2dcda99e91c")),
				common.UnprefixedAddress(common.HexToAddress("0xa034350c8e80eb4d15ac62310657b29c711bb3d5")),
			},
		},
	}
	

	// MainnetTrustedCheckpoint contains the light client trusted checkpoint for the main network.
	MainnetTrustedCheckpoint = &TrustedCheckpoint{
		SectionIndex: 300,
		SectionHead:  common.HexToHash("0x022d252ffcd289444eed5a4b8c58018aecb2afc9ab0da5fe059a69a7fb618702"),
		CHTRoot:      common.HexToHash("0xe7044c70ae068969573c7f5abe58ef23d9d82d4ee9152ec88b7c6d0cc8ee2714"),
		BloomRoot:    common.HexToHash("0xe22600caa25653abaef00d0c112b07b90f4e3395ce0c1f5f7f791cdd6d30a408"),
	}

	// MainnetCheckpointOracle contains a set of configs for the main network oracle.
	MainnetCheckpointOracle = &CheckpointOracleConfig{
		Address: common.HexToAddress("0x9a9070028361F7AAbeB3f2F2Dc07F82C4a98A02a"),
		Signers: []common.Address{
			common.HexToAddress("0x1b2C260efc720BE89101890E4Db589b44E950527"), // Peter
			common.HexToAddress("0x78d1aD571A1A09D60D9BBf25894b44e4C8859595"), // Martin
			common.HexToAddress("0x286834935f4A8Cfb4FF4C77D5770C2775aE2b0E7"), // Zsolt
			common.HexToAddress("0xb86e2B0Ab5A4B1373e40c51A7C712c70Ba2f9f8E"), // Gary
			common.HexToAddress("0x0DF8fa387C602AE62559cC4aFa4972A7045d6707"), // Guillaume
		},
		Threshold: 2,
	}

	// RopstenChainConfig contains the chain parameters to run a node on the Ropsten test network.
	RopstenChainConfig = &ChainConfig{
		ChainID:             big.NewInt(3),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.HexToHash("0x41941023680923e0fe4d74a34bdac8141f2540e3ae90623718e47d66d1ca4a2d"),
		EIP155Block:         big.NewInt(10),
		EIP158Block:         big.NewInt(10),
		ByzantiumBlock:      big.NewInt(1700000),
		ConstantinopleBlock: big.NewInt(4230000),
		PetersburgBlock:     big.NewInt(4939394),
		IstanbulBlock:       big.NewInt(6485846),
		MuirGlacierBlock:    big.NewInt(7117117),
		Ethash:              new(EthashConfig),
	}

	// RopstenTrustedCheckpoint contains the light client trusted checkpoint for the Ropsten test network.
	RopstenTrustedCheckpoint = &TrustedCheckpoint{
		SectionIndex: 234,
		SectionHead:  common.HexToHash("0x34659b817e99e6de868b0d4c5321bcff7e36c2cf79307386a2f5053361794d95"),
		CHTRoot:      common.HexToHash("0x249401cd2b07e3f64892729d3f6198514cd11001231a1c001c2e7245659b26e0"),
		BloomRoot:    common.HexToHash("0x37657aa58a07ac3fa13f421c3e5500a944a76def5a11c6d57f17a85f5b33c129"),
	}

	// RopstenCheckpointOracle contains a set of configs for the Ropsten test network oracle.
	RopstenCheckpointOracle = &CheckpointOracleConfig{
		Address: common.HexToAddress("0xEF79475013f154E6A65b54cB2742867791bf0B84"),
		Signers: []common.Address{
			common.HexToAddress("0x32162F3581E88a5f62e8A61892B42C46E2c18f7b"), // Peter
			common.HexToAddress("0x78d1aD571A1A09D60D9BBf25894b44e4C8859595"), // Martin
			common.HexToAddress("0x286834935f4A8Cfb4FF4C77D5770C2775aE2b0E7"), // Zsolt
			common.HexToAddress("0xb86e2B0Ab5A4B1373e40c51A7C712c70Ba2f9f8E"), // Gary
			common.HexToAddress("0x0DF8fa387C602AE62559cC4aFa4972A7045d6707"), // Guillaume
		},
		Threshold: 2,
	}

	// RinkebyChainConfig contains the chain parameters to run a node on the Rinkeby test network.
	RinkebyChainConfig = &ChainConfig{
		ChainID:             big.NewInt(4),
		HomesteadBlock:      big.NewInt(1),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(2),
		EIP150Hash:          common.HexToHash("0x9b095b36c15eaf13044373aef8ee0bd3a382a5abb92e402afa44b8249c3a90e9"),
		EIP155Block:         big.NewInt(3),
		EIP158Block:         big.NewInt(3),
		ByzantiumBlock:      big.NewInt(1035301),
		ConstantinopleBlock: big.NewInt(3660663),
		PetersburgBlock:     big.NewInt(4321234),
		IstanbulBlock:       big.NewInt(5435345),
		Clique: &CliqueConfig{
			Period: 15,
			Epoch:  30000,
		},
	}

	// RinkebyTrustedCheckpoint contains the light client trusted checkpoint for the Rinkeby test network.
	RinkebyTrustedCheckpoint = &TrustedCheckpoint{
		SectionIndex: 191,
		SectionHead:  common.HexToHash("0xfdf3085848b4126048caf176634fd96a208d8a3b055c643e9e32690420df36d5"),
		CHTRoot:      common.HexToHash("0x48059ceb7e0bd25708cc736e5603d28a6f173a3bb904e6e1b3511a97fa30ca97"),
		BloomRoot:    common.HexToHash("0x3566c2b173c0591d5bb4f3ef7e341d82da7577c125fca94e9b51fb7134a676d7"),
	}

	// RinkebyCheckpointOracle contains a set of configs for the Rinkeby test network oracle.
	RinkebyCheckpointOracle = &CheckpointOracleConfig{
		Address: common.HexToAddress("0xebe8eFA441B9302A0d7eaECc277c09d20D684540"),
		Signers: []common.Address{
			common.HexToAddress("0xd9c9cd5f6779558b6e0ed4e6acf6b1947e7fa1f3"), // Peter
			common.HexToAddress("0x78d1aD571A1A09D60D9BBf25894b44e4C8859595"), // Martin
			common.HexToAddress("0x286834935f4A8Cfb4FF4C77D5770C2775aE2b0E7"), // Zsolt
			common.HexToAddress("0xb86e2B0Ab5A4B1373e40c51A7C712c70Ba2f9f8E"), // Gary
		},
		Threshold: 2,
	}

	// GoerliChainConfig contains the chain parameters to run a node on the Görli test network.
	GoerliChainConfig = &ChainConfig{
		ChainID:             big.NewInt(5),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        nil,
		DAOForkSupport:      true,
		EIP150Block:         big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(1561651),
		Clique: &CliqueConfig{
			Period: 15,
			Epoch:  30000,
		},
	}

	// GoerliTrustedCheckpoint contains the light client trusted checkpoint for the Görli test network.
	GoerliTrustedCheckpoint = &TrustedCheckpoint{
		SectionIndex: 76,
		SectionHead:  common.HexToHash("0xf56ca390d1131767b924d85ee8e039c8a4c4a498cfaf017c1a9abf63ef01ff17"),
		CHTRoot:      common.HexToHash("0x78ffc5eecf514eed42f61e6f6df1bdcd79f9296c462faf6f33bd600f70a2e8b9"),
		BloomRoot:    common.HexToHash("0x5186111a2d6c459cc341319398f7d14fa2c973b1ba846b7f2ec678129c7115fd"),
	}

	// GoerliCheckpointOracle contains a set of configs for the Goerli test network oracle.
	GoerliCheckpointOracle = &CheckpointOracleConfig{
		Address: common.HexToAddress("0x18CA0E045F0D772a851BC7e48357Bcaab0a0795D"),
		Signers: []common.Address{
			common.HexToAddress("0x4769bcaD07e3b938B7f43EB7D278Bc7Cb9efFb38"), // Peter
			common.HexToAddress("0x78d1aD571A1A09D60D9BBf25894b44e4C8859595"), // Martin
			common.HexToAddress("0x286834935f4A8Cfb4FF4C77D5770C2775aE2b0E7"), // Zsolt
			common.HexToAddress("0xb86e2B0Ab5A4B1373e40c51A7C712c70Ba2f9f8E"), // Gary
			common.HexToAddress("0x0DF8fa387C602AE62559cC4aFa4972A7045d6707"), // Guillaume
		},
		Threshold: 2,
	}

	// AllEthashProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Ethash consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllEthashProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, nil, new(EthashConfig), nil,nil}

	// AllCliqueProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Clique consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllCliqueProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, nil, nil, &CliqueConfig{Period: 0, Epoch: 30000},nil}

	AllAlienProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), common.Big0, false, common.Big0, common.Hash{}, common.Big0, common.Big0, common.Big0, common.Big0, common.Big0,common.Big0,nil,nil,nil,nil,&AlienConfig{Period: 3, Epoch: 30000, MaxSignerCount: 21, MinVoterBalance: new(big.Int).Mul(big.NewInt(10000), big.NewInt(1000000000000000000)), GenesisTimestamp: 0, SelfVoteSigners: []common.UnprefixedAddress{}}}

	TestChainConfig = &ChainConfig{big.NewInt(1), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, nil, new(EthashConfig), nil,nil}
	TestRules       = TestChainConfig.Rules(new(big.Int))
)

// TrustedCheckpoint represents a set of post-processed trie roots (CHT and
// BloomTrie) associated with the appropriate section index and head hash. It is
// used to start light syncing from this checkpoint and avoid downloading the
// entire header chain while still being able to securely access old headers/logs.
type TrustedCheckpoint struct {
	SectionIndex uint64      `json:"sectionIndex"`
	SectionHead  common.Hash `json:"sectionHead"`
	CHTRoot      common.Hash `json:"chtRoot"`
	BloomRoot    common.Hash `json:"bloomRoot"`
}

// HashEqual returns an indicator comparing the itself hash with given one.
func (c *TrustedCheckpoint) HashEqual(hash common.Hash) bool {
	if c.Empty() {
		return hash == common.Hash{}
	}
	return c.Hash() == hash
}

// Hash returns the hash of checkpoint's four key fields(index, sectionHead, chtRoot and bloomTrieRoot).
func (c *TrustedCheckpoint) Hash() common.Hash {
	buf := make([]byte, 8+3*common.HashLength)
	binary.BigEndian.PutUint64(buf, c.SectionIndex)
	copy(buf[8:], c.SectionHead.Bytes())
	copy(buf[8+common.HashLength:], c.CHTRoot.Bytes())
	copy(buf[8+2*common.HashLength:], c.BloomRoot.Bytes())
	return crypto.Keccak256Hash(buf)
}

// Empty returns an indicator whether the checkpoint is regarded as empty.
func (c *TrustedCheckpoint) Empty() bool {
	return c.SectionHead == (common.Hash{}) || c.CHTRoot == (common.Hash{}) || c.BloomRoot == (common.Hash{})
}

// CheckpointOracleConfig represents a set of checkpoint contract(which acts as an oracle)
// config which used for light client checkpoint syncing.
type CheckpointOracleConfig struct {
	Address   common.Address   `json:"address"`
	Signers   []common.Address `json:"signers"`
	Threshold uint64           `json:"threshold"`
}

// ChainConfig is the core config which determines the blockchain settings.
//
// ChainConfig is stored in the database on a per block basis. This means
// that any network, identified by its genesis block, can have its own
// set of configuration options.
type ChainConfig struct {
	ChainID *big.Int `json:"chainId"` // chainId identifies the current chain and is used for replay protection

	HomesteadBlock *big.Int `json:"homesteadBlock,omitempty"` // Homestead switch block (nil = no fork, 0 = already homestead)

	DAOForkBlock   *big.Int `json:"daoForkBlock,omitempty"`   // TheDAO hard-fork switch block (nil = no fork)
	DAOForkSupport bool     `json:"daoForkSupport,omitempty"` // Whether the nodes supports or opposes the DAO hard-fork

	// EIP150 implements the Gas price changes (https://github.com/ethereum/EIPs/issues/150)
	EIP150Block *big.Int    `json:"eip150Block,omitempty"` // EIP150 HF block (nil = no fork)
	EIP150Hash  common.Hash `json:"eip150Hash,omitempty"`  // EIP150 HF hash (needed for header only clients as only gas pricing changed)

	EIP155Block *big.Int `json:"eip155Block,omitempty"` // EIP155 HF block
	EIP158Block *big.Int `json:"eip158Block,omitempty"` // EIP158 HF block

	ByzantiumBlock      *big.Int `json:"byzantiumBlock,omitempty"`      // Byzantium switch block (nil = no fork, 0 = already on byzantium)
	ConstantinopleBlock *big.Int `json:"constantinopleBlock,omitempty"` // Constantinople switch block (nil = no fork, 0 = already activated)
	PetersburgBlock     *big.Int `json:"petersburgBlock,omitempty"`     // Petersburg switch block (nil = same as Constantinople)
	IstanbulBlock       *big.Int `json:"istanbulBlock,omitempty"`       // Istanbul switch block (nil = no fork, 0 = already on istanbul)
	MuirGlacierBlock    *big.Int `json:"muirGlacierBlock,omitempty"`    // Eip-2384 (bomb delay) switch block (nil = no fork, 0 = already activated)
	EWASMBlock          *big.Int `json:"ewasmBlock,omitempty"`          // EWASM switch block (nil = no fork, 0 = already activated)

	// Various consensus engines
	Ethash *EthashConfig `json:"ethash,omitempty"`
	Clique *CliqueConfig `json:"clique,omitempty"`
	Alien  *AlienConfig  `json:"alien,omitempty"`
}

// EthashConfig is the consensus engine configs for proof-of-work based sealing.
type EthashConfig struct{}

// String implements the stringer interface, returning the consensus engine details.
func (c *EthashConfig) String() string {
	return "ethash"
}

// CliqueConfig is the consensus engine configs for proof-of-authority based sealing.
type CliqueConfig struct {
	Period uint64 `json:"period"` // Number of seconds between blocks to enforce
	Epoch  uint64 `json:"epoch"`  // Epoch length to reset votes and checkpoint
}

// String implements the stringer interface, returning the consensus engine details.
func (c *CliqueConfig) String() string {
	return "clique"
}

type GenesisAccount struct {
	Balance string `json:"balance"`
}

// AlienLightConfig is the config for light node of alien
type AlienLightConfig struct {
	Alloc map[common.UnprefixedAddress]GenesisAccount `json:"alloc"`
}

// AlienConfig is the consensus engine configs for delegated-proof-of-stake based sealing.
type AlienConfig struct {
	Period           uint64                     `json:"period"`           // Number of seconds between blocks to enforce
	Epoch            uint64                     `json:"epoch"`            // Epoch length to reset votes and checkpoint
	MaxSignerCount   uint64                     `json:"maxSignersCount"`  // Max count of signers
	MinVoterBalance  *big.Int                   `json:"minVoterBalance"`  // Min voter balance to valid this vote
	BlockReward      *big.Int                   `json:"blockReward"` // produce block reward
	GenesisTimestamp uint64                     `json:"genesisTimestamp"` // The LoopStartTime of first Block
	SelfVoteSigners  []common.UnprefixedAddress `json:"signers"`          // Signers vote by themselves to seal the block, make sure the signer accounts are pre-funded
	SideChain        bool                       `json:"sideChain"`        // If side chain or not
	MCRPCClient      *rpc.Client                // Main chain rpc client for side chain
	PBFTEnable       bool                       `json:"pbft"` //

	TrantorBlock  *big.Int          `json:"trantorBlock,omitempty"`  // Trantor switch block (nil = no fork)
	TerminusBlock *big.Int          `json:"terminusBlock,omitempty"` // Terminus switch block (nil = no fork)
	LightConfig   *AlienLightConfig `json:"lightConfig,omitempty"`
}

// String implements the stringer interface, returning the consensus engine details.
func (a *AlienConfig) String() string {
	return "alien"
}

// IsTrantor returns whether num is either equal to the Trantor block or greater.
func (a *AlienConfig) IsTrantor(num *big.Int) bool {
	return isForked(a.TrantorBlock, num)
}

// IsTerminus returns whether num is either equal to the Terminus block or greater.
func (a *AlienConfig) IsTerminus(num *big.Int) bool {
	return isForked(a.TerminusBlock, num)
}


// String implements the fmt.Stringer interface.
func (c *ChainConfig) String() string {
	var engine interface{}
	switch {
	case c.Ethash != nil:
		engine = c.Ethash
	case c.Clique != nil:
		engine = c.Clique
	default:
		engine = "unknown"
	}
	return fmt.Sprintf("{ChainID: %v Homestead: %v DAO: %v DAOSupport: %v EIP150: %v EIP155: %v EIP158: %v Byzantium: %v Constantinople: %v Petersburg: %v Istanbul: %v, Muir Glacier: %v, Engine: %v}",
		c.ChainID,
		c.HomesteadBlock,
		c.DAOForkBlock,
		c.DAOForkSupport,
		c.EIP150Block,
		c.EIP155Block,
		c.EIP158Block,
		c.ByzantiumBlock,
		c.ConstantinopleBlock,
		c.PetersburgBlock,
		c.IstanbulBlock,
		c.MuirGlacierBlock,
		engine,
	)
}

// IsHomestead returns whether num is either equal to the homestead block or greater.
func (c *ChainConfig) IsHomestead(num *big.Int) bool {
	return isForked(c.HomesteadBlock, num)
}

// IsDAOFork returns whether num is either equal to the DAO fork block or greater.
func (c *ChainConfig) IsDAOFork(num *big.Int) bool {
	return isForked(c.DAOForkBlock, num)
}

// IsEIP150 returns whether num is either equal to the EIP150 fork block or greater.
func (c *ChainConfig) IsEIP150(num *big.Int) bool {
	return isForked(c.EIP150Block, num)
}

// IsEIP155 returns whether num is either equal to the EIP155 fork block or greater.
func (c *ChainConfig) IsEIP155(num *big.Int) bool {
	return isForked(c.EIP155Block, num)
}

// IsEIP158 returns whether num is either equal to the EIP158 fork block or greater.
func (c *ChainConfig) IsEIP158(num *big.Int) bool {
	return isForked(c.EIP158Block, num)
}

// IsByzantium returns whether num is either equal to the Byzantium fork block or greater.
func (c *ChainConfig) IsByzantium(num *big.Int) bool {
	return isForked(c.ByzantiumBlock, num)
}

// IsConstantinople returns whether num is either equal to the Constantinople fork block or greater.
func (c *ChainConfig) IsConstantinople(num *big.Int) bool {
	return isForked(c.ConstantinopleBlock, num)
}

// IsMuirGlacier returns whether num is either equal to the Muir Glacier (EIP-2384) fork block or greater.
func (c *ChainConfig) IsMuirGlacier(num *big.Int) bool {
	return isForked(c.MuirGlacierBlock, num)
}

// IsPetersburg returns whether num is either
// - equal to or greater than the PetersburgBlock fork block,
// - OR is nil, and Constantinople is active
func (c *ChainConfig) IsPetersburg(num *big.Int) bool {
	return isForked(c.PetersburgBlock, num) || c.PetersburgBlock == nil && isForked(c.ConstantinopleBlock, num)
}

// IsIstanbul returns whether num is either equal to the Istanbul fork block or greater.
func (c *ChainConfig) IsIstanbul(num *big.Int) bool {
	return isForked(c.IstanbulBlock, num)
}

// IsEWASM returns whether num represents a block number after the EWASM fork
func (c *ChainConfig) IsEWASM(num *big.Int) bool {
	return isForked(c.EWASMBlock, num)
}

// CheckCompatible checks whether scheduled fork transitions have been imported
// with a mismatching chain configuration.
func (c *ChainConfig) CheckCompatible(newcfg *ChainConfig, height uint64) *ConfigCompatError {
	bhead := new(big.Int).SetUint64(height)

	// Iterate checkCompatible to find the lowest conflict.
	var lasterr *ConfigCompatError
	for {
		err := c.checkCompatible(newcfg, bhead)
		if err == nil || (lasterr != nil && err.RewindTo == lasterr.RewindTo) {
			break
		}
		lasterr = err
		bhead.SetUint64(err.RewindTo)
	}
	return lasterr
}

// CheckConfigForkOrder checks that we don't "skip" any forks, geth isn't pluggable enough
// to guarantee that forks can be implemented in a different order than on official networks
func (c *ChainConfig) CheckConfigForkOrder() error {
	type fork struct {
		name  string
		block *big.Int
	}
	var lastFork fork
	for _, cur := range []fork{
		{"homesteadBlock", c.HomesteadBlock},
		{"eip150Block", c.EIP150Block},
		{"eip155Block", c.EIP155Block},
		{"eip158Block", c.EIP158Block},
		{"byzantiumBlock", c.ByzantiumBlock},
		{"constantinopleBlock", c.ConstantinopleBlock},
		{"petersburgBlock", c.PetersburgBlock},
		{"istanbulBlock", c.IstanbulBlock},
		{"muirGlacierBlock", c.MuirGlacierBlock},
	} {
		if lastFork.name != "" {
			// Next one must be higher number
			if lastFork.block == nil && cur.block != nil {
				return fmt.Errorf("unsupported fork ordering: %v not enabled, but %v enabled at %v",
					lastFork.name, cur.name, cur.block)
			}
			if lastFork.block != nil && cur.block != nil {
				if lastFork.block.Cmp(cur.block) > 0 {
					return fmt.Errorf("unsupported fork ordering: %v enabled at %v, but %v enabled at %v",
						lastFork.name, lastFork.block, cur.name, cur.block)
				}
			}
		}
		lastFork = cur
	}
	return nil
}

func (c *ChainConfig) checkCompatible(newcfg *ChainConfig, head *big.Int) *ConfigCompatError {
	if isForkIncompatible(c.HomesteadBlock, newcfg.HomesteadBlock, head) {
		return newCompatError("Homestead fork block", c.HomesteadBlock, newcfg.HomesteadBlock)
	}
	if isForkIncompatible(c.DAOForkBlock, newcfg.DAOForkBlock, head) {
		return newCompatError("DAO fork block", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if c.IsDAOFork(head) && c.DAOForkSupport != newcfg.DAOForkSupport {
		return newCompatError("DAO fork support flag", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if isForkIncompatible(c.EIP150Block, newcfg.EIP150Block, head) {
		return newCompatError("EIP150 fork block", c.EIP150Block, newcfg.EIP150Block)
	}
	if isForkIncompatible(c.EIP155Block, newcfg.EIP155Block, head) {
		return newCompatError("EIP155 fork block", c.EIP155Block, newcfg.EIP155Block)
	}
	if isForkIncompatible(c.EIP158Block, newcfg.EIP158Block, head) {
		return newCompatError("EIP158 fork block", c.EIP158Block, newcfg.EIP158Block)
	}
	if c.IsEIP158(head) && !configNumEqual(c.ChainID, newcfg.ChainID) {
		return newCompatError("EIP158 chain ID", c.EIP158Block, newcfg.EIP158Block)
	}
	if isForkIncompatible(c.ByzantiumBlock, newcfg.ByzantiumBlock, head) {
		return newCompatError("Byzantium fork block", c.ByzantiumBlock, newcfg.ByzantiumBlock)
	}
	if isForkIncompatible(c.ConstantinopleBlock, newcfg.ConstantinopleBlock, head) {
		return newCompatError("Constantinople fork block", c.ConstantinopleBlock, newcfg.ConstantinopleBlock)
	}
	if isForkIncompatible(c.PetersburgBlock, newcfg.PetersburgBlock, head) {
		return newCompatError("Petersburg fork block", c.PetersburgBlock, newcfg.PetersburgBlock)
	}
	if isForkIncompatible(c.IstanbulBlock, newcfg.IstanbulBlock, head) {
		return newCompatError("Istanbul fork block", c.IstanbulBlock, newcfg.IstanbulBlock)
	}
	if isForkIncompatible(c.MuirGlacierBlock, newcfg.MuirGlacierBlock, head) {
		return newCompatError("Muir Glacier fork block", c.MuirGlacierBlock, newcfg.MuirGlacierBlock)
	}
	if isForkIncompatible(c.EWASMBlock, newcfg.EWASMBlock, head) {
		return newCompatError("ewasm fork block", c.EWASMBlock, newcfg.EWASMBlock)
	}
	return nil
}

// isForkIncompatible returns true if a fork scheduled at s1 cannot be rescheduled to
// block s2 because head is already past the fork.
func isForkIncompatible(s1, s2, head *big.Int) bool {
	return (isForked(s1, head) || isForked(s2, head)) && !configNumEqual(s1, s2)
}

// isForked returns whether a fork scheduled at block s is active at the given head block.
func isForked(s, head *big.Int) bool {
	if s == nil || head == nil {
		return false
	}
	return s.Cmp(head) <= 0
}

func configNumEqual(x, y *big.Int) bool {
	if x == nil {
		return y == nil
	}
	if y == nil {
		return x == nil
	}
	return x.Cmp(y) == 0
}

// ConfigCompatError is raised if the locally-stored blockchain is initialised with a
// ChainConfig that would alter the past.
type ConfigCompatError struct {
	What string
	// block numbers of the stored and new configurations
	StoredConfig, NewConfig *big.Int
	// the block number to which the local chain must be rewound to correct the error
	RewindTo uint64
}

func newCompatError(what string, storedblock, newblock *big.Int) *ConfigCompatError {
	var rew *big.Int
	switch {
	case storedblock == nil:
		rew = newblock
	case newblock == nil || storedblock.Cmp(newblock) < 0:
		rew = storedblock
	default:
		rew = newblock
	}
	err := &ConfigCompatError{what, storedblock, newblock, 0}
	if rew != nil && rew.Sign() > 0 {
		err.RewindTo = rew.Uint64() - 1
	}
	return err
}

func (err *ConfigCompatError) Error() string {
	return fmt.Sprintf("mismatching %s in database (have %d, want %d, rewindto %d)", err.What, err.StoredConfig, err.NewConfig, err.RewindTo)
}

// Rules wraps ChainConfig and is merely syntactic sugar or can be used for functions
// that do not have or require information about the block.
//
// Rules is a one time interface meaning that it shouldn't be used in between transition
// phases.
type Rules struct {
	ChainID                                                 *big.Int
	IsHomestead, IsEIP150, IsEIP155, IsEIP158               bool
	IsByzantium, IsConstantinople, IsPetersburg, IsIstanbul bool
}

// Rules ensures c's ChainID is not nil.
func (c *ChainConfig) Rules(num *big.Int) Rules {
	chainID := c.ChainID
	if chainID == nil {
		chainID = new(big.Int)
	}
	return Rules{
		ChainID:          new(big.Int).Set(chainID),
		IsHomestead:      c.IsHomestead(num),
		IsEIP150:         c.IsEIP150(num),
		IsEIP155:         c.IsEIP155(num),
		IsEIP158:         c.IsEIP158(num),
		IsByzantium:      c.IsByzantium(num),
		IsConstantinople: c.IsConstantinople(num),
		IsPetersburg:     c.IsPetersburg(num),
		IsIstanbul:       c.IsIstanbul(num),
	}
}
