// GAuth Educational Demo JavaScript

// Global state
let currentToken = null;
let demoState = {
    tokenCreated: false,
    subscriptionsActive: false,
    auditEntries: []
};

// Utility functions
function addConsoleOutput(containerId, message, type = 'info') {
    const container = document.getElementById(containerId);
    const timestamp = new Date().toISOString().split('T')[1].split('.')[0];
    const typeClass = type === 'success' ? 'console-success' :
                     type === 'error' ? 'console-error' :
                     type === 'warning' ? 'console-warning' :
                     'console-info';
    
    const line = document.createElement('div');
    line.className = `console-line ${typeClass}`;
    line.innerHTML = `<span class="text-gray-400">[${timestamp}]</span> ${message}`;
    
    container.appendChild(line);
    container.scrollTop = container.scrollHeight;
}

function clearConsoleOutput(containerId) {
    const container = document.getElementById(containerId);
    container.innerHTML = `
        <span class="text-gray-500"># ${containerId.replace('-output', '').toUpperCase()} Console</span><br>
        <span class="text-blue-400">gauth-${containerId.replace('-output', '')}></span> <span class="blinking-cursor">_</span>
    `;
}

function generateRandomId() {
    return Math.random().toString(36).substr(2, 9);
}

function simulateApiDelay() {
    return new Promise(resolve => setTimeout(resolve, Math.random() * 1000 + 500));
}

// Tab system
function showTab(tabId) {
    // Hide all tab contents
    const tabContents = document.querySelectorAll('.tab-content');
    tabContents.forEach(content => {
        content.style.display = 'none';
        content.classList.remove('active');
    });
    
    // Remove active class from all tab buttons
    const tabButtons = document.querySelectorAll('.tab-button');
    tabButtons.forEach(button => button.classList.remove('active'));
    
    // Show selected tab content
    const selectedTab = document.getElementById(tabId);
    if (selectedTab) {
        selectedTab.style.display = 'block';
        selectedTab.classList.add('active');
    }
    
    // Add active class to clicked button
    event.target.classList.add('active');
}

// Navigation
function scrollToDemo() {
    document.getElementById('demo').scrollIntoView({ 
        behavior: 'smooth' 
    });
}

// Token Management Demo Functions
async function createToken() {
    addConsoleOutput('token-output', 'Initiating educational token creation...', 'info');
    
    try {
        await simulateApiDelay();
        
        currentToken = {
            id: `token_${generateRandomId()}`,
            type: 'educational_demo',
            issuer: 'gauth-educational-demo',
            subject: 'demo-user@example.com',
            audience: 'learning-environment',
            expiresAt: new Date(Date.now() + 3600000).toISOString(),
            createdAt: new Date().toISOString(),
            claims: {
                scope: 'read write demo',
                educational: true,
                purpose: 'RFC-0150 demonstration'
            }
        };
        
        addConsoleOutput('token-output', `✓ Token created successfully`, 'success');
        addConsoleOutput('token-output', `  ID: ${currentToken.id}`, 'info');
        addConsoleOutput('token-output', `  Subject: ${currentToken.subject}`, 'info');
        addConsoleOutput('token-output', `  Expires: ${currentToken.expiresAt}`, 'info');
        addConsoleOutput('token-output', `  ⚠️ Educational demo token - not cryptographically secure`, 'warning');
        
        demoState.tokenCreated = true;
        
        // Add to audit trail
        demoState.auditEntries.push({
            timestamp: new Date().toISOString(),
            action: 'TOKEN_CREATED',
            tokenId: currentToken.id,
            subject: currentToken.subject,
            educational: true
        });
        
    } catch (error) {
        addConsoleOutput('token-output', `✗ Token creation failed: ${error.message}`, 'error');
    }
}

