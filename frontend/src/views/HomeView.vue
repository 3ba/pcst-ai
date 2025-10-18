<script setup lang="ts">
import { useRouter } from 'vue-router'
import { onMounted, onUnmounted, ref } from 'vue'
import * as anime from 'animejs'
import P5 from 'p5'
import Typed from 'typed.js'

const router = useRouter()
let p5Instance: any = null
let typedInstance: any = null

// Initialize particle background
const initParticles = () => {
  p5Instance = new P5((p: any) => {
    let particles: any[] = []

    p.setup = () => {
      const canvas = p.createCanvas(p.windowWidth, p.windowHeight)
      const container = document.getElementById('particle-container')
      if (container) {
        canvas.parent(container)
      }

      // Create particles
      for (let i = 0; i < 50; i++) {
        particles.push({
          x: p.random(p.width),
          y: p.random(p.height),
          vx: p.random(-0.5, 0.5),
          vy: p.random(-0.5, 0.5),
          size: p.random(1, 3),
          opacity: p.random(0.3, 0.8)
        })
      }
    }

    p.draw = () => {
      p.clear()

      // Update and draw particles
      particles.forEach((particle) => {
        particle.x += particle.vx
        particle.y += particle.vy

        // Wrap around edges
        if (particle.x < 0) particle.x = p.width
        if (particle.x > p.width) particle.x = 0
        if (particle.y < 0) particle.y = p.height
        if (particle.y > p.height) particle.y = 0

        // Draw particle
        p.fill(102, 126, 234, particle.opacity * 255)
        p.noStroke()
        p.circle(particle.x, particle.y, particle.size)
      })

      // Draw connections
      particles.forEach((p1, i) => {
        particles.slice(i + 1).forEach((p2) => {
          const distance = p.dist(p1.x, p1.y, p2.x, p2.y)
          if (distance < 100) {
            p.stroke(102, 126, 234, (1 - distance / 100) * 50)
            p.strokeWeight(0.5)
            p.line(p1.x, p1.y, p2.x, p2.y)
          }
        })
      })
    }

    p.windowResized = () => {
      p.resizeCanvas(p.windowWidth, p.windowHeight)
    }
  })
}

// Initialize typewriter effect
const initTypewriter = () => {
  const element = document.getElementById('hero-text')
  if (element) {
    typedInstance = new Typed('#hero-text', {
      strings: [
        'AI for Process Control Systems',
        'Intelligent Industrial Innovation',
        'Smart Troubleshooting Solutions',
        'Predictive Maintenance AI'
      ],
      typeSpeed: 50,
      backSpeed: 30,
      backDelay: 2000,
      loop: true,
      showCursor: false,
      cursorChar: '|'
    })
  }
}

// Initialize scroll animations
const initScrollAnimations = () => {
  const observerOptions = {
    threshold: 0.1,
    rootMargin: '0px 0px -50px 0px'
  }

  const observer = new IntersectionObserver((entries) => {
    entries.forEach((entry) => {
      if (entry.isIntersecting) {
        const element = entry.target

        anime.animate(element, {
          opacity: [0, 1],
          translateY: [30, 0],
          duration: 800,
          easing: 'out-quad'
        })

        observer.unobserve(element)
      }
    })
  }, observerOptions)

  document.querySelectorAll('.animate-on-scroll').forEach((el: any) => {
    el.style.opacity = '0'
    observer.observe(el)
  })
}

onMounted(() => {
  initParticles()
  initTypewriter()

  setTimeout(() => {
    initScrollAnimations()
  }, 100)

  // Add smooth scrolling for all anchor links
  document.querySelectorAll('a[href^="#"]').forEach((anchor) => {
    anchor.addEventListener('click', function (e) {
      e.preventDefault()
      const target = document.querySelector(this.getAttribute('href') as string)
      if (target) {
        target.scrollIntoView({
          behavior: 'smooth',
          block: 'start'
        })
      }
    })
  })
})

onUnmounted(() => {
  // Clean up animations
  if (p5Instance) {
    p5Instance.remove()
  }
  if (typedInstance) {
    typedInstance.destroy()
  }
})
</script>

