#  GCP Function Deployment Automation

This project automates Google Cloud Function deployment using a custom Go CLI tool (`gcptool`) and GitHub Actions.

---

##  Project Overview

This repository contains:

* `cmd/gcptool/`: A CLI tool written in Go to deploy and describe Google Cloud Functions
* `function/`: The Cloud Function source code
* `.github/workflows/deploy.yml`: GitHub Actions workflow for testing and deployment

---

##  CLI Tool – `gcptool`

A lightweight CLI utility to deploy or describe GCP functions.

###  Usage

```bash
gcptool -h
```

###  Deploy a function

```bash
gcptool deploy <functionName> -e <environment> -v <revision> [-c]
```

**Options:**

* `-e`: Target environment (`sandbox`, `dev`, or `prod`)
* `-v`: Revision/version number
* `-c`: *(Optional)* Clean and rebuild before deploying

**Example:**

```bash
gcptool deploy myFunction -e dev -v 2 -c
```

###  Describe a function

```bash
gcptool describe <functionName>
```

**Example:**

```bash
gcptool describe myFunction
```

---

##  GitHub Actions Workflow

A CI workflow is set up to:

* Run unit tests when code is pushed
* Deploy the function automatically if tests pass

**Workflow path:**

```text
.github/workflows/deploy.yml
```

To activate it, configure the following GitHub Secrets:

* `GCP_SA_KEY`
* `GCP_PROJECT_ID`

---

## Testing

Unit tests cover the deploy and describe logic using mocked `exec.Command`.

**Run tests:**

```bash
cd cmd/gcptool
go test -v ./...
```

---

##  Prerequisites

* Go 1.21+
* `gcloud` CLI
* A GCP project with Cloud Functions API enabled
* A Google Service Account with deploy permissions

---

##  Folder Structure

```text
.
├── cmd/
│   └── gcptool/
│       ├── main.go
│       ├── deploy.go
│       ├── describe.go
│       ├── common.go
│       └── main_test.go
├── function/
│   ├── go.mod
│   └── function.go
├── .github/
│   └── workflows/
│       └── deploy.yml
└── README.md
```

---

##  Author

**Venkatesh P**

