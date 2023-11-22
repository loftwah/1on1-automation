# 🌟 AI-Enhanced One-on-One Meeting Assistant 🌟

![Banner](https://github.com/loftwah/1on1-automation/assets/19922556/bba4de93-2ab2-4843-81cd-ace907f0c12a)

👋 Welcome to the 🚀 **AI-Enhanced One-on-One Meeting Assistant**! This innovative tool is your go-to assistant for sparking engaging and productive one-on-one meetings. Whether you're a manager or an employee, get ready to dive into deep, meaningful conversations with AI-powered assistance!

## 📊 Overview

In the hustle of today’s workplace, forging effective communication is key. 🗝️ Our tool is crafted not just to converse but to inspire! It's here to provide personalized prompts and suggestions, making every manager-employee interaction an avenue for profound connection and growth. 🌱

## 🌈 Features

- **🔍 Customized Conversation Starters**: AI magic at your fingertips! Get tailor-made questions that resonate with your meeting's context.
- **💬 Interactive AI Prep Chat**: Warm up with our AI chatbot! Discuss various workplace scenarios and get in the right mindset.
- **🤔 Suggestion Generation**: After your chat, our AI compiles insightful suggestions on topics to bring up and strategies to approach them. It's like having a conversation coach in your pocket!
- **🔗 Easy Integration**: Seamlessly connect with Slack and email for a smooth, uninterrupted experience.
- **🔐 Top-Notch Privacy**: We take confidentiality seriously, ensuring your discussions stay private and secure.

## 💡 Benefits

- **🎯 Precise Meeting Prep**: Walk into your meetings fully prepared, with a clear idea of what to discuss.
- **⏱️ Time Efficiency**: Cut down on preparation time and get straight to the heart of the matter.
- **🎙️ Clearer Communication**: Foster understanding and openness for healthier workplace dynamics.
- **📈 Practical Insights**: Transform feedback into tangible steps for personal and team development.
- **🔄 Versatile Sharing**: Distribute insights in various formats to suit everyone’s style.

## 🚀 Getting Started

### Prerequisites

- **🔑 OpenAI API Key**: To tap into the power of AI.
- **💻 Go Environment**: Your foundation for running the application.
- **💬 Slack Account**: For easy chat integration.
- **📧 Email Server Access**: To send and receive insights via email.
- **🔊 Echo Framework**: For efficient web server handling.

### Installation

Get the gears turning:

```bash
git clone https://github.com/loftwah/1on1-automation.git
cd 1on1-automation
go get .
```

### Configuration

Set your stage:

```env
OPENAI_KEY=YOUR_API_KEY
AWS_ACCESS_KEY_ID=YOUR_ACCESS_KEY_ID
AWS_SECRET_ACCESS_KEY=YOUT_SECRET_ACCESS_KEY
SLACK_BOT_TOKEN=YOUR_SLACK_BOT_TOKEN
```

### Running the Application

Bring it to life:

```bash
air
```

### Testing the Endpoints

#### 🧐 `generate-questions` Endpoint

Get personalized questions for your next one-on-one.

**Command:**

```bash
curl http://localhost:1323/generate-questions
```

**Expected Response:** A curated list of questions in Markdown, ready to make your meeting a hit.

#### 💡 `generate-suggestions` Endpoint

Turn your chat responses into actionable suggestions.

**Command:**

```bash
curl -X POST http://localhost:1323/generate-suggestions \
     -H "Content-Type: application/x-www-form-urlencoded" \
     -d "responses=<your_responses>"
```

**Expected Response:** A detailed breakdown of topics and approaches, formatted for clarity.

**Note:** Replace `<your_responses>` with real-time chat inputs for best results.

## 📖 Usage

### Initiating the Chat

Jumpstart your meeting prep through Slack commands or an email to the bot. Let the AI guide you with engaging conversation starters.

### Embracing and Sharing the Suggestions

Post-chat, the AI whips up suggestions tailored to your discussion. Share them instantly over Slack, email, or manually, as you prefer.

## 🤝 Contributing

Join our open-source journey! Your contributions add immense value and are **highly cherished**. 🌟

## 📜 License

This gem is under the MIT License. Check out `LICENSE` for the details.

## 🙏 Acknowledgments

Big thanks to:

- **OpenAI**: For the AI brains behind the scenes.
- **Slack**: For their robust platform that keeps us connected.
- **Echo Framework**: For making web requests a breeze.
- **Contributors**: You're the heroes who keep this project soaring!
