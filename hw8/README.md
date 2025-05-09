# README.md
# HW8: Safe Logger — Exploring Synchronization and fsync in Go

**Course:** EECE 4811/5811  
**Due Date:** 5/8/2025  
**Group Members:** Sam Chum 

---
## Overview
This assignment implements three concurrent loggers in Go to explore synchronization, file flushing (`fsync`), and data races. Each logger writes log entries to a file in the format:

