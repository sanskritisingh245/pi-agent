# Pi Terminal

## Vision

Build a production-grade conversational AI agent inspired by Pi AI that runs entirely inside the terminal.

The goal is NOT to clone Pi's UI or proprietary implementation. The goal is to understand and build every system that makes Pi feel intelligent, emotionally aware, conversational, and human.

The project is architected as a modular AI agent framework where Pi is simply one personality. Everything should be extensible. The system should later allow replacing:

- OpenAI with Claude/Gemini/Ollama
- Terminal with Desktop/Web/Voice
- Pi personality with any custom assistant

without rewriting the application.

## Final User Experience

```
$ pi

╭──────────────────────────────────────────╮
│ Pi Terminal                              │
╰──────────────────────────────────────────╯

Pi:
Hi, I'm Pi.
What's on your mind today?

You >
```

The conversation should feel completely natural. The assistant:

- remembers previous conversations
- understands emotions
- understands intent
- plans responses
- asks thoughtful follow-up questions
- avoids robotic responses
- never dumps unnecessary information
- feels calm, curious, conversational, emotionally intelligent

It should feel like talking to another human rather than an LLM.

## Primary Goal

Build the systems around the LLM. The LLM is only responsible for generating language. Everything else — emotion, intent, memory, planning, reflection — is built by us.

## Architecture Philosophy

Never build:

```
User → GPT → Response
```

Instead build:

```
User
 ↓
Conversation Manager
 ↓
Emotion Analysis
 ↓
Intent Classification
 ↓
Memory Retrieval
 ↓
Conversation State
 ↓
Planner
 ↓
Context Builder
 ↓
LLM
 ↓
Reflection
 ↓
Memory Update
 ↓
Response
```

## Core Components

### 1. Terminal Interface
Read user input, display streaming responses, handle commands, display status. Never contains business logic.

### 2. Agent
Central orchestrator. Coordinates every subsystem. Never directly communicates with OpenAI, never stores data, never builds prompts itself.

Workflow: Receive Message → Emotion → Intent → Memory → Planner → Context Builder → LLM → Reflection → Memory Update → Return Response.

### 3. LLM Client
Abstracts all model providers. Handles API communication, retries, streaming, model selection. Returns generated text.

Should never know about emotions, memory, conversations, users, or personality.

Supported providers: OpenAI (Phase 1), later Claude, Gemini, Ollama, LM Studio.

### 4. Emotion Engine
Input: user message. Output: primary emotion, secondary emotion, energy level, confidence, stress, needs-validation flag. Understands emotional state; never generates responses.

### 5. Intent Engine
Input: user message. Output: intent (Learning, Advice, Reflection, Venting, Casual Chat, Coding, Planning). Determines user goal.

### 6. Memory Engine
Responsibilities: store, retrieve, rank, forget.

Types: short-term (current conversation), long-term (persistent facts, emotional memory, relationship memory, preference memory).

Memory should never store everything — it decides whether each item should be remembered (YES → store, NO → ignore).

### 7. Conversation State
Tracks current topic, current goal, current emotional state, conversation mode (Teaching, Planning, Support, Reflection, Brainstorming).

### 8. Planner
Decides HOW to respond. Example — user: "I'm scared I'll fail." → Planner output: validate emotion → avoid giving advice → ask one open question → keep under 80 words → be encouraging. Generation only begins after planning.

### 9. Context Builder
Collects user message, emotion, intent, memory, conversation state, personality, and planner instructions, and merges everything into one prompt. This is the only component that builds prompts.

### 10. Reflection Engine
Runs after every response. Evaluates: Did I answer? Did I ignore emotion? Was I repetitive? Should memory be updated? Did I ask too many questions?

### 11. Personality System
Personality is NOT one prompt — it is behavioural rules, e.g.:
- Always validate emotions before advice
- Never overwhelm the user
- Never sound clinical
- Ask at most one follow-up question
- Avoid excessive positivity
- Prefer curiosity
- Keep responses concise

### 12. Prompt System
Prompts are stored separately, never hardcoded (e.g. `system.md`, `planner.md`, `emotion.md`, `reflection.md`, `personality.md`).

### 13. Storage
Phase 1: SQLite — stores conversations, memory, reflections, preferences.
Later: PostgreSQL + pgvector.

### 14. Configuration
Environment variables: API keys, model, temperature, max tokens, voice provider, database.

### 15. Logging
Logs: API latency, prompt size, token usage, errors, memory retrieval, planner decisions.

### 16. Streaming
Support token streaming; display responses character-by-character.

### 17. Voice (later)
Speech To Text → Conversation Pipeline → Text To Speech.

## Technology Stack

- Language: Go
- LLM: OpenAI (initially)
- Storage: SQLite (later PostgreSQL + pgvector)
- Logging: `slog`
- Configuration: `.env`
- CLI: standard terminal first, Bubble Tea later

## Engineering Principles

- Every package has one responsibility
- Business logic never lives in `main.go`
- Agent never depends directly on OpenAI
- LLM Client never knows about memory
- Memory never knows about prompts
- Planner never generates responses
- Emotion never stores data
- Context Builder is the only component that builds prompts
- Everything should be replaceable
- Every dependency should be injected
- Avoid global variables
- Keep components loosely coupled

## Development Roadmap

- **Phase 0** — Foundation: project structure, config, LLM client, agent, terminal loop
- **Phase 1** — Streaming
- **Phase 2** — Conversation history
- **Phase 3** — Persistent memory
- **Phase 4** — Emotion engine
- **Phase 5** — Intent engine
- **Phase 6** — Planner
- **Phase 7** — Context builder
- **Phase 8** — Reflection
- **Phase 9** — Voice

## Final Goal

By the end of the project, the application should not feel like "ChatGPT in a terminal." It should feel like a genuine conversational companion that listens, remembers, reasons, plans, reflects, and responds naturally through a carefully designed architecture rather than relying solely on the underlying language model.
