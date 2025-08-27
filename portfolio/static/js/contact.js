// Contact Page JavaScript - Terminal Theme
document.addEventListener('DOMContentLoaded', function() {
    // Initialize all functionality
    initTerminalStructure();
    initTypingEffect();
    initContactForm();
    initTerminalCommands();
});

// Create terminal structure and wrap content
function initTerminalStructure() {
    const contactSection = document.querySelector('.contact-section');
    const existingContent = contactSection.querySelector('.contact-content');
    
    if (existingContent && !contactSection.querySelector('.terminal-container')) {
        // Store the original content
        const content = existingContent.outerHTML;
        
        // Create terminal structure
        const terminalHTML = `
            <div class="terminal-container">
                <div class="terminal-header">
                    <div class="terminal-controls">
                        <span class="control-btn close"></span>
                        <span class="control-btn minimize"></span>
                        <span class="control-btn maximize"></span>
                    </div>
                    <div class="terminal-title">musaiyab@portfolio:~/contact$</div>
                </div>
                <div class="terminal-body">
                    <div class="command-line">
                        <span class="prompt">user@portfolio:~$</span>
                        <span class="command">./contact_me.sh</span>
                    </div>
                    ${content}
                    <div id="messageContainer"></div>
                    <div class="command-line">
                        <span class="prompt">user@portfolio:~$</span>
                        <span class="command"></span>
                        <span class="blinking-cursor"></span>
                    </div>
                </div>
            </div>
        `;
        
        // Replace existing content with terminal
        existingContent.remove();
        contactSection.insertAdjacentHTML('beforeend', terminalHTML);
    }
}

// Terminal typing effect for the main heading
function initTypingEffect() {
    const heading = document.querySelector('.contact-section h1');
    if (!heading) return;
    
    const text = heading.textContent;
    heading.textContent = '';
    
    let i = 0;
    const typeSpeed = 100;
    
    function typeWriter() {
        if (i < text.length) {
            heading.textContent += text.charAt(i);
            i++;
            setTimeout(typeWriter, typeSpeed);
        }
    }
    
    // Start typing effect after a short delay
    setTimeout(typeWriter, 500);
}

// Initialize contact form functionality
function initContactForm() {
    setTimeout(() => {
        const form = document.getElementById('contactForm');
        const messageContainer = document.getElementById('messageContainer');
        
        if (form) {
            form.addEventListener('submit', handleFormSubmit);
            
            // Add terminal-style focus effects to form inputs
            const inputs = form.querySelectorAll('input, textarea');
            inputs.forEach(input => {
                input.addEventListener('focus', function() {
                    this.style.borderColor = 'var(--ubuntu-orange)';
                    this.style.boxShadow = '0 0 10px rgba(233, 84, 32, 0.3)';
                });
                
                input.addEventListener('blur', function() {
                    this.style.borderColor = 'rgba(0, 255, 0, 0.3)';
                    this.style.boxShadow = 'none';
                });
                
                // Terminal-style typing effect
                input.addEventListener('input', function() {
                    simulateTerminalInput(this);
                });
            });
        }
    }, 100);
}

// Handle form submission
function handleFormSubmit(e) {
    e.preventDefault();
    
    const form = e.target;
    const messageContainer = document.getElementById('messageContainer');
    const submitBtn = form.querySelector('.btn-submit');
    
    // Get form data
    const formData = new FormData(form);
    const name = formData.get('name');
    const email = formData.get('email');
    const message = formData.get('message');
    
    // Validate form
    if (!name || !email || !message) {
        showMessage('All fields are required!', 'error', messageContainer);
        return;
    }
    
    // Disable submit button and show loading
    submitBtn.disabled = true;
    submitBtn.textContent = '$ Sending...';
    
    // Create mailto URL
    const subject = encodeURIComponent(`Portfolio Contact from ${name}`);
    const body = encodeURIComponent(
        `Name: ${name}\n` +
        `Email: ${email}\n\n` +
        `Message:\n${message}\n\n` +
        `---\n` +
        `Sent from Portfolio Contact Form`
    );
    
    const mailtoUrl = `mailto:musaiyab2003@gmail.com?subject=${subject}&body=${body}`;
    
    // Simulate sending delay for better UX
    setTimeout(() => {
        // Open default mail client
        window.location.href = mailtoUrl;
        
        // Show success message
        showMessage('Mail client opened! Please send the email to complete your message.', 'success', messageContainer);
        
        // Reset form
        form.reset();
        
        // Re-enable submit button
        submitBtn.disabled = false;
        submitBtn.textContent = '$ Send Message';
        
        // Add terminal command log
        addTerminalLog(`Email composed for musaiyab2003@gmail.com`);
        
    }, 1500);
}

// Show success/error messages
function showMessage(text, type, container) {
    // Remove existing messages
    const existingMessages = container.querySelectorAll('.message');
    existingMessages.forEach(msg => msg.remove());
    
    // Create new message
    const message = document.createElement('div');
    message.className = `message ${type}`;
    message.textContent = text;
    
    container.appendChild(message);
    
    // Trigger animation
    setTimeout(() => {
        message.classList.add('show');
    }, 100);
    
    // Auto-remove after 5 seconds
    setTimeout(() => {
        if (message.parentNode) {
            message.classList.remove('show');
            setTimeout(() => {
                if (message.parentNode) {
                    message.remove();
                }
            }, 300);
        }
    }, 5000);
}

