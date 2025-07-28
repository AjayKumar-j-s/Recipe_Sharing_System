 🧠 Recipe Sharing App with LLM-Based Nutrition Analysis

This is a recipe sharing web application built in Go using the Gin framework. Users can register, log in, and share their recipes securely. It also includes an exciting feature: Nutrition Analysis using a Local LLM or API (like Hugging Face or your own inference server).

---

 🚀 Features

- 🔐 JWT-based Authentication
- 🍳 Add, View, Update, Delete Recipes
- 📊 Nutrition Analysis using Language Model (LLM)
- 🧪 Secure API integration using environment variables
- 🖥️ Built with Go + Gin

---
 🧠 LLM Integration for Nutrition Analysis

We’ve integrated a Local or Remote LLM that analyzes a recipe's ingredients and returns estimated nutritional values (calories, proteins, fats, vitamins, etc.).

🔌 Modes of LLM Usage

You can configure the app to use one of the following:

✅ 1. Hugging Face Inference API
- Uses remote LLMs like Mistral via Hugging Face.
- Requires a free [Hugging Face API token](https://huggingface.co/settings/tokens).
- Set the token using an environment variable:
  ```bash
  export HF_API_KEY=hf_your_token_here
✅ 2. Local LLM via Python Inference Server
Run a FastAPI or Flask server on localhost:8000.

Use any transformer-based model locally.

🧪 Endpoint: Nutrition Analysis
http
POST /recipe/:id/nutrition
Authorization: Bearer <your JWT token>
✅ Response
json
{
  "nutrition": "This recipe contains approximately 400 calories, 20g protein, 15g fat, 50g carbs..."
}


🔐 Environment Variables
Create a .env file (optional, for local dev):
HF_API_KEY=hf_your_token_here


🛠️ Tech Stack
Backend: Go (Gin)

Authentication: JWT

LLM Inference: Hugging Face or local Python server

Database: (Your choice — SQLite/PostgreSQL/MySQL)

📧 Contact
Author: Ajay Kumar
Email: ajayak2306@gmail.com
GitHub: AjayKumar-j-s

