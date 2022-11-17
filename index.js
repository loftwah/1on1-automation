// read questions.json file and set it to a constant
const questions = require('./questions.json');

// create a function that will return a random question
function getRandomQuestion() {
    const randomIndex = Math.floor(Math.random() * questions.length);
    return questions[randomIndex];
    }

// read all of the category values from the questions.json file and set it to a constant
const categories = questions.map(question => question.category);

// print all of the categories to the console
console.log(categories);

// create a function that will return a random question from the conversation starters category
function getRandomConversationStarter(category) {
    const conversationStarters = questions.filter(question => question.category === category);
    const randomIndex = Math.floor(Math.random() * conversationStarters.length);
    return conversationStarters[randomIndex];
    }

// create the same as above but if the category isn't present, return a random question
function getRandomQuestionByCategory(category) {
    const conversationStarters = questions.filter(question => question.category === category);
    if (conversationStarters.length === 0) {
        return getRandomQuestion();
    }
    const randomIndex = Math.floor(Math.random() * conversationStarters.length);
    return conversationStarters[randomIndex];
    }

// print two random questions from each category to the console
categories.forEach(category => {
    console.log(getRandomQuestionByCategory(category));
    console.log(getRandomQuestionByCategory(category));
    }
);
