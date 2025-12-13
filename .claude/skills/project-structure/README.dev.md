---
Note for Claude: This is developer documentation only. Read SKILL.md for execution instructions.
---

# Project Structure Skill - Design Decisions

## Why this skill

- Claude does not know what it does not know and has a bad 'gut feeling' for when it's lacking information
- Claude kept reimplementing existing helpers which were already there or in installed libraries
- Claude has suboptimal intuition for where to put stuff

## Why it's a skill

- I wanted it to be predictable with procedural instruction
- This way it can focus on generating comprehensive documentation
- No need for higher reasoning, it's just for Claude to navigate the repo

## Why own `PROJECT-STRUCTURE.md`

- Users use Claude for more than feature implementation, so no need to blow up context if it's not needed
- Users can instruct Claude to read PROJECT-STRUCTURE.md before planning/searching e.g. via CLAUDE.md

## Why comment everything (all files and directories)

- May seem verbose but could save tokens when Claude searches the repo

## Best usage

- Generate fresh documentation when needed (always regenerates from scratch)
- Read PROJECT-STRUCTURE.md before refactoring/implementing features/vague searching
- Great tool for a planner agent
