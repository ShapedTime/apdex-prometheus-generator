# Apdex Rule Generator

## Overview
This application is a rule generator for Prometheus, designed to calculate Application Performance Index (Apdex) scores. It takes a configuration file in YAML format to create Prometheus recording rules, which can then be used to track the Apdex scores of different flows, processes or applications.

## Features
- **Configurable Apdex Scores**: Set thresholds and labels for satisfactory (Good), tolerable (Tolerating), and frustrating response times.
- **Custom Labels**: Enhance metrics with additional labels for more detailed monitoring.
- **Multiple Time Windows**: Supports various time windows for flexible and comprehensive Apdex score calculations.

## Prerequisites
- Go (Version 1.16 or later)
- Access to a Prometheus server (for deploying the generated rules)

## Installation
1. Clone the repository to your local machine.
2. Ensure Go is installed and set up.
3. Navigate to the project directory.

## Configuration
Configure your Apdex rules in the `config/config.yml` file. Here's an example structure based on your provided sample:

```yaml
apdex:
  - name: Name of this apdex
    subsystem: subsystem for this apdex  # Optional: specify a subsystem if applicable
    system: "system of this apdex"
    good_bucket: bucket that indicate good part of the apdex value
    tolerating_bucket: bucket that indicates tolerating part of the apdex value
    total: total bucket
    windows: 
      - windows for this alert to be generated
    extra_labels:
      extra: labels
      that: you want to add to apdex metrics
```

### Example Config:

```yaml
apdex:
  - name: "hello_world"
    system: "Universe"
    good_bucket: response_time_bucket{status="success",name="account_controller",le="360"}
    tolerating_bucket: "response_time_bucket{status="success",name="account_controller",le="1500"}
    total: "response_time_bucket{status="success",name="account_controller", le="+Inf"}
    windows: 
      - "5m"
      - "10m"
      - "30m"
      - "1h"
      - "6h"
      - "12h"
      - "1d"
      - "3d"
      - "7d"
      - "30d"
    extra_labels:
      controller: "account_controller"
```

# Running the Application
Generate the Prometheus rule file with the following command:
```bash
go run main.go --config=<path to your config.yml> --output=<path to your desired output.yml>
```
The application will default to config.yml and output.yml if no paths are specified.

# Output
The application outputs a Prometheus recording rule file (output.yml) based on the provided configuration. This file can then be used within Prometheus to start monitoring Apdex scores across your specified time windows.

# Contributing
Contributions are very welcome! If you have suggestions for improvements or encounter any issues, please feel free to open an issue or submit a pull request.