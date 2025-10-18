<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

const router = useRouter()
const task = ref('')
const loading = ref(false)
const results = ref<any>(null)
const error = ref('')

const API_BASE_URL = '/api'

// Sample tasks for demonstration
const sampleTasks = {
  electrical: `Maintenance technician performing electrical troubleshooting on 480V motor control center while plant is in operation.
Task involves voltage testing, component replacement, and system verification.
Environment contains energized equipment, arc flash hazards, and potential for electrical shock.`,

  confined: `Entry into 12-foot deep process vessel for inspection and cleaning.
Vessel diameter is 8 feet with single 24-inch manway access.
Atmosphere testing required, continuous ventilation needed, and rescue team on standby.
Entry duration approximately 4 hours with two-person team.`,

  hotwork: `Welding repair on 6-inch carbon steel pipeline in process area.
Pipe contains residual hydrocarbon vapors, located 15 feet from active process equipment.
Hot work permit required, fire watch needed, and area must be cleared of flammable materials.
Work duration estimated at 6 hours including setup and cooldown.`
}

const insertSampleTask = (type: keyof typeof sampleTasks) => {
  task.value = sampleTasks[type]
}

const handleSubmit = async () => {
  if (!task.value.trim()) {
    error.value = 'Please describe the task to evaluate.'
    return
  }

  loading.value = true
  error.value = ''
  results.value = null

  try {
    // API Request Format:
    // POST /api/safety/analyze
    // Body: { task: string }
    const response = await axios.post(`${API_BASE_URL}/safety/analyze`, {
      task: task.value
    })

    // Expected Response Format:
    // {
    //   success: boolean,
    //   response: {
    //     hazardLevel: "HIGH" | "MEDIUM" | "LOW",
    //     hazards: [
    //       { name: string, severity: string, probability: string }
    //     ],
    //     mitigations: string[],
    //     standards: string[]
    //   }
    // }

    if (response.data.success) {
      results.value = response.data.response
      setTimeout(() => {
        document.getElementById('safety-results')?.scrollIntoView({ behavior: 'smooth' })
      }, 100)
    }
  } catch (err: any) {
    console.error('Safety analysis error:', err)
    error.value = err.response?.data?.error || err.message || 'An error occurred during analysis'
  } finally {
    loading.value = false
  }
}

const getHazardColor = (hazardLevel: string) => {
  switch (hazardLevel) {
    case 'HIGH': return '#ff2d55'
    case 'MEDIUM': return '#ff6b35'
    case 'LOW': return '#00ff88'
    default: return '#00d4ff'
  }
}

const getSeverityClass = (severity: string) => {
  return `severity-${severity.toLowerCase()}`
}

const goHome = () => {
  router.push('/')
}

// Make insertSampleTask available globally for inline onclick handlers
if (typeof window !== 'undefined') {
  (window as any).insertSampleTask = insertSampleTask
}
</script>

