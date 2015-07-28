// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package status

import (
	"encoding/json"
	"fmt"

	"github.com/juju/juju/api"
	"github.com/juju/juju/apiserver/params"
	"github.com/juju/juju/instance"
	"github.com/juju/juju/network"
	"github.com/juju/juju/state/multiwatcher"
)

type formattedStatus struct {
	Environment string                   `json:"environment"`
	Machines    map[string]machineStatus `json:"machines"`
	Services    map[string]serviceStatus `json:"services"`
	Networks    map[string]networkStatus `json:"networks,omitempty" yaml:",omitempty"`
}

type errorStatus struct {
	StatusError string `json:"status-error" yaml:"status-error"`
}

type machineStatus struct {
	Err            error                    `json:"-" yaml:",omitempty"`
	AgentState     params.Status            `json:"agent-state,omitempty" yaml:"agent-state,omitempty"`
	AgentStateInfo string                   `json:"agent-state-info,omitempty" yaml:"agent-state-info,omitempty"`
	AgentVersion   string                   `json:"agent-version,omitempty" yaml:"agent-version,omitempty"`
	DNSName        string                   `json:"dns-name,omitempty" yaml:"dns-name,omitempty"`
	InstanceId     instance.Id              `json:"instance-id,omitempty" yaml:"instance-id,omitempty"`
	InstanceState  string                   `json:"instance-state,omitempty" yaml:"instance-state,omitempty"`
	Life           string                   `json:"life,omitempty" yaml:"life,omitempty"`
	Series         string                   `json:"series,omitempty" yaml:"series,omitempty"`
	Id             string                   `json:"-" yaml:"-"`
	Containers     map[string]machineStatus `json:"containers,omitempty" yaml:"containers,omitempty"`
	Hardware       string                   `json:"hardware,omitempty" yaml:"hardware,omitempty"`
	HAStatus       string                   `json:"state-server-member-status,omitempty" yaml:"state-server-member-status,omitempty"`
}

// A goyaml bug means we can't declare these types
// locally to the GetYAML methods.
type machineStatusNoMarshal machineStatus

func (s machineStatus) MarshalJSON() ([]byte, error) {
	if s.Err != nil {
		return json.Marshal(errorStatus{s.Err.Error()})
	}
	return json.Marshal(machineStatusNoMarshal(s))
}

func (s machineStatus) GetYAML() (tag string, value interface{}) {
	if s.Err != nil {
		return "", errorStatus{s.Err.Error()}
	}
	// TODO(rog) rename mNoMethods to noMethods (and also in
	// the other GetYAML methods) when people are using the non-buggy
	// goyaml version. // TODO(jw4) however verify that gccgo does not
	// complain about symbol already defined.
	type mNoMethods machineStatus
	return "", mNoMethods(s)
}

type serviceStatus struct {
	Err           error                 `json:"-" yaml:",omitempty"`
	Charm         string                `json:"charm" yaml:"charm"`
	CanUpgradeTo  string                `json:"can-upgrade-to,omitempty" yaml:"can-upgrade-to,omitempty"`
	Exposed       bool                  `json:"exposed" yaml:"exposed"`
	Life          string                `json:"life,omitempty" yaml:"life,omitempty"`
	StatusInfo    statusInfoContents    `json:"service-status,omitempty" yaml:"service-status,omitempty"`
	Relations     map[string][]string   `json:"relations,omitempty" yaml:"relations,omitempty"`
	Networks      map[string][]string   `json:"networks,omitempty" yaml:"networks,omitempty"`
	SubordinateTo []string              `json:"subordinate-to,omitempty" yaml:"subordinate-to,omitempty"`
	Units         map[string]unitStatus `json:"units,omitempty" yaml:"units,omitempty"`
}

type serviceStatusNoMarshal serviceStatus

func (s serviceStatus) MarshalJSON() ([]byte, error) {
	if s.Err != nil {
		return json.Marshal(errorStatus{s.Err.Error()})
	}
	type ssNoMethods serviceStatus
	return json.Marshal(ssNoMethods(s))
}

func (s serviceStatus) GetYAML() (tag string, value interface{}) {
	if s.Err != nil {
		return "", errorStatus{s.Err.Error()}
	}
	type ssNoMethods serviceStatus
	return "", ssNoMethods(s)
}

type meterStatus struct {
	Color   string `json:"color,omitempty" yaml:"color,omitempty"`
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}

