<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

const router = useRouter()
const logs = ref('')
const loading = ref(false)
const results = ref<any>(null)
const error = ref('')

const API_BASE_URL = '/api'

// Sample logs for demonstration
const sampleLogs = {
  pressure: `14:23:12 ALARM: HIGH PRESSURE - Line 3 - 78.5 bar
14:23:45 WARNING: PRESSURE TRENDING UP - Line 3 - 79.2 bar
14:24:15 ALARM: PRESSURE SETPOINT EXCEEDED - Line 3 - 81.1 bar
14:24:30 EVENT: SAFETY VALVE LIFTING - PSV-301 - OPEN
14:25:00 WARNING: FLOW REDUCTION - Line 3 - 65% normal
14:25:30 ALARM: UPSTREAM PRESSURE DROP - Station A - 15% loss`,

  temperature: `13:45:00 WARNING: TEMP RISING - Reactor R-2 - 142°C
13:46:15 WARNING: TEMP HIGH - Reactor R-2 - 148°C
13:47:30 ALARM: TEMP SETPOINT EXCEEDED - Reactor R-2 - 152°C
13:48:00 EVENT: COOLING WATER INCREASE - CW-201 - 85% open
13:49:15 WARNING: PRESSURE INCREASE - Reactor R-2 - 3.2 bar
13:50:30 EVENT: TEMP STABILIZING - Reactor R-2 - 149°C`,

  vibration: `15:10:23 ALARM: VIBRATION HIGH - Pump P-101 - 8.7 mm/s
15:11:45 WARNING: VIBRATION INCREASING - Pump P-101 - 10.2 mm/s
15:12:30 EVENT: SPEED REDUCTION - Pump P-101 - 80% setpoint
15:13:20 WARNING: BEARING TEMP HIGH - Pump P-101 - 75°C
15:14:15 ALARM: VIBRATION CRITICAL - Pump P-101 - 12.5 mm/s
15:15:00 EVENT: EMERGENCY SHUTDOWN - Pump P-101 - INITIATED`
}

const insertSampleLogs = (type: keyof typeof sampleLogs) => {
  logs.value = sampleLogs[type]
}

const handleSubmit = async () => {
  if (!logs.value.trim()) {
    error.value = 'Please enter control room logs to analyze.'
    return
  }

  loading.value = true
  error.value = ''
  results.value = null

  try {
    // API Request Format:
    // POST /api/vcra/analyze
    // Body: { logs: string }
    const response = await axios.post(`${API_BASE_URL}/vcra/analyze`, {
      logs: logs.value
    })

    if (response.data.success) {
      results.value = response.data.response
      console.log('data: ', results.value)
      setTimeout(() => {
        document.getElementById('vca-results')?.scrollIntoView({ behavior: 'smooth' })
      }, 100)
    }
  } catch (err: any) {
    console.error('VCA Analysis error:', err)
    error.value = err.response?.data?.error || err.message || 'An error occurred during analysis'
  } finally {
    loading.value = false
  }
}

const getRiskColor = (riskLevel: string) => {
  switch (riskLevel) {
    case 'HIGH': return '#ff2d55'
    case 'MEDIUM': return '#ff6b35'
    case 'LOW': return '#00ff88'
    default: return '#00d4ff'
  }
}

const goHome = () => {
  router.push('/')
}

// Make insertSampleLogs available globally for inline onclick handlers
if (typeof window !== 'undefined') {
  (window as any).insertSampleLogs = insertSampleLogs
}
</script>

