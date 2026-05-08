# Architecture Decision Records

This directory contains **active** Architecture Decision Records (ADRs) for the KubeOpenCode project.
ADRs document significant architectural and design decisions along with their context and consequences.

> **For AI agents**: Only ADRs in this directory are relevant for current development.
> The `archived/` subdirectory contains superseded, fully-implemented, and research-only ADRs
> that are kept for historical reference but should NOT be loaded into context unless explicitly needed.

## Active ADRs

### Recently Implemented (still useful as reference)

| # | Title | Date |
|---|-------|------|
| 0033 | [Agent Share Link — Token-based Terminal Sharing](0033-agent-share-link.md) | 2026-04-15 |
| 0034 | [First-Class Plugin Support and Slack Integration](0034-plugin-support-and-slack-integration.md) | 2026-04-18 |
| 0035 | [Exposing OpenCode Sessions in Task Status](0035-task-session-integration.md) | 2026-04-22 |
| 0037 | [Task Timeout](0037-task-timeout.md) | 2026-05-06 |
| 0038 | [Extra Environment Variables and Volume Mounts for System Containers](0038-extra-env-init-containers.md) | 2026-05-07 |

### Accepted (active design contracts)

| # | Title | Notes |
|---|-------|-------|
| 0024 | [Agent Standby — Unified Suspend/Resume Lifecycle](0024-agent-standby.md) | |
| 0025 | [CronTask — Scheduled Task Execution](0025-crontask.md) | |
| 0026 | [Skills as a Top-Level Agent Field](0026-skills.md) | |
| 0027 | [Git Context Auto-Sync for Agents](0027-git-sync.md) | |
| 0028 | [Connection-Aware Standby — Heartbeat-based Idle Detection](0028-connection-aware-standby.md) | |
| 0029 | [Container-in-Container (DinD) Support for Agent Workloads](0029-sysbox-dind-support.md) | |

### Partially Implemented

| # | Title | Notes |
|---|-------|-------|
| 0016 | [Human-in-the-Loop (HITL) Design](0016-human-in-the-loop-design.md) | Simplified model adopted; Phase 2 UI superseded by ADR 0018; Phase 3 not started |

### Accepted (Deferred)

| # | Title | Notes |
|---|-------|-------|
| 0012 | [Defer Session API to Post-v0.1](0012-defer-session-api.md) | Blocked on Server Mode production validation and security model |

### Proposed (future work)

| # | Title | Date |
|---|-------|------|
| 0020 | [Enterprise Readiness Roadmap](0020-enterprise-readiness-roadmap.md) | |
| 0026b | [MCP Server Support in Agent API](0026-mcp-support.md) | |
| 0030 | [Graceful Task Termination on Deletion](0030-task-deletion-graceful-stop.md) | |
| 0031 | [OpenTelemetry Observability for Tasks and Agents](0031-opentelemetry-observability.md) | |
| 0036 | [Agent Registry — Enterprise Agent Asset Management and Visual Agent Assembly](0036-agent-registry.md) | |

## Archived ADRs

Archived ADRs are in [`archived/`](archived/). They include:

- **Superseded**: Decisions replaced by newer ADRs (0002, 0011, 0013, 0015, 0017, 0023)
- **Fully Implemented**: Decisions fully absorbed into code and AGENTS.md (0001, 0006-0010, 0014, 0018-0019, 0021-0022)
- **Research/Informational**: Analysis documents with no code impact (0032, 0039)

### Archive Policy

An ADR is archived when:
1. It is **superseded** by a newer ADR
2. It is **fully implemented** and its key decisions are summarized in `AGENTS.md`
3. It is a **research/informational** document with no actionable code changes

Archived ADRs are never deleted — they remain available for historical context.
When a new feature fully implements an Accepted ADR, move it to `archived/` and ensure
the key decisions are captured in `AGENTS.md`.
