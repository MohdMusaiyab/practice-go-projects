window.addEventListener("load", function () {
  setTimeout(() => {
    document.getElementById("bootSequence").classList.add("hidden");
    setTimeout(() => {
      document.getElementById("bootSequence").style.display = "none";
    }, 500);
  }, 3500);
});

// Terminal command typing animation
function typeCommand(element, commands, index = 0) {
  if (index >= commands.length) {
    // Restart animation
    setTimeout(() => typeCommand(element, commands, 0), 3000);
    return;
  }

  const command = commands[index];
  element.textContent = "";
  let charIndex = 0;

  const typeInterval = setInterval(() => {
    if (charIndex < command.length) {
      element.textContent += command[charIndex];
      charIndex++;
    } else {
      clearInterval(typeInterval);
      setTimeout(() => typeCommand(element, commands, index + 1), 2000);
    }
  }, 80);
}

// Start command typing after boot
setTimeout(() => {
  const commandLine = document.getElementById("commandLine");
  const commands = [
    "./run_portfolio.sh --mode=awesome",
    "sudo apt update && apt upgrade skills",
    "git push origin main",
    "npm run build:future",
  ];
  typeCommand(commandLine, commands);
}, 4000);

// Skill package interactions
document.querySelectorAll(".skill-package").forEach((pkg) => {
  pkg.addEventListener("click", function () {
    const packageName = this.getAttribute("data-package");
    showPackageInfo(packageName, this);
  });

  pkg.addEventListener("mouseenter", function () {
    this.style.transform = "translateY(-5px) scale(1.02)";
  });

  pkg.addEventListener("mouseleave", function () {
    this.style.transform = "translateY(0) scale(1)";
  });
});

// Package info display
function showPackageInfo(packageName, element) {
  const descriptions = {
    html5:
      "Status: Active\nMemory Usage: 2.4MB\nDescription: Foundation of web development",
    css3: "Status: Running\nMemory Usage: 1.8MB\nDescription: Styling and layout engine",
    javascript:
      "Status: Active\nMemory Usage: 15.2MB\nDescription: Dynamic web programming",
    golang:
      "Status: Optimized\nMemory Usage: 8.1MB\nDescription: High-performance backend",
    react:
      "Status: Rendering\nMemory Usage: 12.5MB\nDescription: Modern UI framework",
    nodejs:
      "Status: Listening\nMemory Usage: 22.3MB\nDescription: JavaScript server runtime",
  };

  // Create terminal popup
  const popup = document.createElement("div");
  popup.style.cssText = `
                position: fixed;
                top: 50%;
                left: 50%;
                transform: translate(-50%, -50%);
                background: var(--terminal-bg);
                color: var(--terminal-green);
                padding: 20px;
                border: 2px solid var(--ubuntu-orange);
                border-radius: 8px;
                font-family: 'Ubuntu Mono', monospace;
                white-space: pre-line;
                z-index: 1000;
                box-shadow: 0 0 30px rgba(233, 84, 32, 0.5);
                min-width: 300px;
            `;

  popup.innerHTML = `
                <div style="color: var(--ubuntu-orange); margin-bottom: 10px;">$ apt show ${packageName}</div>
                <div>${descriptions[packageName]}</div>
                <div style="margin-top: 15px; color: var(--warm-grey);">Press ESC to close</div>
            `;

  document.body.appendChild(popup);

  // Close on ESC or click outside
  function closePopup(e) {
    if (e.key === "Escape" || !popup.contains(e.target)) {
      popup.remove();
      document.removeEventListener("keydown", closePopup);
      document.removeEventListener("click", closePopup);
    }
  }

  setTimeout(() => {
    document.addEventListener("keydown", closePopup);
    document.addEventListener("click", closePopup);
  }, 100);
}

// Contact email interaction
document.getElementById("contactInfo").addEventListener("click", function () {
  this.innerHTML = "Email copied to clipboard! 📋";
  this.style.background = "var(--terminal-green)";
  this.style.color = "var(--terminal-bg)";

  setTimeout(() => {
    this.innerHTML = "> musaiyab@example.com";
    this.style.background = "var(--terminal-bg)";
    this.style.color = "var(--terminal-green)";
  }, 2000);

  // Simulate clipboard copy
  const tempInput = document.createElement("input");
  tempInput.value = "musaiyab@example.com";
  document.body.appendChild(tempInput);
  tempInput.select();
  document.execCommand("copy");
  document.body.removeChild(tempInput);
});

