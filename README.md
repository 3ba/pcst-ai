# PCST-AI Simple Setup Guide

### This is a simplified version with just Python, Flask, and Google Gemini API.

## PREREQUISITES
- Python 3.8 or higher
- Google Cloud account with Gemini API access

## SETUP STEPS

1. Install Python dependencies:
   pip install -r requirements.txt

2. Create environment file:
   Copy env.example to .env and add your Gemini API key:
   GEMINI_API_KEY=your_actual_api_key_here

3. Run the application:
   python app.py

4. Open your browser and go to:
   http://localhost:5000

## HOW IT WORKS

1. User selects equipment type from dropdown
2. User describes the problem in the text area
3. User can optionally add an error code
4. User clicks "Search for Solution"
5. The app sends the request to Google Gemini API
6. Gemini AI analyzes the problem and provides:
   - Problem analysis
   - Possible causes
   - Step-by-step troubleshooting
   - Safety warnings
   - Equipment-specific notes
7. Results are displayed in a clean, organized format

## FEATURES
- Simple search bar interface
- Equipment type selection
- Problem description input
- Optional error code input
- AI-powered troubleshooting solutions
- Safety-focused responses
- Clean results display

## EQUIPMENTS SUPPORTED
- Control systems (valves, panels, SCADA, PLC, DCS)
- Sensors and transmitters (flow, pressure, temperature, level)
- Safety systems (SIS, ESD, fire detection)
- Process equipment (compressors, pumps, turbines, etc.)


-----
#### This project addresses real challenges faced by PCSTs at Aramco:
- Provides immediate access to expert troubleshooting knowledge
- Ensures safety protocols are followed during maintenance
- Reduces training time for new technicians
- Demonstrates practical application of AI in industrial settings

The simple, focused approach shows how AI can be effectively integrated into critical industrial processes while maintaining ease of use and safety focus.

Built for Aramco ITC Students and the Future of Process Control

