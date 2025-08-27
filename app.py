from flask import Flask, render_template, request, jsonify
import google.generativeai as genai
import os
from dotenv import load_dotenv

# load environment variables
load_dotenv()

app = Flask(__name__)

genai.configure(api_key=os.getenv('GEMINI_API_KEY'))
model = genai.GenerativeModel('gemini-1.5-flash')

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/search', methods=['POST'])
def search():
    try:
        data = request.get_json()
        equipment = data.get('equipment', '')
        problem = data.get('problem', '')
        error_code = data.get('error_code', '')
        
        if not equipment or not problem:
            return jsonify({'error': 'Equipment and problem description are required'}), 400
        
        prompt = f"""You are an expert Process Control System Technician at Aramco. 

Analyze the following troubleshooting request and provide a clear, structured response:

EQUIPMENT: {equipment}
PROBLEM: {problem}
ERROR CODE: {error_code if error_code else 'None provided'}

IMPORTANT SAFETY GUIDELINES:
- Always prioritize safety over production
- Follow Aramco safety protocols
- Verify equipment isolation before maintenance
- Use proper PPE and safety equipment
- Never bypass safety systems
- Follow lockout/tagout procedures

Please provide your response in this exact format:

ANALYSIS:
[Your analysis of the problem]

POSSIBLE CAUSES:
1. [Cause 1]
2. [Cause 2]
3. [Cause 3]

TROUBLESHOOTING STEPS:
1. [Step 1 - include safety precautions]
2. [Step 2 - include safety precautions]
3. [Step 3 - include safety precautions]
4. [Continue as needed]

SAFETY WARNINGS:
- [Important safety warning 1]
- [Important safety warning 2]
- [Important safety warning 3]

EQUIPMENT NOTES:
[Specific considerations for this equipment type]

Provide your response in a clear, structured format that a technician can follow safely."""

        # generate a response from gemini
        response = model.generate_content(prompt)
        
        response_text = response.text

        # print(response_text)
        
        sections = {
            'analysis': '',
            'causes': [],
            'steps': [],
            'safety_warnings': [],
            'equipment_notes': ''
        }
        
        lines = response_text.split('\n')
        current_section = None
        
        for line in lines:
            line = line.strip()
            if not line:
                continue
                
            if 'ANALYSIS:' in line.upper():
                current_section = 'analysis'
                sections['analysis'] = line.replace('ANALYSIS:', '').strip()
            elif 'POSSIBLE CAUSES:' in line.upper():
                current_section = 'causes'
            elif 'TROUBLESHOOTING STEPS:' in line.upper():
                current_section = 'steps'
            elif 'SAFETY WARNINGS:' in line.upper():
                current_section = 'safety_warnings'
            elif 'EQUIPMENT NOTES:' in line.upper():
                current_section = 'equipment_notes'
            elif line.startswith(('1.', '2.', '3.', '4.', '5.', '6.', '7.', '8.', '9.')):
                if current_section == 'causes':
                    sections['causes'].append(line)
                elif current_section == 'steps':
                    sections['steps'].append(line)
            elif line.startswith('-') and current_section == 'safety_warnings':
                sections['safety_warnings'].append(line)
            elif current_section == 'equipment_notes' and line:
                sections['equipment_notes'] += line + ' '
            elif current_section == 'analysis' and line and not sections['analysis']:
                sections['analysis'] = line
        
        sections['equipment_notes'] = sections['equipment_notes'].strip()
        
        return jsonify({
            'success': True,
            'equipment': equipment,
            'response': sections
        })
        
    except Exception as e:
        print(f'An error occurred: {str(e)}')
        return jsonify({'error': f'An error occurred: {str(e)}'}), 500

@app.route('/equipment')
def get_equipment():
    equipment_list = [
        "Control Valve",
        "Pressure Gauge",
        "Displacement Level Transmitter",
        "Flow Transmitter", 
        "Pressure Transmitter",
        "Temperature Transmitter",
        "Level Transmitter",
        "Control Panel",
        "SCADA System",
        "PLC",
        "DCS",
        "Safety Instrumented System (SIS)",
        "Emergency Shutdown System (ESD)",
        "Fire and Gas Detection System",
        "Compressor",
        "Pump",
        "Turbine",
        "Heat Exchanger",
        "Separator",
        "Storage Tank",
        "Pipeline",
        "Valve Actuator"
    ]
    return jsonify({'equipment': equipment_list})

if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=5000)
