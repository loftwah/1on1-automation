# One-on-One Weekly Question Generator

![1on1](images/1on1.png)

This repository is for making 1on1 meetings easy, and automated. It is inspired by a combination of a couple of different related repositories and I have mostly just joined the ideas together.

I have a video example [here](https://www.youtube.com/watch?v=ZpLRhfRjJWQ) of how it works.

This project is designed to automate the process of generating weekly 1-on-1 questions for use in GitHub issues. It selects random questions from a set of categories and updates a template file used for creating weekly 1-on-1 issues. This project includes a Ruby script for selecting questions and two GitHub Actions workflows to automate the process.

## Table of Contents

* [Overview](#overview)
* [Prerequisites](#prerequisites)
* [Getting Started](#getting-started)
* [File Structure](#file-structure)
* [Customization](#customization)
* [License](#license)

## Overview

The project uses a Ruby script (`app.rb`) to read questions from a JSON file (`questions.json`) and randomly selects a specified number of questions from each category. The selected questions are then inserted into the 1-on-1 issue template (`1-on-1-template.md`) found in the `.github/ISSUE_TEMPLATE` directory.

Two GitHub Actions workflows are used to automate this process:

1. `weekly-update.yml` - Updates the issue template with newly selected questions each week.
2. `weekly-1on1.yml` - Creates a new issue using the updated template each week.

These workflows are scheduled to run at specific times on Sundays.

## Prerequisites

* A GitHub account
* A repository for storing the project files
* Basic knowledge of Ruby and GitHub Actions

## Getting Started

1. Fork or clone this repository to your GitHub account.
2. Customize the `questions.json` file with your own set of questions and categories.
3. Modify the `.github/workflows/weekly-update.yml` and `.github/workflows/weekly-1on1.yml` files if necessary, e.g., to change the scheduled times.
4. Push your changes to your repository.
5. The GitHub Actions workflows will start running at the specified times, updating the issue template and creating new issues weekly.

## File Structure

* `app.rb` - Ruby script for selecting random questions and updating the issue template.
* `questions.json` - JSON file containing the questions and their categories.
* `.github/ISSUE_TEMPLATE/1-on-1-template.md` - The issue template used for creating 1-on-1 issues.
* `.github/workflows/weekly-update.yml` - GitHub Actions workflow for updating the issue template.
* `.github/workflows/weekly-1on1.yml` - GitHub Actions workflow for creating new issues using the updated template.

## Customization

To customize the project for your specific use case, you can:

1. Update the `questions.json` file with your own set of questions and categories.
2. Modify the `how_many` variable in the `app.rb` script to change the number of questions selected per category.
3. Adjust the scheduled times for the GitHub Actions workflows in the `.github/workflows/weekly-update.yml` and `.github/workflows/weekly-1on1.yml` files.
4. Customize the issue template in the `.github/ISSUE_TEMPLATE/1-on-1-template.md` file as needed.

## License

This project is released under the [MIT License](LICENSE).
