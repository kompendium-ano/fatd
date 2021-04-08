// MIT License
//
// Copyright 2020-2021 Kompendium, LLC
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

package engine

import (
	"fmt"
	"sync"
)

// Peer represents a connected remote node.
type Peer struct {
	addr    uint32
	hash    string
	version string
}

// PeerStore holds active Peers, managing them in a concurrency safe
// manner and providing lookup via various functions
type PeerStore struct {
	mtx       sync.RWMutex
	peers     map[string]*Peer // hash -> peer
	connected map[string]int   // (ip|ip:port) -> count
	current   []*Peer
	incoming  int
	outgoing  int
}

// NewPeerStore initializes a new peer store
func NewPeerStore() *PeerStore {
	ps := new(PeerStore)
	ps.peers = make(map[string]*Peer)
	ps.connected = make(map[string]int)
	return ps
}

// Add a peer to be managed. Returns an error if a peer with that hash
// is already tracked
func (ps *PeerStore) Add(p *Peer) error {
	if p == nil {
		return fmt.Errorf("trying to add nil")
	}
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	if _, ok := ps.peers[p.hash]; ok {
		return fmt.Errorf("peer already exists")
	}
	ps.current = nil
	ps.peers[p.hash] = p
	ps.connected[fmt.Sprint(p.addr)]++
	ps.connected[p.version]++

	return nil
}

// Get list of all peers available on the Network
func GetPeers() []Peer {
	// do network discovery

	return []Peer{}
}
