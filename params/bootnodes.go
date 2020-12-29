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
var MainnetBootnodes = []string{
	"enode://ab459fc90922b2afa61a05fdb5f14c7399839cecf7f682a0b4ec3df726259568c584d9e06f376a1de12b78b28c8ce5612e7186aad5fc2604a9d586705050b050@101.133.151.154:26691",
	"enode://1fd588aeb01f3501b39f9c63661312362980bb215fba78ce0de0808afb4650818f656294aa937cd3e1876088ebc402a9e79617b5795ea51084adc6046d375cf3@101.133.225.179:26698",
	"enode://f987588021270aadcd19a79b2f45285150c9e7924322e40c50339c6f636ed37edb5e44b7afc243162e70064af3b61b26b0d1a69d0fab8ac91893230e7de5c562@47.103.38.41:26691",
	"enode://d38a70907b42f2b92c1485126cce91c09f9a66400babfcfbf8c47ba58740c86824c6a7fc0a7c2841c516b93fdd0ebd6a21de2ba902bc3b288cd73b903d4041b6@161.189.9.100:26691",
}
var TestnetBootnodes = []string{
	"enode://ab459fc90922b2afa61a05fdb5f14c7399839cecf7f682a0b4ec3df726259568c584d9e06f376a1de12b78b28c8ce5612e7186aad5fc2604a9d586705050b050@101.133.151.154:26691",
	"enode://1fd588aeb01f3501b39f9c63661312362980bb215fba78ce0de0808afb4650818f656294aa937cd3e1876088ebc402a9e79617b5795ea51084adc6046d375cf3@101.133.225.179:26698",
	"enode://f987588021270aadcd19a79b2f45285150c9e7924322e40c50339c6f636ed37edb5e44b7afc243162e70064af3b61b26b0d1a69d0fab8ac91893230e7de5c562@47.103.38.41:26691",
	"enode://d38a70907b42f2b92c1485126cce91c09f9a66400babfcfbf8c47ba58740c86824c6a7fc0a7c2841c516b93fdd0ebd6a21de2ba902bc3b288cd73b903d4041b6@161.189.9.100:26691",
}

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
