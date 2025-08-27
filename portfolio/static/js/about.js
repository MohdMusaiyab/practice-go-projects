// About Page JavaScript - Terminal Theme
document.addEventListener('DOMContentLoaded', function() {
    // Initialize all functionality
    initTerminalStructure();
    initScrollAnimations();
    initTypingEffect();
    initTerminalCommands();
});

// Create terminal structure and wrap content
function initTerminalStructure() {
    const aboutSection = document.querySelector('.about-section');
    const existingContent = aboutSection.querySelector('.about-content');
    
    if (existingContent && !aboutSection.querySelector('.terminal-container')) {
        // Store the original content
        const content = existingContent.innerHTML;
        
        // Create terminal structure
        const terminalHTML = `
            <div class="terminal-container">
                <div class="terminal-header">
                    <div class="terminal-controls">
                        <span class="control-btn close"></span>
                        <span class="control-btn minimize"></span>
                        <span class="control-btn maximize"></span>
                    </div>
                    <div class="terminal-title">musaiyab@portfolio:~/about$</div>
                </div>
                <div class="terminal-body">
                    <div class="command-line">
                        <span class="prompt">user@portfolio:~$</span>
                        <span class="command">cat about.txt</span>
                    </div>
                    <div class="about-content">
                        ${content}
                    </div>
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
        aboutSection.insertAdjacentHTML('beforeend', terminalHTML);
    }
}

// Scroll-based animations
function initScrollAnimations() {
    const observerOptions = {
        threshold: 0.1,
        rootMargin: '0px 0px -50px 0px'
    };

    const observer = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                entry.target.classList.add('animate');
                // Add staggered delay for multiple elements
                const delay = Array.from(entry.target.parentNode.children).indexOf(entry.target) * 100;
                setTimeout(() => {
                    entry.target.style.animationDelay = delay + 'ms';
                }, delay);
            }
        });
    }, observerOptions);

    // Observe all content elements for scroll animations
    setTimeout(() => {
        const elements = document.querySelectorAll('.about-content p, .about-content h2');
        elements.forEach(el => {
            el.classList.add('scroll-animate');
            observer.observe(el);
        });
    }, 100);
}

// Terminal typing effect for the main heading
function initTypingEffect() {
    const heading = document.querySelector('.about-section h1');
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

// Add interactive hover effects
document.addEventListener('DOMContentLoaded', function() {
    setTimeout(() => {
        const paragraphs = document.querySelectorAll('.about-content p');
        const headings = document.querySelectorAll('.about-content h2');
        
        // Terminal-style hover effects for paragraphs
        paragraphs.forEach(p => {
            p.addEventListener('mouseenter', function() {
                this.style.backgroundColor = 'rgba(0, 255, 0, 0.1)';
                this.style.padding = '5px';
                this.style.borderRadius = '3px';
                this.style.transition = 'all 0.3s ease';
            });
            
            p.addEventListener('mouseleave', function() {
                this.style.backgroundColor = 'transparent';
                this.style.padding = '0';
            });
        });
        
        // Interactive headings with terminal effects
        headings.forEach(h => {
            h.addEventListener('click', function() {
                // Add a terminal-style blink effect
                this.style.animation = 'blink 0.6s ease 3';
                setTimeout(() => {
                    this.style.animation = '';
                }, 1800);
            });
        });
    }, 200);
});

// Simulate terminal typing for commands
function simulateTyping(element, text, speed = 50) {
    element.textContent = '';
    let i = 0;
    
    function type() {
        if (i < text.length) {
            element.textContent += text.charAt(i);
            i++;
            setTimeout(type, speed);
        }
    }
    
    type();
}

// Enhanced terminal experience
document.addEventListener('keydown', function(e) {
    const terminalBody = document.querySelector('.terminal-body');
    
    if (terminalBody && e.key === 'Enter') {
        // Add new command line on Enter
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
`;
document.head.appendChild(terminalFocusStyle);