# Copilot instructions for opa-test-problem

## Purpose of this repository

This repo exists solely to **reproduce and demonstrate a problematic behaviour in Open Policy Agent (OPA)** that occurs when running `opa test` with large JSON data files located in the directory tree under an OPA bundle root.

The goal is to produce a **minimal, shareable reproduction** that can be discussed publicly with the OPA project (e.g., in a GitHub issue). It must therefore:

- Contain **no real/work-related data**. All JSON data, package names, policy content, and identifiers must be synthetic/generated placeholder data.
- Be small and self-contained, with clear steps to reproduce the behaviour.

## Confidentiality rule (important)

The original problem was observed with sensitive, work-related data that **cannot be disclosed**. Never ask for, copy in, or reference the original data. When generating test data, use obviously fake/generated content (e.g., generated records, lorem-ipsum-like values, numeric sequences).

## Tech stack

- **Go** (used for tooling, e.g., generating large JSON files or driving OPA programmatically)
- **Open Policy Agent (OPA)**: Rego policies, `opa test`, bundle directory layout

## Working conventions

- Prefer minimal reproductions: the smallest bundle layout, policy, and data size that still triggers the behaviour.
- Make data generation reproducible (scripted/parameterized, e.g., configurable file size/record count) so file size can be varied to find thresholds.
- Document exact commands, OPA version, and observed vs. expected behaviour in the README as the investigation progresses.
- When discussing the issue, describe behaviour in generic terms (file sizes, counts, timings, memory usage) rather than data content.

## Status

Early stage — the behaviour is being described by the repo owner before any code is written. Do not scaffold or implement anything until asked.
