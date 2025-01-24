# Codename: Scavenger Hunt

## A local LLM chatbot that helps you find things.

A work in progress to be sure. The purpose of this project is to develop a chatbot that can
execute functions. My plan is to set up a sort of scavenger hunt in a file system and have the
bot open files and follow instructions. The bot will be able to read and write to files.

This project is built in Go with the Fyne GUI toolkit and uses the model of your choice through
Ollama. You just need to set up a .env with the model you want to use.

For example:
`model = deepseek-r1:latest`