// Window control interactions
document.querySelectorAll(".control-btn, .terminal-btn").forEach((btn) => {
  btn.addEventListener("click", function (e) {
    e.preventDefault();

    if (this.classList.contains("close")) {
      const window = this.closest(".ubuntu-window") || this.closest("header");
      if (window) {
        window.style.animation = "windowSlideOut 0.5s ease-in forwards";
        setTimeout(() => {
          window.style.display = "none";
          // Show restoration hint
          showRestoreHint();
        }, 500);
      }
    } else if (this.classList.contains("minimize")) {
      const window = this.closest(".ubuntu-window") || this.closest("header");
      if (window) {
        window.style.transform = "scale(0.1)";
        window.style.opacity = "0.5";
        setTimeout(() => {
          window.style.transform = "scale(1)";
          window.style.opacity = "1";
        }, 1000);
      }
    }
  });
});

function showRestoreHint() {
  const hint = document.createElement("div");
  hint.style.cssText = `
                position: fixed;
                bottom: 20px;
                right: 20px;
                background: var(--ubuntu-orange);
                color: white;
                padding: 10px 15px;
                border-radius: 5px;
                font-family: 'Ubuntu Mono', monospace;
                z-index: 1000;
            `;
  hint.textContent = "Ctrl+Z to restore windows";
  document.body.appendChild(hint);

  setTimeout(() => hint.remove(), 3000);
}

// Window slide out animation
const style = document.createElement("style");
style.textContent = `
            @keyframes windowSlideOut {
                to {
                    transform: translateY(-100px);
                    opacity: 0;
                }
            }
        `;
document.head.appendChild(style);

// Keyboard shortcuts
document.addEventListener("keydown", function (e) {
  if (e.ctrlKey && e.key === "z") {
    // Restore hidden windows
    document
      .querySelectorAll(
        '.ubuntu-window[style*="display: none"], header[style*="display: none"]'
      )
      .forEach((window) => {
        window.style.display = "";
        window.style.animation = "windowSlideIn 0.5s ease-out";
      });
  }
});

// Smooth scrolling
document.documentElement.style.scrollBehavior = "smooth";

// Add subtle background animations
setTimeout(() => {
  const floatingLogos = document.querySelectorAll(".floating-logo");
  floatingLogos.forEach((logo, index) => {
    logo.addEventListener("mouseenter", function () {
      this.style.transform = "scale(1.5) rotate(360deg)";
      this.style.opacity = "0.3";
    });

    logo.addEventListener("mouseleave", function () {
      this.style.transform = "scale(1) rotate(0deg)";
      this.style.opacity = "0.1";
    });
  });
}, 5000);

// Terminal cursor blink effect
setInterval(() => {
  const commandLine = document.getElementById("commandLine");
  if (commandLine) {
    commandLine.style.borderRight =
      commandLine.style.borderRight === "2px solid transparent"
        ? "2px solid var(--terminal-green)"
        : "2px solid transparent";
  }
}, 500);

// Add matrix rain effect (subtle)
function createMatrixRain() {
  const matrix = document.createElement("div");
  matrix.style.cssText = `
                position: fixed;
                top: 0;
                left: ${Math.random() * 100}%;
                color: var(--terminal-green);
                font-family: 'Ubuntu Mono', monospace;
                font-size: 14px;
                opacity: 0.1;
                z-index: -1;
                pointer-events: none;
            `;

  const chars = "01";
  matrix.textContent = chars[Math.floor(Math.random() * chars.length)];
  document.body.appendChild(matrix);

  let position = 0;
  const fall = setInterval(() => {
    position += 2;
    matrix.style.top = position + "px";

    if (position > window.innerHeight) {
      clearInterval(fall);
      matrix.remove();
    }
  }, 50);
}

// Occasional matrix rain
setInterval(createMatrixRain, 2000);

