const questions = require("./questions.json");
const categories = [...new Set(questions.map((question) => question.category))];
const howMany = 3;

// function that will return a random question
function getRandomQuestion() {
  const randomIndex = Math.floor(Math.random() * questions.length);
  return questions[randomIndex];
}

// as above but if the category isn't present, return a random question
function getRandomQuestionByCategory(category) {
  const randomQuestion = questions.filter(
    (question) => question.category === category
  );
  if (randomQuestion.length === 0) {
    return getRandomQuestion();
  }
  const randomIndex = Math.floor(Math.random() * randomQuestion.length);
  return randomQuestion[randomIndex];
}

// set payload to x random questions from each of the categories following the above format
let payload = categories.map((category) => {
  const questions = [];
  for (let i = 0; i < howMany; i++) {
    questions.push(getRandomQuestionByCategory(category).question);
  }
  return {
    category,
    questions,
  };
});

// clear anything between <!-- start of questions --> and <!-- end of questions --> with a new line between them in .github/ISSUE_TEMPLATE/1-on-1-directreport-template.md
// then add the payload to the file
const fs = require("fs");
const path = require("path");
const filePath = path.join(
  __dirname,
  ".github/ISSUE_TEMPLATE/1-on-1-directreport-template.md"
);
const file = fs
  .readFileSync(filePath, "utf8")
  .replace(
    /<!-- start of questions -->[\s\S]*<!-- end of questions -->/g,
    `<!-- start of questions -->\n\n${payload
      .map(
        (category) =>
          `## ${category.category}\n\n${category.questions
            .map((question) => `- ${question}`)
            .join("\n\n")}`
      )
      .join("\n\n")}\n\n<!-- end of questions -->`
  );

fs.writeFileSync(filePath, file, "utf8");
