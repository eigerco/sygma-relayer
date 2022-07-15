package topology_test

import (
	"os"
	"testing"

	"github.com/ChainSafe/sygma/topology"
	"github.com/stretchr/testify/suite"
)

type TopologyStoreTestSuite struct {
	suite.Suite
	topologyStore *topology.TopologyStore
	path          string
}

func TestRunTopologyStoreTestSuite(t *testing.T) {
	suite.Run(t, new(TopologyStoreTestSuite))
}

func (s *TopologyStoreTestSuite) SetupSuite()    {}
func (s *TopologyStoreTestSuite) TearDownSuite() {}
func (s *TopologyStoreTestSuite) SetupTest() {
	s.path = "topology.json"
	s.topologyStore = topology.NewTopologyStore(s.path)
}
func (s *TopologyStoreTestSuite) TearDownTest() {
	os.Remove(s.path)
}

func (s *TopologyStoreTestSuite) Test_RetrieveNonExistentFile_Error() {
	_, err := s.topologyStore.Topology()
	s.NotNil(err)
}

func (s *TopologyStoreTestSuite) Test_StoreAndRetrieveTopology() {
	networkTopology, _ := topology.ProcessRawTopology(&topology.RawTopology{
		Peers: []topology.RawPeer{
			{PeerAddress: "/dns4/relayer2/tcp/9001/p2p/QmeTuMtdpPB7zKDgmobEwSvxodrf5aFVSmBXX3SQJVjJaT"},
			{PeerAddress: "/dns4/relayer3/tcp/9002/p2p/QmYAYuLUPNwYEBYJaKHcE7NKjUhiUV8txx2xDXHvcYa1xK"},
			{PeerAddress: "/dns4/relayer1/tcp/9000/p2p/QmcvEg7jGvuxdsUFRUiE4VdrL2P1Yeju5L83BsJvvXz7zX"},
		},
	})

	err := s.topologyStore.StoreTopology(networkTopology)
	s.Nil(err)

	storedTopology, err := s.topologyStore.Topology()
	s.Nil(err)

	s.Equal(networkTopology, storedTopology)
}
