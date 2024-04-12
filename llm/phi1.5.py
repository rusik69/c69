import torch
from transformers import AutoModelForCausalLM, AutoTokenizer
from flask import Flask, request
import logging

app = Flask(__name__)

torch.set_default_device("cpu")

model = AutoModelForCausalLM.from_pretrained("microsoft/phi-1_5", torch_dtype="auto", trust_remote_code=True)
tokenizer = AutoTokenizer.from_pretrained("microsoft/phi-1_5", trust_remote_code=True)

@app.route('/generate', methods=['POST'])
def generate():
    input = request.get_json()['input']
    logging.info(f"Received input: {input}")
    output = generate_output(input)
    logging.info(f"Generated output: {output}")
    return {'output': output}

def generate_output(text):
    inputs = tokenizer(text, return_tensors="pt", return_attention_mask=False)
    outputs = model.generate(**inputs, max_length=200)
    text = tokenizer.batch_decode(outputs)[0]
    return text

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=80)