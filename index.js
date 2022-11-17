const questions = require('./questions.json');
const categories = [...new Set(questions.map(question => question.category))];
const howMany = 3;

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

// print out x random questions from each of the categories
categories.forEach(category => {
    console.log(`## ${category}`);
    console.log('');
    for (let i = 0; i < howMany; i++) {
        console.log(getRandomQuestionByCategory(category).question);
        console.log('');
    }
}
);