// Add terminal log entry
function addTerminalLog(text) {
    const terminalBody = document.querySelector('.terminal-body');
    const lastCommandLine = terminalBody.querySelector('.command-line:last-child');
    
    // Create log entry
    const logEntry = document.createElement('div');
    logEntry.className = 'command-line';
    logEntry.style.color = 'var(--terminal-green)';
    logEntry.innerHTML = `<span class="prompt">[INFO]</span><span class="command">${text}</span>`;
    
    // Insert before the last command line
    terminalBody.insertBefore(logEntry, lastCommandLine);
    
    // Scroll to bottom
    terminalBody.scrollTop = terminalBody.scrollHeight;
}

// Simulate terminal input effect
function simulateTerminalInput(input) {
    input.style.color = 'var(--terminal-green)';
    
    // Add subtle glow effect while typing
    input.style.boxShadow = '0 0 5px rgba(0, 255, 0, 0.3)';
    
    // Remove glow after a short delay
    clearTimeout(input.glowTimeout);
    input.glowTimeout = setTimeout(() => {
        input.style.boxShadow = '0 0 10px rgba(233, 84, 32, 0.3)';
    }, 500);
}

// Terminal command interactions
function initTerminalCommands() {
    // Add click handlers for terminal controls
    const closeBtn = document.querySelector('.control-btn.close');
    const minimizeBtn = document.querySelector('.control-btn.minimize');
    const maximizeBtn = document.querySelector('.control-btn.maximize');
    const terminalBody = document.querySelector('.terminal-body');
    
    if (closeBtn) {
        closeBtn.addEventListener('click', function() {
            if (terminalBody) {
                terminalBody.style.display = terminalBody.style.display === 'none' ? 'block' : 'none';
            }
        });
    }
    
    if (minimizeBtn) {
        minimizeBtn.addEventListener('click', function() {
            if (terminalBody) {
                terminalBody.style.height = terminalBody.style.height === '50px' ? 'auto' : '50px';
                terminalBody.style.overflow = terminalBody.style.overflow === 'hidden' ? 'visible' : 'hidden';
            }
        });
    }
    
    if (maximizeBtn) {
        maximizeBtn.addEventListener('click', function() {
            const terminal = document.querySelector('.terminal-container');
            if (terminal) {
                terminal.classList.toggle('maximized');
                if (terminal.classList.contains('maximized')) {
                    terminal.style.position = 'fixed';
                    terminal.style.top = '0';
                    terminal.style.left = '0';
                    terminal.style.width = '100vw';
                    terminal.style.height = '100vh';
                    terminal.style.zIndex = '9999';
                    terminal.style.borderRadius = '0';
                } else {
                    terminal.style.position = 'relative';
                    terminal.style.top = 'auto';
                    terminal.style.left = 'auto';
                    terminal.style.width = 'auto';
                    terminal.style.height = 'auto';
                    terminal.style.zIndex = 'auto';
                    terminal.style.borderRadius = '8px';
                }
            }
        });
    }
}

// Enhanced terminal experience
document.addEventListener('keydown', function(e) {
    const terminalBody = document.querySelector('.terminal-body');
    
    if (terminalBody && e.key === 'Enter' && !e.target.matches('input, textarea, button')) {
        // Add new command line on Enter (only when not in form fields)
        const newCommandLine = document.createElement('div');
        newCommandLine.className = 'command-line';
        newCommandLine.innerHTML = `
            <span class="prompt">user@portfolio:~$</span>
            <span class="command"></span>
            <span class="blinking-cursor"></span>
        `;
        
        // Remove cursor from previous line
        const previousCursor = terminalBody.querySelector('.blinking-cursor');
        if (previousCursor && previousCursor.parentElement !== newCommandLine) {
            previousCursor.remove();
        }
        
        terminalBody.appendChild(newCommandLine);
        
        // Scroll to bottom
        terminalBody.scrollTop = terminalBody.scrollHeight;
    }
});

// Add terminal-style hover effects
document.addEventListener('DOMContentLoaded', function() {
    setTimeout(() => {
        const contactItems = document.querySelectorAll('.contact-item');
        const links = document.querySelectorAll('.contact-link');
        
        // Terminal-style hover effects for contact items
        contactItems.forEach(item => {
            item.addEventListener('mouseenter', function() {
                this.style.backgroundColor = 'rgba(0, 255, 0, 0.1)';
                this.style.borderLeft = '3px solid var(--ubuntu-orange)';
                this.style.paddingLeft = '15px';
                this.style.transition = 'all 0.3s ease';
            });
            
            item.addEventListener('mouseleave', function() {
                this.style.backgroundColor = 'transparent';
                this.style.borderLeft = 'none';
                this.style.paddingLeft = '10px';
            });
        });
        
        // Enhanced link interactions
        links.forEach(link => {
            link.addEventListener('click', function(e) {
                // Add terminal log for link clicks
                const linkText = this.textContent;
                addTerminalLog(`Opening: ${linkText}`);
            });
        });
    }, 200);
});

// Performance optimization: Throttle scroll events
function throttle(func, wait) {
    let timeout;
    return function executedFunction(...args) {
        const later = () => {
            clearTimeout(timeout);
            func(...args);
        };
        clearTimeout(timeout);
        timeout = setTimeout(later, wait);
    };
}

// Apply throttling to scroll events for better performance
const throttledScrollHandler = throttle(() => {
    // Terminal-specific scroll effects can go here
}, 16); // ~60fps

window.addEventListener('scroll', throttledScrollHandler);

// Add terminal-style focus effects for accessibility
const terminalFocusStyle = document.createElement('style');
terminalFocusStyle.textContent = `
    .terminal-body *:focus {
        outline: 2px solid var(--terminal-green);
        outline-offset: 2px;
    }
    
    .control-btn:focus {
        outline: 2px solid white;
        outline-offset: 2px;
    }
    
    .contact-link:focus {
        outline: 2px solid var(--ubuntu-orange);
        outline-offset: 2px;
    }
`;
document.head.appendChild(terminalFocusStyle);