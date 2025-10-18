<script setup lang="ts">
import { ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

const router = useRouter()

// API Configuration
// Backend endpoint: POST /api/corrosion/analyze
// Request body: { material: string, temperature: number, ph: number, pressure: number, velocity: number }
// Response: { success: boolean, response: { riskLevel: "HIGH" | "MEDIUM" | "LOW", corrosionRate: number,
//             mechanisms: string[], recommendations: string[], estimatedLife: string } }

// Form data
const material = ref('')
const temperature = ref(25)
const ph = ref(7)
const pressure = ref(10)
const velocity = ref(2)

// Display values for sliders
const temperatureDisplay = ref('25°C')
const phDisplay = ref('7.0')
const pressureDisplay = ref('10 bar')
const velocityDisplay = ref('2.0 m/s')

// State management
const loading = ref(false)
const error = ref('')
const results = ref<any>(null)

// Update slider display values
const updateTemperature = () => {
  temperatureDisplay.value = `${temperature.value}°C`
}

const updatePh = () => {
  phDisplay.value = ph.value.toFixed(1)
}

const updatePressure = () => {
  pressureDisplay.value = `${pressure.value} bar`
}

const updateVelocity = () => {
  velocityDisplay.value = `${velocity.value.toFixed(1)} m/s`
}

// Get risk level color
const getRiskColor = (level: string): string => {
  switch (level) {
    case 'HIGH': return 'var(--error-red)'
    case 'MEDIUM': return 'var(--accent-orange)'
    case 'LOW': return 'var(--success-green)'
    default: return 'var(--text-secondary)'
  }
}

// Handle form submission
const handleSubmit = async (e: Event) => {
  e.preventDefault()

  if (!material.value) {
    error.value = 'Please select a material type'
    return
  }

  loading.value = true
  error.value = ''
  results.value = null

  try {
    const response = await axios.post('/api/corrosion/analyze', {
      material: material.value,
      temperature: temperature.value,
      ph: ph.value,
      pressure: pressure.value,
      velocity: velocity.value
    })

    if (response.data.success) {
      results.value = response.data.response
    } else {
      error.value = 'Failed to analyze corrosion risk'
    }
  } catch (err: any) {
    error.value = err.response?.data?.error || 'Failed to connect to analysis service'
    console.error('Corrosion analysis error:', err)
  } finally {
    loading.value = false
  }
}

// Navigate back to home
const goHome = () => {
  router.push('/')
}
</script>

<template>
    <head>
        <title>CorrosionAI - PCST-AI Portal</title>
        <link href="https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700&family=JetBrains+Mono:wght@400;500&family=Orbitron:wght@400;700;900&display=swap" rel="stylesheet">
    </head>

    <body>
        <!-- Background Effects -->
        <div class="bg-grid"></div>
        
        <!-- Navigation -->
        <nav class="navbar">
            <div class="nav-container">
                <a href="/" class="logo">PCST-AI</a>
                <div id="breadcrumbs" class="breadcrumbs">
                    <a href="/">Home</a> / <span>CorrosionAI</span>
                </div>
            </div>
        </nav>
        
        <!-- Main Content -->
        <main class="main-content">
            <!-- Page Header -->
            <div class="page-header">
                <h1 class="page-title">CorrosionAI</h1>
                <p class="page-subtitle">
                    Advanced AI-powered corrosion risk assessment using process data, material properties, and environmental conditions to predict corrosion rates and optimize maintenance strategies.
                </p>
            </div>
            
            <!-- Application Interface -->
            <div class="app-interface">
                <!-- Input Section -->
                <div class="input-section">
                    <h2 class="section-title">
                        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                            <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2z"/>
                            <path d="M12 6v6l4 2"/>
                        </svg>
                        Process Parameters
                    </h2>
                    
                    <form @submit.prevent="handleSubmit">
                        <div class="form-group">
                            <label for="material" class="form-label">Material Type</label>
                            <select id="material" v-model="material" class="form-select" required>
                                <option value="">Select Material</option>
                                <option value="carbon-steel">Carbon Steel</option>
                                <option value="stainless-316">Stainless Steel 316</option>
                                <option value="stainless-304">Stainless Steel 304</option>
                                <option value="inconel-625">Inconel 625</option>
                                <option value="inconel-825">Inconel 825</option>
                                <option value="duplex-2205">Duplex 2205</option>
                                <option value="super-duplex">Super Duplex</option>
                                <option value="titanium">Titanium</option>
                            </select>
                        </div>

                        <div class="form-group">
                            <label for="temperature" class="form-label">Operating Temperature</label>
                            <div class="slider-container">
                                <input
                                    type="range"
                                    id="temperature"
                                    v-model.number="temperature"
                                    @input="updateTemperature"
                                    class="slider"
                                    min="-50"
                                    max="200"
                                >
                                <span class="slider-value">{{ temperatureDisplay }}</span>
                            </div>
                        </div>

                        <div class="form-group">
                            <label for="ph" class="form-label">pH Level</label>
                            <div class="slider-container">
                                <input
                                    type="range"
                                    id="ph"
                                    v-model.number="ph"
                                    @input="updatePh"
                                    class="slider"
                                    min="0"
                                    max="14"
                                    step="0.1"
                                >
                                <span class="slider-value">{{ phDisplay }}</span>
                            </div>
                        </div>

                        <div class="form-group">
                            <label for="pressure" class="form-label">Operating Pressure (bar)</label>
                            <div class="slider-container">
                                <input
                                    type="range"
                                    id="pressure"
                                    v-model.number="pressure"
                                    @input="updatePressure"
                                    class="slider"
                                    min="1"
                                    max="100"
                                >
                                <span class="slider-value">{{ pressureDisplay }}</span>
                            </div>
                        </div>

                        <div class="form-group">
                            <label for="velocity" class="form-label">Fluid Velocity (m/s)</label>
                            <div class="slider-container">
                                <input
                                    type="range"
                                    id="velocity"
                                    v-model.number="velocity"
                                    @input="updateVelocity"
                                    class="slider"
                                    min="0.1"
                                    max="10"
                                    step="0.1"
                                >
                                <span class="slider-value">{{ velocityDisplay }}</span>
                            </div>
                        </div>

                        <button type="submit" class="assess-btn" :disabled="loading || !material">
                            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" style="margin-right: 0.5rem;">
                                <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2z"/>
                                <path d="M12 6v6l4 2"/>
                            </svg>
                            {{ loading ? 'Analyzing...' : 'Assess Risk' }}
                        </button>
                    </form>
                </div>
                
                <!-- Results Section -->
                <div class="results-section">
                    <h2 class="section-title">
                        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                            <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2z"/>
                            <path d="M12 6v6l4 2"/>
                        </svg>
                        Corrosion Analysis
                    </h2>
                    
                    <!-- Loading State -->
                    <div v-if="loading" style="text-align: center; padding: 3rem;">
                        <div class="loading-spinner"></div>
                        <p style="color: var(--text-secondary); margin-top: 1rem;">Analyzing corrosion risk...</p>
                    </div>

                    <!-- Error State -->
                    <div v-else-if="error" class="error-message">
                        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" style="margin-right: 0.5rem;">
                            <circle cx="12" cy="12" r="10"/>
                            <path d="M12 8v4m0 4h.01"/>
                        </svg>
                        {{ error }}
                    </div>

                    <!-- Results State -->
                    <div v-else-if="results" id="corrosion-results">
                        <!-- Risk Level Badge -->
                        <div style="margin-bottom: 2rem;">
                            <div :class="['risk-badge', `risk-${results.riskLevel.toLowerCase()}`]">
                                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                    <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2z"/>
                                </svg>
                                {{ results.riskLevel }} RISK
                            </div>
                        </div>

                        <!-- Corrosion Rate -->
                        <div class="result-section">
                            <h3 class="result-heading">
                                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                    <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2z"/>
                                    <path d="M12 6v6l4 2"/>
                                </svg>
                                Corrosion Rate
                            </h3>
                            <div class="corrosion-rate-box">
                                <div class="rate-value">{{ results.corrosionRate }}</div>
                                <div class="rate-unit">mm/year</div>
                            </div>
                        </div>

                        <!-- Estimated Lifetime -->
                        <div class="result-section">
                            <h3 class="result-heading">
                                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                    <rect x="3" y="4" width="18" height="18" rx="2" ry="2"/>
                                    <line x1="16" y1="2" x2="16" y2="6"/>
                                    <line x1="8" y1="2" x2="8" y2="6"/>
                                    <line x1="3" y1="10" x2="21" y2="10"/>
                                </svg>
                                Estimated Equipment Life
                            </h3>
                            <div class="lifetime-badge">
                                {{ results.estimatedLife }}
                            </div>
                        </div>

                        <!-- Corrosion Mechanisms -->
                        <div class="result-section">
                            <h3 class="result-heading">
                                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                    <path d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"/>
                                    <path d="M12 9v4m0 4h.01"/>
                                </svg>
                                Identified Corrosion Mechanisms
                            </h3>
                            <div class="mechanisms-list">
                                <div
                                    v-for="(mechanism, index) in results.mechanisms"
                                    :key="index"
                                    class="mechanism-item"
                                >
                                    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                        <polyline points="9 11 12 14 22 4"/>
                                        <path d="M21 12v7a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11"/>
                                    </svg>
                                    {{ mechanism }}
                                </div>
                            </div>
                        </div>

                        <!-- Recommendations -->
                        <div class="result-section">
                            <h3 class="result-heading">
                                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                    <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/>
                                    <path d="M9 12l2 2 4-4"/>
                                </svg>
                                Recommendations
                            </h3>
                            <ul class="recommendations-list">
                                <li v-for="(rec, index) in results.recommendations" :key="index">
                                    {{ rec }}
                                </li>
                            </ul>
                        </div>
                    </div>

                    <!-- Empty State -->
                    <div v-else style="text-align: center; color: var(--text-secondary); padding: 2rem;">
                        <svg width="64" height="64" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1" style="margin: 0 auto 1rem; opacity: 0.5;">
                            <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2z"/>
                            <path d="M12 6v6l4 2"/>
                        </svg>
                        <p>Adjust the process parameters above and click "Assess Risk" to receive AI-powered corrosion analysis and recommendations.</p>
                    </div>
                </div>
            </div>
            
            <!-- Back to Projects -->
            <div class="back-section">
                <button @click="goHome" class="back-btn">
                    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M19 12H5M12 19l-7-7 7-7"/>
                    </svg>
                    Back to Projects
                </button>
            </div>
        </main>
    </body>
</template>

<style>
    /* Background Effects */
    .bg-grid {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background-image: 
            linear-gradient(rgba(0, 212, 255, 0.1) 1px, transparent 1px),
            linear-gradient(90deg, rgba(0, 212, 255, 0.1) 1px, transparent 1px);
        background-size: 50px 50px;
        z-index: -2;
        opacity: 0.3;
    }

    /* Navigation */
    .navbar {
        position: fixed;
        top: 0;
        left: 0;
        right: 0;
        background: rgba(26, 26, 26, 0.9);
        backdrop-filter: blur(10px);
        border-bottom: 1px solid rgba(0, 212, 255, 0.2);
        z-index: 1000;
        padding: 1rem 0;
    }

    .nav-container {
        max-width: 1400px;
        margin: 0 auto;
        padding: 0 2rem;
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    .logo {
        font-family: 'Orbitron', monospace;
        font-size: 1.5rem;
        font-weight: 700;
        color: var(--accent-cyan);
        text-decoration: none;
    }

    .breadcrumbs {
        color: var(--text-secondary);
        font-size: 0.9rem;
    }

    .breadcrumbs a {
        color: var(--accent-cyan);
        text-decoration: none;
    }

    .breadcrumbs a:hover {
        text-decoration: underline;
    }

    /* Main Content */
    .main-content {
        margin-top: 80px;
        padding: 2rem;
        max-width: 1400px;
        margin-left: auto;
        margin-right: auto;
    }

    .page-header {
        text-align: center;
        margin-bottom: 3rem;
        padding: 3rem 0;
        background: linear-gradient(135deg, rgba(0, 212, 255, 0.1), rgba(255, 107, 53, 0.1));
        border-radius: 12px;
        border: 1px solid rgba(0, 212, 255, 0.2);
    }

    .page-title {
        font-family: 'Orbitron', monospace;
        font-size: 2.5rem;
        font-weight: 900;
        color: var(--accent-cyan);
        margin-bottom: 1rem;
    }

    .page-subtitle {
        font-size: 1.1rem;
        color: var(--text-secondary);
        max-width: 600px;
        margin: 0 auto;
    }

    /* Application Interface */
    .app-interface {
        display: grid;
        grid-template-columns: 1fr 1fr;
        gap: 3rem;
        margin-bottom: 3rem;
    }

    .input-section {
        background: rgba(45, 45, 45, 0.8);
        border: 1px solid rgba(0, 212, 255, 0.2);
        border-radius: 12px;
        padding: 2rem;
    }

    .section-title {
        font-size: 1.5rem;
        font-weight: 700;
        color: var(--accent-cyan);
        margin-bottom: 1.5rem;
        display: flex;
        align-items: center;
        gap: 0.5rem;
    }

    .form-group {
        margin-bottom: 1.5rem;
    }

    .form-label {
        display: block;
        margin-bottom: 0.5rem;
        font-weight: 600;
        color: var(--text-primary);
    }

    .form-select, .form-input {
        width: 100%;
        padding: 0.75rem 1rem;
        background: rgba(26, 26, 26, 0.8);
        border: 1px solid rgba(0, 212, 255, 0.3);
        border-radius: 8px;
        color: var(--text-primary);
        font-family: 'Inter', sans-serif;
        font-size: 1rem;
        transition: all 0.3s ease;
    }

    .form-select:focus, .form-input:focus {
        outline: none;
        border-color: var(--accent-cyan);
        box-shadow: 0 0 0 3px rgba(0, 212, 255, 0.1);
    }

    .slider-container {
        margin-top: 0.5rem;
    }

    .slider {
        width: 100%;
        height: 6px;
        border-radius: 3px;
        background: rgba(0, 212, 255, 0.2);
        outline: none;
        -webkit-appearance: none;
    }

    .slider::-webkit-slider-thumb {
        -webkit-appearance: none;
        appearance: none;
        width: 20px;
        height: 20px;
        border-radius: 50%;
        background: var(--accent-cyan);
        cursor: pointer;
        box-shadow: 0 2px 6px rgba(0, 212, 255, 0.3);
    }

    .slider::-moz-range-thumb {
        width: 20px;
        height: 20px;
        border-radius: 50%;
        background: var(--accent-cyan);
        cursor: pointer;
        border: none;
        box-shadow: 0 2px 6px rgba(0, 212, 255, 0.3);
    }

    .slider-value {
        display: inline-block;
        margin-left: 1rem;
        font-weight: 600;
        color: var(--accent-cyan);
        min-width: 60px;
    }

    .assess-btn {
        background: linear-gradient(135deg, var(--accent-cyan), #0099cc);
        color: white;
        border: none;
        padding: 1rem 2rem;
        border-radius: 8px;
        font-size: 1rem;
        font-weight: 600;
        cursor: pointer;
        transition: all 0.3s ease;
        width: 100%;
        margin-top: 1rem;
    }

    .assess-btn:hover {
        transform: translateY(-2px);
        box-shadow: 0 10px 25px rgba(0, 212, 255, 0.3);
    }

    .assess-btn:disabled {
        opacity: 0.6;
        cursor: not-allowed;
        transform: none;
        box-shadow: none;
    }

    /* Results Section */
    .results-section {
        background: rgba(45, 45, 45, 0.8);
        border: 1px solid rgba(0, 212, 255, 0.2);
        border-radius: 12px;
        padding: 2rem;
        min-height: 400px;
    }

    .results-container {
        display: none;
    }

    .results-container.show {
        display: block;
        animation: fadeInUp 0.6s ease;
    }

    @keyframes fadeInUp {
        from {
            opacity: 0;
            transform: translateY(20px);
        }
        to {
            opacity: 1;
            transform: translateY(0);
        }
    }

    .result-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 2rem;
        padding-bottom: 1rem;
        border-bottom: 1px solid rgba(0, 212, 255, 0.2);
    }

    .result-header h3 {
        font-size: 1.5rem;
        font-weight: 700;
        color: var(--accent-cyan);
    }

    .risk-badge {
        display: inline-flex;
        align-items: center;
        gap: 0.5rem;
        padding: 0.75rem 1.5rem;
        border-radius: 25px;
        font-weight: 700;
        font-size: 1rem;
        color: white;
        letter-spacing: 0.5px;
    }

    .risk-high {
        background: linear-gradient(135deg, var(--error-red), #cc0000);
        border: 2px solid var(--error-red);
        box-shadow: 0 4px 15px rgba(255, 77, 77, 0.3);
    }

    .risk-medium {
        background: linear-gradient(135deg, var(--accent-orange), #cc5200);
        border: 2px solid var(--accent-orange);
        box-shadow: 0 4px 15px rgba(255, 107, 53, 0.3);
    }

    .risk-low {
        background: linear-gradient(135deg, var(--success-green), #00994d);
        border: 2px solid var(--success-green);
        box-shadow: 0 4px 15px rgba(0, 204, 102, 0.3);
    }

    .result-section {
        margin-bottom: 2.5rem;
        padding: 1.5rem;
        background: rgba(26, 26, 26, 0.5);
        border-radius: 10px;
        border: 1px solid rgba(0, 212, 255, 0.2);
    }

    .result-heading {
        font-size: 1.2rem;
        font-weight: 700;
        color: var(--accent-cyan);
        margin-bottom: 1.5rem;
        display: flex;
        align-items: center;
        gap: 0.5rem;
    }

    .corrosion-rate-box {
        display: flex;
        align-items: baseline;
        gap: 0.75rem;
        background: rgba(0, 212, 255, 0.1);
        border: 2px solid rgba(0, 212, 255, 0.3);
        border-radius: 10px;
        padding: 2rem;
        justify-content: center;
    }

    .rate-unit {
        font-size: 1.2rem;
        color: var(--text-secondary);
        font-weight: 600;
    }

    .lifetime-badge {
        background: linear-gradient(135deg, rgba(0, 212, 255, 0.2), rgba(0, 212, 255, 0.1));
        border: 2px solid var(--accent-cyan);
        border-radius: 10px;
        padding: 1.5rem 2rem;
        text-align: center;
        font-size: 1.5rem;
        font-weight: 700;
        color: var(--accent-cyan);
    }

    .mechanisms-list {
        display: flex;
        flex-direction: column;
        gap: 0.75rem;
    }

    .mechanism-item {
        display: flex;
        align-items: center;
        gap: 0.75rem;
        background: rgba(255, 107, 53, 0.1);
        border: 1px solid rgba(255, 107, 53, 0.3);
        border-radius: 8px;
        padding: 1rem 1.25rem;
        transition: all 0.3s ease;
    }

    .mechanism-item:hover {
        background: rgba(255, 107, 53, 0.15);
        transform: translateX(5px);
    }

    .mechanism-item svg {
        color: var(--accent-orange);
        flex-shrink: 0;
    }

    .recommendations-list {
        list-style: none;
        padding: 0;
        counter-reset: rec-counter;
    }

    .recommendations-list li {
        counter-increment: rec-counter;
        background: rgba(0, 212, 255, 0.1);
        border-left: 4px solid var(--accent-cyan);
        padding: 1rem 1rem 1rem 3rem;
        margin-bottom: 0.75rem;
        border-radius: 0 8px 8px 0;
        position: relative;
        transition: all 0.3s ease;
    }

    .recommendations-list li:hover {
        background: rgba(0, 212, 255, 0.15);
        transform: translateX(5px);
    }

    .recommendations-list li::before {
        content: counter(rec-counter);
        position: absolute;
        left: 1rem;
        top: 50%;
        transform: translateY(-50%);
        background: var(--accent-cyan);
        color: var(--primary-bg);
        width: 26px;
        height: 26px;
        border-radius: 50%;
        display: flex;
        align-items: center;
        justify-content: center;
        font-weight: 700;
        font-size: 0.85rem;
    }

    .error-message {
        background: rgba(255, 77, 77, 0.15);
        border: 2px solid var(--error-red);
        border-radius: 10px;
        padding: 1.5rem;
        color: var(--error-red);
        display: flex;
        align-items: center;
        font-weight: 600;
    }

    .loading-spinner {
        width: 50px;
        height: 50px;
        border: 4px solid rgba(0, 212, 255, 0.2);
        border-top: 4px solid var(--accent-cyan);
        border-radius: 50%;
        animation: spin 1s linear infinite;
        margin: 0 auto;
    }

    .analysis-section {
        margin-bottom: 2rem;
    }

    .analysis-section h4 {
        font-size: 1.1rem;
        font-weight: 600;
        color: var(--text-primary);
        margin-bottom: 1rem;
    }

    .corrosion-rate {
        display: flex;
        align-items: center;
        gap: 1rem;
        background: rgba(0, 212, 255, 0.1);
        border: 1px solid rgba(0, 212, 255, 0.3);
        border-radius: 8px;
        padding: 1.5rem;
        margin-bottom: 1rem;
    }

    .rate-value {
        font-size: 2rem;
        font-weight: 900;
        color: var(--accent-cyan);
        font-family: 'JetBrains Mono', monospace;
    }

    .material {
        color: var(--text-secondary);
        font-size: 1rem;
    }

    .prevention-list, .maintenance-list {
        list-style: none;
        padding: 0;
    }

    .prevention-list li, .maintenance-list li {
        background: rgba(0, 212, 255, 0.05);
        border-left: 3px solid var(--accent-cyan);
        padding: 0.75rem 1rem;
        margin-bottom: 0.5rem;
        border-radius: 0 6px 6px 0;
    }

    .chart-container {
        margin-top: 2rem;
        background: rgba(26, 26, 26, 0.5);
        border-radius: 8px;
        padding: 1rem;
    }

    /* Loading State */
    .loading-spinner {
        display: flex;
        align-items: center;
        justify-content: center;
        gap: 1rem;
        padding: 2rem;
        color: var(--accent-cyan);
    }

    .spinner {
        width: 30px;
        height: 30px;
        border: 3px solid rgba(0, 212, 255, 0.3);
        border-top: 3px solid var(--accent-cyan);
        border-radius: 50%;
        animation: spin 1s linear infinite;
    }

    @keyframes spin {
        0% { transform: rotate(0deg); }
        100% { transform: rotate(360deg); }
    }

    /* Back Button */
    .back-section {
        text-align: center;
        margin-top: 3rem;
    }

    .back-btn {
        background: transparent;
        color: var(--accent-cyan);
        border: 2px solid var(--accent-cyan);
        padding: 1rem 2rem;
        border-radius: 8px;
        font-weight: 600;
        text-decoration: none;
        display: inline-flex;
        align-items: center;
        gap: 0.5rem;
        cursor: pointer;
        transition: all 0.3s ease;
    }

    .back-btn:hover {
        background: var(--accent-cyan);
        color: var(--primary-bg);
        transform: translateY(-2px);
    }

    /* Responsive Design */
    @media (max-width: 768px) {
        .app-interface {
            grid-template-columns: 1fr;
            gap: 2rem;
        }
        
        .page-title {
            font-size: 2rem;
        }
        
        .nav-container {
            padding: 0 1rem;
        }
        
        .main-content {
            padding: 1rem;
        }
    }
</style>