async function validateToken() {
    if (!currentToken) {
        addConsoleOutput('token-output', '✗ No token available for validation. Create a token first.', 'error');
        return;
    }
    
    addConsoleOutput('token-output', 'Validating educational token...', 'info');
    
    try {
        await simulateApiDelay();
        
        const now = new Date();
        const expiry = new Date(currentToken.expiresAt);
        
        if (now > expiry) {
            addConsoleOutput('token-output', '✗ Token validation failed: Token has expired', 'error');
            addConsoleOutput('token-output', `  Expired at: ${currentToken.expiresAt}`, 'error');
            currentToken = null;
            demoState.tokenCreated = false;
        } else {
            addConsoleOutput('token-output', '✓ Token validation successful', 'success');
            addConsoleOutput('token-output', `  Valid until: ${currentToken.expiresAt}`, 'info');
            addConsoleOutput('token-output', `  Claims verified: ${Object.keys(currentToken.claims).join(', ')}`, 'info');
        }
        
        // Add to audit trail
        demoState.auditEntries.push({
            timestamp: new Date().toISOString(),
            action: 'TOKEN_VALIDATED',
            tokenId: currentToken?.id,
            valid: now <= expiry,
            educational: true
        });
        
    } catch (error) {
        addConsoleOutput('token-output', `✗ Token validation failed: ${error.message}`, 'error');
    }
}

async function revokeToken() {
    if (!currentToken) {
        addConsoleOutput('token-output', '✗ No token available for revocation.', 'error');
        return;
    }
    
    addConsoleOutput('token-output', `Revoking token ${currentToken.id}...`, 'info');
    
    try {
        await simulateApiDelay();
        
        const revokedTokenId = currentToken.id;
        
        addConsoleOutput('token-output', `✓ Token ${revokedTokenId} revoked successfully`, 'success');
        addConsoleOutput('token-output', '  Token added to blacklist', 'info');
        addConsoleOutput('token-output', '  All dependent sessions invalidated', 'info');
        
        // Add to audit trail
        demoState.auditEntries.push({
            timestamp: new Date().toISOString(),
            action: 'TOKEN_REVOKED',
            tokenId: revokedTokenId,
            educational: true
        });
        
        currentToken = null;
        demoState.tokenCreated = false;
        
    } catch (error) {
        addConsoleOutput('token-output', `✗ Token revocation failed: ${error.message}`, 'error');
    }
}

// Authorization Demo Functions
async function checkAuthorization() {
    const action = document.getElementById('resource-action').value;
    
    addConsoleOutput('authz-output', `Checking authorization for action: ${action}`, 'info');
    
    try {
        await simulateApiDelay();
        
        if (!currentToken) {
            addConsoleOutput('authz-output', '✗ Authorization denied: No valid token', 'error');
            addConsoleOutput('authz-output', '  Create and validate a token first', 'warning');
            return;
        }
        
        // Simulate authorization logic
        const authzDecision = simulateAuthzDecision(action, currentToken);
        
        if (authzDecision.allowed) {
            addConsoleOutput('authz-output', `✓ Authorization granted for ${action}`, 'success');
            addConsoleOutput('authz-output', `  Policy: ${authzDecision.policy}`, 'info');
            addConsoleOutput('authz-output', `  Delegation chain: ${authzDecision.delegationChain.join(' → ')}`, 'info');
        } else {
            addConsoleOutput('authz-output', `✗ Authorization denied for ${action}`, 'error');
            addConsoleOutput('authz-output', `  Reason: ${authzDecision.reason}`, 'error');
        }
        
        addConsoleOutput('authz-output', `  ⚠️ Educational simulation - not production authorization`, 'warning');
        
        // Add to audit trail
        demoState.auditEntries.push({
            timestamp: new Date().toISOString(),
            action: 'AUTHORIZATION_CHECK',
            resource: action,
            allowed: authzDecision.allowed,
            policy: authzDecision.policy,
            educational: true
        });
        
    } catch (error) {
        addConsoleOutput('authz-output', `✗ Authorization check failed: ${error.message}`, 'error');
    }
}