type unitStatus struct {
	// New Juju Health Status fields.
	WorkloadStatusInfo statusInfoContents `json:"workload-status,omitempty" yaml:"workload-status,omitempty"`
	AgentStatusInfo    statusInfoContents `json:"agent-status,omitempty" yaml:"agent-status,omitempty"`
	MeterStatus        *meterStatus       `json:"meter-status,omitempty" yaml:"meter-status,omitempty"`

	// Legacy status fields, to be removed in Juju 2.0
	AgentState     params.Status `json:"agent-state,omitempty" yaml:"agent-state,omitempty"`
	AgentStateInfo string        `json:"agent-state-info,omitempty" yaml:"agent-state-info,omitempty"`
	Err            error         `json:"-" yaml:",omitempty"`
	AgentVersion   string        `json:"agent-version,omitempty" yaml:"agent-version,omitempty"`
	Life           string        `json:"life,omitempty" yaml:"life,omitempty"`

	Charm         string                `json:"upgrading-from,omitempty" yaml:"upgrading-from,omitempty"`
	Machine       string                `json:"machine,omitempty" yaml:"machine,omitempty"`
	OpenedPorts   []string              `json:"open-ports,omitempty" yaml:"open-ports,omitempty"`
	PublicAddress string                `json:"public-address,omitempty" yaml:"public-address,omitempty"`
	Subordinates  map[string]unitStatus `json:"subordinates,omitempty" yaml:"subordinates,omitempty"`
}

type statusInfoContents struct {
	Err     error         `json:"-" yaml:",omitempty"`
	Current params.Status `json:"current,omitempty" yaml:"current,omitempty"`
	Message string        `json:"message,omitempty" yaml:"message,omitempty"`
	Since   string        `json:"since,omitempty" yaml:"since,omitempty"`
	Version string        `json:"version,omitempty" yaml:"version,omitempty"`
}

type statusInfoContentsNoMarshal statusInfoContents

func (s statusInfoContents) MarshalJSON() ([]byte, error) {
	if s.Err != nil {
		return json.Marshal(errorStatus{s.Err.Error()})
	}
	return json.Marshal(statusInfoContentsNoMarshal(s))
}

func (s statusInfoContents) GetYAML() (tag string, value interface{}) {
	if s.Err != nil {
		return "", errorStatus{s.Err.Error()}
	}
	type sicNoMethods statusInfoContents
	return "", sicNoMethods(s)
}

type unitStatusNoMarshal unitStatus

func (s unitStatus) MarshalJSON() ([]byte, error) {
	if s.Err != nil {
		return json.Marshal(errorStatus{s.Err.Error()})
	}
	return json.Marshal(unitStatusNoMarshal(s))
}

func (s unitStatus) GetYAML() (tag string, value interface{}) {
	if s.Err != nil {
		return "", errorStatus{s.Err.Error()}
	}
	type usNoMethods unitStatus
	return "", usNoMethods(s)
}

type networkStatus struct {
	Err        error      `json:"-" yaml:",omitempty"`
	ProviderId network.Id `json:"provider-id" yaml:"provider-id"`
	CIDR       string     `json:"cidr,omitempty" yaml:"cidr,omitempty"`
	VLANTag    int        `json:"vlan-tag,omitempty" yaml:"vlan-tag,omitempty"`
}

type networkStatusNoMarshal networkStatus

func (n networkStatus) MarshalJSON() ([]byte, error) {
	if n.Err != nil {
		return json.Marshal(errorStatus{n.Err.Error()})
	}
	type nNoMethods networkStatus
	return json.Marshal(nNoMethods(n))
}

func (n networkStatus) GetYAML() (tag string, value interface{}) {
	if n.Err != nil {
		return "", errorStatus{n.Err.Error()}
	}
	type nNoMethods networkStatus
	return "", nNoMethods(n)
}

type statusFormatter struct {
	status        *api.Status
	relations     map[int]api.RelationStatus
	isoTime       bool
	compatVersion int
}

func newStatusFormatter(status *api.Status, compatVersion int, isoTime bool) *statusFormatter {
	sf := statusFormatter{
		status:        status,
		relations:     make(map[int]api.RelationStatus),
		compatVersion: compatVersion,
		isoTime:       isoTime,
	}
	for _, relation := range status.Relations {
		sf.relations[relation.Id] = relation
	}
	return &sf
}