<template>
  <head>
    
    <!-- Fonts -->
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700&family=JetBrains+Mono:wght@400;500&family=Orbitron:wght@400;700;900&display=swap" rel="stylesheet">
  </head>
  <body>
    <!-- Background Effects -->
    <div id="particle-container"></div>
    <div class="bg-grid"></div>
    
    <!-- Navigation -->
    <nav class="navbar">
        <div class="nav-container">
            <a href="/" class="logo">PCST-AI</a>
            <ul class="nav-links">
                <li><a href="#about">About</a></li>
                <li><a href="#projects">Projects</a></li>
                <li><a href="#vision">Vision</a></li>
                <li><a href="#contact">Contact</a></li>
            </ul>
        </div>
    </nav>
    
    <!-- Hero Section -->
    <section class="hero">
        <div class="hero-content">
            <h1 id="hero-text">AI for Process Control Systems</h1>
            <p class="hero-subtitle">
                Empowering technicians with AI-driven tools for smarter troubleshooting, safer operations, and faster decisions.
            </p>
            <div class="hero-buttons">
                <a href="#projects" class="btn btn-primary">
                    <span>Explore Projects</span>
                    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M7 17L17 7M17 7H7M17 7V17"/>
                    </svg>
                </a>
                <a href="#about" class="btn btn-secondary">
                    <span>Learn About AI in PCST</span>
                    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <circle cx="12" cy="12" r="10"/>
                        <path d="M12 16v-4M12 8h.01"/>
                    </svg>
                </a>
            </div>
        </div>
    </section>
    
    <!-- About Section -->
    <section id="about" class="about animate-on-scroll">
        <h2 class="section-title">Where Industry Meets Intelligence</h2>
        <div class="about-content">
            <div class="about-text">
                <p>
                    Process Control Systems Technicians (PCSTs) in oil, gas, and industrial environments face increasingly complex challenges. 
                    Our AI-powered portal transforms how technicians diagnose problems, assess risks, and maintain critical infrastructure.
                </p>
                <p>
                    By integrating artificial intelligence with decades of industrial expertise, we provide intelligent tools that enhance 
                    diagnostic accuracy, improve safety outcomes, and accelerate decision-making in high-stakes environments.
                </p>
                <p>
                    From predictive maintenance to real-time safety assessment, our suite of AI applications empowers technicians with 
                    the insights they need to keep operations running smoothly and safely.
                </p>
            </div>
            <div class="about-image">
                <img src="../assets/hero-industrial.jpg" alt="Industrial AI Control Center" />
            </div>
        </div>
    </section>
    
    <!-- Projects Section -->
    <section id="projects" class="projects animate-on-scroll">
        <div class="projects-container">
            <h2 class="section-title">AI-Powered Industrial Solutions</h2>
            <div class="projects-grid">
                <!-- Virtual Control Room Advisor -->
                <div class="project-card">
                    <div class="project-icon">🎛️</div>
                    <h3 class="project-title">Virtual Control Room Advisor</h3>
                    <p class="project-description">
                        Analyzes control room alarms and operational data to recommend corrective actions, 
                        predict equipment failures, and optimize process parameters in real-time.
                    </p>
                    <a href="/vcra" class="try-it-btn">Try It Now</a>
                </div>
                
                <!-- AI Safety Advisor -->
                <div class="project-card">
                    <div class="project-icon">🛡️</div>
                    <h3 class="project-title">AI Safety Advisor</h3>
                    <p class="project-description">
                        Evaluates job tasks for potential hazards and suggests comprehensive safety mitigations, 
                        PPE requirements, and procedural improvements.
                    </p>
                    <a href="/safetyadvisor" class="try-it-btn">Try It Now</a>
                </div>
                
                <!-- CorrosionAI -->
                <div class="project-card">
                    <div class="project-icon">🔬</div>
                    <h3 class="project-title">CorrosionAI</h3>
                    <p class="project-description">
                        Predicts corrosion risk using process data, material properties, and environmental conditions 
                        to optimize maintenance schedules and prevent pipeline failures.
                    </p>
                    <a href="/corrosionai" class="try-it-btn">Try It Now</a>
                </div>
                
                <!-- PCST Troubleshooting Assistant -->
                <div class="project-card">
                    <div class="project-icon">🔧</div>
                    <h3 class="project-title">PCST Troubleshooting Assistant</h3>
                    <p class="project-description">
                        AI-driven assistant that helps troubleshoot field equipment step-by-step, providing detailed 
                        diagnostics, safety warnings, and maintenance procedures.
                    </p>
                    <a href="/troubleshooting" class="try-it-btn">Try It Now</a>
                </div>
            </div>
        </div>
    </section>
    
    <!-- Vision Section -->
    <section id="vision" class="vision animate-on-scroll">
        <h2 class="section-title">Integrating AI into Industrial Workflows</h2>
        <div class="vision-grid">
            <div class="vision-card">
                <div class="vision-icon">🧠</div>
                <h3 class="vision-title">Predictive Intelligence</h3>
                <p class="vision-description">
                    Advanced machine learning algorithms analyze historical and real-time data to predict equipment failures, 
                    optimize maintenance schedules, and prevent costly downtime.
                </p>
            </div>
            
            <div class="vision-card">
                <div class="vision-icon">⚡</div>
                <h3 class="vision-title">Safety Automation</h3>
                <p class="vision-description">
                    Automated safety assessments and real-time hazard detection ensure worker protection and regulatory compliance 
                    across all industrial operations.
                </p>
            </div>
            
            <div class="vision-card">
                <div class="vision-icon">💾</div>
                <h3 class="vision-title">Knowledge Retention</h3>
                <p class="vision-description">
                    AI systems capture and preserve institutional knowledge, ensuring critical expertise is retained and accessible 
                    to new generations of technicians.
                </p>
            </div>
            
            <div class="vision-card">
                <div class="vision-icon">🎯</div>
                <h3 class="vision-title">Operational Efficiency</h3>
                <p class="vision-description">
                    Intelligent automation reduces manual tasks, streamlines workflows, and enables technicians to focus on 
                    high-value activities that require human expertise.
                </p>
            </div>
        </div>
    </section>
    
    <!-- Footer -->
    <footer id="contact" class="footer">
        <div class="footer-content">
            <p class="footer-text">
                Developed to explore the fusion of AI and Process Control Systems.
            </p>
            <div class="social-links">
                <a href="https://github.com/3ba" target="_blank" class="social-link" title="GitHub">
                    <svg width="20" height="20" viewBox="0 0 24 24" fill="currentColor">
                        <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
                    </svg>
                </a>
                <a href="#" target="_blank" class="social-link" title="LinkedIn">
                    <svg width="20" height="20" viewBox="0 0 24 24" fill="currentColor">
                        <path d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433c-1.144 0-2.063-.926-2.063-2.065 0-1.138.92-2.063 2.063-2.063 1.14 0 2.064.925 2.064 2.063 0 1.139-.925 2.065-2.064 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z"/>
                    </svg>
                </a>
                <a href="#" target="_blank" class="social-link" title="Contact">
                    <svg width="20" height="20" viewBox="0 0 24 24" fill="currentColor">
                        <path d="M0 3v18h24v-18h-24zm6.623 7.929l-4.623 5.712v-9.458l4.623 3.746zm-4.141-5.929h19.035l-9.517 7.713-9.518-7.713zm5.694 7.188l3.824 3.099 3.83-3.104 5.612 6.817h-18.779l5.513-6.812zm9.208-1.264l4.616-3.741v9.348l-4.616-5.607z"/>
                    </svg>
                </a>
            </div>
        </div>
    </footer>
  </body>
