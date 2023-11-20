# 🌟 AI-Enhanced One-on-One Meeting Assistant 🌟

![Banner](https://github.com/loftwah/1on1-automation/assets/19922556/bba4de93-2ab2-4843-81cd-ace907f0c12a)

Welcome to the 🚀 AI-Enhanced One-on-One Meeting Assistant, a groundbreaking tool designed to transform the dynamics of manager-employee interactions into opportunities for deep engagement and professional development.

## 📊 Overview

In today's fast-paced workplace, effective communication is crucial for a productive and harmonious work environment. Our AI-Enhanced One-on-One Meeting Assistant is crafted to enhance this communication by equipping both managers and employees with insightful, context-driven conversation starters and comprehensive reports. It's not just a tool; it's a catalyst for meaningful dialogue and actionable insights.

## 🌈 Features

- **🔍 Personalized Question Generation**: Leverage AI to generate contextually relevant questions that guide and enrich your conversations.
- **💬 Interactive AI Chat**: Engage in preparatory discussions with an AI bot, helping you explore various workplace topics.
- **📝 Insightful Reports**: Receive post-chat summaries encapsulating key discussion points, action items, and insights.
- **🔗 Multi-platform Integration**: Seamlessly integrate with Slack and email for easy access and sharing.
- **🔐 Privacy and Security**: Highest confidentiality standards in handling sensitive discussions.

## 💡 Benefits

- **🎯 Optimized Meeting Preparation**: Enter your meetings well-prepared, ensuring efficient and focused discussions.
- **⏱️ Time-Saving**: Streamline meeting preparations and follow-ups, saving valuable time.
- **🎙️ Enhanced Communication**: Promote clarity and understanding, leading to improved workplace relationships.
- **📈 Actionable Feedback**: Gain practical insights for professional development and team dynamics.
- **🔄 Flexible Interaction**: Share insights in diverse communication styles, accommodating preferences.

## 🚀 Getting Started

### Prerequisites

- **🔑 OpenAI API Key**: An OpenAI key to access AI functionalities.
- **💻 Go Environment**: Setup for Go programming language.
- **💬 Slack Account**: For Slack channel integration.
- **📧 Email Server Access**: For email report distribution.
- **🔊 Echo Framework**: Setup Echo framework for handling web server requests.

### Installation

Clone the repository:

```bash
git clone https://github.com/loftwah/1on1-automation.git
cd 1on1-automation
```

Install dependencies:

```bash
go get .
```

### Configuration

Set up your environment variables:

```env
OPENAI_KEY=YOUR_API_KEY
AWS_ACCESS_KEY_ID=YOUR_ACCESS_KEY_ID
AWS_SECRET_ACCESS_KEY=YOUT_SECRET_ACCESS_KEY
SLACK_BOT_TOKEN=YOUR_SLACK_BOT_TOKEN
```

### Running the Application

Execute the following command to run the application:

```bash
go run main.go
```

### Testing the AI-Enhanced One-on-One Meeting Assistant Endpoints

#### Testing `generate-questions` Endpoint

This endpoint generates a set of questions for a one-on-one meeting. It does not require any input parameters and can be tested using a simple GET request.

**Command:**

```bash
curl http://localhost:1323/generate-questions
```

**Expected Response:** A list of questions formatted in Markdown, designed to facilitate a one-on-one meeting.

#### Testing `generate-report` Endpoint

This endpoint generates a summary report based on the responses from a one-on-one meeting. It requires the meeting responses to be sent as URL-encoded text in the POST request.

**Command:**

```bash
curl -X POST http://localhost:1323/generate-report \
     -H "Content-Type: application/x-www-form-urlencoded" \
     -d "responses=The%20new%20marketing%20project%20is%20shaping%20up%20well%2C%20facing%20some%20team%20alignment%20challenges.%20Interested%20in%20improving%20leadership%20skills.%20Creating%20content%20strategies%20is%20most%20satisfying.%20Time%20management%20is%20a%20current%20hurdle.%20Enjoy%20the%20team%27s%20collaborative%20spirit.%20Keen%20on%20exploring%20product%20management.%20Hiking%20helps%20maintain%20work-life%20balance.%20Interested%20in%20learning%20about%20R%26D.%20Proud%20of%20the%20positive%20client%20feedback%20on%20our%20last%20campaign.%20Would%20advise%20past%20self%20to%20be%20open%20to%20feedback.%20Interested%20in%20a%20monthly%20team%20recognition%20program.%20Digital%20marketing%20analytics%20was%20a%20fascinating%20learning%20area.%20Planning%20a%20road%20trip%20next%20month.%20Feel%20empowered%20at%20level%208%2C%20seeking%20more%20decision-making%20opportunities.%20%27Never%20stop%20learning%27%20has%20been%20pivotal%20advice."
```

**Expected Response:** A comprehensive report summarizing key points and actionable items from the provided responses, formatted in plain text or Markdown.

**Note:** The responses in the `generate-report` command are an example. You should replace them with actual responses from a one-on-one meeting when testing this endpoint.

## 📖 Usage

### Starting a Conversation

Invoke the AI chatbot through Slack commands or by sending an email to the bot's address. Use the provided conversation starters to kick off your interaction.

### Receiving and Sharing Reports

After your chat, the AI will craft a report summarizing the key points. Share this automatically via Slack or email, or choose to distribute it manually.

## 🤝 Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

## 📜 License

Distributed under the MIT License. See `LICENSE` for more information.

## 🙏 Acknowledgments

- **OpenAI**: For the AI engine powering our application.
- **Slack**: For their robust API and platform support.
- **Echo Framework**: For managing web server requests efficiently.
- **Contributors**: To all who contribute to making this project a success.