func (sf *statusFormatter) format() formattedStatus {
	if sf.status == nil {
		return formattedStatus{}
	}
	out := formattedStatus{
		Environment: sf.status.EnvironmentName,
		Machines:    make(map[string]machineStatus),
		Services:    make(map[string]serviceStatus),
	}
	for k, m := range sf.status.Machines {
		out.Machines[k] = sf.formatMachine(m)
	}
	for sn, s := range sf.status.Services {
		out.Services[sn] = sf.formatService(sn, s)
	}
	for k, n := range sf.status.Networks {
		if out.Networks == nil {
			out.Networks = make(map[string]networkStatus)
		}
		out.Networks[k] = sf.formatNetwork(n)
	}
	return out
}

func (sf *statusFormatter) formatMachine(machine api.MachineStatus) machineStatus {
	var out machineStatus

	if machine.Agent.Status == "" {
		// Older server
		// TODO: this will go away at some point (v1.21?).
		out = machineStatus{
			AgentState:     machine.AgentState,
			AgentStateInfo: machine.AgentStateInfo,
			AgentVersion:   machine.AgentVersion,
			Life:           machine.Life,
			Err:            machine.Err,
			DNSName:        machine.DNSName,
			InstanceId:     machine.InstanceId,
			InstanceState:  machine.InstanceState,
			Series:         machine.Series,
			Id:             machine.Id,
			Containers:     make(map[string]machineStatus),
			Hardware:       machine.Hardware,
		}
	} else {
		// New server
		agent := machine.Agent
		out = machineStatus{
			AgentState:     machine.AgentState,
			AgentStateInfo: adjustInfoIfMachineAgentDown(machine.AgentState, agent.Status, agent.Info),
			AgentVersion:   agent.Version,
			Life:           agent.Life,
			Err:            agent.Err,
			DNSName:        machine.DNSName,
			InstanceId:     machine.InstanceId,
			InstanceState:  machine.InstanceState,
			Series:         machine.Series,
			Id:             machine.Id,
			Containers:     make(map[string]machineStatus),
			Hardware:       machine.Hardware,
		}
	}

	for k, m := range machine.Containers {
		out.Containers[k] = sf.formatMachine(m)
	}

	for _, job := range machine.Jobs {
		if job == multiwatcher.JobManageEnviron {
			out.HAStatus = makeHAStatus(machine.HasVote, machine.WantsVote)
			break
		}
	}
	return out
}

func (sf *statusFormatter) formatService(name string, service api.ServiceStatus) serviceStatus {
	out := serviceStatus{
		Err:           service.Err,
		Charm:         service.Charm,
		Exposed:       service.Exposed,
		Life:          service.Life,
		Relations:     service.Relations,
		Networks:      make(map[string][]string),
		CanUpgradeTo:  service.CanUpgradeTo,
		SubordinateTo: service.SubordinateTo,
		Units:         make(map[string]unitStatus),
		StatusInfo:    sf.getServiceStatusInfo(service),
	}
	if len(service.Networks.Enabled) > 0 {
		out.Networks["enabled"] = service.Networks.Enabled
	}
	if len(service.Networks.Disabled) > 0 {
		out.Networks["disabled"] = service.Networks.Disabled
	}
	for k, m := range service.Units {
		out.Units[k] = sf.formatUnit(unitFormatInfo{
			unit:          m,
			unitName:      k,
			serviceName:   name,
			meterStatuses: service.MeterStatuses,
		})
	}
	return out
}

func (sf *statusFormatter) getServiceStatusInfo(service api.ServiceStatus) statusInfoContents {
	info := statusInfoContents{
		Err:     service.Status.Err,
		Current: service.Status.Status,
		Message: service.Status.Info,
		Version: service.Status.Version,
	}
	if service.Status.Since != nil {
		info.Since = formatStatusTime(service.Status.Since, sf.isoTime)
	}
	return info
}

type unitFormatInfo struct {
	unit          api.UnitStatus
	unitName      string
	serviceName   string
	meterStatuses map[string]api.MeterStatus
}

