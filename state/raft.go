package state

import hsm "github.com/hhkbp2/go-hsm"

type RaftState struct {
    *hsm.StateHead
}

func NewRaftState(super hsm.State) *RaftState {
    object := &RaftState{
        hsm.NewStateHead(super),
    }
    super.AddChild(object)
    return object
}

func (*RaftState) ID() string {
    return RaftStateID
}

func (self *RaftState) Entry(sm hsm.HSM, event hsm.Event) (state hsm.State) {
    // ignore events
    return nil
}

func (self *RaftState) Init(sm hsm.HSM, event hsm.Event) (state hsm.State) {
    sm.QInit(FollowerID)
    return nil
}

func (self *RaftState) Exit(sm hsm.HSM, event hsm.Event) (state hsm.State) {
    // ignore events
    return nil
}

func (self *RaftState) Handle(sm hsm.HSM, event hsm.Event) (state hsm.State) {
    // TODO add event handling if needed
    return nil
}