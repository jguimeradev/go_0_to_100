# go_0_to_100

# 🟢 LEVEL 0 — Language & Core Foundations (1–10)

> Syntax, types, memory, stdlib, tooling, build system

1. `hello-go` — Hello World + build/run script
2. `args-parser` — CLI args parser
3. `env-reader` — Read env vars + config loader
4. `file-io-tool` — Read/write files
5. `json-tool` — Encode/decode JSON
6. `csv-parser`
7. `log-system` — Custom logger
8. `config-manager` — YAML/JSON config loader
9. `error-wrapper` — Custom error handling package
10. `module-playground` — Multi-module Go project

**Skills:**
Go toolchain, `go mod`, packages, error handling, interfaces, structs, build lifecycle

---

# 🟢 LEVEL 1 — CLI Engineering (11–20)

> CLI UX, commands, flags, structure

11. `todo-cli`
12. `notes-cli`
13. `password-generator`
14. `file-search`
15. `dir-tree`
16. `log-analyzer`
17. `system-info-cli`
18. `git-like-cli`
19. `port-scanner`
20. `task-runner`

**Skills:**
CLI UX, flags, subcommands, config files, IO streams, cross-platform

---

# 🟡 LEVEL 2 — Networking + HTTP (21–30)

21. `http-client`
22. `http-server`
23. `rest-api-basic`
24. `middleware-engine`
25. `router`
26. `rate-limiter`
27. `api-gateway-mini`
28. `proxy-server`
29. `tcp-chat`
30. `udp-monitor`

**Skills:**
net/http, TCP/UDP, concurrency, sockets, protocols, routing

---

# 🟡 LEVEL 3 — REST APIs & Backend Systems (31–40)

31. `crud-api`
32. `auth-api`
33. `jwt-auth`
34. `rbac-system`
35. `session-manager`
36. `file-upload-api`
37. `api-logging`
38. `metrics-endpoint`
39. `healthcheck-system`
40. `api-versioning-system`

**Skills:**
API design, auth, security, middleware chains, architecture

---

# 🟠 LEVEL 4 — Persistence & Data (41–50)

41. `in-memory-db`
42. `kv-store`
43. `cache-engine`
44. `file-db`
45. `sql-layer`
46. `migration-tool`
47. `backup-tool`
48. `replication-simulator`
49. `indexing-engine`
50. `query-engine`

**Skills:**
storage models, consistency, durability, indexing, serialization

---

# 🟠 LEVEL 5 — Concurrency & Systems (51–60)

51. `worker-pool`
52. `job-queue`
53. `task-scheduler`
54. `event-bus`
55. `pub-sub`
56. `message-broker`
57. `distributed-lock`
58. `leader-election`
59. `heartbeat-system`
60. `failover-simulator`

**Skills:**
goroutines, channels, synchronization, distributed patterns

---

# 🔴 LEVEL 6 — Observability & Reliability (61–70)

61. `log-aggregator`
62. `metrics-collector`
63. `tracing-system`
64. `alert-engine`
65. `health-monitor`
66. `service-discovery`
67. `config-distributor`
68. `secrets-manager`
69. `feature-flags`
70. `chaos-engineering-tool`

**Skills:**
SRE patterns, monitoring, resilience engineering

---

# 🔴 LEVEL 7 — Platform Engineering (71–80)

71. `deployment-engine`
72. `orchestrator-mini`
73. `service-mesh-mini`
74. `api-management-platform`
75. `infra-provisioner`
76. `pipeline-runner`
77. `artifact-repo`
78. `build-system`
79. `image-builder`
80. `runtime-manager`

**Skills:**
CI/CD engines, infra orchestration, pipelines, runtimes

---

# 🧠 LEVEL 8 — Containerization Internals (81–90)

81. `namespace-manager`
82. `cgroup-manager`
83. `process-isolator`
84. `fs-overlay`
85. `layered-fs`
86. `image-format`
87. `container-runtime`
88. `container-network`
89. `container-scheduler`
90. `container-daemon`

**Skills:**
Linux namespaces, cgroups, filesystem layers, isolation, kernel interfaces

---

# 🧠 LEVEL 9 — Docker-like Tool (91–100)

91. `image-parser`
92. `build-engine`
93. `layer-cache`
94. `registry-client`
95. `registry-server`
96. `image-pusher`
97. `image-puller`
98. `runtime-cli`
99. `container-manager`
100. **`dockermind`** — full docker-like system:

* image build
* layer caching
* isolated runtime
* networking
* volume mounting
* registry
* CLI
* daemon
* API

---

# 🏗 Final Project Architecture (100)

```
dockermind/
├── cli/
├── daemon/
├── runtime/
├── image/
├── network/
├── storage/
├── registry/
├── api/
├── scheduler/
├── config/
├── security/
├── observability/
```

---

# 🧭 Learning Strategy (Critical)

### Mode:

* **Docs-first**
* **No tutorials**
* **Spec-first**
* **RFC-driven**
* **Code reading**

### Sources:

* Go stdlib
* Kubernetes source
* containerd
* runc
* Docker
* Hashicorp tools
* Linux kernel docs