func (sf *statusFormatter) formatUnit(info unitFormatInfo) unitStatus {
	// TODO(Wallyworld) - this should be server side but we still need to support older servers.
	sf.updateUnitStatusInfo(&info.unit, info.serviceName)

	out := unitStatus{
		WorkloadStatusInfo: sf.getWorkloadStatusInfo(info.unit),
		AgentStatusInfo:    sf.getAgentStatusInfo(info.unit),
		Machine:            info.unit.Machine,
		OpenedPorts:        info.unit.OpenedPorts,
		PublicAddress:      info.unit.PublicAddress,
		Charm:              info.unit.Charm,
		Subordinates:       make(map[string]unitStatus),
	}

	if ms, ok := info.meterStatuses[info.unitName]; ok {
		out.MeterStatus = &meterStatus{
			Color:   ms.Color,
			Message: ms.Message,
		}
	}

	// These legacy fields will be dropped for Juju 2.0.
	if sf.compatVersion < 2 || out.AgentStatusInfo.Current == "" {
		out.Err = info.unit.Err
		out.AgentState = info.unit.AgentState
		out.AgentStateInfo = info.unit.AgentStateInfo
		out.Life = info.unit.Life
		out.AgentVersion = info.unit.AgentVersion
	}

	for k, m := range info.unit.Subordinates {
		out.Subordinates[k] = sf.formatUnit(unitFormatInfo{
			unit:          m,
			unitName:      k,
			serviceName:   info.serviceName,
			meterStatuses: info.meterStatuses,
		})
	}
	return out
}

func (sf *statusFormatter) getWorkloadStatusInfo(unit api.UnitStatus) statusInfoContents {
	info := statusInfoContents{
		Err:     unit.Workload.Err,
		Current: unit.Workload.Status,
		Message: unit.Workload.Info,
		Version: unit.Workload.Version,
	}
	if unit.Workload.Since != nil {
		info.Since = formatStatusTime(unit.Workload.Since, sf.isoTime)
	}
	return info
}

func (sf *statusFormatter) getAgentStatusInfo(unit api.UnitStatus) statusInfoContents {
	info := statusInfoContents{
		Err:     unit.UnitAgent.Err,
		Current: unit.UnitAgent.Status,
		Message: unit.UnitAgent.Info,
		Version: unit.UnitAgent.Version,
	}
	if unit.UnitAgent.Since != nil {
		info.Since = formatStatusTime(unit.UnitAgent.Since, sf.isoTime)
	}
	return info
}

func (sf *statusFormatter) updateUnitStatusInfo(unit *api.UnitStatus, serviceName string) {
	// This logic has no business here but can't be moved until Juju 2.0.
	statusInfo := unit.Workload.Info
	if unit.Workload.Status == "" {
		// Old server that doesn't support this field and others.
		// Just use the info string as-is.
		statusInfo = unit.AgentStateInfo
	}
	if unit.Workload.Status == params.StatusError {
		if relation, ok := sf.relations[getRelationIdFromData(unit)]; ok {
			// Append the details of the other endpoint on to the status info string.
			if ep, ok := findOtherEndpoint(relation.Endpoints, serviceName); ok {
				unit.Workload.Info = statusInfo + " for " + ep.String()
				unit.AgentStateInfo = unit.Workload.Info
			}
		}
	}
}

func (sf *statusFormatter) formatNetwork(network api.NetworkStatus) networkStatus {
	return networkStatus{
		Err:        network.Err,
		ProviderId: network.ProviderId,
		CIDR:       network.CIDR,
		VLANTag:    network.VLANTag,
	}
}

func makeHAStatus(hasVote, wantsVote bool) string {
	var s string
	switch {
	case hasVote && wantsVote:
		s = "has-vote"
	case hasVote && !wantsVote:
		s = "removing-vote"
	case !hasVote && wantsVote:
		s = "adding-vote"
	case !hasVote && !wantsVote:
		s = "no-vote"
	}
	return s
}

func getRelationIdFromData(unit *api.UnitStatus) int {
	if relationId_, ok := unit.Workload.Data["relation-id"]; ok {
		if relationId, ok := relationId_.(float64); ok {
			return int(relationId)
		} else {
			logger.Infof("relation-id found status data but was unexpected "+
				"type: %q. Status output may be lacking some detail.", relationId_)
		}
	}
	return -1
}

// findOtherEndpoint searches the provided endpoints for an endpoint
// that *doesn't* match serviceName. The returned bool indicates if
// such an endpoint was found.
func findOtherEndpoint(endpoints []api.EndpointStatus, serviceName string) (api.EndpointStatus, bool) {
	for _, endpoint := range endpoints {
		if endpoint.ServiceName != serviceName {
			return endpoint, true
		}
	}
	return api.EndpointStatus{}, false
}

// adjustInfoIfMachineAgentDown modifies the agent status info string if the
// agent is down. The original status and info is included in
// parentheses.
func adjustInfoIfMachineAgentDown(status, origStatus params.Status, info string) string {
	if status == params.StatusDown {
		if info == "" {
			return fmt.Sprintf("(%s)", origStatus)
		}
		return fmt.Sprintf("(%s: %s)", origStatus, info)
	}
	return info
}