function simulateAuthzDecision(action, token) {
    const decisions = {
        'read': {
            allowed: true,
            policy: 'allow_authenticated_reads',
            delegationChain: ['user', 'demo-session'],
            reason: 'Valid token with read scope'
        },
        'write': {
            allowed: token.claims.scope.includes('write'),
            policy: 'require_write_permission',
            delegationChain: ['user', 'demo-session', 'write-delegate'],
            reason: token.claims.scope.includes('write') ? 'Valid write permission' : 'Missing write scope'
        },
        'admin': {
            allowed: false,
            policy: 'deny_admin_in_demo',
            delegationChain: ['user', 'demo-session'],
            reason: 'Administrative actions not permitted in educational demo'
        },
        'delegate': {
            allowed: token.claims.educational === true,
            policy: 'allow_educational_delegation',
            delegationChain: ['user', 'demo-session', 'delegation-authority'],
            reason: token.claims.educational ? 'Educational delegation permitted' : 'Delegation requires educational context'
        }
    };
    
    return decisions[action] || {
        allowed: false,
        policy: 'default_deny',
        delegationChain: ['user'],
        reason: 'Unknown action'
    };
}

// Event System Demo Functions
async function publishEvent() {
    addConsoleOutput('event-output', 'Publishing typed event...', 'info');
    
    try {
        await simulateApiDelay();
        
        const event = {
            id: `event_${generateRandomId()}`,
            type: 'gauth.demo.user_action',
            version: '1.0',
            source: 'educational-demo',
            subject: currentToken?.subject || 'anonymous',
            time: new Date().toISOString(),
            data: {
                action: 'demo_interaction',
                educational: true,
                metadata: {
                    session_id: `demo_session_${generateRandomId()}`,
                    user_agent: 'GAuth Educational Demo',
                    interaction_type: 'event_publish'
                }
            }
        };
        
        addConsoleOutput('event-output', `✓ Event published successfully`, 'success');
        addConsoleOutput('event-output', `  Event ID: ${event.id}`, 'info');
        addConsoleOutput('event-output', `  Type: ${event.type}`, 'info');
        addConsoleOutput('event-output', `  Subject: ${event.subject}`, 'info');
        addConsoleOutput('event-output', `  Metadata enriched with session context`, 'info');
        addConsoleOutput('event-output', `  ⚠️ Educational event - not persisted to production systems`, 'warning');
        
        // Simulate event handlers
        setTimeout(() => {
            addConsoleOutput('event-output', `📨 Event handler triggered: audit_logger`, 'info');
            addConsoleOutput('event-output', `📨 Event handler triggered: metrics_collector`, 'info');
        }, 1000);
        
    } catch (error) {
        addConsoleOutput('event-output', `✗ Event publication failed: ${error.message}`, 'error');
    }
}

async function subscribeEvents() {
    if (demoState.subscriptionsActive) {
        addConsoleOutput('event-output', '✗ Event subscriptions already active', 'warning');
        return;
    }
    
    addConsoleOutput('event-output', 'Subscribing to event streams...', 'info');
    
    try {
        await simulateApiDelay();
        
        addConsoleOutput('event-output', `✓ Subscribed to: gauth.demo.*`, 'success');
        addConsoleOutput('event-output', `✓ Subscribed to: gauth.audit.*`, 'success');
        addConsoleOutput('event-output', `✓ Subscribed to: gauth.token.*`, 'success');
        addConsoleOutput('event-output', `  Pattern matching enabled`, 'info');
        addConsoleOutput('event-output', `  Typed event handlers registered`, 'info');
        
        demoState.subscriptionsActive = true;
        
        // Simulate receiving events
        let eventCount = 0;
        const eventInterval = setInterval(() => {
            if (eventCount >= 3 || !demoState.subscriptionsActive) {
                clearInterval(eventInterval);
                return;
            }
            
            const eventTypes = ['token.created', 'auth.granted', 'audit.logged'];
            const eventType = eventTypes[eventCount];
            const eventId = generateRandomId();
            
            addConsoleOutput('event-output', `📩 Received: gauth.demo.${eventType} [${eventId}]`, 'success');
            eventCount++;
        }, 2000);
        
        addConsoleOutput('event-output', `  ⚠️ Educational subscription - simulated events only`, 'warning');
        
    } catch (error) {
        addConsoleOutput('event-output', `✗ Event subscription failed: ${error.message}`, 'error');
    }
}

