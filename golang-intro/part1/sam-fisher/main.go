package main

import (
	"fmt"
	"strings"
)

const (
	SkillStealth = iota
	SkillHandToHandCombat
)

type Skill int

func (s Skill) String() string {
	switch s {
	case SkillStealth:
		return "stealth"
	case SkillHandToHandCombat:
		return "hand to hand combat"
	default:
		return "unknown skill"
	}
}

type Player struct {
	name   string
	skills []Skill
	energy float32
}

func (p *Player) Describe() string {
	var skillsStr []string
	for _, skill := range p.skills {
		skillsStr = append(skillsStr, skill.String())
	}
	return fmt.Sprintf("My name is %s, and my skills are %s", p.name, strings.Join(skillsStr, ", "))
}

func NewPlayer(name string) *Player {
	return &Player{name: name, energy: 100, skills: []Skill{SkillStealth}}
}

func (p *Player) GainSkill(skill Skill) {
	p.skills = append(p.skills, skill)
}

func main() {
	p := NewPlayer("Sam Fisher")
	p.GainSkill(SkillHandToHandCombat)
	fmt.Print(p.Describe())
}