<template>
    <head>
        <title>AI Safety Advisor - PCST-AI Portal</title>
    
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
                    <a href="/">Home</a> / <span>AI Safety Advisor</span>
                </div>
            </div>
        </nav>
        
        <!-- Main Content -->
        <main class="main-content">
            <!-- Page Header -->
            <div class="page-header">
                <h1 class="page-title">AI Safety Advisor</h1>
                <p class="page-subtitle">
                    Advanced AI-powered safety assessment tool that evaluates job tasks for potential hazards and provides comprehensive safety recommendations and mitigations.
                </p>
            </div>
            
            <!-- Application Interface -->
            <div class="app-interface">
                <!-- Input Section -->
                <div class="input-section">
                    <h2 class="section-title">
                        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                            <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/>
                            <path d="M12 8v4l3 3"/>
                        </svg>
                        Task Description
                    </h2>
                    
                    <form id="safety-form" @submit.prevent="handleSubmit">
                        <div class="form-group">
                            <label for="task" class="form-label">Describe the Job Task or Operation</label>
                            <textarea
                                id="task"
                                v-model="task"
                                name="task"
                                class="form-textarea"
                                placeholder="Example: Maintenance technician performing electrical troubleshooting on 480V motor control center while plant is in operation. Task involves voltage testing, component replacement, and system verification."
                                required
                            ></textarea>
                        </div>

                        <div class="sample-tasks">
                            <button type="button" class="sample-btn" @click="insertSampleTask('electrical')">Sample: Electrical Work</button>
                            <button type="button" class="sample-btn" @click="insertSampleTask('confined')">Sample: Confined Space</button>
                            <button type="button" class="sample-btn" @click="insertSampleTask('hotwork')">Sample: Hot Work</button>
                        </div>

                        <button type="submit" class="evaluate-btn" :disabled="loading">
                            <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" style="margin-right: 0.5rem;">
                                <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/>
                                <path d="M12 8v4l3 3"/>
                            </svg>
                            {{ loading ? 'Evaluating...' : 'Evaluate Safety' }}
                        </button>
                    </form>
                </div>
                
                <!-- Results Section -->
                <div class="results-section">
                    <h2 class="section-title">
                        <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                            <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/>
                            <path d="M12 8v4l3 3"/>
                        </svg>
                        Safety Assessment
                    </h2>
                    
                    <!-- Loading State -->
                    <div v-if="loading" style="text-align: center; padding: 3rem;">
                        <div class="loading-spinner"></div>
                        <p style="color: var(--text-secondary); margin-top: 1rem;">Analyzing safety hazards...</p>
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
                    <div v-else-if="results" id="safety-results">
                        <!-- Hazard Level Badge -->
                        <div style="margin-bottom: 2rem;">
                            <div :class="['hazard-badge', `hazard-${results.hazardLevel.toLowerCase()}`]">
                                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                    <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/>
                                </svg>
                                {{ results.hazardLevel }} HAZARD LEVEL
                            </div>
                        </div>

                        <!-- Identified Hazards -->
                        <div class="result-section">
                            <h3 class="result-heading">
                                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                    <path d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"/>
                                    <path d="M12 9v4m0 4h.01"/>
                                </svg>
                                Identified Hazards
                            </h3>
                            <div class="hazards-grid">
                                <div
                                    v-for="(hazard, index) in results.hazards"
                                    :key="index"
                                    class="hazard-card"
                                >
                                    <div class="hazard-name">{{ hazard.name }}</div>
                                    <div class="hazard-details">
                                        <span :class="['severity-tag', getSeverityClass(hazard.severity)]">
                                            {{ hazard.severity }}
                                        </span>
                                        <span class="probability-tag">
                                            {{ hazard.probability }} Probability
                                        </span>
                                    </div>
                                </div>
                            </div>
                        </div>

                        <!-- Recommended Mitigations -->
                        <div class="result-section">
                            <h3 class="result-heading">
                                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                    <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/>
                                    <path d="M9 12l2 2 4-4"/>
                                </svg>
                                Recommended Mitigations
                            </h3>
                            <ul class="steps-list">
                                <li v-for="(mitigation, index) in results.mitigations" :key="index">
                                    {{ mitigation }}
                                </li>
                            </ul>
                        </div>

                        <!-- Relevant Standards -->
                        <div class="result-section">
                            <h3 class="result-heading">
                                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                    <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
                                    <path d="M14 2v6h6M16 13H8m8 4H8m8 4H8"/>
                                </svg>
                                Relevant Standards
                            </h3>
                            <div class="standards-tags">
                                <span
                                    v-for="(standard, index) in results.standards"
                                    :key="index"
                                    class="standard-tag"
                                >
                                    {{ standard }}
                                </span>
                            </div>
                        </div>
                    </div>

                    <!-- Empty State -->
                    <div v-else style="text-align: center; color: var(--text-secondary); padding: 2rem;">
                        <svg width="64" height="64" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1" style="margin: 0 auto 1rem; opacity: 0.5;">
                            <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"/>
                            <path d="M12 8v4l3 3"/>
                        </svg>
                        <p>Describe a job task above and click "Evaluate Safety" to receive AI-powered safety assessment and recommendations.</p>
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
        background: linear-gradient(135deg, rgba(255, 107, 53, 0.1), rgba(0, 212, 255, 0.1));
        border-radius: 12px;
        border: 1px solid rgba(255, 107, 53, 0.2);
    }
    
    .page-title {
        font-family: 'Orbitron', monospace;
        font-size: 2.5rem;
        font-weight: 900;
        color: var(--accent-orange);
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
        border: 1px solid rgba(255, 107, 53, 0.2);
        border-radius: 12px;
        padding: 2rem;
    }
    
    .section-title {
        font-size: 1.5rem;
        font-weight: 700;
        color: var(--accent-orange);
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
        min-height: 200px;
        padding: 1rem;
        background: rgba(26, 26, 26, 0.8);
        border: 1px solid rgba(255, 107, 53, 0.3);
        border-radius: 8px;
        color: var(--text-primary);
        font-family: 'Inter', sans-serif;
        font-size: 1rem;
        resize: vertical;
        transition: all 0.3s ease;
    }
    
    .form-textarea:focus {
        outline: none;
        border-color: var(--accent-orange);
        box-shadow: 0 0 0 3px rgba(255, 107, 53, 0.1);
    }
    
    .form-textarea::placeholder {
        color: var(--text-secondary);
        font-style: italic;
    }
    
    .sample-tasks {
        margin-top: 1rem;
    }
    
    .sample-btn {
        background: rgba(0, 212, 255, 0.2);
        color: var(--accent-cyan);
        border: 1px solid var(--accent-cyan);
        padding: 0.5rem 1rem;
        border-radius: 6px;
        font-size: 0.9rem;
        cursor: pointer;
        transition: all 0.3s ease;
        margin-right: 0.5rem;
        margin-bottom: 0.5rem;
    }
    
    .sample-btn:hover {
        background: var(--accent-cyan);
        color: white;
    }
    
    .evaluate-btn {
        background: linear-gradient(135deg, var(--accent-orange), #cc5200);
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
    
    .evaluate-btn:hover {
        transform: translateY(-2px);
        box-shadow: 0 10px 25px rgba(255, 107, 53, 0.3);
    }
    
    .evaluate-btn:disabled {
        opacity: 0.6;
        cursor: not-allowed;
        transform: none;
        box-shadow: none;
    }
    
    /* Results Section */
    .results-section {
        background: rgba(45, 45, 45, 0.8);
        border: 1px solid rgba(255, 107, 53, 0.2);
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
        border-bottom: 1px solid rgba(255, 107, 53, 0.2);
    }
    
    .result-header h3 {
        font-size: 1.5rem;
        font-weight: 700;
        color: var(--accent-orange);
    }
    
    .hazard-badge {
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

    .hazard-high {
        background: linear-gradient(135deg, var(--error-red), #cc0000);
        border: 2px solid var(--error-red);
        box-shadow: 0 4px 15px rgba(255, 77, 77, 0.3);
    }

    .hazard-medium {
        background: linear-gradient(135deg, var(--accent-orange), #cc5200);
        border: 2px solid var(--accent-orange);
        box-shadow: 0 4px 15px rgba(255, 107, 53, 0.3);
    }

    .hazard-low {
        background: linear-gradient(135deg, var(--success-green), #00994d);
        border: 2px solid var(--success-green);
        box-shadow: 0 4px 15px rgba(0, 204, 102, 0.3);
    }

    .result-section {
        margin-bottom: 2.5rem;
        padding: 1.5rem;
        background: rgba(26, 26, 26, 0.5);
        border-radius: 10px;
        border: 1px solid rgba(255, 107, 53, 0.2);
    }

    .result-heading {
        font-size: 1.2rem;
        font-weight: 700;
        color: var(--accent-orange);
        margin-bottom: 1.5rem;
        display: flex;
        align-items: center;
        gap: 0.5rem;
    }

    .hazards-grid {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
        gap: 1rem;
    }

    .hazard-card {
        background: rgba(255, 107, 53, 0.1);
        border: 1px solid rgba(255, 107, 53, 0.3);
        border-radius: 10px;
        padding: 1.25rem;
        transition: all 0.3s ease;
    }

    .hazard-card:hover {
        transform: translateY(-2px);
        box-shadow: 0 4px 15px rgba(255, 107, 53, 0.2);
        border-color: var(--accent-orange);
    }

    .hazard-details {
        display: flex;
        gap: 0.75rem;
        margin-top: 0.75rem;
        flex-wrap: wrap;
    }

    .severity-tag {
        padding: 0.35rem 0.9rem;
        border-radius: 20px;
        font-size: 0.8rem;
        font-weight: 700;
        text-transform: uppercase;
        letter-spacing: 0.5px;
    }

    .severity-critical,
    .severity-high {
        background: var(--error-red);
        color: white;
    }

    .severity-medium {
        background: var(--accent-orange);
        color: white;
    }

    .severity-low {
        background: var(--success-green);
        color: var(--primary-bg);
    }

    .probability-tag {
        padding: 0.35rem 0.9rem;
        background: rgba(0, 212, 255, 0.2);
        border: 1px solid var(--accent-cyan);
        border-radius: 20px;
        font-size: 0.8rem;
        font-weight: 600;
        color: var(--accent-cyan);
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

    .standards-tags {
        display: flex;
        flex-wrap: wrap;
        gap: 0.75rem;
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
        border: 4px solid rgba(255, 107, 53, 0.2);
        border-top: 4px solid var(--accent-orange);
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
    
    .hazard-list {
        display: flex;
        flex-direction: column;
        gap: 0.75rem;
    }
    
    .hazard-item {
        display: grid;
        grid-template-columns: 2fr 1fr 1fr;
        gap: 1rem;
        background: rgba(255, 107, 53, 0.1);
        border: 1px solid rgba(255, 107, 53, 0.3);
        border-radius: 8px;
        padding: 1rem;
        align-items: center;
    }
    
    .hazard-name {
        font-weight: 600;
        color: var(--text-primary);
    }
    
    .hazard-severity {
        padding: 0.25rem 0.75rem;
        border-radius: 20px;
        font-size: 0.8rem;
        font-weight: 600;
        text-align: center;
    }
    
    .severity-critical {
        background: var(--error-red);
        color: white;
    }
    
    .severity-high {
        background: var(--accent-orange);
        color: white;
    }
    
    .severity-medium {
        background: #ffaa00;
        color: var(--primary-bg);
    }
    
    .severity-low {
        background: var(--success-green);
        color: var(--primary-bg);
    }
    
    .hazard-probability {
        padding: 0.25rem 0.75rem;
        background: rgba(0, 212, 255, 0.2);
        border-radius: 20px;
        font-size: 0.8rem;
        font-weight: 600;
        text-align: center;
        color: var(--accent-cyan);
    }
    
    .mitigation-list {
        list-style: none;
        padding: 0;
    }
    
    .mitigation-list li {
        background: rgba(0, 212, 255, 0.1);
        border-left: 3px solid var(--accent-cyan);
        padding: 0.75rem 1rem;
        margin-bottom: 0.5rem;
        border-radius: 0 6px 6px 0;
    }
    
    .standards-list {
        display: flex;
        flex-wrap: wrap;
        gap: 0.5rem;
    }
    
    .standard-tag {
        background: rgba(255, 107, 53, 0.2);
        color: var(--accent-orange);
        padding: 0.5rem 1rem;
        border-radius: 20px;
        font-size: 0.9rem;
        font-weight: 600;
        border: 1px solid var(--accent-orange);
    }
    
    /* Loading State */
    .loading-spinner {
        display: flex;
        align-items: center;
        justify-content: center;
        gap: 1rem;
        padding: 2rem;
        color: var(--accent-orange);
    }
    
    .spinner {
        width: 30px;
        height: 30px;
        border: 3px solid rgba(255, 107, 53, 0.3);
        border-top: 3px solid var(--accent-orange);
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
        
        .hazard-item {
            grid-template-columns: 1fr;
            gap: 0.5rem;
        }
    }
</style>