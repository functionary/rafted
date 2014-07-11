package rafted

import hsm "github.com/hhkbp2/go-hsm"

const (
    EventTerm hsm.EventType = hsm.EventUser + 100 + iota
    EventRaftBegin
    EventAppendEntriesRequest
    EventAppendEntriesResponse
    EventRequestVoteRequest
    EventRequestVoteResponse
    EventPrepareInstallSnapshotRequest
    EventPrepareInstallSnapshotResponse
    EventInstallSnapshotRequest
    EventInstallSnapshotResponse
    EventRaftEnd
    EventTimeoutBegin
    EventTimeoutHeartBeat
    EventTimeoutElection
    EventTimeoutEnd
    EventLeaderRedirectResponse
    EventLeaderUnknownResponse
    EventClientUser = hsm.EventUser + 1000 + iota
)

func IsEventBetween(eventType, beginEvent, endEvent hsm.EventType) bool {
    if (eventType > beginEvent) && (eventType < endEvent) {
        return true
    }
    return false
}

func IsRaftEvent(eventType hsm.EventType) bool {
    return IsEventBetween(eventType, EventRaftBegin, EventRaftEnd)
}

func IsTimeoutEvent(eventType hsm.EventType) bool {
    return IsEventBetween(eventType, EventTimeoutBegin, EventTimeoutEnd)
}

func IsClientEvent(eventType hsm.EventType) bool {
    return (eventType >= EventClientUser)
}

type RaftEvent interface {
    hsm.Event
    Message() interface{}
}

type ResponsiveEvent interface {
    SendResponse(RaftEvent)
    RecvResponse() RaftEvent
}

type RequestEvent interface {
    RaftEvent
    ResponsiveEvent
}

type RequestEventHead struct {
    *hsm.StdEvent
    resultChan chan RaftEvent
}

func (self *RequestEventHead) SendResponse(event RaftEvent) {
    self.resultChan <- event
}

func (self *RequestEventHead) RecvResponse() RaftEvent {
    response := <-self.resultChan
    return response
}

func NewRequestEventHead(
    eventType hsm.EventType) *RequestEvent {
    return &RequestEventHead{
        hsm.NewStdEvent(eventType),
        make(chan RaftEvent, 1),
    }
}

type AppendEntriesReqeustEvent struct {
    *RequestEventHead
    request *AppendEntriesRequest
}

func NewAppendEntriesRequestEvent(
    request *AppendEntriesRequest) *AppendEntriesReqeustEvent {
    return &AppendEntriesReqeustEvent{
        NewRequestEventHead(EventAppendEntriesRequest),
        request,
    }
}
func (self *AppendEntriesReqeustEvent) Message() interface{} {
    return self.request
}

type AppendEntriesResponseEvent struct {
    *hsm.StdEvent
    response *AppendEntriesResponse
}

func NewAppendEntriesResponseEvent(
    response *AppendEntriesResponse) *AppendEntriesResponseEvent {
    return &AppendEntriesResponseEvent{
        hsm.NewStdEvent(EventAppendEntriesResponse),
        response,
    }
}

func (self *AppendEntriesResponseEvent) Message() interface{} {
    return self.response
}

type RequestVoteRequestEvent struct {
    *RequestEventHead
    request *RequestVoteRequest
}

func NewRequestVoteRequestEvent(
    request *RequestVoteRequest) *RequestVoteRequestEvent {
    return &RequestVoteRequestEvent{
        NewRequestEventHead(EventRequestVoteRequest),
        request,
    }
}

func (self *RequestVoteRequestEvent) Message() interface{} {
    return self.request
}

type RequestVoteResponseEvent struct {
    *hsm.StdEvent
    response *RequestVoteResponse
}

func NewRequestVoteResponseEvent(
    response *RequestVoteResponse) *RequestVoteResponseEvent {
    return &RequestVoteResponseEvent{
        hsm.NewStdEvent(EventRequestVoteResponse),
        response,
    }
}

func (self *RequestVoteRequestEvent) Message() interface{} {
    return self.response
}

type PrepareInstallSnapshotRequestEvent struct {
    *RequestEventHead
    request *PrepareInstallSnapshotRequest
}

func NewPrepareInstallSnapshotRequestEvent(
    request *PrepareInstallSnapshotRequest) *PrepareInstallSnapshotRequestEvent {
    return &PrepareInstallSnapshotRequestEvent{
        NewRequestEventHead(EventPrepareInstallSnapshotRequest),
        request,
    }
}

func (self *PrepareInstallSnapshotRequest) Message() interface{} {
    return self.request
}

type PrepareInstallSnapshotResponseEvent struct {
    *hsm.StdEvent
    response *PrepareInstallSnapshotResponse
}

func NewPrepareInstallSnapshotResponseEvent(
    response *PrepareInstallSnapshotResponse) *PrepareInstallSnapshotResponseEvent {
    return &PrepareInstallSnapshotResponseEvent{
        hsm.NewStdEvent(EventPrepareInstallSnapshotResponse),
        response,
    }
}

func (self *PrepareInstallSnapshotResponseEvent) Message() interface{} {
    return self.response
}

type InstallSnapshotRequestEvent struct {
    *RequestEventHead
    request *InstallSnapshotRequest
}

func NewInstallSnapshotRequestEvent(
    request *InstallSnapshotRequest) *InstallSnapshotRequestEvent {
    return &InstallSnapshotRequestEvent{
        NewRequestEventHead(EventInstallSnapshotRequest),
        request,
    }
}

func (self *InstallSnapshotRequestEvent) Message() interface{} {
    return self.request
}

type InstallSnapshotResponseEvent struct {
    *hsm.StdEvent
    response *PrepareInstallSnapshotResponse
}

func NewInstallSnapshotResponseEvent(
    response *InstallSnapshotResponse) *InstallSnapshotRequestEvent {
    return &InstallSnapshotResponseEvent{
        hsm.NewStdEvent(EventInstallSnapshotResponse),
        response,
    }
}

func (self *InstallSnapshotResponseEvent) Message() interface{} {
    return self.response
}

type HeartbeatTimeoutEvent struct {
    *hsm.StdEvent
}

func NewHeartbeatTimeoutEvent() *HeartbeatTimeoutEvent {
    return &HeartbeatTimeoutEvent{hsm.NewStdEvent(EventTimeoutHeartBeat)}
}

type ElectionTimeoutEvent struct {
    *hsm.StdEvent
}

func NewElectionTimeoutEvent() *ElectionTimeoutEvent {
    return &ElectionTimeoutEvent{hsm.NewStdEvent(EventTimeoutElection)}
}

type LeaderRedirectResponseEvent struct {
    *hsm.StdEvent
    response *LeaderRedirectResponse
}

func NewLeaderRedirectResponseEvent(
    response *LeaderRedirectResponse) *LeaderRedirectResponseEvent {
    return &LeaderRedirectResponseEvent{
        hsm.NewStdEvent(EventRedirectResponse),
        response,
    }
}

func (self *LeaderRedirectResponseEvent) Message() interface{} {
    return self.response
}

type LeaderUnknownResponseEvent struct {
    *hsm.StdEvent
    response *LeaderUnknownResponse
}

func NewLeaderUnknownResponseEvent(
    response *LeaderUnknownResponse) *LeaderUnknownResponseEvent {
    return &LeaderUnknownResponseEvent{
        hsm.NewStdEvent(EventLeaderUnknownResponse),
        response,
    }
}

func (self *LeaderUnknownResponseEvent) Message() interface{} {
    return self.response
}
