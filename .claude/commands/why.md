---
description: "Suggest 'why' comments for git changes or specified file"
---

# Task: Suggest "Why" Comments

## Target Code

If no file is specified below, analyze the git diff. If a file is specified, analyze the entire file.

**Specified file/arguments:** $ARGUMENTS

**Current git changes:**
!git diff HEAD 2>/dev/null || echo "(No git changes found)"

## Instructions

**Analysis Mode:**

- If a file was specified in arguments above, read and analyze that entire file using the Read tool
- If no file was specified (arguments empty), analyze the git diff shown above
- Do NOT analyze both - choose based on whether arguments were provided

Once you have the appropriate code, suggest "why" comments where appropriate. Format your suggestions as code comments that the developer should add to explain their reasoning.

**Core Principle: Comment the WHY, not the WHAT.**
Add a comment only when a future reader couldn't reliably infer intent, constraints, or trade-offs.

---

## WHERE to Add Comments (Criteria)

### 1. Domain rules & business logic

- Policies, regulatory quirks, accounting/tax rules
- Product decisions, stakeholder requirements that motivated the code
- Example: "Why must refunds be processed within 24 hours?"

### 2. Invariants & contracts

- Preconditions, postconditions, non-obvious guarantees
- Idempotency, ordering requirements
- "Must stay in sync with X"
- Example: "Why must items be processed in chronological order?"

### 3. Assumptions & limits

- Data ranges, units, time zones
- Performance limits, "we only support ≤ N items"
- Fallbacks for partial data
- Example: "Why do we cap at 1000 items?"

### 4. Security & privacy rationale

- Threat model, trust boundaries
- Why a check/encoding is needed
- Why something isn't logged or must be PII-safe
- Example: "Why must we sanitize this input before logging?"

### 5. Performance choices

- Hot paths, complexity notes
- Micro-optimizations, caching strategies
- Why we avoided a simpler but slower approach
- Example: "Why use a Map instead of array.find()?"

### 6. Concurrency & lifecycle constraints

- Locks, atomicity, races avoided
- Event ordering, cleanup requirements
- Async gotchas
- Example: "Why must we acquire the lock before checking the flag?"

### 7. Workarounds & tech debt

- HACK/WORKAROUND with link to issue/PR
- Library bugs, environment quirks
- Migration phases, feature flags and removal plan
- Example: "Why this workaround? [Link to issue]"

### 8. Public API behavior

- Anything exported/consumed across boundaries
- Inputs/outputs, error semantics
- Versioning/compatibility guarantees
- Example: "Why do we return null instead of throwing?"

### 9. Non-obvious algorithms or math

- Proof sketch, reference, derivation
- Chosen heuristic and why it beats alternatives
- Example: "Why exponential backoff with jitter?"

### 10. Cross-module coupling

- "This relies on X selecting Y"
- Data shape contracts between services
- Schema evolution notes
- Example: "Why does this depend on the order of middleware registration?"

### 11. Generated/derived code

- Point back to the source (schema, tool, command)
- How to regenerate safely
- Example: "Why is this generated from schema.graphql?"

---

## WHERE NOT to Comment (Skip These)

### ❌ Narrating the obvious WHAT

- "// loop over items", "// increment i", "// call API"
- The code already tells us this

### ❌ Restating names

- "// the activeUsers map of active users"
- If a good identifier says it, don't echo it

### ❌ Explaining language/library defaults

- "// map returns a new array"
- Assume baseline competence

### ❌ Outdated, speculative, or joking comments

- Noise that ages poorly or contradicts code

---

## Interactive Workflow

This is a **one-step batch process**. Follow these steps:

### Step 1: Initial Analysis

- Analyze the code and identify ALL locations that would benefit from "why" comments
- Create an internal list of these locations with numbers
- Be selective—only include locations where the "why" is genuinely non-obvious

### Step 2: Present ALL Locations at Once

Display ALL identified locations in a single message:

1. Number each location (1, 2, 3, etc.)
2. For each location, show:
   - File path and line number range
   - Code with **context** (~5 lines before and ~5 lines after)
   - The "WHY:" question **directly in the code** as a comment
   - Which criterion applies

3. After showing all locations, prompt the user to select which ones to apply

Example format:

```
Found 3 locations that could benefit from "why" comments:

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

📍 1. src/auth.ts:40-50

Code:
  function handleLogin(username, password) {
    const user = findUser(username);

    if (!user) {
      return null;
    }

    const attempts = getLoginAttempts(user.id);
    // WHY: Why do we lock after 3 attempts specifically❓
    if (attempts > 3) {
      lockAccount(userId);
    }

    return validatePassword(user, password);
  }

📋 Criterion: Security & privacy rationale

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

📍 2. src/cache.ts:120-128

Code:
  function getFromCache(key) {
    const cached = cache.get(key);

    // WHY: Why do we use 5 minutes as the TTL❓
    if (cached && Date.now() - cached.timestamp < 5 * 60 * 1000) {
      return cached.value;
    }

    return null;
  }

📋 Criterion: Performance choices / Assumptions & limits

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

📍 3. src/api.ts:78-85

Code:
  async function fetchData(endpoint) {
    // WHY: Why do we retry 3 times❓
    for (let i = 0; i < 3; i++) {
      try {
        return await fetch(endpoint);
      } catch (err) {
        if (i === 2) throw err;
      }
    }
  }

📋 Criterion: Domain rules & business logic

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Provide your answers below using this format:
1. <your answer for question 1>
3. <your answer for question 3>

Questions not answered will be skipped. Example:
1. Prevents brute-force attacks per security policy SEC-2024-01
3. Retries handle transient network failures per SRE-2024

(If you don't want to add any comments, just say "none" or "skip")
```

