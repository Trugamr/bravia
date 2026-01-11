// Utility Functions
function showToast(message, type = 'success') {
    const toast = document.getElementById('toast');
    toast.textContent = message;
    toast.className = `toast ${type} show`;

    // Clear after 3 seconds
    setTimeout(() => {
        toast.classList.remove('show');
    }, 3000);
}

function showStatus(message, isError = false) {
    showToast(message, isError ? 'error' : 'success');
}

async function apiCall(url, options = {}) {
    try {
        const response = await fetch(url, {
            headers: {
                'Content-Type': 'application/json',
                ...options.headers
            },
            ...options
        });

        const data = await response.json();

        if (!response.ok) {
            throw new Error(data.error || 'Request failed');
        }

        return data;
    } catch (error) {
        showStatus(`Error: ${error.message}`, true);
        throw error;
    }
}

// IRCC Command Handler
async function sendIRCC(command, buttonElement = null) {
    try {
        // Add visual feedback
        if (buttonElement) {
            buttonElement.classList.add('btn-loading');
        }

        await apiCall('/api/ircc/send', {
            method: 'POST',
            body: JSON.stringify({ command })
        });

        // Remove loading state
        if (buttonElement) {
            setTimeout(() => {
                buttonElement.classList.remove('btn-loading');
            }, 300);
        }
    } catch (error) {
        console.error('Failed to send IRCC command:', error);
        if (buttonElement) {
            buttonElement.classList.remove('btn-loading');
        }
    }
}

// Add click listeners to all IRCC buttons
document.querySelectorAll('[data-ircc]').forEach(button => {
    button.addEventListener('click', function() {
        const command = button.getAttribute('data-ircc');
        sendIRCC(command, this);
    });
});

// Fallback icon for apps without proper icon URL
function getFallbackIcon() {
    return '<svg class="w-8 h-8 text-slate-400" viewBox="0 0 24 24" fill="currentColor"><path d="M4 8h4V4H4v4zm6 12h4v-4h-4v4zm-6 0h4v-4H4v4zm0-6h4v-4H4v4zm6 0h4v-4h-4v4zm6-10v4h4V4h-4zm-6 4h4V4h-4v4zm6 6h4v-4h-4v4zm0 6h4v-4h-4v4z"></path></svg>';
}

// Apps loading
async function loadApps() {
    const appsGrid = document.getElementById('apps-grid');
    const loadButton = document.getElementById('load-apps');

    // Show loading skeleton
    appsGrid.innerHTML = Array(6).fill(0).map(() => `
        <div class="btn bg-slate-100 py-4 flex-col border border-slate-200 animate-pulse">
            <div class="w-8 h-8 bg-slate-200 rounded"></div>
            <div class="w-16 h-3 bg-slate-200 rounded mt-1"></div>
        </div>
    `).join('');

    loadButton.style.display = 'none';

    try {
        const result = await apiCall('/api/apps');
        const apps = result.data;

        appsGrid.innerHTML = '';

        if (apps.length === 0) {
            appsGrid.innerHTML = '<div class="col-span-full text-center py-4 text-slate-500 text-sm">No apps found</div>';
            loadButton.style.display = 'block';
            return;
        }

        apps.forEach((app, index) => {
            const appEl = document.createElement('button');
            appEl.className = 'btn bg-gradient-to-br from-slate-50 to-slate-100 hover:from-slate-100 hover:to-slate-200 active:from-slate-200 active:to-slate-300 text-slate-700 py-4 flex-col border border-slate-200';

            // Create image element for app icon
            const iconEl = document.createElement('img');
            iconEl.className = 'app-icon';
            iconEl.alt = app.title;

            // Use icon from API if available, otherwise show fallback
            if (app.icon) {
                iconEl.src = app.icon;
                iconEl.onerror = () => {
                    // Replace failed image with fallback SVG
                    const container = iconEl.parentElement;
                    if (container) {
                        iconEl.remove();
                        const fallback = document.createElement('div');
                        fallback.innerHTML = getFallbackIcon();
                        container.insertBefore(fallback.firstElementChild, container.firstChild);
                    }
                };
            } else {
                // No icon provided, use fallback SVG
                const fallbackContainer = document.createElement('div');
                fallbackContainer.innerHTML = getFallbackIcon();
                appEl.appendChild(fallbackContainer.firstElementChild);
            }

            const titleEl = document.createElement('span');
            titleEl.className = 'text-xs font-medium mt-1 line-clamp-2 text-center';
            titleEl.textContent = app.title;

            if (app.icon) {
                appEl.appendChild(iconEl);
            }
            appEl.appendChild(titleEl);

            appEl.addEventListener('click', async function() {
                try {
                    this.classList.add('btn-loading');
                    await apiCall('/api/apps/open', {
                        method: 'POST',
                        body: JSON.stringify({ uri: app.uri })
                    });
                    showToast(`Opened ${app.title}`, 'success');
                    setTimeout(() => this.classList.remove('btn-loading'), 300);
                } catch (error) {
                    console.error('Failed to open app:', error);
                    this.classList.remove('btn-loading');
                }
            });

            appsGrid.appendChild(appEl);
        });

        showToast(`Loaded ${apps.length} apps`, 'success');
    } catch (error) {
        appsGrid.innerHTML = '<div class="col-span-full text-center py-4 text-red-500 text-sm">Failed to load apps. <button class="text-blue-600 underline" onclick="loadApps()">Retry</button></div>';
        console.error('Failed to load apps:', error);
    }
}

// Auto-load apps on page load
document.getElementById('load-apps').addEventListener('click', loadApps);
setTimeout(loadApps, 500); // Auto-load after page loads

