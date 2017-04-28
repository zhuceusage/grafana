package clustering

import (
	"errors"
	"fmt"

	"github.com/grafana/grafana/pkg/bus"
	"github.com/grafana/grafana/pkg/log"
	m "github.com/grafana/grafana/pkg/models"
	"github.com/grafana/grafana/pkg/setting"
)

type ClusterNodeMgmt interface {
	GetNodeId() (string, error)
	CheckIn() error
	GetNode(heartbeat int64) (*m.ActiveNode, error)
	CheckInNodeProcessingMissingAlerts() error
	GetActiveNodesCount(heartbeat int64) (int, error)
	GetLastHeartbeat() (int64, error)
}

type ClusterNode struct {
	nodeId string
	log    log.Logger
}

var (
	clusterNodeMgmt ClusterNodeMgmt = nil
)

func getCurrentNodeId() string {
	if setting.InstanceName == "" || setting.HttpPort == "" {
		panic("InstanceName or HttpPort is empty. Check your configuration.")
	}
	return fmt.Sprintf("%v:%v", setting.InstanceName, setting.HttpPort)
}

func getClusterNode() ClusterNodeMgmt {
	if clusterNodeMgmt != nil {
		return clusterNodeMgmt
	}
	node := &ClusterNode{
		nodeId: getCurrentNodeId(),
		log:    log.New("clustering.clusterNode"),
	}
	clusterNodeMgmt = node
	return node

}

func (node *ClusterNode) GetNodeId() (string, error) {
	if node == nil {
		return "", errors.New("Cluster node object is nil")
	}
	return node.nodeId, nil
}

func (node *ClusterNode) CheckIn() error {
	if node == nil {
		return errors.New("Cluster node object is nil")
	}
	cmd := &m.SaveActiveNodeCommand{
		Node:        &m.ActiveNode{NodeId: node.nodeId, AlertRunType: m.NORMAL_ALERT},
		FetchResult: true,
	}
	node.log.Debug("Sending command ", "SaveActiveNodeCommand:Node", cmd.Node)
	if err := bus.Dispatch(cmd); err != nil {

		errmsg := fmt.Sprintf("Failed to save heartbeat - node %v", cmd.Node)
		node.log.Error(errmsg, "error", err)
		return err
	}
	node.log.Debug("Command executed successfully")
	return nil
}

func (node *ClusterNode) GetNode(heartbeat int64) (*m.ActiveNode, error) {
	if node == nil {
		return nil, errors.New("Cluster node object is nil")
	}
	cmd := &m.GetActiveNodeByIdHeartbeatQuery{
		NodeId:    node.nodeId,
		Heartbeat: heartbeat,
	}
	node.log.Debug("Sending command ", "GetActiveNodeByIdHeartbeatQuery", cmd)
	if err := bus.Dispatch(cmd); err != nil {

		errmsg := fmt.Sprintf("Failed to get node %v", cmd)
		node.log.Error(errmsg, "error", err)
		return nil, err
	}
	node.log.Debug("Command executed successfully")
	return cmd.Result, nil
}

func (node *ClusterNode) CheckInNodeProcessingMissingAlerts() error {
	if node == nil {
		return errors.New("Cluster node object is nil")
	}
	cmd := &m.SaveNodeProcessingMissingAlertCommand{}
	cmd.Node = &m.ActiveNode{NodeId: node.nodeId}
	node.log.Debug("Sending command ", "SaveNodeProcessingMissingAlertCommand:Node", cmd.Node)
	if err := bus.Dispatch(cmd); err != nil {

		errmsg := fmt.Sprintf("Failed to save node processing missing alert %v", cmd.Node)
		node.log.Error(errmsg, "error", err)
		return err
	}
	node.log.Debug("Command executed successfully")
	return nil
}

func (node *ClusterNode) GetActiveNodesCount(heartbeat int64) (int, error) {
	if node == nil {
		return 0, errors.New("Cluster node object is nil")
	}
	cmd := &m.GetActiveNodesCountCommand{
		NodeId:    node.nodeId,
		Heartbeat: heartbeat,
	}	
	node.log.Debug("Sending command ", "GetActiveNodesCountCommand:Node", cmd.NodeId)
	if err := bus.Dispatch(cmd); err != nil {
		errmsg := fmt.Sprintf("Failed to get active node count %v", cmd.NodeId)
		node.log.Error(errmsg, "error", err)
		return 0, err
	}
	node.log.Debug("GetActiveNodesCountCommand executed successfully")
	return cmd.Result, nil
}

func (node *ClusterNode) GetLastHeartbeat() (int64, error) {
	if node == nil {
		return 0, errors.New("Cluster node object is nil")
	}
	cmd := &m.GetLastHeartbeatCommand{}
	cmd.Node = &m.ActiveNode{NodeId: node.nodeId}
	node.log.Debug("Sending command ", "GetLastHeartbeatCommand:Node", cmd.Node)
	if err := bus.Dispatch(cmd); err != nil {
		errmsg := fmt.Sprintf("Failed to get last heartbeat %v", cmd.Node)
		node.log.Error(errmsg, "error", err)
		return 0, err
	}
	node.log.Debug("GetLastHeartbeatCommand executed successfully")
	return cmd.Result, nil
}
