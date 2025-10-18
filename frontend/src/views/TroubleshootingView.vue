<script setup lang="ts">
import { ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

const router = useRouter()

const selectedEquipment = ref('')
const problem = ref('')
const errorCode = ref('')
const loading = ref(false)
const error = ref('')
const results = ref<any>(null)

const selectEquipment = (equipment: string) => {
  selectedEquipment.value = equipment
}

const handleSubmit = async (e: Event) => {
  e.preventDefault()

  if (!selectedEquipment.value) {
    error.value = 'Please select an equipment type'
    return
  }

  if (!problem.value.trim()) {
    error.value = 'Please describe the problem'
    return
  }

  loading.value = true
  error.value = ''
  results.value = null

  try {
    const response = await axios.post('/api/search', {
      equipment: selectedEquipment.value,
      problem: problem.value,
      error_code: errorCode.value
    })

    if (response.data.success) {
      results.value = response.data.response
    } else {
      error.value = 'Failed to get troubleshooting analysis'
    }
  } catch (err: any) {
    error.value = err.response?.data?.error || 'Failed to connect to troubleshooting service'
    console.error('Troubleshooting error:', err)
  } finally {
    loading.value = false
  }
}

const goHome = () => {
  router.push('/')
}
</script>

<template>
  <!-- Fonts -->
  <link href="https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700&family=JetBrains+Mono:wght@400;500&family=Orbitron:wght@400;700;900&display=swap" rel="stylesheet">

  <body>
    <!-- Background Effects -->
    <div class="bg-grid"></div>
    
    <!-- Navigation -->
    <nav class="navbar">
        <div class="nav-container">
            <a href="/" class="logo">PCST-AI</a>
            <div id="breadcrumbs" class="breadcrumbs">
                <a href="/">Home</a> / <span>PCST Troubleshooting Assistant</span>
            </div>
        </div>
    </nav>
    
    <!-- Main Content -->
    <main class="main-content">
        <!-- Page Header -->
        <div class="page-header">
            <h1 class="page-title">PCST Troubleshooting Assistant</h1>
            <p class="page-subtitle">
                AI-powered troubleshooting assistant that helps Process Control Systems Technicians diagnose equipment problems, find solutions, and access step-by-step repair procedures with safety guidance.
            </p>
        </div>
        
        <!-- Application Interface -->
        <div class="app-interface">
            <!-- Input Section -->
            <div class="input-section">
                <h2 class="section-title">
                    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M14.7 6.3a1 1 0 0 0 0 1.4l1.6 1.6a1 1 0 0 0 1.4 0l3.77-3.77a6 6 0 0 1-7.94 7.94l-6.91 6.91a2.12 2.12 0 0 1-3-3l6.91-6.91a6 6 0 0 1 7.94-7.94l-3.76 3.76z"/>
                        <path d="m8.5 12.5 2.5 2.5"/>
                    </svg>
                    Equipment Information
                </h2>
                
                <form @submit.prevent="handleSubmit">
                    <div class="form-group">
                        <label class="form-label">Select Equipment Type</label>
                        <div class="equipment-icons">
                            <div
                                class="equipment-option"
                                :class="{ selected: selectedEquipment === 'Control Valve' }"
                                @click="selectEquipment('Control Valve')"
                            >
                                <div class="equipment-icon">🎛️</div>
                                <div class="equipment-info">
                                    <div class="equipment-name">Control Valve</div>
                                    <div class="equipment-type">Actuated Valve</div>
                                </div>
                            </div>

                            <div
                                class="equipment-option"
                                :class="{ selected: selectedEquipment === 'Pressure Transmitter' }"
                                @click="selectEquipment('Pressure Transmitter')"
                            >
                                <div class="equipment-icon">📊</div>
                                <div class="equipment-info">
                                    <div class="equipment-name">Pressure Transmitter</div>
                                    <div class="equipment-type">4-20mA Output</div>
                                </div>
                            </div>

                            <div
                                class="equipment-option"
                                :class="{ selected: selectedEquipment === 'Flow Meter' }"
                                @click="selectEquipment('Flow Meter')"
                            >
                                <div class="equipment-icon">🌊</div>
                                <div class="equipment-info">
                                    <div class="equipment-name">Flow Meter</div>
                                    <div class="equipment-type">Magnetic/Coriolis</div>
                                </div>
                            </div>

                            <div
                                class="equipment-option"
                                :class="{ selected: selectedEquipment === 'Temperature Sensor' }"
                                @click="selectEquipment('Temperature Sensor')"
                            >
                                <div class="equipment-icon">🌡️</div>
                                <div class="equipment-info">
                                    <div class="equipment-name">Temperature Sensor</div>
                                    <div class="equipment-type">RTD/Thermocouple</div>
                                </div>
                            </div>
                        </div>
                    </div>

                    <div class="form-group">
                        <label for="problem" class="form-label">Problem Description</label>
                        <textarea
                            id="problem"
                            v-model="problem"
                            class="form-textarea"
                            placeholder="Describe the problem you're experiencing with the equipment. Include any symptoms, error messages, or unusual behavior you've observed."
                            required
                        ></textarea>
                    </div>

                    <div class="form-group">
                        <label for="errorCode" class="form-label">Error Code (Optional)</label>
                        <input
                            type="text"
                            id="errorCode"
                            v-model="errorCode"
                            class="form-input"
                            placeholder="e.g., ERR-101, FAULT-23, ALARM-45"
                        />
                    </div>

                    <button type="submit" class="search-btn" :disabled="loading || !selectedEquipment">
                        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" style="margin-right: 0.5rem;">
                            <path d="M14.7 6.3a1 1 0 0 0 0 1.4l1.6 1.6a1 1 0 0 0 1.4 0l3.77-3.77a6 6 0 0 1-7.94 7.94l-6.91 6.91a2.12 2.12 0 0 1-3-3l6.91-6.91a6 6 0 0 1 7.94-7.94l-3.76 3.76z"/>
                            <path d="m8.5 12.5 2.5 2.5"/>
                        </svg>
                        {{ loading ? 'Analyzing...' : 'Search for Solution' }}
                    </button>
                </form>
            </div>
            
            <!-- Results Section -->
            <div class="results-section">
                <h2 class="section-title">
                    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M14.7 6.3a1 1 0 0 0 0 1.4l1.6 1.6a1 1 0 0 0 1.4 0l3.77-3.77a6 6 0 0 1-7.94 7.94l-6.91 6.91a2.12 2.12 0 0 1-3-3l6.91-6.91a6 6 0 0 1 7.94-7.94l-3.76 3.76z"/>
                        <path d="m8.5 12.5 2.5 2.5"/>
                    </svg>
                    Troubleshooting Analysis
                </h2>
                
                <div v-if="loading" style="text-align: center; padding: 3rem;">
                    <div class="loading-spinner"></div>
                    <p style="color: var(--text-secondary); margin-top: 1rem;">Analyzing troubleshooting data...</p>
                </div>

                <div v-else-if="error" class="error-message">
                    <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" style="margin-right: 0.5rem;">
                        <circle cx="12" cy="12" r="10"/>
                        <path d="M12 8v4m0 4h.01"/>
                    </svg>
                    {{ error }}
                </div>

                <div v-else-if="results" id="troubleshooting-results">
                    <div class="result-section">
                        <h3 class="result-heading">
                            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <path d="M9 11l3 3L22 4"/>
                                <path d="M21 12v7a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11"/>
                            </svg>
                            Analysis
                        </h3>
                        <p class="analysis-text">{{ results.analysis }}</p>
                    </div>

                    <div class="result-section">
                        <h3 class="result-heading">
                            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <circle cx="12" cy="12" r="10"/>
                                <path d="M12 16v-4m0-4h.01"/>
                            </svg>
                            Possible Causes
                        </h3>
                        <ul class="causes-list">
                            <li v-for="(cause, index) in results.causes" :key="index">
                                {{ cause }}
                            </li>
                        </ul>
                    </div>

                    <div class="result-section">
                        <h3 class="result-heading">
                            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <path d="M14.7 6.3a1 1 0 0 0 0 1.4l1.6 1.6a1 1 0 0 0 1.4 0l3.77-3.77a6 6 0 0 1-7.94 7.94l-6.91 6.91a2.12 2.12 0 0 1-3-3l6.91-6.91a6 6 0 0 1 7.94-7.94l-3.76 3.76z"/>
                            </svg>
                            Troubleshooting Steps
                        </h3>
                        <ul class="steps-list">
                            <li v-for="(step, index) in results.steps" :key="index">
                                {{ step }}
                            </li>
                        </ul>
                    </div>

                    <div class="result-section">
                        <h3 class="result-heading">
                            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <path d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"/>
                                <path d="M12 9v4m0 4h.01"/>
                            </svg>
                            Safety Warnings
                        </h3>
                        <ul class="safety-warnings">
                            <li v-for="(warning, index) in results.safety_warnings" :key="index">
                                {{ warning }}
                            </li>
                        </ul>
                    </div>

                    <div class="result-section">
                        <h3 class="result-heading">
                            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
                                <path d="M14 2v6h6M16 13H8m8 4H8"/>
                            </svg>
                            Equipment Notes
                        </h3>
                        <p class="equipment-notes">{{ results.equipment_notes }}</p>
                    </div>
                </div>

                <div v-else style="text-align: center; color: var(--text-secondary); padding: 2rem;">
                    <svg width="64" height="64" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1" style="margin: 0 auto 1rem; opacity: 0.5;">
                        <path d="M14.7 6.3a1 1 0 0 0 0 1.4l1.6 1.6a1 1 0 0 0 1.4 0l3.77-3.77a6 6 0 0 1-7.94 7.94l-6.91 6.91a2.12 2.12 0 0 1-3-3l6.91-6.91a6 6 0 0 1 7.94-7.94l-3.76 3.76z"/>
                        <path d="m8.5 12.5 2.5 2.5"/>
                    </svg>
                    <p>Select equipment type, describe the problem, and click "Search for Solution" to receive AI-powered troubleshooting guidance.</p>
                    <p style="margin-top: 1rem; font-size: 0.9rem;">Powered by Google Gemini AI</p>
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
    
    .form-select, .form-input, .form-textarea {
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
    
    .form-select:focus, .form-input:focus, .form-textarea:focus {
        outline: none;
        border-color: var(--accent-cyan);
        box-shadow: 0 0 0 3px rgba(0, 212, 255, 0.1);
    }
    
    .form-textarea {
        min-height: 120px;
        resize: vertical;
    }
    
    .form-textarea::placeholder {
        color: var(--text-secondary);
        font-style: italic;
    }
    
    .equipment-icons {
        display: grid;
        grid-template-columns: repeat(2, 1fr);
        gap: 1rem;
        margin-top: 1rem;
    }
    
    .equipment-option {
        display: flex;
        align-items: center;
        gap: 0.75rem;
        padding: 0.75rem;
        background: rgba(26, 26, 26, 0.6);
        border: 1px solid rgba(0, 212, 255, 0.2);
        border-radius: 8px;
        cursor: pointer;
        transition: all 0.3s ease;
    }
    
    .equipment-option:hover {
        border-color: var(--accent-cyan);
        background: rgba(0, 212, 255, 0.1);
    }
    
    .equipment-option.selected {
        border-color: var(--accent-cyan);
        background: rgba(0, 212, 255, 0.2);
    }
    
    .equipment-icon {
        width: 40px;
        height: 40px;
        background: linear-gradient(135deg, var(--accent-cyan), var(--accent-orange));
        border-radius: 8px;
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 1.2rem;
    }
    
    .equipment-info {
        flex: 1;
    }
    
    .equipment-name {
        font-weight: 600;
        color: var(--text-primary);
    }
    
    .equipment-type {
        font-size: 0.8rem;
        color: var(--text-secondary);
    }
    
    .search-btn {
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
    
    .search-btn:hover {
        transform: translateY(-2px);
        box-shadow: 0 10px 25px rgba(0, 212, 255, 0.3);
    }
    
    .search-btn:disabled {
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
    
    .ai-badge {
        background: linear-gradient(135deg, var(--accent-orange), #cc5200);
        color: white;
        padding: 0.5rem 1rem;
        border-radius: 20px;
        font-weight: 600;
        font-size: 0.8rem;
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
    
    .problem-analysis {
        background: rgba(0, 212, 255, 0.1);
        border-left: 4px solid var(--accent-cyan);
        padding: 1rem;
        border-radius: 0 8px 8px 0;
        margin-bottom: 1rem;
        font-weight: 500;
    }
    
    .causes-list, .prevention-list, .maintenance-list {
        list-style: none;
        padding: 0;
    }
    
    .causes-list li {
        background: rgba(255, 107, 53, 0.1);
        border-left: 3px solid var(--accent-orange);
        padding: 0.75rem 1rem;
        margin-bottom: 0.5rem;
        border-radius: 0 6px 6px 0;
    }
    
    .troubleshooting-steps {
        display: flex;
        flex-direction: column;
        gap: 1rem;
    }
    
    .step-item {
        display: flex;
        gap: 1rem;
        background: rgba(0, 212, 255, 0.05);
        border: 1px solid rgba(0, 212, 255, 0.2);
        border-radius: 8px;
        padding: 1rem;
    }
    
    .step-number {
        width: 30px;
        height: 30px;
        background: var(--accent-cyan);
        color: var(--primary-bg);
        border-radius: 50%;
        display: flex;
        align-items: center;
        justify-content: center;
        font-weight: 700;
        font-size: 0.9rem;
        flex-shrink: 0;
    }
    
    .step-content {
        flex: 1;
        font-weight: 500;
    }
    
    .safety-warnings {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
    }
    
    .warning-item {
        background: rgba(255, 45, 85, 0.1);
        border-left: 3px solid var(--error-red);
        padding: 0.75rem 1rem;
        border-radius: 0 6px 6px 0;
        font-weight: 500;
    }
    
    .equipment-notes {
        background: rgba(0, 212, 255, 0.1);
        border-left: 4px solid var(--accent-cyan);
        padding: 1rem;
        border-radius: 0 8px 8px 0;
        font-style: italic;
        color: var(--text-secondary);
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

    .analysis-text {
        background: rgba(0, 212, 255, 0.1);
        border-left: 4px solid var(--accent-cyan);
        padding: 1rem;
        border-radius: 0 8px 8px 0;
        line-height: 1.6;
    }

    .steps-list {
        list-style: none;
        padding: 0;
        counter-reset: step-counter;
    }

    .steps-list li {
        counter-increment: step-counter;
        background: rgba(0, 212, 255, 0.1);
        border-left: 4px solid var(--accent-cyan);
        padding: 1rem 1rem 1rem 3rem;
        margin-bottom: 0.75rem;
        border-radius: 0 8px 8px 0;
        position: relative;
        transition: all 0.3s ease;
    }

    .steps-list li:hover {
        background: rgba(0, 212, 255, 0.15);
        transform: translateX(5px);
    }

    .steps-list li::before {
        content: counter(step-counter);
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

    .safety-warnings {
        list-style: none;
        padding: 0;
    }

    .safety-warnings li {
        background: rgba(255, 77, 77, 0.1);
        border-left: 4px solid var(--error-red);
        padding: 0.75rem 1rem;
        margin-bottom: 0.75rem;
        border-radius: 0 6px 6px 0;
        font-weight: 500;
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
        
        .equipment-icons {
            grid-template-columns: 1fr;
        }
    }
</style>