// Input icon helper
function getInputIcon(inputTitle) {
    const title = inputTitle.toLowerCase();

    if (title.includes('hdmi')) {
        return '<svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="2" y="7" width="20" height="15" rx="2" ry="2"></rect><polyline points="17 2 12 7 7 2"></polyline></svg>';
    } else if (title.includes('component') || title.includes('composite')) {
        return '<svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"></circle><line x1="12" y1="8" x2="12" y2="16"></line><line x1="8" y1="12" x2="16" y2="12"></line></svg>';
    } else if (title.includes('tv')) {
        return '<svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="2" y="3" width="20" height="14" rx="2"></rect><line x1="8" y1="21" x2="16" y2="21"></line><line x1="12" y1="17" x2="12" y2="21"></line></svg>';
    } else {
        return '<svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><rect x="3" y="3" width="18" height="18" rx="2"></rect><circle cx="8.5" cy="8.5" r="1.5"></circle><polyline points="21 15 16 10 5 21"></polyline></svg>';
    }
}

// Inputs
document.getElementById('load-inputs').addEventListener('click', async () => {
    const inputsGrid = document.getElementById('inputs-grid');
    inputsGrid.innerHTML = '<div class="text-center py-4 text-slate-500 text-sm">Loading inputs...</div>';

    try {
        const result = await apiCall('/api/inputs');
        const inputs = result.data;

        inputsGrid.innerHTML = '';

        if (inputs.length === 0) {
            inputsGrid.innerHTML = '<div class="text-center py-4 text-slate-500 text-sm">No inputs found</div>';
            return;
        }

        inputs.forEach((input, index) => {
            const inputEl = document.createElement('button');
            const label = input.label ? ` (${input.label})` : '';
            const displayText = `${input.title}${label}`;

            inputEl.className = 'btn bg-gradient-to-br from-purple-50 to-purple-100 hover:from-purple-100 hover:to-purple-200 active:from-purple-200 active:to-purple-300 text-purple-700 py-3 px-4 justify-start border border-purple-200';
            inputEl.innerHTML = `
                ${getInputIcon(input.title)}
                <span class="text-sm font-medium">${displayText}</span>
            `;

            inputEl.addEventListener('click', async function() {
                try {
                    this.classList.add('btn-loading');
                    await apiCall('/api/inputs/select', {
                        method: 'POST',
                        body: JSON.stringify({ uri: input.uri })
                    });
                    showToast(`Selected ${input.title}`, 'success');
                    setTimeout(() => this.classList.remove('btn-loading'), 300);
                } catch (error) {
                    console.error('Failed to select input:', error);
                    this.classList.remove('btn-loading');
                }
            });

            inputsGrid.appendChild(inputEl);
        });

        showToast(`Loaded ${inputs.length} inputs`, 'success');
    } catch (error) {
        inputsGrid.innerHTML = '<div class="text-center py-4 text-red-500 text-sm">Failed to load inputs</div>';
        console.error('Failed to load inputs:', error);
    }
});

// Keyboard shortcuts
document.addEventListener('keydown', (e) => {
    // Prevent shortcuts when typing in input fields
    if (e.target.tagName === 'INPUT') return;

    const irccCommands = {
        'ArrowUp': 'AAAAAQAAAAEAAAB0Aw==',
        'ArrowDown': 'AAAAAQAAAAEAAAB1Aw==',
        'ArrowLeft': 'AAAAAQAAAAEAAAA0Aw==',
        'ArrowRight': 'AAAAAQAAAAEAAAAzAw==',
        'Enter': 'AAAAAQAAAAEAAABlAw==',
        ' ': 'AAAAAgAAAJcAAAAaAw==', // Space = Play/Pause
    };

    if (irccCommands[e.key]) {
        e.preventDefault();
        sendIRCC(irccCommands[e.key]);
    }
});

// SSE Connection for real-time TV state updates
function connectSSE() {
    const eventSource = new EventSource('/api/sse');

    eventSource.onopen = () => {
        console.log('SSE connection established');
    };

    eventSource.onmessage = (event) => {
        try {
            const data = JSON.parse(event.data);

            // Skip connection messages
            if (data.type === 'connected') return;

            // Update header status bar
            const statusEl = document.getElementById('status');
            const isActive = data.powerStatus === 'active';
            const powerColor = isActive ? 'text-green-600' : 'text-slate-400';
            const volumeColor = data.muted ? 'text-red-600' : 'text-slate-700';

            statusEl.innerHTML = `
                <div class="flex items-center gap-4">
                    <div class="flex items-center gap-1.5 ${powerColor}">
                        <svg class="w-3.5 h-3.5" viewBox="0 0 24 24" fill="currentColor">
                            <circle cx="12" cy="12" r="10"></circle>
                        </svg>
                        <span class="text-xs font-semibold">${isActive ? 'ON' : 'OFF'}</span>
                    </div>
                    <div class="flex items-center gap-1.5 ${volumeColor}">
                        <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                            ${data.muted
                                ? '<polygon points="11 5 6 9 2 9 2 15 6 15 11 19 11 5"></polygon><line x1="23" y1="9" x2="17" y2="15"></line><line x1="17" y1="9" x2="23" y2="15"></line>'
                                : '<polygon points="11 5 6 9 2 9 2 15 6 15 11 19 11 5"></polygon><path d="M15.54 8.46a5 5 0 0 1 0 7.07"></path>'
                            }
                        </svg>
                        <span class="text-xs font-semibold">${data.volume}</span>
                    </div>
                </div>
            `;

        } catch (error) {
            console.error('Error parsing SSE data:', error);
        }
    };

    eventSource.onerror = (error) => {
        console.error('SSE error:', error);
        eventSource.close();

        // Reconnect after 5 seconds
        setTimeout(() => {
            connectSSE();
        }, 5000);
    };

    return eventSource;
}

// Start SSE connection
connectSSE();
