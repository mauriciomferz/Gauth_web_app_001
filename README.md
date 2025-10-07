# GAuth Educational Web Application

⚠️ **EDUCATIONAL IMPLEMENTATION NOTICE**  
This web application is designed for learning and demonstration purposes only. It is NOT intended for production use and should not be deployed in any production environment.

## 🌐 Modern Educational Interface for GAuth RFC-0150

This repository contains a comprehensive educational web application that demonstrates the GAuth authorization framework concepts through interactive learning experiences.

### 🎯 Purpose

Created to provide hands-on learning for:
- **Authorization Patterns** - Modern RBAC/ABAC with power-of-attorney flows
- **Token Management** - Complete lifecycle from creation to revocation
- **Event-Driven Architecture** - Typed events and pub/sub patterns
- **Audit and Compliance** - Comprehensive logging and reporting

### 🚀 Quick Start

```bash
# Clone this repository
git clone https://github.com/mauriciomferz/Gauth_web_app_001.git
cd Gauth_web_app_001

# Start the educational server
./start-web-demo.sh

# Open in browser
open http://localhost:8080
```

### 📋 Requirements

- **Go 1.21+** - Backend server implementation
- **Modern Browser** - Chrome, Firefox, Safari, Edge
- **Local Environment** - Designed for local development and learning

### 🎓 Learning Modules

1. **Token Management Demo**
   - Create educational tokens with RFC-compliant structure
   - Validate token integrity and expiration handling
   - Revoke tokens and understand blacklist management

2. **Authorization Flow Demo**
   - Experience power-of-attorney delegation patterns
   - Test different resource access scenarios
   - Understand policy evaluation and decision making

3. **Event System Demo**
   - Publish typed events with structured metadata
   - Subscribe to event streams with pattern matching
   - See real-time event processing and handling

4. **Audit Trail Demo**
   - View comprehensive audit logs from all interactions
   - Generate compliance reports with detailed analysis
   - Understand regulatory requirements and patterns

### 🛠 Technical Implementation

- **Backend**: Go with Gin web framework providing educational APIs
- **Frontend**: Modern HTML5, Tailwind CSS, and Vanilla JavaScript
- **Interactive**: Real-time console outputs and visual feedback
- **Responsive**: Mobile-first design that works on all devices

### 🔒 Educational Safety

This implementation uses:
- ✅ **Simulated Data** - All tokens and responses are educational
- ✅ **Local Development** - No external services or network exposure
- ✅ **Clear Warnings** - Prominent disclaimers throughout interface
- ✅ **Safe Algorithms** - Simplified logic for learning clarity

### 📚 What You'll Learn

- **GAuth Framework Principles** - Core authorization concepts
- **RFC-0150 Compliance** - Standards implementation patterns
- **Power-of-Attorney Flows** - Explicit delegation and verification
- **Event-Driven Patterns** - Modern architectural approaches
- **Audit and Security** - Compliance and monitoring practices

### 🔗 Related Projects

- [Gimel Foundation Repositories](https://github.com/Gimel-Foundation)
- [Original GAuth Implementation](https://github.com/Gimel-Foundation/GiFo-RFC-0150-Go-Implementation-of-GAuth-1.0)
- [Gimel Foundation Official](https://gimelfoundation.com)

### ⚠️ Important Notice

This is an **educational implementation** designed for learning purposes. It demonstrates GAuth concepts and RFC-0150 patterns but should **never be used in production environments** or for real security applications.

---

**Educational Web Application** by Mauricio Fernandez  
Based on GAuth RFC-0150 Implementation by Gimel Foundation  
Licensed under Apache 2.0