</template>

<style>
  /* Background effects */
  #particle-container {
      position: fixed;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
      z-index: -1;
      pointer-events: none;
  }

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

  .nav-links {
      display: flex;
      gap: 2rem;
      list-style: none;
  }

  .nav-links a {
      color: var(--text-secondary);
      text-decoration: none;
      transition: color 0.3s ease;
      font-weight: 500;
  }

  .nav-links a:hover {
      color: var(--accent-cyan);
  }

  .hero {
      min-height: 100vh;
      display: flex;
      align-items: center;
      justify-content: center;
      text-align: center;
      position: relative;
      padding: 2rem;
      background: linear-gradient(135deg, rgba(26, 26, 26, 0.9) 0%, rgba(45, 45, 45, 0.8) 100%);
  }

  .hero-content {
      max-width: 1000px;
      z-index: 2;
  }

  .hero h1 {
      font-family: 'Orbitron', monospace;
      font-size: clamp(2.5rem, 5vw, 4rem);
      font-weight: 900;
      margin-bottom: 1rem;
      background: linear-gradient(135deg, var(--accent-cyan), var(--accent-orange));
      -webkit-background-clip: text;
      -webkit-text-fill-color: transparent;
      background-clip: text;
  }

  .hero-subtitle {
      font-size: 1.25rem;
      color: var(--text-secondary);
      margin-bottom: 2rem;
      max-width: 600px;
      margin-left: auto;
      margin-right: auto;
  }

  .hero-buttons {
      display: flex;
      gap: 1rem;
      justify-content: center;
      flex-wrap: wrap;
  }

  .btn {
      padding: 1rem 2rem;
      border: none;
      border-radius: 8px;
      font-size: 1rem;
      font-weight: 600;
      text-decoration: none;
      display: inline-flex;
      align-items: center;
      gap: 0.5rem;
      cursor: pointer;
      transition: all 0.3s ease;
      position: relative;
      overflow: hidden;
  }

  .btn-primary {
      background: linear-gradient(135deg, var(--accent-cyan), #0099cc);
      color: white;
  }

  .btn-secondary {
      background: transparent;
      color: var(--accent-cyan);
      border: 2px solid var(--accent-cyan);
  }

  .btn:hover {
      transform: translateY(-2px);
      box-shadow: 0 10px 25px rgba(0, 212, 255, 0.3);
  }

  .btn-primary:hover {
      background: linear-gradient(135deg, #0099cc, var(--accent-cyan));
  }

  .btn-secondary:hover {
      background: var(--accent-cyan);
      color: var(--primary-bg);
  }

  /* About section */
  .about {
      padding: 6rem 2rem;
      max-width: 1200px;
      margin: 0 auto;
  }

  .section-title {
      font-family: 'Orbitron', monospace;
      font-size: 2.5rem;
      text-align: center;
      margin-bottom: 3rem;
      color: var(--accent-cyan);
  }

  .about-content {
      display: grid;
      grid-template-columns: 1fr 1fr;
      gap: 4rem;
      align-items: center;
  }

  .about-text {
      font-size: 1.1rem;
      line-height: 1.8;
      color: var(--text-secondary);
  }

  .about-image {
      position: relative;
      border-radius: 12px;
      overflow: hidden;
      box-shadow: 0 20px 40px rgba(0, 212, 255, 0.1);
  }

  .about-image img {
      width: 100%;
      height: 300px;
      object-fit: cover;
  }

  .about-image::after {
      content: '';
      position: absolute;
      top: 0;
      left: 0;
      right: 0;
      bottom: 0;
      background: linear-gradient(135deg, rgba(0, 212, 255, 0.1), rgba(255, 107, 53, 0.1));
  }

  /* Projects section */
  .projects {
      padding: 6rem 2rem;
      background: rgba(45, 45, 45, 0.3);
  }

  .projects-container {
      max-width: 1400px;
      margin: 0 auto;
  }

  .projects-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
      gap: 2rem;
      margin-top: 3rem;
  }

  .project-card {
      background: rgba(45, 45, 45, 0.8);
      border: 1px solid rgba(0, 212, 255, 0.2);
      border-radius: 12px;
      padding: 2rem;
      transition: all 0.3s ease;
      cursor: pointer;
      position: relative;
      overflow: hidden;
  }

  .project-card::before {
      content: '';
      position: absolute;
      top: 0;
      left: -100%;
      width: 100%;
      height: 100%;
      background: linear-gradient(90deg, transparent, rgba(0, 212, 255, 0.1), transparent);
      transition: left 0.5s ease;
  }

  .project-card:hover::before {
      left: 100%;
  }

  .project-card:hover {
      transform: translateY(-5px);
      border-color: var(--accent-cyan);
      box-shadow: 0 15px 35px rgba(0, 212, 255, 0.2);
  }

  .project-icon {
      width: 60px;
      height: 60px;
      background: linear-gradient(135deg, var(--accent-cyan), var(--accent-orange));
      border-radius: 12px;
      display: flex;
      align-items: center;
      justify-content: center;
      margin-bottom: 1.5rem;
      font-size: 1.5rem;
  }

  .project-title {
      font-size: 1.5rem;
      font-weight: 700;
      margin-bottom: 1rem;
      color: var(--text-primary);
  }

  .project-description {
      color: var(--text-secondary);
      margin-bottom: 2rem;
      line-height: 1.6;
  }

  .try-it-btn {
      background: linear-gradient(135deg, var(--accent-cyan), #0099cc);
      color: white;
      border: none;
      padding: 0.75rem 1.5rem;
      border-radius: 6px;
      font-weight: 600;
      cursor: pointer;
      transition: all 0.3s ease;
      text-decoration: none;
      display: inline-block;
  }

  .try-it-btn:hover {
      transform: translateY(-2px);
      box-shadow: 0 8px 20px rgba(0, 212, 255, 0.3);
  }

  /* Vision section */
  .vision {
      padding: 6rem 2rem;
      max-width: 1200px;
      margin: 0 auto;
  }

  .vision-grid {
      display: grid;
      grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
      gap: 2rem;
      margin-top: 3rem;
  }

  .vision-card {
      text-align: center;
      padding: 2rem;
      background: rgba(45, 45, 45, 0.5);
      border-radius: 12px;
      border: 1px solid rgba(0, 212, 255, 0.1);
      transition: all 0.3s ease;
  }

  .vision-card:hover {
      border-color: var(--accent-cyan);
      transform: translateY(-3px);
  }

  .vision-icon {
      width: 80px;
      height: 80px;
      background: linear-gradient(135deg, var(--accent-orange), #cc5200);
      border-radius: 50%;
      display: flex;
      align-items: center;
      justify-content: center;
      margin: 0 auto 1.5rem;
      font-size: 2rem;
  }

  .vision-title {
      font-size: 1.25rem;
      font-weight: 700;
      margin-bottom: 1rem;
      color: var(--text-primary);
  }

  .vision-description {
      color: var(--text-secondary);
      line-height: 1.6;
  }

  /* Footer */
  .footer {
      background: rgba(26, 26, 26, 0.9);
      border-top: 1px solid rgba(0, 212, 255, 0.2);
      padding: 2rem;
      text-align: center;
  }

  .footer-content {
      max-width: 1200px;
      margin: 0 auto;
  }

  .footer-text {
      color: var(--text-secondary);
      margin-bottom: 1rem;
  }

  .social-links {
      display: flex;
      justify-content: center;
      gap: 1rem;
  }

  .social-link {
      width: 40px;
      height: 40px;
      background: rgba(0, 212, 255, 0.1);
      border: 1px solid var(--accent-cyan);
      border-radius: 50%;
      display: flex;
      align-items: center;
      justify-content: center;
      color: var(--accent-cyan);
      text-decoration: none;
      transition: all 0.3s ease;
  }

  .social-link:hover {
      background: var(--accent-cyan);
      color: var(--primary-bg);
      transform: translateY(-2px);
  }

  /* Responsive design */
  @media (max-width: 768px) {
      .nav-links {
          display: none;
      }
      
      .hero h1 {
          font-size: 2.5rem;
      }
      
      .about-content {
          grid-template-columns: 1fr;
          gap: 2rem;
      }
      
      .hero-buttons {
          flex-direction: column;
          align-items: center;
      }
      
      .projects-grid {
          grid-template-columns: 1fr;
      }
  }

  /* Animation classes */
  .animate-on-scroll {
      opacity: 0;
      transform: translateY(30px);
  }

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

  #hero-text {
      display: inline-block;
      white-space: nowrap;
  }
</style>
