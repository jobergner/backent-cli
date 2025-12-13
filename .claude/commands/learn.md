---
description: Capture knowledge from real corrections to build lasting principles
argument-hint: [optional: correction or lesson learned]
tags: [learning, principles, feedback, knowledge-capture]
---

# Task:

Extract lasting principles from corrections and lessons learned, then package them for future use.

**Usage modes:**
- `/learn` - Automatically find corrections from this session
- `/learn [context]` - Manually provide a correction or lesson (e.g., `/learn always use Read instead of cat for files`)

## Steps

1. **Identify or Gather Corrections**

   **If `$ARGUMENTS` is provided:**
   - Use `$ARGUMENTS` as the initial correction/lesson context (freeform - could be correction text, description, or notes)
   - Store this as the first item to process
   - Also scan conversation history (see below) to find additional learnings

   **Always scan conversation history:**
   - Review the **last 10 user messages** for potential learnings, including:
     - User corrections/feedback (e.g., "actually do it this way instead", "no, use X not Y")
     - Design decisions with explanations (e.g., "we use X because...", "the reason is...")
     - Rejected approaches with reasoning (e.g., "don't use Y because...", "that won't work because...")
     - Explicit rejections (e.g., "No, and tell Claude what to do differently")
   - Look for messages where the user is teaching, correcting, or explaining their preferences
   - Collect all potential learnings found

   **If no learnings found:**
   - Use AskUserQuestion to prompt: "No corrections or learnings found in recent conversation. Would you like to describe one now?"
   - Options:
     1. "Yes, let me describe it" - proceed to gather freeform input
     2. "No, exit" - exit the command
   - If user chooses to describe, gather the correction context and proceed to Step 2
   - If user chooses to exit, exit gracefully

   **If multiple learnings found:**
   - Process them one by one in sequence
   - For each learning, give the user the option to skip it before proceeding to Step 2
   - Use AskUserQuestion with options:
     1. "Capture this as a principle" - proceed to Step 2 for this learning
     2. "Skip this one" - move to the next learning
     3. "Skip all remaining" - exit the command
   - After processing all learnings, notify user of completion

2. **Ask WHY with Context**

   - **Always ask WHY** - even if reasoning seems obvious or was provided in `$ARGUMENTS`
   - Use AskUserQuestion to understand the deeper reasoning behind the correction
   - Include 1 option drawn from the correction context
   - Add 1-2 thoughtful alternate interpretations
   - Make the user articulate the real reason, not just confirm your assumptions
   - Include option to skip this correction if they change their mind

3. **Categorize the Principle**

   - Scan `.claude/principles/` for existing category files (e.g., `general-coding.md`, `unit-tests.md`, `css-styling.md`)
   - Use AskUserQuestion to ask which category file this principle belongs in, options are the existing ones, plus the "Type something" option to create a new category file
   - If user chooses "Type something":
     - They specify a new category name (e.g., `documentation.md`, `api-design.md`)
     - Use AskUserQuestion to define keywords for when these principles should be loaded
     - Auto-suggest comma-separated keywords based on the category name, but let the user edit them
     - These keywords will be added to `.claude/principles/principles-index.md` so Claude can recognize when to load these principles
     - Examples: "test, testing, jest, vitest", "documentation, docs, readme", "typescript, ts, type"

4. **Check for Similar Principles**

   - Read the selected category file to check for similar existing principles
   - If a similar principle exists, show the user both the existing and new principle
   - Use AskUserQuestion to give the user these options:
     1. **Merge** - Combine the existing and new principles into one comprehensive principle
     2. **Overwrite** - Replace the existing principle with the new one
     3. **Add anyway** - Keep both principles as separate entries
     4. **Abort** - Don't add this new principle (keep the existing one as-is)
   - If no similar principle exists, proceed to write the new principle

5. **Write the Principle**

   - Add the principle to the appropriate category file in `.claude/principles/`
   - Create the directory structure if it doesn't exist yet
   - Write principles that are **precise, concise, and LLM-friendly**: use clear section labels, actionable directives, and scannable format
   - Follow this template for every principle:

   ```markdown
   ## [Title]

   **Context:** [Problem this solves and user's reasoning]

   **Rule:** [Clear, actionable directive]

   **Good:**
   ```

   [Correct approach - actual code/text]

   ```

   **Bad:**
   ```

   [What to avoid - what went wrong]

   ```

   **Apply when:** [Specific situations]

   ---
   ```

   - Append to existing file or create new category file as needed
   - Use horizontal rules `---` to separate principles
   - If a new category file was created, update `.claude/principles/principles-index.md`:
     - Read the principles-index.md file and find the section "Principle Categories and their Trigger-Keywords"
     - create or add to the list of all category files and their keywords under "Principle Categories and their Trigger-Keywords"
     - Format the entry as: "- Read **[filename.md]** when you detect the keywords [keyword1, keyword2, keyword3] - either in what the user says or in your own thinking while analyzing the task"
     - Example: "- Read **unit-tests.md** when you detect the keywords test, testing, jest, vitest, spec - either in what the user says or in your own thinking while analyzing the task"
     - This ensures principles automatically activate when Claude detects these keywords during conversation or internal analysis

## Examples

**Automatic mode (finding corrections in chat history):**
```
Earlier in the session, user provided feedback like "actually use X not Y" and explained a design decision
User: /learn
Claude: Found 3 potential learnings in recent conversation:
1. Correction about using X instead of Y
2. Design decision explanation for Z
3. Explanation of why approach A won't work

Let's start with the first one. Would you like to capture this as a principle, skip it, or skip all remaining?
```

**Manual mode with argument:**
```
User: /learn always use the Read tool instead of cat for reading files
Claude: I'll capture this principle. I also found 1 additional learning in recent conversation.

First, let me understand the reasoning behind "always use Read tool instead of cat"...
[Proceeds to ask WHY even though context was provided]
```

**Manual mode with additional history learnings:**
```
User: /learn it should be done like X and not like Y
Claude: I'll capture this principle. I also scanned recent conversation and found 2 additional potential learnings.

Starting with your provided learning: "it should be done like X and not like Y"
Let me understand the reasoning...
```

**No learnings found:**
```
User: /learn
Claude: No corrections or learnings found in recent conversation. Would you like to describe one now?
User: Yes, let me describe it
Claude: Please describe the correction or lesson learned...
```

## Critical Rules

- **NO assumptions** - Always ask for context, never guess the reasoning
- **Use their words** - Include the actual correction text in question options
- **Focus on patterns** - Extract principles that apply beyond this specific instance
- **Make it actionable** - Rules should be concrete enough to follow