// Audit Demo Functions
async function viewAuditLog() {
    addConsoleOutput('audit-output', 'Retrieving audit log entries...', 'info');
    
    try {
        await simulateApiDelay();
        
        if (demoState.auditEntries.length === 0) {
            addConsoleOutput('audit-output', '📋 No audit entries found', 'warning');
            addConsoleOutput('audit-output', '  Interact with other demos to generate audit entries', 'info');
            return;
        }
        
        addConsoleOutput('audit-output', `✓ Found ${demoState.auditEntries.length} audit entries`, 'success');
        
        demoState.auditEntries.forEach((entry, index) => {
            const timestamp = new Date(entry.timestamp).toLocaleTimeString();
            addConsoleOutput('audit-output', `  [${index + 1}] ${timestamp} - ${entry.action}`, 'info');
            if (entry.tokenId) {
                addConsoleOutput('audit-output', `      Token: ${entry.tokenId}`, 'info');
            }
            if (entry.resource) {
                addConsoleOutput('audit-output', `      Resource: ${entry.resource}`, 'info');
            }
        });
        
        addConsoleOutput('audit-output', `  ⚠️ Educational audit log - not compliance-grade storage`, 'warning');
        
    } catch (error) {
        addConsoleOutput('audit-output', `✗ Audit log retrieval failed: ${error.message}`, 'error');
    }
}

async function generateReport() {
    addConsoleOutput('audit-output', 'Generating compliance report...', 'info');
    
    try {
        await simulateApiDelay();
        
        const report = {
            generated: new Date().toISOString(),
            period: 'Educational Demo Session',
            entries: demoState.auditEntries.length,
            breakdown: {
                token_operations: demoState.auditEntries.filter(e => e.action.includes('TOKEN')).length,
                authorization_checks: demoState.auditEntries.filter(e => e.action === 'AUTHORIZATION_CHECK').length,
                successful_operations: demoState.auditEntries.filter(e => e.allowed !== false).length
            }
        };
        
        addConsoleOutput('audit-output', `✓ Report generated successfully`, 'success');
        addConsoleOutput('audit-output', `  Total entries: ${report.entries}`, 'info');
        addConsoleOutput('audit-output', `  Token operations: ${report.breakdown.token_operations}`, 'info');
        addConsoleOutput('audit-output', `  Authorization checks: ${report.breakdown.authorization_checks}`, 'info');
        addConsoleOutput('audit-output', `  Successful operations: ${report.breakdown.successful_operations}`, 'info');
        addConsoleOutput('audit-output', `  Generated: ${report.generated}`, 'info');
        addConsoleOutput('audit-output', `  ⚠️ Educational report - not regulatory compliance`, 'warning');
        
    } catch (error) {
        addConsoleOutput('audit-output', `✗ Report generation failed: ${error.message}`, 'error');
    }
}

// Example viewing functions
function viewExample(exampleType) {
    const examples = {
        'rfc-0115': 'examples/rfc_0115_poa_definition/',
        'typed-events': 'examples/typed_events/',
        'token-revocation': 'examples/token/',
        'resilience': 'examples/resilience/',
        'cascade': 'examples/cascade/',
        'microservices': 'examples/microservices/'
    };
    
    const examplePath = examples[exampleType];
    if (examplePath) {
        // In a real implementation, this would navigate to the example
        alert(`🎓 Educational Example\n\nThis would open: ${examplePath}\n\nIn the actual implementation, this links to the GitHub repository with working code examples.`);
    }
}

// Initialize the application
document.addEventListener('DOMContentLoaded', function() {
    // Show the first tab by default
    const firstTab = document.querySelector('.tab-content');
    if (firstTab) {
        firstTab.style.display = 'block';
        firstTab.classList.add('active');
    }
    
    // Add educational notices to all consoles
    const consoleIds = ['token-output', 'authz-output', 'event-output', 'audit-output'];
    consoleIds.forEach(id => {
        const container = document.getElementById(id);
        if (container) {
            container.innerHTML += '<br><span class="text-yellow-400">⚠️ Educational Demo Environment</span>';
        }
    });
    
    console.log('🎓 GAuth Educational Demo initialized');
    console.log('⚠️ This is an educational implementation for learning purposes only');
});