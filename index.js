const questions = require('./questions.json');
const categories = [...new Set(questions.map(question => question.category))];

// function that will return a random question
function getRandomQuestion() {
    const randomIndex = Math.floor(Math.random() * questions.length);
    return questions[randomIndex];
    }

// as above but if the category isn't present, return a random question
function getRandomQuestionByCategory(category) {
    const randomQuestion = questions.filter(question => question.category === category);
    if (randomQuestion.length === 0) {
        return getRandomQuestion();
    }
    const randomIndex = Math.floor(Math.random() * randomQuestion.length);
    return randomQuestion[randomIndex];
    }

// ## Category
// 
// - question
//
// - question
//
// print two random questions from each category to the console following the above format
categories.forEach(category => {
    console.log(`## ${category}`);
    console.log('');
    console.log(getRandomQuestionByCategory(category).question);
    console.log('');
    console.log(getRandomQuestionByCategory(category).question);
    console.log('');
    }
);