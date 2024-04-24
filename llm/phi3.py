import torch
from transformers import AutoModelForCausalLM, AutoTokenizer
from flask import Flask, request
import logging

app = Flask(__name__)

torch.set_default_device("cpu")

model = AutoModelForCausalLM.from_pretrained("microsoft/Phi-3-mini-4k-instruct", torch_dtype=torch.float32, device_map="cpu", trust_remote_code=True)
tokenizer = AutoTokenizer.from_pretrained("microsoft/Phi-3-mini-4k-instruct", trust_remote_code=True)

@app.route('/generate', methods=['POST'])
def generate():
    input = request.data.decode('utf-8')
    logging.error(f"Received input: {input}")
    output = generate_output(input)
    logging.error(f"Generated output: {output}")
    return {'output': output}

def generate_output(text):
    inputs = tokenizer(text, return_tensors="pt", return_attention_mask=False)
    outputs = model.generate(**inputs, max_length=1000)
    text = tokenizer.batch_decode(outputs)[0]
    return text

@app.route('/health')
def health():
    return {'status': 'healthy'}

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=80)