// System notification simulation
function showNotification(message, type = "info") {
  const notification = document.createElement("div");
  notification.style.cssText = `
                position: fixed;
                top: 20px;
                right: 20px;
                background: var(--terminal-bg);
                color: var(--terminal-green);
                border: 2px solid var(--ubuntu-orange);
                border-radius: 6px;
                padding: 15px 20px;
                font-family: 'Ubuntu Mono', monospace;
                z-index: 1000;
                max-width: 300px;
                animation: slideInFromRight 0.5s ease-out;
            `;

  notification.innerHTML = `
                <div style="color: var(--ubuntu-orange); margin-bottom: 5px;">System Notification</div>
                <div>${message}</div>
            `;

  document.body.appendChild(notification);

  setTimeout(() => {
    notification.style.animation = "slideOutToRight 0.5s ease-in";
    setTimeout(() => notification.remove(), 500);
  }, 3000);
}

// Add notification animations
const notificationStyle = document.createElement("style");
notificationStyle.textContent = `
            @keyframes slideInFromRight {
                from {
                    transform: translateX(100%);
                    opacity: 0;
                }
                to {
                    transform: translateX(0);
                    opacity: 1;
                }
            }
            
            @keyframes slideOutToRight {
                to {
                    transform: translateX(100%);
                    opacity: 0;
                }
            }
        `;
document.head.appendChild(notificationStyle);

// Show welcome notification after boot
setTimeout(() => {
  showNotification("Portfolio system initialized successfully!");
}, 4500);

// Random system messages
const systemMessages = [
  "Background processes optimized",
  "Cache cleared successfully",
  "All systems operational",
  "Performance: Excellent",
  "Memory usage: Optimal",
];

setInterval(() => {
  if (Math.random() < 0.3) {
    // 30% chance
    const message =
      systemMessages[Math.floor(Math.random() * systemMessages.length)];
    showNotification(message);
  }
}, 30000); // Every 30 seconds

// Add Ubuntu-style loading spinner
function createLoadingSpinner() {
  const spinner = document.createElement("div");
  spinner.style.cssText = `
                width: 20px;
                height: 20px;
                border: 2px solid var(--warm-grey);
                border-top: 2px solid var(--ubuntu-orange);
                border-radius: 50%;
                animation: spin 1s linear infinite;
                display: inline-block;
                margin-right: 10px;
            `;
  return spinner;
}

const spinnerStyle = document.createElement("style");
spinnerStyle.textContent = `
            @keyframes spin {
                0% { transform: rotate(0deg); }
                100% { transform: rotate(360deg); }
            }
        `;
document.head.appendChild(spinnerStyle);

// Add context menu simulation
document.addEventListener("contextmenu", function (e) {
  e.preventDefault();

  const contextMenu = document.createElement("div");
  contextMenu.style.cssText = `
                position: fixed;
                top: ${e.clientY}px;
                left: ${e.clientX}px;
                background: var(--terminal-bg);
                border: 1px solid var(--ubuntu-orange);
                border-radius: 4px;
                padding: 5px 0;
                z-index: 1000;
                font-family: 'Ubuntu', sans-serif;
                font-size: 14px;
                box-shadow: 0 4px 12px rgba(0,0,0,0.3);
            `;

  contextMenu.innerHTML = `
                <div style="padding: 8px 15px; color: var(--ubuntu-white); cursor: pointer; hover: background: var(--ubuntu-orange);">Open Terminal</div>
                <div style="padding: 8px 15px; color: var(--ubuntu-white); cursor: pointer;">View Source</div>
                <div style="padding: 8px 15px; color: var(--ubuntu-white); cursor: pointer;">Inspect Element</div>
                <hr style="margin: 5px 0; border-color: var(--ubuntu-orange);">
                <div style="padding: 8px 15px; color: var(--warm-grey); cursor: pointer;">Properties</div>
            `;

  document.body.appendChild(contextMenu);

  // Close context menu on click outside
  setTimeout(() => {
    document.addEventListener("click", function closeMenu(e) {
      if (!contextMenu.contains(e.target)) {
        contextMenu.remove();
        document.removeEventListener("click", closeMenu);
      }
    });
  }, 100);

  // Auto-close after 5 seconds
  setTimeout(() => {
    if (contextMenu.parentNode) {
      contextMenu.remove();
    }
  }, 5000);
});