<template>
    <head>
        <title>Virtual Control Room Advisor - PCST-AI Portal</title>
    
        <!-- Fonts -->
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
                    <a href="/">Home</a> / <span>Virtual Control Room Advisor</span>
                </div>
            </div>
        </nav>
        
        <!-- Main Content -->
        <main class="main-content">
            <!-- Page Header -->
            <div class="page-header">
                <h1 class="page-title">Virtual Control Room Advisor</h1>
                <p class="page-subtitle">
                    AI-powered analysis of control room alarms and operational data to provide intelligent recommendations and predictive insights for process optimization.
                </p>
            </div>
            
            <!-- Application Interface -->
            <div class="app-interface">
                <!-- Input Section -->
                <div class="input-section">
                    <h2 class="section-title">
                        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                            <path d="M9 11H5a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-7a2 2 0 0 0-2-2z"/>
                            <path d="M19 7h-4a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2z"/>
                            <path d="M5 11V7a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v4"/>
                        </svg>
                        Control Room Logs
                    </h2>
                    
                    <form id="vca-form" @submit.prevent="handleSubmit">
                        <div class="form-group">
                            <label for="logs" class="form-label">Paste Control Room Logs or Alarm Data</label>
                            <textarea
                                id="logs"
                                v-model="logs"
                                name="logs"
                                class="form-textarea"
                                placeholder="Example logs:
    14:23:12 ALARM: HIGH PRESSURE - Line 3 - 78.5 bar
    14:24:45 WARNING: TEMP HIGH - Reactor R-2 - 145°C
    14:25:30 EVENT: FLOW LOW - Pump P-101 - 85% setpoint
    14:26:15 ALARM: VIBRATION HIGH - Compressor C-2 - 12.3 mm/s"
                                required
                            ></textarea>
                        </div>

                        <div class="sample-logs">
                            <button type="button" class="sample-btn" @click="insertSampleLogs('pressure')">Sample: Pressure Alarm</button>
                            <button type="button" class="sample-btn" @click="insertSampleLogs('temperature')">Sample: Temperature Warning</button>
                            <button type="button" class="sample-btn" @click="insertSampleLogs('vibration')">Sample: Vibration Alert</button>
                        </div>

                        <button type="submit" class="analyze-btn" :disabled="loading">
                            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" style="margin-right: 0.5rem;">
                                <path d="M9 11H5a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-7a2 2 0 0 0-2-2z"/>
                                <path d="M19 7h-4a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2z"/>
                                <path d="M5 11V7a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v4"/>
                            </svg>
                            {{ loading ? 'Analyzing...' : 'Analyze Logs' }}
                        </button>
                    </form>
                </div>
                
                <!-- Results Section -->
                <div class="results-section">
                    <h2 class="section-title">
                        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                            <path d="M9 11H5a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-7a2 2 0 0 0-2-2z"/>
                            <path d="M19 7h-4a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2z"/>
                            <path d="M5 11V7a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v4"/>
                        </svg>
                        AI Analysis Results
                    </h2>
                    
                    <div id="vca-results">
                        <!-- Loading State -->
                        <div v-if="loading" class="loading-spinner" style="text-align: center; padding: 2rem;">
                            <div class="spinner"></div>
                            <!-- <span>Analyzing control room logs...</span> -->
                        </div>

                        <!-- Error State -->
                        <div v-else-if="error" class="error-message" style="background: #ff2d5520; border-left: 4px solid #ff2d55; padding: 1rem; border-radius: 0.5rem; color: #ff2d55;">
                            <strong>Error:</strong> {{ error }}
                        </div>

                        <!-- Results -->
                        <div v-else-if="results" id="results-container">
                            <div class="result-header" style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 2rem;">
                                <h3>Analysis Results</h3>
                                <div class="confidence-badge" style="background: var(--accent-cyan); color: #000; padding: 0.5rem 1rem; border-radius: 2rem; font-weight: 600;">
                                    {{ Math.round(results.confidence * 100) }}% Confidence
                                </div>
                            </div>

                            <div id="analysis-section" style="margin-bottom: 2rem; padding: 1.5rem; background: rgba(26, 26, 26, 0.5); border-radius: 0.5rem; border: 1px solid rgba(0, 212, 255, 0.2);">
                                <h4 style="margin-bottom: 1rem;">Root Cause Analysis</h4>
                                <p class="root-cause" style="font-size: 1.1rem; line-height: 1.6;">{{ results.rootCause }}</p>
                                <div class="risk-indicator" :style="{ backgroundColor: getRiskColor(results.riskLevel) + '20', borderLeft: '4px solid ' + getRiskColor(results.riskLevel), padding: '1rem', marginTop: '1rem', borderRadius: '0.5rem' }">
                                    <strong>Risk Level: {{ results.riskLevel }}</strong>
                                </div>
                            </div>

                            <div id="analysis-section" style="margin-bottom: 2rem; padding: 1.5rem; background: rgba(26, 26, 26, 0.5); border-radius: 0.5rem; border: 1px solid rgba(0, 212, 255, 0.2);">
                                <h4 style="margin-bottom: 1rem;">Suggested Actions</h4>
                                <ol class="action-list" style="padding-left: 1.5rem;">
                                    <li v-for="(action, index) in results.actions" :key="index" style="margin-bottom: 0.75rem;">{{ action }}</li>
                                </ol>
                            </div>

                            <div id="analysis-section" style="margin-bottom: 2rem; padding: 1.5rem; background: rgba(26, 26, 26, 0.5); border-radius: 0.5rem; border: 1px solid rgba(0, 212, 255, 0.2);">
                                <h4 style="margin-bottom: 1rem;">Event Timeline</h4>
                                <div class="timeline">
                                    <div v-for="(event, index) in results.timeline" :key="index" class="timeline-item" style="padding: 0.75rem; border-left: 2px solid var(--accent-cyan); margin-bottom: 0.5rem; padding-left: 1rem; font-family: 'JetBrains Mono', monospace; font-size: 0.9rem;">
                                        {{ event }}
                                    </div>
                                </div>
                            </div>
                        </div>

                        <!-- Empty State -->
                        <div v-else style="text-align: center; color: var(--text-secondary); padding: 2rem;">
                            <svg width="64" height="64" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1" style="margin: 0 auto 1rem; opacity: 0.5;">
                                <path d="M9 11H5a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-7a2 2 0 0 0-2-2z"/>
                                <path d="M19 7h-4a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2z"/>
                                <path d="M5 11V7a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v4"/>
                            </svg>
                            <p>Enter control room logs above and click "Analyze Logs" to see AI-powered analysis and recommendations.</p>
                        </div>
                    </div>
                </div>
            </div>
            
            <!-- Back to Projects -->
            <div class="back-section">
                <a href="/" class="back-btn">
                    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M19 12H5M12 19l-7-7 7-7"/>
                    </svg>
                    Back to Projects
                </a>
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
    
    .form-textarea {
        width: 100%;
        min-height: 300px;
        padding: 1rem;
        background: rgba(26, 26, 26, 0.8);
        border: 1px solid rgba(0, 212, 255, 0.3);
        border-radius: 8px;
        color: var(--text-primary);
        font-family: 'JetBrains Mono', monospace;
        font-size: 0.9rem;
        resize: vertical;
        transition: all 0.3s ease;
    }
    
    .form-textarea:focus {
        outline: none;
        border-color: var(--accent-cyan);
        box-shadow: 0 0 0 3px rgba(0, 212, 255, 0.1);
    }
    
    .form-textarea::placeholder {
        color: var(--text-secondary);
        font-style: italic;
    }
    
    .sample-logs {
        margin-top: 1rem;
    }
    
    .sample-btn {
        background: rgba(255, 107, 53, 0.2);
        color: var(--accent-orange);
        border: 1px solid var(--accent-orange);
        padding: 0.5rem 1rem;
        border-radius: 6px;
        font-size: 0.9rem;
        cursor: pointer;
        transition: all 0.3s ease;
        margin-right: 0.5rem;
        margin-bottom: 0.5rem;
    }
    
    .sample-btn:hover {
        background: var(--accent-orange);
        color: white;
    }
    
    .analyze-btn {
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
    
    .analyze-btn:hover {
        transform: translateY(-2px);
        box-shadow: 0 10px 25px rgba(0, 212, 255, 0.3);
    }
    
    .analyze-btn:disabled {
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
    
    .confidence-badge {
        background: var(--success-green);
        color: var(--primary-bg);
        padding: 0.5rem 1rem;
        border-radius: 20px;
        font-weight: 600;
        font-size: 0.9rem;
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
    
    .root-cause {
        background: rgba(0, 212, 255, 0.1);
        border-left: 4px solid var(--accent-cyan);
        padding: 1rem;
        border-radius: 0 8px 8px 0;
        margin-bottom: 1rem;
        font-weight: 500;
    }
    
    .risk-indicator {
        padding: 1rem;
        border-radius: 8px;
        margin-bottom: 1rem;
        font-weight: 600;
    }
    
    .action-list, .timeline {
        list-style: none;
        padding: 0;
    }
    
    .action-list li {
        background: rgba(0, 212, 255, 0.05);
        border-left: 3px solid var(--accent-cyan);
        padding: 0.75rem 1rem;
        margin-bottom: 0.5rem;
        border-radius: 0 6px 6px 0;
    }
    
    .timeline-item {
        background: rgba(255, 107, 53, 0.05);
        border-left: 3px solid var(--accent-orange);
        padding: 0.5rem 1rem;
        margin-bottom: 0.5rem;
        border-radius: 0 6px 6px 0;
        font-family: 'JetBrains Mono', monospace;
        font-size: 0.9rem;
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