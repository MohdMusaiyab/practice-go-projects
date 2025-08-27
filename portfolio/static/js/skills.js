document.addEventListener('DOMContentLoaded', function() {
            // Typewriter effect for the command
            const commandElement = document.querySelector('.command');
            const originalText = "cat skills.txt";
            let i = 0;
            
            function typeWriter() {
                if (i < originalText.length) {
                    commandElement.textContent += originalText.charAt(i);
                    i++;
                    setTimeout(typeWriter, 100);
                } else {
                    // After typing is done, animate skill bars
                    setTimeout(animateSkillBars, 500);
                }
            }
            
            // Start typing effect
            typeWriter();
            
            // Animate skill bars
            function animateSkillBars() {
                const skillLevels = document.querySelectorAll('.skill-level');
                
                skillLevels.forEach((level, index) => {
                    setTimeout(() => {
                        const width = level.getAttribute('data-level') + '%';
                        level.style.width = width;
                    }, index * 200);
                });
                
                // Start blinking cursor at the end
                setTimeout(() => {
                    document.querySelector('.blinking-cursor').style.display = 'inline-block';
                }, skillLevels.length * 200 + 500);
            }
            
            // Add hover effect to skill bars
            const skillItems = document.querySelectorAll('.skill-item');
            skillItems.forEach(item => {
                item.addEventListener('mouseenter', function() {
                    this.style.backgroundColor = 'rgba(0, 255, 0, 0.1)';
                    this.style.transform = 'translateX(5px)';
                    this.style.transition = 'all 0.3s ease';
                });
                
                item.addEventListener('mouseleave', function() {
                    this.style.backgroundColor = '';
                    this.style.transform = '';
                });
            });
            
            // Make terminal controls interactive
            const closeBtn = document.querySelector('.control-btn.close');
            const minimizeBtn = document.querySelector('.control-btn.minimize');
            const maximizeBtn = document.querySelector('.control-btn.maximize');
            
            closeBtn.addEventListener('click', function() {
                alert('Closing terminal... Just kidding! This is a simulation.');
            });
            
            minimizeBtn.addEventListener('click', function() {
                alert('Minimizing terminal... This is a simulation.');
            });
            
            maximizeBtn.addEventListener('click', function() {
                document.querySelector('.terminal-container').classList.toggle('maximized');
                if (document.querySelector('.terminal-container').classList.contains('maximized')) {
                    document.querySelector('.terminal-container').style.width = '100%';
                    document.querySelector('.terminal-container').style.maxWidth = '100%';
                } else {
                    document.querySelector('.terminal-container').style.width = '';
                    document.querySelector('.terminal-container').style.maxWidth = '';
                }
            });
        });