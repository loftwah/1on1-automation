require "json"
require "set"

questions = JSON.parse(File.read("questions.json"))
categories = questions.map { |question| question["category"] }.to_set.to_a
how_many = 3
template_path = ".github/ISSUE_TEMPLATE/1-on-1-directreport-template.md"

def get_random_question(questions)
    random_index = rand(questions.length)
    questions[random_index]
    end

def get_random_question_by_category(questions, category)
    random_question = questions.select { |question| question["category"] == category }
    if random_question.length == 0
        return get_random_question(questions)
    end
    random_index = rand(random_question.length)
    random_question[random_index]
end

payload = categories.map do |category|
    these_questions = []
    how_many.times do
        these_questions << get_random_question_by_category(questions, category)
    end
    { category => these_questions }
end

template = File.read(template_path)
template.gsub!(/<!-- start of questions -->.*<!-- end of questions -->/m, "<!-- start of questions -->\n#{payload.map do |category|
    "## #{category.keys[0]}\n\n#{category.values[0].map { |question| "- #{question["question"]}" }.join("\n\n")}\n"
end
.join("\n")}\n<!-- end of questions -->")

File.write(template_path, template)