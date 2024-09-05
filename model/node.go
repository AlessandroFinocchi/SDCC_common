package model

import (
	"fmt"
	"github.com/AlessandroFinocchi/sdcc_common/pb"
)

func ProtoNodeMembershipAddress(n *pb.Node) string {
	return fmt.Sprintf("%s:%d", n.GetMembershipIp(), n.GetMembershipPort())
}

func ProtoNodeVivaldiAddress(n *pb.Node) string {
	return fmt.Sprintf("%s:%d", n.GetVivaldiIp(), n.GetVivaldiPort())
}

func ProtoNodeGossipAddress(n *pb.Node) string {
	return fmt.Sprintf("%s:%d", n.GetGossipIp(), n.GetGossipPort())
}
