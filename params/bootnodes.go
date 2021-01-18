// Copyright 2015 The go-ethereum Authors
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

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{}
var TestnetBootnodes = []string{
	"enode://843d26b28043b426017c4f7f43e77af6e5ef048667f1b5a09a13e9da3b1012d504cf16a5fb9e078d9136149e9009ced14ea98510466522133f4720c782507d32@47.91.180.165:26681",
	"enode://b00af2b2f8d3754f37d122d960fbd0c3ebf4649a4da5ad30b57245c47a2a765995d548a98cc7c59486052b69d666ab22266c7870460b5f5b2da5288e50799c3e@101.133.225.179:26681",
}

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
