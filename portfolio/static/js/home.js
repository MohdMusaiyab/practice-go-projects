// Load header and footer
document.addEventListener('DOMContentLoaded', function() {
    // Load header
    fetch('/static/html/header.html')
        .then(response => response.text())
        .then(html => {
            document.getElementById('header-container').innerHTML = html;
        })
        .catch(err => console.log('Header load failed:', err));

    // Load footer
    fetch('/static/html/footer.html')
        .then(response => response.text())
        .then(html => {
            document.getElementById('footer-container').innerHTML = html;
        })
        .catch(err => console.log('Footer load failed:', err));

    // Initialize page animations
    initializeAnimations();
});

function initializeAnimations() {
    // Simple typewriter effect for command
    const commandElement = document.querySelector('.command');
    if (commandElement) {
        const originalText = commandElement.textContent;
        commandElement.textContent = '';
        
        let i = 0;
        const typeInterval = setInterval(() => {
            if (i < originalText.length) {
                commandElement.textContent += originalText.charAt(i);
                i++;
            } else {
                clearInterval(typeInterval);
            }
        }, 100);
    }

    // Fade in intro card
    const introCard = document.querySelector('.intro-card');
    if (introCard) {
        introCard.style.opacity = '0';
        introCard.style.transform = 'translateY(30px)';
        introCard.style.transition = 'all 0.8s ease-out';
        
        setTimeout(() => {
            introCard.style.opacity = '1';
            introCard.style.transform = 'translateY(0)';
        }, 500);
    }
}

// Tech badge hover effects
document.addEventListener('click', function(e) {
    if (e.target.classList.contains('tech-badge')) {
        const badge = e.target;
        const originalText = badge.textContent;
        
        badge.textContent = '✓ Selected';
        badge.style.background = 'var(--terminal-green)';
        badge.style.color = 'var(--ubuntu-dark)';
        
        setTimeout(() => {
            badge.textContent = originalText;
            badge.style.background = 'var(--ubuntu-orange)';
            badge.style.color = 'white';
        }, 1000);
    }
});

// Simple terminal blink effect
setInterval(() => {
    const prompt = document.querySelector('.prompt');
    if (prompt) {
        prompt.style.opacity = prompt.style.opacity === '0.5' ? '1' : '0.5';
    }
}, 1000);