**Important formatting guidelines:**

- Use the Read tool to get the full file context (not just git diff lines)
- Show approximately 5 lines before and 5 lines after (adjust if needed for clarity)
- Place the "WHY:" question exactly where the actual comment should go in the code
- **ALWAYS end the why question with a question mark emoji ❓** for visibility
- Use proper comment syntax for the language (// for JS/TS, # for Python, etc.)
- Match the indentation level of the code being commented
- Show line numbers in the file path (e.g., "src/auth.ts:40-50")
- Use separator lines (━━━) between locations for clarity

### Step 3: Parse User Response

The user will provide answers in this format:
```
1. <answer for question 1>
2. <answer for question 2>
4. <answer for question 4>
```

**Parse the response:**
1. Look for lines that start with a number followed by a period
2. Extract the number and the answer text after the period
3. For each numbered answer provided:
   - Locate the corresponding location from your internal list
   - Apply the user's answer as a comment to the file
4. Skip any locations not mentioned by the user

**Special cases:**
- If user writes "none", "skip", or provides no numbered answers → skip all comments
- The user can answer any subset of questions (e.g., just 1 and 3, skipping 2)

### Step 4: Apply All Comments

For each numbered answer the user provided:

1. Use the Edit tool to add the comment to the correct file and line
2. Insert the comment at the exact location where you showed the "WHY:" placeholder
3. **Transform the user's answer into a complete, standalone comment:**
   - Combine the WHY question context with the user's answer
   - Create a grammatically complete sentence
   - Ensure proper capitalization and punctuation
   - Make it readable without needing to see the original question
   - Preserve the user's core reasoning and terminology
4. Match the indentation of the surrounding code
5. Use appropriate comment syntax for the language (// for JS/TS, # for Python, etc.)

**Transformation Examples:**

| WHY Question | User Answer | Final Comment |
|--------------|-------------|---------------|
| "Why lock after 3 attempts❓" | "because of security policy SEC-2024-01" | `// Lock after 3 attempts per security policy SEC-2024-01` |
| "Why 5 minute TTL❓" | "balances freshness with API rate limits" | `// 5 minute TTL balances data freshness with API rate limits` |
| "Why is contract terminated here❓" | "because this is the first step of the leasing process" | `// Contract is terminated here as the first step of the leasing process` |
| "Why use Map instead of array❓" | "O(1) lookup vs O(n), matters for 1000+ items" | `// Use Map for O(1) lookup instead of O(n) array scan; critical for 1000+ items` |

**Key principles:**
- Create complete, standalone sentences
- Preserve user's technical terms, policy references, and specific details
- Remove conversational fillers ("because", "well", "so")
- Ensure the comment makes sense to future readers who won't see your question

Apply all comments in sequence, then provide a summary.

### Step 5: Complete the Process

After applying all comments, provide a summary:
```
✅ Applied 2 comments:
   - src/auth.ts:45
   - src/api.ts:80

⏭️  Skipped 1 location
```

## Critical Instructions

- ⚠️ **SHOW ALL AT ONCE**: Display ALL identified locations in a single message with numbers
- ⚠️ **SHOW CONTEXT**: Always show ~5 lines before and ~5 lines after the target line for natural reading
- ⚠️ **COMMENT IN CODE**: Place the "WHY:" question directly in the code block where the comment would go
- ⚠️ **END WITH EMOJI**: Always end why questions with ❓ emoji
- ⚠️ **USE READ TOOL**: Always use Read tool to get full file context, not just git diff snippets
- ⚠️ **WAIT FOR ANSWERS**: After showing all locations, wait for user to provide numbered answers
- ⚠️ **PARSE NUMBERED ANSWERS**: Extract lines starting with "number. text" format (e.g., "1. answer here")
- ⚠️ **TRANSFORM ANSWERS**: Convert user's answers into complete, grammatically correct standalone comments that incorporate the question context
- ⚠️ **PRESERVE MEANING**: Keep the user's core reasoning, terminology, and specific references intact
- ⚠️ **APPLY ALL AT ONCE**: Apply all provided answers in one batch using Edit tool for each
- ⚠️ **SKIP UNANSWERED**: Any question number not provided by user is automatically skipped
- ⚠️ **BE SELECTIVE**: Only suggest comments for genuinely non-obvious "why" scenarios

---

Begin your analysis now:

1. Identify all locations that need comments
2. Display ALL locations at once with numbers
3. Wait for user to provide answers in numbered format (e.g., "1. answer\n3. answer")
4. Parse the numbered answers and apply all comments
5. Provide summary of what was applied and skipped
