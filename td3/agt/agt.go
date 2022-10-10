package agt

import (
	"reflect"
)

type AgentID int
type Alternative int

type Agent struct {
	ID    AgentID
	Name  string
	Prefs []Alternative
}

type AgentI interface {
	Equal(ag AgentI) bool
	DeepEqual(ag AgentI) bool
	Clone() AgentI
	String() string
	Prefers(a Alternative, b Alternative)
	Start()
}

func (agent *Agent) Equal(ag AgentI) bool {
	agt, ok := ag.(Agent)
	if ok && agent.ID == agt.ID {
		return true
	}
	return false
}

func (agent *Agent) DeepEqual(ag AgentI) bool {
	agt, ok := ag.(Agent)
	if ok && agent.Name == agt.Name && agent.ID == agt.ID && reflect.DeepEqual(agent.Prefs, agt.Prefs) {
		return true
	}
	return false
}

func (agent Agent) Clone() AgentI {
	newPrefs := make([]Alternative, len(agent.Prefs))
	copy(newPrefs, agent.Prefs)
	return Agent{agent.ID, agent.Name, newPrefs}
}

func (agent *Agent) String() string {
	return "123"
}
func (agent *Agent) Prefers(a Alternative, b Alternative) {
	if !contains(agent.Prefs, a) {
		agent.Prefs = append(agent.Prefs, a)
	}

}
func (agent *Agent) Start()

func contains(s []Alternative, e Alternative) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}