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
	"enode://27752d8b5284e48cbc43903ebaa1cab212174507d8ac0a94936b97c668baa31781864ea91b1c9286561a0e296f681422e626e0890f8ff3e7fe20de028dbedb4a@101.133.151.154:26691",
	"enode://f987588021270aadcd19a79b2f45285150c9e7924322e40c50339c6f636ed37edb5e44b7afc243162e70064af3b61b26b0d1a69d0fab8ac91893230e7de5c562@47.103.38.41:26691",
	"enode://9c8ace6b8facb413c1bcc8aa8ade295009ce186a63ad616ed5ecf9400293e572b4173fb1520ede0ca7b90bc3f086db5126154ef9cb908b7556cdafc5724620aa@161.189.9.100:26691",
	"enode://24decef4fa4ab3dfd84c331a3a5e626d671ffec76d505524bc2e169436c90d510eb3aec366f5a49cfea30984f21b24095bf58e09ed171ceab2d2559b536337e1@101.133.225.179:26698",
}
var TestnetBootnodes = []string{
	"enode://27752d8b5284e48cbc43903ebaa1cab212174507d8ac0a94936b97c668baa31781864ea91b1c9286561a0e296f681422e626e0890f8ff3e7fe20de028dbedb4a@101.133.151.154:26691",
	"enode://f987588021270aadcd19a79b2f45285150c9e7924322e40c50339c6f636ed37edb5e44b7afc243162e70064af3b61b26b0d1a69d0fab8ac91893230e7de5c562@47.103.38.41:26691",
	"enode://9c8ace6b8facb413c1bcc8aa8ade295009ce186a63ad616ed5ecf9400293e572b4173fb1520ede0ca7b90bc3f086db5126154ef9cb908b7556cdafc5724620aa@161.189.9.100:26691",
	"enode://24decef4fa4ab3dfd84c331a3a5e626d671ffec76d505524bc2e169436c90d510eb3aec366f5a49cfea30984f21b24095bf58e09ed171ceab2d2559b536337e1@101.133.225.179:26698",
}

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
