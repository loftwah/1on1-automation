# 🌟 AI-Enhanced One-on-One Meeting Assistant 🌟

![Banner](https://github.com/loftwah/1on1-automation/assets/19922556/bba4de93-2ab2-4843-81cd-ace907f0c12a)

Step into the future of workplace communication with the AI-Enhanced One-on-One Meeting Assistant. This transformative tool is engineered to revolutionize the dynamics of manager-employee interactions, paving the way for deeper engagement and impactful professional development.

## 📊 Project Overview

In the rapidly evolving work environment of today, effective communication is paramount. Our AI-Enhanced One-on-One Meeting Assistant is a game-changer, providing insightful, AI-driven conversation starters and comprehensive reports. It's more than a tool - it's your partner in fostering meaningful dialogue and actionable insights.

## 🌈 Distinguished Features

- **🔍 AI-Driven Question Generation**: Experience AI's prowess in generating contextually relevant questions, enriching your conversations.
- **💬 Interactive AI Chat**: Prepare for meetings with AI guidance, exploring varied workplace topics.
- **📝 Insightful Reports**: Receive detailed summaries post-discussion, highlighting crucial points and actionable insights.
- **🔗 Multi-Platform Integration**: Effortlessly integrate with Slack and email for streamlined access and sharing.
- **🔐 Robust Privacy and Security**: Prioritize confidentiality in sensitive discussions with the highest security standards.

## 💡 Key Benefits

- **🎯 Precision in Meeting Preparation**: Enter meetings well-prepared, ensuring focused and efficient discussions.
- **⏱️ Time Efficiency**: Streamline meeting preparations and follow-ups, saving precious time.
- **🎙️ Enhanced Communication**: Enhance clarity and understanding, leading to improved workplace relationships.
- **📈 Actionable Feedback**: Garner valuable insights for professional and team development.
- **🔄 Adaptive Interaction**: Cater to diverse communication styles, accommodating individual preferences.

## 🚀 Getting Started

### Setting Up Your Environment

- **🔑 OpenAI API Key**: Obtain an OpenAI key to access AI functionalities.
- **💻 Go Environment**: Ensure Go programming language is set up.
- **💬 Slack Account**: Have a Slack account for channel integration.
- **📧 Email Access**: Ensure access to an email server for report distribution.
- **🔊 Echo Framework**: Install the Echo framework for web server request handling.

### Installation & Local Setup

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/loftwah/1on1-automation.git
   cd 1on1-automation
   ```

2. **Install Dependencies**:

   ```bash
   go get .
   ```

3. **Environment Configuration**: Set up your `.env` file with necessary API keys and tokens:

   ```env
   OPENAI_KEY=YOUR_API_KEY
   SLACK_BOT_TOKEN=YOUR_SLACK_BOT_TOKEN
   ```

4. **Running Locally**: Start the application by executing:

   ```bash
   go run cmd/server/main.go
   ```

### Docker Development Setup

1. **Build and Run with Docker**:

   ```bash
   docker build -t oneonone-assistant .
   docker run -p 1323:1323 oneonone-assistant
   ```

2. **Using Docker Compose for Development**:

   ```bash
   docker-compose up
   ```

   This will set up the development environment with hot reloading using Air.

### Testing Endpoints

- **Generate Questions**:

  ```bash
  curl http://localhost:1323/generate-questions
  ```

- **Generate Reports**:

  ```bash
  curl -X POST http://localhost:1323/generate-report -d "responses=Your encoded responses"
  ```

## 📖 Extensive Usage

### Starting Conversations

- **Through Slack**: Invoke the AI chatbot via Slack commands for an interactive prep session.
- **Via Email**: Send an email to the bot's address to kickstart the AI-guided conversation.

### Sharing Insights

- **Automatically**: Reports generated post-chat can be shared through Slack or email.
- **Manually**: Opt to distribute these insights manually for a more personal touch.

## 🤝 How to Contribute

Join the vibrant open-source community by contributing to this project. Every contribution, whether big or small, is immensely valued.

## 📜 License

This project is licensed under the MIT License. For more information, see the `LICENSE` file.

## 🙏 Acknowledgments

- **OpenAI**: For providing the AI engine.
- **Slack**: For their robust API support.
- **Echo Framework**: For efficient web server management.
- **Contributors**: Thank you to all who contribute to our collective success.

---

### Deploying to Production with AWS ECS

For deploying this application in a production environment, refer to the comprehensive guide in `ecs/README.md`. It covers the steps for setting up the infrastructure on AWS using ECS with Fargate, including Terraform scripts for a seamless deployment experience.

In this setup, you'll utilize essential AWS services like ECS, ECR, VPC, and ALB to ensure a scalable and secure deployment in the `ap-southeast-2` region. The guide provides detailed instructions for building and pushing Docker images, deploying infrastructure with Terraform, and verifying your setup post-deployment.

With this infrastructure-as-code approach, you can confidently manage and scale your AI-Enhanced One-on-One Meeting Assistant in a cloud environment.

## Conclusion

Embrace the innovation of AI in enhancing workplace interactions. With the AI-Enhanced One-on-One Meeting Assistant, you're not just adopting a tool; you're stepping into a new era of meaningful professional communication. Join us in transforming the landscape of manager-employee dialogues!
